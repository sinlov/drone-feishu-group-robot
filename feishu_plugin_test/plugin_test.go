package feishu_plugin_test

import (
	"fmt"
	"github.com/sinlov/drone-feishu-group-robot/feishu_plugin"
	"github.com/sinlov/drone-info-tools/drone_info"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlugin(t *testing.T) {
	// mock FeishuPlugin
	if envCheck(t) {
		return
	}
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

	p.Config.Webhook = envFeishuWebHook
	if envFeishuSecret != "" {
		p.Config.Secret = envFeishuSecret
	}

	p.Config.MsgType = "mock" // not support type
	err = p.Exec()
	if nil == err {
		t.Error("feishu msg type not support error should be catch!")
	}

	p.Config.MsgType = feishu_plugin.MsgTypeInteractive

	p.Config.Debug = envDebug
	p.Config.FeishuEnableForward = false
	pagePasswd := mockOssPagePasswd

	p.Drone = *drone_info.MockDroneInfo(drone_info.DroneBuildStatusSuccess)
	checkCardOssRenderByPlugin(&p, pagePasswd, false)
	p.Config.CardOss.InfoSendResult = ""
	// verify FeishuPlugin
	err = p.Exec()

	if err != nil {
		t.Fatalf("send failure error at %v", err)
	}

	p.Drone = *drone_info.MockDroneInfo(drone_info.DroneBuildStatusSuccess)
	p.Drone.Commit.Message = "build success but oss send failure and render RenderOssCard show"
	p.Config.RenderOssCard = feishu_plugin.RenderStatusShow
	checkCardOssRenderByPlugin(&p, pagePasswd, false)
	// verify FeishuPlugin
	err = p.Exec()

	if err != nil {
		t.Fatalf("send failure error at %v", err)
	}

	p.Drone = *drone_info.MockDroneInfo(drone_info.DroneBuildStatusSuccess)
	p.Drone.Commit.Message = "send success and render OssStatus"
	checkCardOssRenderByPlugin(&p, pagePasswd, true)

	assert.Equal(t, "sinlov", p.Drone.Repo.OwnerName)
	// verify FeishuPlugin
	err = p.Exec()

	if err != nil {
		t.Fatalf("send error at %v", err)
	}

	p.Drone = *drone_info.MockDroneInfo(drone_info.DroneBuildStatusFailure)
	p.Drone.Commit.Message = "build failure and hide Oss settings and render OssStatus"
	p.Config.FeishuEnableForward = true
	p.Config.RenderOssCard = feishu_plugin.RenderStatusHide
	pagePasswd = ""
	checkCardOssRenderByPlugin(&p, pagePasswd, true)
	// verify FeishuPlugin
	err = p.Exec()

	if err != nil {
		t.Fatalf("send failure error at %v", err)
	}

	drone, err := drone_info.MockDroneInfoDroneSystemRefs(
		envDroneSystemProto,
		envDroneSystemHost,
		envDroneSystemHostName,
		drone_info.DroneBuildStatusSuccess,
		fmt.Sprintf("refs/heads/%s", envDroneBranch))

	p.Config.IgnoreLastSuccessByBadges = true

	p.Drone = *drone
	p.Drone.Commit.Message = "build success and hide Oss settings and render OssStatus, open IgnoreLastSuccessByBadges"
	p.Config.FeishuEnableForward = true
	p.Config.RenderOssCard = feishu_plugin.RenderStatusHide
	pagePasswd = ""
	checkCardOssRenderByPlugin(&p, pagePasswd, true)
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
