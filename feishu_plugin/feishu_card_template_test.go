package feishu_plugin

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RenderFeishuCard(t *testing.T) {
	// mock _RenderFeishuCard
	plugin := Plugin{}
	t.Logf("~> mock _RenderFeishuCard")
	// do _RenderFeishuCard
	renderFeishuCard, err := RenderFeishuCard("abc", &plugin)
	t.Logf("~> do _RenderFeishuCard")
	// verify _RenderFeishuCard
	if err != nil {
		t.Logf("RenderFeishuCard err %v", err)
	}
	assert.Equal(t, "abc", renderFeishuCard)
}
