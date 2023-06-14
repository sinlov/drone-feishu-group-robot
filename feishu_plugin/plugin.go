package feishu_plugin

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/antchfx/xmlquery"
	"github.com/sinlov/drone-feishu-group-robot/feishu_message"
	"github.com/sinlov/drone-feishu-group-robot/tools/drone_api"
	"github.com/sinlov/drone-info-tools/drone_info"
	tools "github.com/sinlov/drone-info-tools/tools/str_tools"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

const (
	EnvDroneSystemAdminToken = "PLUGIN_DRONE_SYSTEM_ADMIN_TOKEN"

	EnvPluginFeishuIgnoreLastSuccessByAdminTokenDistance = "PLUGIN_FEISHU_IGNORE_LAST_SUCCESS_BY_ADMIN_TOKEN_DISTANCE"
	EnvPluginFeishuIgnoreLastSuccessByBadges             = "PLUGIN_FEISHU_IGNORE_LAST_SUCCESS_BY_BADGES"
	EnvPluginFeishuIgnoreLastSuccessBranch               = "PLUGIN_FEISHU_IGNORE_LAST_SUCCESS_BRANCH"
	EnvPluginFeishuWebhook                               = "PLUGIN_FEISHU_WEBHOOK"
	EnvPluginFeishuSecret                                = "PLUGIN_FEISHU_SECRET"
	EnvPluginFeishuMsgTitle                              = "PLUGIN_FEISHU_MSG_TITLE"
	EnvPluginFeishuEnableForward                         = "PLUGIN_FEISHU_ENABLE_FORWARD"
	EnvPluginFeishuMsgType                               = "PLUGIN_FEISHU_MSG_TYPE"
	EnvPluginFeishuMsgPoweredByImageKey                  = "PLUGIN_FEISHU_MSG_POWERED_BY_IMAGE_KEY"
	EnvPluginFeishuMsgPoweredByImageAlt                  = "PLUGIN_FEISHU_MSG_POWERED_BY_IMAGE_ALT"

	EnvPluginFeishuOssHost           = "PLUGIN_FEISHU_OSS_HOST"
	EnvPluginFeishuOssInfoSendResult = "PLUGIN_FEISHU_OSS_INFO_SEND_RESULT"
	EnvPluginFeishuOssInfoUser       = "PLUGIN_FEISHU_OSS_INFO_USER"
	EnvPluginFeishuOssInfoPath       = "PLUGIN_FEISHU_OSS_INFO_PATH"
	EnvPluginFeishuOssResourceUrl    = "PLUGIN_FEISHU_OSS_RESOURCE_URL"
	EnvPluginFeishuOssPageUrl        = "PLUGIN_FEISHU_OSS_PAGE_URL"
	EnvPluginFeishuOssPagePasswd     = "PLUGIN_FEISHU_OSS_PAGE_PASSWD"
)

type (
	// FeishuPlugin plugin all config
	FeishuPlugin struct {
		Name                   string
		Version                string
		Drone                  drone_info.Drone
		Config                 Config
		SendTarget             SendTarget
		FeishuRobotMsgTemplate feishu_message.FeishuRobotMsgTemplate
		droneApiClient         *drone_api.DroneApiClient
	}
)

