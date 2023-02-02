package feishu_plugin

import (
	"github.com/sinlov/drone-info-tools/drone_info"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const (
	mockVersion        = "v0.0.0"
	mockOssHost        = "https://docs.aws.amazon.com/s3/index.html"
	mockOssUser        = "ossAdmin"
	mockOssPath        = "dist/demo/pass.tar.gz"
	mockOssResourceUrl = "https://docs.aws.amazon.com/s/dist/demo/pass.tar.gz"
	mockOssPageUrl     = "https://docs.aws.amazon.com/p/dist/demo/pass.tar.gz"
	mockOssPagePasswd  = "abc-zxy"
)

func TestPlugin(t *testing.T) {
	// mock FeishuPlugin
	t.Logf("~> mock FeishuPlugin")
	p := FeishuPlugin{
		Version: mockVersion,
	}
	// do FeishuPlugin
	t.Logf("~> do FeishuPlugin")
	err := p.Exec()
	if nil == err {
		t.Error("feishu webhook empty error should be catch!")
	}

	envFeishuWebHook := os.Getenv(EnvPluginFeishuWebhook)
	if envFeishuWebHook == "" {
		t.Errorf("please set env:%s", EnvPluginFeishuWebhook)
	}

	p.Config.Webhook = envFeishuWebHook
	if os.Getenv(EnvPluginFeishuSecret) != "" {
		p.Config.Secret = os.Getenv(EnvPluginFeishuSecret)
	}

	p.Config.MsgType = "mock" // not support type
	err = p.Exec()
	if nil == err {
		t.Error("feishu msg type not support error should be catch!")
	}

	envMsgType := msgTypeInteractive // only support this type now
	p.Config.MsgType = envMsgType

	if os.Getenv("PLUGIN_DEBUG") == "true" {
		p.Config.Debug = true
	}
	p.Config.FeishuEnableForward = false
	p.Config.RenderOssCard = RenderStatusShow
	pagePasswd := mockOssPagePasswd
	checkCardOssRenderByPlugin(&p, pagePasswd)
	p.Drone = *drone_info.MockDroneInfo("success")

	assert.Equal(t, "sinlov", p.Drone.Repo.OwnerName)
	// verify FeishuPlugin
	err = p.Exec()

	if err != nil {
		t.Fatalf("send error at %v", err)
	}

	p.Config.FeishuEnableForward = true
	p.Config.RenderOssCard = RenderStatusShow
	pagePasswd = ""
	checkCardOssRenderByPlugin(&p, pagePasswd)
	p.Drone = *drone_info.MockDroneInfo("failure")
	// verify FeishuPlugin
	err = p.Exec()

	if err != nil {
		t.Fatalf("send failure error at %v", err)
	}
}

func checkCardOssRenderByPlugin(p *FeishuPlugin, pagePasswd string) {
	p.Config.CardOss.PagePasswd = pagePasswd
	if p.Config.CardOss.PagePasswd == "" {
		p.Config.CardOss.RenderResourceUrl = RenderStatusShow
	} else {
		p.Config.CardOss.RenderResourceUrl = RenderStatusHide
	}
	p.Config.CardOss.Host = mockOssHost
	p.Config.CardOss.InfoUser = mockOssUser
	p.Config.CardOss.InfoPath = mockOssPath
	p.Config.CardOss.ResourceUrl = mockOssResourceUrl
	p.Config.CardOss.PageUrl = mockOssPageUrl
}
