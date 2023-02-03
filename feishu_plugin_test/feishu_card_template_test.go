package feishu_plugin_test

import (
	"github.com/sinlov/drone-feishu-group-robot/feishu_plugin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RenderFeishuCard(t *testing.T) {
	// mock _RenderFeishuCard
	plugin := feishu_plugin.FeishuPlugin{}
	t.Logf("~> mock _RenderFeishuCard")
	// do _RenderFeishuCard
	renderFeishuCard, err := feishu_plugin.RenderFeishuCard("abc", &plugin)
	t.Logf("~> do _RenderFeishuCard")
	// verify _RenderFeishuCard
	if err != nil {
		t.Logf("RenderFeishuCard err %v", err)
	}
	assert.Equal(t, "abc", renderFeishuCard)
}
