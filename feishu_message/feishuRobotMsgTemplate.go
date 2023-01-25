package feishu_message

const (
	defaultLogoImgKey = "img_v2_18a207ce-2428-4b5c-8fc0-f791c05d938g"
	defaultLogoAltStr = "sinlov"
)

const (
	defaultMsgTitle = "Drone CI Notification"
)

type (
	// CardTemp
	// @doc https://open.feishu.cn/document/ukTMukTMukTM/uMjNwUjLzYDM14yM2ATN
	CardTemp struct {
		// EnableForward
		// @doc https://open.feishu.cn/document/ukTMukTMukTM/uAjNwUjLwYDM14CM2ATN
		EnableForward bool

		// @doc https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/image/create
		LogoImgKey string
		LogoAltStr string

		CardTitle string
	}

	CtxTemp struct {
		CardTemp CardTemp
	}

	FeishuRobotMsgTemplate struct {
		Timestamp int64  `json:"timestamp"`
		Sign      string `json:"sign"`
		// MsgType
		// @doc https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN?lang=zh-CN#8b0f2a1b
		MsgType string `json:"msg_type"`

		CtxTemp CtxTemp
	}
)

func (CardTemp) Build(
	cardTitle, poweredByImgKey, poweredByImgAlt string,
) CardTemp {
	//override default card setting
	if cardTitle == "" {
		cardTitle = defaultMsgTitle
	}
	if poweredByImgKey == "" {
		poweredByImgKey = defaultLogoImgKey
	}
	if poweredByImgAlt == "" {
		poweredByImgAlt = defaultLogoAltStr
	}

	var cardTemp = CardTemp{
		CardTitle:  cardTitle,
		LogoImgKey: poweredByImgKey,
		LogoAltStr: poweredByImgAlt,
	}

	return cardTemp
}