func (p *FeishuPlugin) Exec() error {
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
		p.Config.MsgType = MsgTypeInteractive
	}
	if !(tools.StrInArr(p.Config.MsgType, supportMsgType)) {
		return fmt.Errorf("config [ msg_type ] only support %v", supportMsgType)
	}

	// change p.Config.CardOss.InfoSendResult to default
	if p.Config.RenderOssCard == RenderStatusShow {
		if !(tools.StrInArr(p.Config.CardOss.InfoSendResult, drone_info.DroneBuildStatusStatusOptSupport())) {
			return fmt.Errorf("config [ feishu_oss_info_send_result ] only support %v", supportRenderStatus)
		}
		if p.Config.CardOss.InfoSendResult != RenderStatusShow {
			if p.Config.Debug {
				log.Printf("debug: in p.Config.RenderOssCard mode [ %s ] will set p.Config.CardOss.InfoSendResult to [ %s ] and change p.Drone.Build.Status to [ %s ]\n",
					RenderStatusShow, RenderStatusHide, RenderStatusHide,
				)
			}
			p.Config.CardOss.InfoSendResult = RenderStatusHide
			p.Drone.Build.Status = RenderStatusHide
		}
	}

	var err error
	if p.Config.DroneSystemAdminToken != "" {
		errFetchBuildInfo := p.fetchBuildInfoByAdminToken()
		if errFetchBuildInfo != nil {
			log.Printf("trt fetchBuildInfoByAdminToken fail and send message err: %v\n", errFetchBuildInfo)
			p.Drone.Build.Status = drone_info.DroneBuildStatusFailure
			p.Drone.Commit.Message = fmt.Sprintf("%s\n%s", p.Drone.Commit.Message, "fetchBuildInfoByAdminToken fail pleae check config")
			err = p.fetchInfoAndSend()
			if err != nil {
				return err
			}
			return errFetchBuildInfo
		}
	}

	if p.Config.IgnoreLastSuccessByAdminTokenDistance > 0 {
		if p.Config.DroneSystemAdminToken == "" {
			return fmt.Errorf("config [ %s ] is empty", EnvDroneSystemAdminToken)
		}
		pass, errIgnoreByAdminToken := p.checkIgnoreLastSuccessByAdminToken()
		if errIgnoreByAdminToken != nil {
			log.Printf("trt checkIgnoreLastSuccessByAdminToken fail and send message err: %v\n", errIgnoreByAdminToken)
		}
		if pass {
			log.Printf("=> plugin %s version %s", p.Name, p.Version)
			log.Printf("ignore send feishu group robot message at success by LastSuccessByAdminTokenDistance %d.\n", p.Config.IgnoreLastSuccessByAdminTokenDistance)
			return nil
		}
	}

	if p.Config.IgnoreLastSuccessByBadges {
		ignoreLastSuccess, errIgnore := p.ignoreLastBadgesSuccess()
		if errIgnore != nil {
			log.Printf("trt IgnoreLastSuccessByBadges fail and send message err: %v\n", errIgnore)
		}

		if ignoreLastSuccess {
			log.Printf("=> plugin %s version %s", p.Name, p.Version)
			log.Printf("ignore send feishu group robot message at success branch %s.\n", p.Config.IgnoreLastSuccessBadgesBranch)
			return nil
		}
	}

	err = p.fetchInfoAndSend()
	if err != nil {
		return err
	}
	return err
}

func (p *FeishuPlugin) fetchInfoAndSend() error {
	sendTarget, err := p.fetchSendTarget()
	if err != nil {
		return err
	}

	p.SendTarget = sendTarget

	// try use ntpd to sync time
	if p.Config.NtpTarget != "" {

		if p.Config.Debug {
			log.Printf("NtpTarget sync before by [%v] Unix time: %v\n", p.Config.NtpTarget, time.Now().Unix())
		}

		log.Printf("try to sync ntp by taget [%v]\n", p.Config.NtpTarget)
		command := exec.Command("ntpd", "-d", "-q", "-n", "-p", p.Config.NtpTarget)
		var stdOut bytes.Buffer
		var stdErr bytes.Buffer
		command.Stdout = &stdOut
		command.Stderr = &stdErr

		err = command.Run()
		if err != nil {
			return fmt.Errorf("run ntpd target %v stderr %v\nerr: %v", p.Config.NtpTarget, stdErr.String(), err)
		}

		if p.Config.Debug {
			log.Printf("NtpTarget sync after by [%v] Unix time: %v\n", p.Config.NtpTarget, time.Now().Unix())
		}
	}

	err = p.sendMessage()
	if err != nil {
		return err
	}
	log.Printf("=> plugin %s version %s", p.Name, p.Version)
	log.Printf("send feishu group robot message finish.\n")
	return nil
}

func (p *FeishuPlugin) checkIgnoreLastSuccessByAdminToken() (bool, error) {
	if p.Drone.Build.Status != drone_info.DroneBuildStatusSuccess {
		if p.Config.Debug {
			log.Printf("checkIgnoreLastSuccessByAdminToken false by now Drone Build status [ %v ]\n", p.Drone.Build.Status)
		}
		return false, nil
	}
	if p.droneApiClient == nil {
		return false, fmt.Errorf("not fetch droneApiClient")
	}
	distance := p.Drone.Build.Number - p.droneApiClient.LastRecordBuildSuccessNumber()
	if p.Config.Debug {
		log.Printf("checkIgnoreLastSuccessByAdminToken LastRecordBuildSuccessNumber [ %v ]\n", p.droneApiClient.LastRecordBuildSuccessNumber())
		log.Printf("checkIgnoreLastSuccessByAdminToken distance [ %v ] config [ %v ]\n", distance, uint64(p.Config.IgnoreLastSuccessByAdminTokenDistance))
		lastRecordRepoBuildSuccess := *p.droneApiClient.LastRecordRepoBuildSuccess()
		lastSuccessInfo, errMarshal := json.Marshal(lastRecordRepoBuildSuccess)
		if errMarshal != nil {
			return false, nil
		}
		var str bytes.Buffer
		err := json.Indent(&str, []byte(lastSuccessInfo), "", "    ")
		if err != nil {
			return false, nil
		}
		log.Printf("checkIgnoreLastSuccessByAdminToken lastSuccesBuild info\n%v\n", str.String())
	}
	if distance <= uint64(p.Config.IgnoreLastSuccessByAdminTokenDistance) {
		if p.Config.Debug {
			log.Printf("checkIgnoreLastSuccessByAdminToken true by distance [ %v < %v ]\n", distance, p.Config.IgnoreLastSuccessByAdminTokenDistance)
		}
		return true, nil
	}

	return false, nil
}

