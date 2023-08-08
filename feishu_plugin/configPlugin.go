package feishu_plugin

const (
	NamePluginDebug   = "config.debug"
	EnvPluginTimeOut  = "PLUGIN_TIMEOUT_SECOND"
	NamePluginTimeOut = "config.timeout_second"

	RenderStatusShow = "success"
	RenderStatusHide = "failure"

	MsgTypeText        = "text"
	MsgTypePost        = "post"
	MsgTypeInteractive = "interactive"
)

var (
	// supportMsgType
	supportRenderStatus = []string{
		RenderStatusShow,
		RenderStatusHide,
	}

	// supportMsgType
	// @doc https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN?lang=zh-CN#8b0f2a1b
	supportMsgType = []string{
		MsgTypeInteractive,
	}
)

type (
	// SendTarget send feishu target
	SendTarget struct {
		Webhook        string
		Secret         string
		FeishuRobotMeg []byte
	}

	CardOss struct {
		Host string
		// InfoTagResult
		// tag result [ success or failure]
		InfoTagResult string
		// InfoSendResult
		// send result [ success or failure]
		InfoSendResult string

		// pull request [ success or failure]
		InfoPullRequestResult string

		InfoUser string
		InfoPath string

		RenderResourceUrl string
		ResourceUrl       string
		PageUrl           string
		PagePasswd        string
	}

	// Config plugin private config
	Config struct {
		Debug                 bool
		DroneSystemAdminToken string
		NtpTarget             string
		Webhook               string
		Secret                string
		FeishuEnableForward   bool
		TimeoutSecond         int
		MsgType               string
		Title                 string
		PoweredByImageKey     string
		PoweredByImageAlt     string

		IgnoreLastSuccessByAdminTokenDistance uint

		IgnoreLastSuccessByBadges     bool
		IgnoreLastSuccessBadgesBranch string

		RenderOssCard string
		CardOss       CardOss
	}
)
