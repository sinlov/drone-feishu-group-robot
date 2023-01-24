package feishu_message

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

// GenSign 生成签名
//
// @doc https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN?lang=zh-CN#348211be
func GenSign(secret string, timestamp int64) (string, error) {
	//timestamp + "\n" + key 做sha256, 再进行base64 encode
	stringToSign := fmt.Sprintf("%v\n%s", timestamp, secret)
	var data []byte
	h := hmac.New(sha256.New, []byte(stringToSign))
	_, err := h.Write(data)
	if err != nil {
		return "", err
	}
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return signature, nil
}