func (p *FeishuPlugin) ignoreLastBadgesSuccess() (bool, error) {
	if p.Drone.Build.Status != drone_info.DroneBuildStatusSuccess {
		return false, fmt.Errorf("drone build status is not success: %v", p.Drone.Build.Status)
	}

	var ignoreLastSuccessBranch = p.Drone.Build.Branch
	if ignoreLastSuccessBranch == "" {
		if p.Drone.Build.Tag != "" {
			log.Printf("ignoreLastBadgesSuccess false by tag: %v Branch is empty\n", p.Drone.Build.Tag)
			return false, nil
		}
	}
	if p.Config.IgnoreLastSuccessBadgesBranch == "" { // replace by now Build Branch
		p.Config.IgnoreLastSuccessBadgesBranch = p.Drone.Build.Branch
	}

	badgesStatus, err := p.fetchBadgesStatus()
	if err != nil {
		return false, err
	}

	// see https://github.com/harness/drone/blob/master/handler/api/badge/badge.go
	if badgesStatus == drone_info.DroneBuildStatusSuccess {
		if p.Config.Debug {
			log.Printf("ignoreLastBadgesSuccess by badgesStatus: %v", badgesStatus)
		}
		return true, nil
	}
	if badgesStatus == "started" {
		log.Printf("ignoreLastBadgesSuccess false badgesStatus: %v", badgesStatus)
		return false, nil
	}
	if badgesStatus == "none" {
		if p.Config.Debug {
			log.Printf("ignoreLastBadgesSuccess false badgesStatus: %v", badgesStatus)
		}
		return false, nil
	}

	if p.Config.Debug {
		log.Printf("ignoreLastBadgesSuccess false by badgesStatus: %v", badgesStatus)
	}

	return false, nil
}

func (p *FeishuPlugin) fetchBadgesStatus() (string, error) {

	if p.Config.IgnoreLastSuccessBadgesBranch == "" {
		return "", nil
	}
	badgeURL := fmt.Sprintf(
		"%s://%s/api/badges/%s/%s/status.svg?ref=refs/heads/%s",
		p.Drone.DroneSystem.Proto, p.Drone.DroneSystem.Host,
		p.Drone.Repo.GroupName, p.Drone.Repo.ShortName, p.Config.IgnoreLastSuccessBadgesBranch,
	)
	doc, err := xmlquery.LoadURL(badgeURL)
	if err != nil {
		return "", err
	}
	nodeText4, err := xmlquery.QueryAll(doc, "/svg/g/text[4]")
	if err != nil {
		return "", err
	}
	if len(nodeText4) == 0 {
		return "", fmt.Errorf("can not find /svg/g/text[4] in %s", badgeURL)
	}
	statusText := nodeText4[0].InnerText()
	return statusText, nil
}

func (p *FeishuPlugin) fetchSendTarget() (SendTarget, error) {
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
	case MsgTypeInteractive:
		cardTemp := (feishu_message.CardTemp{}).Build(
			p.Config.Title,
			p.Config.PoweredByImageKey,
			p.Config.PoweredByImageAlt,
		)

		cardTemp.EnableForward = p.Config.FeishuEnableForward
		ctxTemp.CardTemp = cardTemp
		robotMsgTemplate.CtxTemp = ctxTemp
		p.FeishuRobotMsgTemplate = robotMsgTemplate

		renderFeishuCard, err := RenderFeishuCard(DefaultCardTemplate, p)
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

func (p *FeishuPlugin) sendMessage() error {
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
		errBody, _ := ioutil.ReadAll(resp.Body)
		if errBody != nil {
			return fmt.Errorf("sendMessage http status code: %v , body: %v", statusCode, string(errBody))
		}
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

func (p *FeishuPlugin) fetchBuildInfoByAdminToken() error {
	droneApiClient := drone_api.NewDroneApiClient(
		fmt.Sprintf("%s://%s", p.Drone.DroneSystem.Proto, p.Drone.DroneSystem.Host),
		p.Config.DroneSystemAdminToken, uint(p.Config.TimeoutSecond), p.Config.Debug)
	err := droneApiClient.FetchBuildInfo(p.Drone.Repo.OwnerName, p.Drone.Repo.ShortName, 0)
	if err != nil {
		return err
	}
	p.droneApiClient = &droneApiClient
	return nil
}
