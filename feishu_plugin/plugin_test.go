package feishu_plugin

import (
	"github.com/sinlov/drone-feishu-group-robot/drone_info"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestPlugin(t *testing.T) {
	// mock Plugin
	t.Logf("~> mock Plugin")
	p := Plugin{}
	// do Plugin
	t.Logf("~> do Plugin")
	err := p.Exec()
	if nil == err {
		t.Error("feishu webhook empty error should be catch!")
	}

	envFeishuWebHook := os.Getenv("PLUGIN_FEISHU_WEBHOOK")
	if envFeishuWebHook == "" {
		t.Error("please set env:PLUGIN_FEISHU_WEBHOOK")
	}

	p.Config.Webhook = envFeishuWebHook
	if os.Getenv("PLUGIN_FEISHU_SECRET") != "" {
		p.Config.Secret = os.Getenv("PLUGIN_FEISHU_SECRET")
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

	p.Drone = *drone_info.MockDroneInfo("success")
	assert.Equal(t, "sinlov", p.Drone.Repo.OwnerName)
	// verify Plugin
	err = p.Exec()

	if err != nil {
		t.Fatalf("send error at %v", err)
	}

	p.Drone = *drone_info.MockDroneInfo("failure")
	// verify Plugin
	err = p.Exec()

	if err != nil {
		t.Fatalf("send failure error at %v", err)
	}
}
