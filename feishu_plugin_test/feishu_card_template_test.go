package feishu_plugin_test

import (
	"fmt"
	"github.com/sinlov/drone-feishu-group-robot/feishu_plugin"
	"github.com/sinlov/drone-info-tools/drone_info"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RenderFeishuCard(t *testing.T) {
	// mock _RenderFeishuCard
	// mock FeishuPlugin
	if envCheck(t) {
		return
	}
	p := feishu_plugin.FeishuPlugin{
		Name:    mockName,
		Version: mockVersion,
	}
	p.Config.Title = mockTitle
	p.FeishuRobotMsgTemplate.CtxTemp.CardTemp.CardTitle = mockTitle

	p.Config.Webhook = envFeishuWebHook
	if envFeishuSecret != "" {
		p.Config.Secret = envFeishuSecret
	}
	p.Config.Debug = envDebug
	p.Config.FeishuEnableForward = false
	p.Drone = *drone_info.MockDroneInfo(drone_info.DroneBuildStatusSuccess)

	t.Logf("~> mock _RenderFeishuCard")
	// do _RenderFeishuCard
	renderFeishuCard, err := feishu_plugin.RenderFeishuCard(feishu_plugin.DefaultCardTemplate, &p)
	t.Logf("~> do _RenderFeishuCard")
	// verify _RenderFeishuCard
	if err != nil {
		t.Logf("RenderFeishuCard err %v", err)
	}
	t.Logf("renderFeishuCard\n%s", renderFeishuCard)
	assert.NotEqual(t, "", renderFeishuCard)
}

func Test_RenderFeishuCardTag(t *testing.T) {
	// mock _RenderFeishuCardTag

	t.Logf("~> mock _RenderFeishuCardTag")
	// mock FeishuPlugin
	if envCheck(t) {
		return
	}
	p := feishu_plugin.FeishuPlugin{
		Name:    mockName,
		Version: mockVersion,
	}
	p.Config.Title = mockTitle
	p.FeishuRobotMsgTemplate.CtxTemp.CardTemp.CardTitle = mockTitle

	p.Config.Webhook = envFeishuWebHook
	if envFeishuSecret != "" {
		p.Config.Secret = envFeishuSecret
	}
	p.Config.Debug = envDebug
	p.Config.FeishuEnableForward = false
	droneInfoRefs, err := drone_info.MockDroneInfoRefs(
		drone_info.DroneBuildStatusSuccess,
		fmt.Sprintf("refs/tags/%s", "v1.2.3"),
	)
	if err != nil {
		t.Fatal(err)
	}
	p.Drone = *droneInfoRefs
	assert.Equal(t, "", p.Drone.Build.Branch)
	assert.Equal(t, "v1.2.3", p.Drone.Build.Tag)

	t.Logf("~> do _RenderFeishuCardTag")
	renderFeishuCard, err := feishu_plugin.RenderFeishuCard(feishu_plugin.DefaultCardTemplate, &p)
	// verify _RenderFeishuCard
	if err != nil {
		t.Logf("_RenderFeishuCardTag err %v", err)
	}
	t.Logf("_RenderFeishuCardTag\n%s", renderFeishuCard)
	assert.NotEqual(t, "", renderFeishuCard)
	// verify _RenderFeishuCardTag
}
