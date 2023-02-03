package feishu_plugin

const (
	RenderStatusShow = "success"
	RenderStatusHide = "failure"

	msgTypeText        = "text"
	msgTypePost        = "post"
	msgTypeInteractive = "interactive"
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
		msgTypeInteractive,
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
		// InfoSendResult
		// send result [ success or failure]
		InfoSendResult string
		InfoUser       string
		InfoPath       string

		RenderResourceUrl string
		ResourceUrl       string
		PageUrl           string
		PagePasswd        string
	}

	// Config plugin private config
	Config struct {
		Debug               bool
		NtpTarget           string
		Webhook             string
		Secret              string
		FeishuEnableForward bool
		TimeoutSecond       int
		MsgType             string
		Title               string
		PoweredByImageKey   string
		PoweredByImageAlt   string

		RenderOssCard string
		CardOss       CardOss
	}
)
