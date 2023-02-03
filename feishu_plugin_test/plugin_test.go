package feishu_plugin_test

import (
	"github.com/sinlov/drone-feishu-group-robot/feishu_plugin"
	"github.com/sinlov/drone-info-tools/drone_info"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const (
	mockName           = "drone-feishu-group-robot"
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
	p := feishu_plugin.FeishuPlugin{
		Name:    mockName,
		Version: mockVersion,
	}
	// do FeishuPlugin
	t.Logf("~> do FeishuPlugin")
	err := p.Exec()
	if nil == err {
		t.Error("feishu webhook empty error should be catch!")
	}

	envFeishuWebHook := os.Getenv(feishu_plugin.EnvPluginFeishuWebhook)
	if envFeishuWebHook == "" {
		t.Errorf("please set env:%s", feishu_plugin.EnvPluginFeishuWebhook)
	}

	p.Config.Webhook = envFeishuWebHook
	if os.Getenv(feishu_plugin.EnvPluginFeishuSecret) != "" {
		p.Config.Secret = os.Getenv(feishu_plugin.EnvPluginFeishuSecret)
	}

	p.Config.MsgType = "mock" // not support type
	err = p.Exec()
	if nil == err {
		t.Error("feishu msg type not support error should be catch!")
	}

	envMsgType := "interactive" // only support this type now
	p.Config.MsgType = envMsgType

	if os.Getenv("PLUGIN_DEBUG") == "true" {
		p.Config.Debug = true
	}
	p.Config.FeishuEnableForward = false
	p.Config.RenderOssCard = feishu_plugin.RenderStatusShow
	pagePasswd := mockOssPagePasswd

	p.Drone = *drone_info.MockDroneInfo("success")
	checkCardOssRenderByPlugin(&p, pagePasswd, false)
	// verify FeishuPlugin
	err = p.Exec()

	if err != nil {
		t.Fatalf("send failure error at %v", err)
	}

	checkCardOssRenderByPlugin(&p, pagePasswd, true)
	p.Drone = *drone_info.MockDroneInfo("success")

	assert.Equal(t, "sinlov", p.Drone.Repo.OwnerName)
	// verify FeishuPlugin
	err = p.Exec()

	if err != nil {
		t.Fatalf("send error at %v", err)
	}

	p.Config.FeishuEnableForward = true
	p.Config.RenderOssCard = feishu_plugin.RenderStatusShow
	pagePasswd = ""
	checkCardOssRenderByPlugin(&p, pagePasswd, true)
	p.Drone = *drone_info.MockDroneInfo("failure")
	// verify FeishuPlugin
	err = p.Exec()

	if err != nil {
		t.Fatalf("send failure error at %v", err)
	}
}

func checkCardOssRenderByPlugin(p *feishu_plugin.FeishuPlugin, pagePasswd string, sendOssSucc bool) {
	p.Config.CardOss.PagePasswd = pagePasswd
	if p.Config.CardOss.PagePasswd == "" {
		p.Config.CardOss.RenderResourceUrl = feishu_plugin.RenderStatusShow
	} else {
		p.Config.CardOss.RenderResourceUrl = feishu_plugin.RenderStatusHide
	}
	if sendOssSucc {
		p.Config.CardOss.InfoSendResult = feishu_plugin.RenderStatusShow
	} else {
		p.Config.CardOss.InfoSendResult = feishu_plugin.RenderStatusHide
	}
	p.Config.CardOss.Host = mockOssHost
	p.Config.CardOss.InfoUser = mockOssUser
	p.Config.CardOss.InfoPath = mockOssPath
	p.Config.CardOss.ResourceUrl = mockOssResourceUrl
	p.Config.CardOss.PageUrl = mockOssPageUrl
}
