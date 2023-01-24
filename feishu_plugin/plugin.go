package feishu_plugin

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sinlov/drone-feishu-group-robot/drone_info"
	"github.com/sinlov/drone-feishu-group-robot/feishu_message"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

type (
	// Plugin plugin all config
	Plugin struct {
		Drone                  drone_info.Drone
		Config                 Config
		SendTarget             SendTarget
		FeishuRobotMsgTemplate feishu_message.FeishuRobotMsgTemplate
	}
)

func (p *Plugin) Exec() error {
	if p.Config.Debug {
		for _, e := range os.Environ() {
			log.Println(e)
		}
	}

	if p.Config.Webhook == "" {
		msg := "missing feishu webhook, please set feishu webhook"
		return errors.New(msg)
	}

	// set default TimeoutSecond
	if p.Config.TimeoutSecond == 0 {
		p.Config.TimeoutSecond = 10
	}

	// set default MsgType
	if p.Config.MsgType == "" {
		p.Config.MsgType = msgTypeInteractive
	}

	if !(StrInArr(p.Config.MsgType, supportMsgType)) {
		return fmt.Errorf("feishu msg type only support %v", supportMsgType)
	}
	var err error

	sendTarget, err := p.fetchSendTarget()
	if err != nil {
		return err
	}

	p.SendTarget = sendTarget

	// try use ntpd to sync time
	if p.Config.NtpTarget != "" {
		command := exec.Command("ntpd", "-d", "-q", "-n", "-p", p.Config.NtpTarget)
		var stdOut bytes.Buffer
		var stdErr bytes.Buffer
		command.Stdout = &stdOut
		command.Stderr = &stdErr

		err = command.Run()
		if err != nil {
			return fmt.Errorf("run ntpd target %v stderr %v\nerr: %v", p.Config.NtpTarget, stdErr.String(), err)
		}
	}

	err = p.sendMessage()
	if err != nil {
		return err
	}

	return err
}

func (p *Plugin) fetchSendTarget() (SendTarget, error) {
	nowTimestamp := time.Now().Unix()
	if p.Config.Debug {
		log.Printf("fetchSendTarget nowTimestamp: %v\n", nowTimestamp)
	}
	sendTarget := SendTarget{
		Webhook: p.Config.Webhook,
		Secret:  p.Config.Secret,
	}

	ctxTemp := feishu_message.CtxTemp{}

	robotMsgTemplate := feishu_message.FeishuRobotMsgTemplate{
		Timestamp: nowTimestamp,
		MsgType:   p.Config.MsgType,
	}
	if sendTarget.Secret != "" {
		sign, err := feishu_message.GenSign(sendTarget.Secret, nowTimestamp)
		if err != nil {
			return sendTarget, err
		}
		robotMsgTemplate.Sign = sign
	}

	switch p.Config.MsgType {
	default:
		log.Printf("fetchSend msg type now not support %v", p.Config.MsgType)
		return sendTarget, fmt.Errorf("fetchSend msg type now not support %v", p.Config.MsgType)
	case msgTypeInteractive:
		cardTemp := (feishu_message.CardTemp{}).Build(
			p.Config.Title,
			p.Config.PoweredByImageKey,
			p.Config.PoweredByImageAlt,
		)

		ctxTemp.CardTemp = cardTemp
		robotMsgTemplate.CtxTemp = ctxTemp
		p.FeishuRobotMsgTemplate = robotMsgTemplate

		renderFeishuCard, err := RenderFeishuCard(defaultCardTemplate, p)
		if err != nil {
			return sendTarget, err
		}
		if p.Config.Debug {
			log.Printf("fetchSendTarget renderFeishuCard: %v\n", renderFeishuCard)
		}
		if renderFeishuCard != "" {
			sendTarget.FeishuRobotMeg = []byte(renderFeishuCard)
		}
	}

	return sendTarget, nil
}

func (p *Plugin) sendMessage() error {
	sendTarget := p.SendTarget
	var feishuUrl = fmt.Sprintf("%s/%s", feishu_message.ApiFeishuBotV2(), sendTarget.Webhook)
	if p.Config.Debug {
		log.Printf("sendMessage url: %v", feishuUrl)
	}
	req, err := http.NewRequest("POST", feishuUrl, bytes.NewBuffer(sendTarget.FeishuRobotMeg))
	if err != nil {
		return fmt.Errorf("sendMessage http NewRequest err: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Timeout: time.Second * time.Duration(p.Config.TimeoutSecond),
	}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("sendMessage http Do err: %v", err)
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Fatalf("sendMessage panic err: %v", err)
		}
	}()
	statusCode := resp.StatusCode
	if statusCode != http.StatusOK {
		return fmt.Errorf("sendMessage http status code: %v", statusCode)
	}
	body, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		return fmt.Errorf("sendMessage http read err: %v", errRead)
	}
	if p.Config.Debug {
		log.Println(statusCode, string(body))
	}
	var respApi feishu_message.ApiRespRotV2
	errUnmarshal := json.Unmarshal(body, &respApi)

	if errUnmarshal != nil {
		return fmt.Errorf("sendMessage http Unmarshal err: %v", errUnmarshal)
	}
	if respApi.Code != 0 {
		return fmt.Errorf("feishu message can not send by code [ %v ] err: %v", respApi.Code, respApi.Msg)
	}
	return nil
}
