package feishu_plugin

import (
	"github.com/sinlov/drone-info-tools/drone_info"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func BindFlag(c *cli.Context, cliVersion, cliName string, drone drone_info.Drone) FeishuPlugin {
	config := Config{
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

	ossHost := findStrFromCliOrCoverByEnv(c, "config.feishu_oss_host", EnvPluginFeishuOssHost)
	cardOss := CardOss{
		Host: ossHost,
	}
	if ossHost == "" {
		config.RenderOssCard = RenderStatusHide
	} else {
		config.RenderOssCard = RenderStatusShow
		cardOss.InfoSendResult = findStrFromCliOrCoverByEnv(c, "config.feishu_oss_info_send_result", EnvPluginFeishuOssInfoSendResult)
		cardOss.InfoUser = findStrFromCliOrCoverByEnv(c, "config.feishu_oss_info_user", EnvPluginFeishuOssInfoUser)
		cardOss.InfoPath = findStrFromCliOrCoverByEnv(c, "config.feishu_oss_info_path", EnvPluginFeishuOssInfoPath)
		cardOss.ResourceUrl = findStrFromCliOrCoverByEnv(c, "config.feishu_oss_resource_url", EnvPluginFeishuOssResourceUrl)
		cardOss.PageUrl = findStrFromCliOrCoverByEnv(c, "config.feishu_oss_page_url", EnvPluginFeishuOssPageUrl)
		ossPagePasswd := findStrFromCliOrCoverByEnv(c, "config.feishu_oss_page_passwd", EnvPluginFeishuOssPagePasswd)
		if ossPagePasswd == "" {
			cardOss.RenderResourceUrl = RenderStatusShow
		} else {
			cardOss.RenderResourceUrl = RenderStatusHide
			cardOss.PagePasswd = ossPagePasswd
		}
	}
	config.CardOss = cardOss

	if config.Debug {
		log.Printf("config.timeout_second: %v", config.TimeoutSecond)
	}

	p := FeishuPlugin{
		Name:    cliName,
		Version: cliVersion,
		Drone:   drone,
		Config:  config,
	}
	return p
}

func findStrFromCliOrCoverByEnv(c *cli.Context, ctxKey, envKey string) string {
	val := c.String(ctxKey)
	envVal, lookupEnv := os.LookupEnv(envKey)
	if lookupEnv {
		val = envVal
	}
	return val
}

// Flag
// set plugin flag at here
func Flag() []cli.Flag {
	return []cli.Flag{
		// plugin start
		&cli.StringFlag{
			Name:    "config.webhook,feishu_webhook",
			Usage:   "feishu webhook for send message",
			EnvVars: []string{EnvPluginFeishuWebhook},
		},
		&cli.StringFlag{
			Name:    "config.secret,feishu_secret",
			Usage:   "feishu secret",
			EnvVars: []string{EnvPluginFeishuSecret},
		},
		&cli.BoolFlag{
			Name:    "config.feishu_enable_forward,feishu_enable_forward",
			Usage:   "feishu message enable forward, default false",
			EnvVars: []string{EnvPluginFeishuEnableForward},
		},
		&cli.StringFlag{
			Name:    "config.msg_type,feishu_msg_type",
			Usage:   "feishu message type",
			EnvVars: []string{EnvPluginFeishuMsgType},
		},
		&cli.StringFlag{
			Name:    "config.msg_title,feishu_msg_title",
			Usage:   "feishu message title",
			EnvVars: []string{EnvPluginFeishuMsgTitle},
		},
		&cli.StringFlag{
			Name:    "config.msg_powered_by_image_key,feishu_msg_powered_by_image_key",
			Usage:   "feishu message powered by image key",
			EnvVars: []string{EnvPluginFeishuMsgPoweredByImageKey},
		},
		&cli.StringFlag{
			Name:    "config.msg_powered_by_image_alt,feishu_msg_powered_by_image_alt",
			Usage:   "feishu message powered by image alt",
			EnvVars: []string{EnvPluginFeishuMsgPoweredByImageAlt},
		},

		// oss card start
		&cli.StringFlag{
			Name:    "config.feishu_oss_host",
			Usage:   "feishu OSS host for show oss info, if empty will not show oss info",
			EnvVars: []string{EnvPluginFeishuOssHost},
		},
		&cli.StringFlag{
			Name:    "config.feishu_oss_info_send_result",
			Usage:   "feishu OSS user for show at card",
			EnvVars: []string{EnvPluginFeishuOssInfoSendResult},
		},
		&cli.StringFlag{
			Name:    "config.feishu_oss_info_user",
			Usage:   "feishu OSS user for show at card",
			EnvVars: []string{EnvPluginFeishuOssInfoUser},
		},
		&cli.StringFlag{
			Name:    "config.feishu_oss_info_path",
			Usage:   "feishu OSS path for show at card",
			EnvVars: []string{EnvPluginFeishuOssInfoPath},
		},
		&cli.StringFlag{
			Name:    "config.feishu_oss_resource_url",
			Usage:   "feishu OSS resource url",
			EnvVars: []string{EnvPluginFeishuOssResourceUrl},
		},
		&cli.StringFlag{
			Name:    "config.feishu_oss_page_url",
			Usage:   "feishu OSS page url",
			EnvVars: []string{EnvPluginFeishuOssPageUrl},
		},
		&cli.StringFlag{
			Name:    "config.feishu_oss_page_passwd",
			Usage:   "OSS password at page url, will hide PLUGIN_FEISHU_OSS_RESOURCE_URL when PAGE_PASSWD not empty",
			EnvVars: []string{EnvPluginFeishuOssPagePasswd},
		},
		// oss card end
		// plugin end
	}
}

// HideFlag
// hide flags
func HideFlag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "config.ntp_target,ntp_target",
			Hidden:  true,
			Usage:   "ntp target like: pool.ntp.org, time1.google.com,time.pool.aliyun.com, default not use ntpd to sync",
			EnvVars: []string{"PLUGIN_NTP_TARGET"},
		},
	}
}

// CommonFlag
// Other modules also have flags
func CommonFlag() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:    "config.debug,debug",
			Usage:   "debug mode",
			Value:   false,
			EnvVars: []string{"PLUGIN_DEBUG"},
		},
		&cli.UintFlag{
			Name:    "config.timeout_second,timeout_second",
			Usage:   "do request timeout setting second. gather than 10",
			Hidden:  true,
			Value:   10,
			EnvVars: []string{"PLUGIN_TIMEOUT_SECOND"},
		},
	}
}
