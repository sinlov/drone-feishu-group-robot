package main

import (
	"fmt"
	"github.com/sinlov/drone-feishu-group-robot/feishu_plugin"
	"github.com/sinlov/drone-info-tools/drone_urfave_cli_v2"
	"github.com/sinlov/drone-info-tools/template"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

const (
	// Version of cli
	Version = "v1.3.1"
	Name    = "drone-feishu-group-robot"
)

func action(c *cli.Context) error {

	isDebug := c.Bool("config.debug")

	drone := drone_urfave_cli_v2.UrfaveCliBindDroneInfo(c)

	if isDebug {
		log.Printf("debug: cli version is %s", Version)
		log.Printf("debug: load droneInfo finish at link: %v\n", drone.Build.Link)
	}
	config := feishu_plugin.Config{
		Debug:               c.Bool("config.debug"),
		TimeoutSecond:       c.Int("config.timeout_second"),
		NtpTarget:           c.String("config.ntp_target"),
		Webhook:             c.String("config.webhook"),
		Secret:              c.String("config.secret"),
		FeishuEnableForward: c.Bool("config.feishu_enable_forward"),
		MsgType:             c.String("config.msg_type"),
		Title:               c.String("config.msg_title"),
		PoweredByImageKey:   c.String("config.msg_powered_by_image_key"),
		PoweredByImageAlt:   c.String("config.msg_powered_by_image_alt"),
	}

	ossHost := findStrFromCliOrCoverByEnv(c, "config.feishu_oss_host", feishu_plugin.EnvPluginFeishuOssHost)
	cardOss := feishu_plugin.CardOss{}
	if ossHost == "" {
		config.RenderOssCard = feishu_plugin.RenderStatusHide
	} else {
		config.RenderOssCard = feishu_plugin.RenderStatusShow
		cardOss.InfoSendResult = findStrFromCliOrCoverByEnv(c, "config.feishu_oss_info_send_result", feishu_plugin.EnvPluginFeishuOssInfoSendResult)
		cardOss.InfoUser = findStrFromCliOrCoverByEnv(c, "config.feishu_oss_info_user", feishu_plugin.EnvPluginFeishuOssInfoUser)
		cardOss.InfoPath = findStrFromCliOrCoverByEnv(c, "config.feishu_oss_info_path", feishu_plugin.EnvPluginFeishuOssInfoPath)
		cardOss.ResourceUrl = findStrFromCliOrCoverByEnv(c, "config.feishu_oss_resource_url", feishu_plugin.EnvPluginFeishuOssResourceUrl)
		cardOss.PageUrl = findStrFromCliOrCoverByEnv(c, "config.feishu_oss_page_url", feishu_plugin.EnvPluginFeishuOssPageUrl)
		ossPagePasswd := findStrFromCliOrCoverByEnv(c, "config.feishu_oss_page_passwd", feishu_plugin.EnvPluginFeishuOssPagePasswd)
		if ossPagePasswd == "" {
			cardOss.RenderResourceUrl = feishu_plugin.RenderStatusShow
		} else {
			cardOss.RenderResourceUrl = feishu_plugin.RenderStatusHide
			cardOss.PagePasswd = ossPagePasswd
		}
	}
	config.CardOss = cardOss

	if isDebug {
		log.Printf("config.timeout_second: %v", config.TimeoutSecond)
	}

	p := feishu_plugin.FeishuPlugin{
		Name:    Name,
		Version: Version,
		Drone:   drone,
		Config:  config,
	}
	err := p.Exec()

	if err != nil {
		log.Fatalf("err: %v", err)
		return err
	}
	return nil
}

func findStrFromCliOrCoverByEnv(c *cli.Context, ctxKey, envKey string) string {
	val := c.String(ctxKey)
	envVal, lookupEnv := os.LookupEnv(envKey)
	if lookupEnv {
		val = envVal
	}
	return val
}

// pluginFlag
// set plugin flag at here
func pluginFlag() []cli.Flag {
	return []cli.Flag{
		// plugin start
		&cli.BoolFlag{
			Name:    "config.debug,debug",
			Usage:   "debug mode",
			EnvVars: []string{"PLUGIN_DEBUG"},
		},
		&cli.IntFlag{
			Name:    "config.timeout_second,timeout_second",
			Usage:   "do request timeout setting second",
			EnvVars: []string{"PLUGIN_TIMEOUT_SECOND"},
		},
		&cli.StringFlag{
			Name:    "config.ntp_target,ntp_target",
			Usage:   "ntp target like: pool.ntp.org, time1.google.com,time.pool.aliyun.com, default not use ntpd to sync",
			EnvVars: []string{"PLUGIN_NTP_TARGET"},
		},
		&cli.StringFlag{
			Name:    "config.webhook,feishu_webhook",
			Usage:   "feishu webhook for send message",
			EnvVars: []string{feishu_plugin.EnvPluginFeishuWebhook},
		},
		&cli.StringFlag{
			Name:    "config.secret,feishu_secret",
			Usage:   "feishu secret",
			EnvVars: []string{feishu_plugin.EnvPluginFeishuSecret},
		},
		&cli.BoolFlag{
			Name:    "config.feishu_enable_forward,feishu_enable_forward",
			Usage:   "feishu message enable forward, default false",
			EnvVars: []string{feishu_plugin.EnvPluginFeishuEnableForward},
		},
		&cli.StringFlag{
			Name:    "config.msg_type,feishu_msg_type",
			Usage:   "feishu message type",
			EnvVars: []string{feishu_plugin.EnvPluginFeishuMsgType},
		},
		&cli.StringFlag{
			Name:    "config.msg_title,feishu_msg_title",
			Usage:   "feishu message title",
			EnvVars: []string{feishu_plugin.EnvPluginFeishuMsgTitle},
		},
		&cli.StringFlag{
			Name:    "config.msg_powered_by_image_key,feishu_msg_powered_by_image_key",
			Usage:   "feishu message powered by image key",
			EnvVars: []string{feishu_plugin.EnvPluginFeishuMsgPoweredByImageKey},
		},
		&cli.StringFlag{
			Name:    "config.msg_powered_by_image_alt,feishu_msg_powered_by_image_alt",
			Usage:   "feishu message powered by image alt",
			EnvVars: []string{feishu_plugin.EnvPluginFeishuMsgPoweredByImageAlt},
		},

		// oss card start
		&cli.StringFlag{
			Name:    "config.feishu_oss_host",
			Usage:   "feishu OSS host for show oss info, if empty will not show oss info",
			EnvVars: []string{feishu_plugin.EnvPluginFeishuOssHost},
		},
		&cli.StringFlag{
			Name:    "config.feishu_oss_info_send_result",
			Usage:   "feishu OSS user for show at card",
			EnvVars: []string{feishu_plugin.EnvPluginFeishuOssInfoSendResult},
		},
		&cli.StringFlag{
			Name:    "config.feishu_oss_info_user",
			Usage:   "feishu OSS user for show at card",
			EnvVars: []string{feishu_plugin.EnvPluginFeishuOssInfoUser},
		},
		&cli.StringFlag{
			Name:    "config.feishu_oss_info_path",
			Usage:   "feishu OSS path for show at card",
			EnvVars: []string{feishu_plugin.EnvPluginFeishuOssInfoPath},
		},
		&cli.StringFlag{
			Name:    "config.feishu_oss_resource_url",
			Usage:   "feishu OSS resource url",
			EnvVars: []string{feishu_plugin.EnvPluginFeishuOssResourceUrl},
		},
		&cli.StringFlag{
			Name:    "config.feishu_oss_page_url",
			Usage:   "feishu OSS page url",
			EnvVars: []string{feishu_plugin.EnvPluginFeishuOssPageUrl},
		},
		&cli.StringFlag{
			Name:    "config.feishu_oss_page_passwd",
			Usage:   "OSS password at page url, will hide PLUGIN_FEISHU_OSS_RESOURCE_URL when PAGE_PASSWD not empty",
			EnvVars: []string{feishu_plugin.EnvPluginFeishuOssPagePasswd},
		},
		// oss card end
		// plugin end
	}
}

func main() {
	template.RegisterSettings(template.DefaultFunctions)
	app := cli.NewApp()
	app.Version = Version
	app.Name = "Drone feishu Message FeishuPlugin"
	app.Usage = "Sending message to feishu group by robot using WebHook"
	year := time.Now().Year()
	app.Copyright = fmt.Sprintf("© 2022-%d sinlov", year)
	authorSinlov := &cli.Author{
		Name:  "sinlov",
		Email: "sinlovgmppt@gmail.com",
	}
	app.Authors = []*cli.Author{
		authorSinlov,
	}

	app.Action = action
	flags := drone_urfave_cli_v2.UrfaveCliAppendCliFlag(drone_urfave_cli_v2.DroneInfoUrfaveCliFlag(), pluginFlag())
	app.Flags = flags

	// kubernetes runner patch
	if _, err := os.Stat("/run/drone/env"); err == nil {
		errDotEnv := godotenv.Overload("/run/drone/env")
		if errDotEnv != nil {
			log.Fatalf("load /run/drone/env err: %v", errDotEnv)
		}
	}

	// app run as urfave
	if err := app.Run(os.Args); nil != err {
		log.Println(err)
	}
}
