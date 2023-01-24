package feishu_message

import (
	"fmt"
)

const (
	ApiFeishuBase string = "https://open.feishu.cn"
)

var apiBotV2 string

// ApiFeishuBotV2
// @doc https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN?lang=zh-CN#383d6e48
func ApiFeishuBotV2() string {
	if apiBotV2 == "" {
		apiBotV2 = fmt.Sprintf("%s/%s", ApiFeishuBase, "open-apis/bot/v2/hook")
	}
	return apiBotV2
}

type (
	ApiRespRotV2 struct {
		Code          uint64 `json:"code,omitempty"`
		Msg           string `json:"msg,omitempty"`
		StatusCode    uint64 `json:"StatusCode,omitempty"`
		StatusMessage string `json:"StatusMessage,omitempty"`
		Extra         interface{}
	}
)
