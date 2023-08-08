package feishu_plugin_test

import (
	"encoding/json"
	"fmt"
	"github.com/sebdah/goldie/v2"
	"github.com/sinlov/drone-feishu-group-robot/feishu_plugin"
	"github.com/sinlov/drone-info-tools/drone_info"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRenderFeishuCard(t *testing.T) {
	// mock RenderFeishuCard

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

	var sampleRender feishu_plugin.FeishuPlugin
	deepCopyByPlugin(&p, &sampleRender)

	var sampleFailRender feishu_plugin.FeishuPlugin
	deepCopyByPlugin(&p, &sampleFailRender)
	sampleFailRender.Drone = *drone_info.MockDroneInfo(drone_info.DroneBuildStatusFailure)

	var tagMessageRender feishu_plugin.FeishuPlugin
	deepCopyByPlugin(&p, &tagMessageRender)
	droneTagInfoRefs, err := drone_info.MockDroneInfoRefs(
		drone_info.DroneBuildStatusSuccess,
		fmt.Sprintf("refs/tags/%s", "v1.2.3"),
	)
	if err != nil {
		t.Fatal(err)
	}
	tagMessageRender.Drone = *droneTagInfoRefs
	assert.Equal(t, "", tagMessageRender.Drone.Build.Branch)
	assert.Equal(t, "v1.2.3", tagMessageRender.Drone.Build.Tag)

	var prMessageRender feishu_plugin.FeishuPlugin
	deepCopyByPlugin(&p, &prMessageRender)
	prMessageRender.Drone.Build.Event = "pull_request"
	prMessageRender.Drone.Build.SourceBranch = "feat-new"
	prMessageRender.Drone.Build.TargetBranch = "main"
	prMessageRender.Drone.Build.PR = "3"
	prMessageRender.Drone.Commit.Ref = "refs/pull/3/head"
	prMessageRender.Drone.Commit.Link = fmt.Sprintf("%s/pulls/%s", prMessageRender.Drone.Repo.Link, prMessageRender.Drone.Build.PR)

	tests := []struct {
		name    string
		p       feishu_plugin.FeishuPlugin
		wantErr bool
	}{
		{
			name: "sample", // testdata/TestRenderFeishuCard/sample.golden
			p:    sampleRender,
		},
		{
			name: "sample_fail", // testdata/TestRenderFeishuCard/sample_fail.golden
			p:    sampleFailRender,
		},
		{
			name: "tag", // testdata/TestRenderFeishuCard/tag.golden
			p:    tagMessageRender,
		},
		{
			name: "pull_request", // testdata/TestRenderFeishuCard/pull_request.golden
			p:    prMessageRender,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ClassicDiff),
			)

			// do RenderFeishuCard
			renderFeishuCard, gotErr := feishu_plugin.RenderFeishuCard(feishu_plugin.DefaultCardTemplate, &tc.p)
			assert.Equal(t, tc.wantErr, gotErr != nil)
			if tc.wantErr {
				t.Logf("~> RenderFeishuCard error: %s", gotErr.Error())
				return
			}
			// verify RenderFeishuCard
			g.Assert(t, t.Name(), []byte(renderFeishuCard))
		})
	}
}

func deepCopyByPlugin(src, dst *feishu_plugin.FeishuPlugin) {
	if tmp, err := json.Marshal(&src); err != nil {
		return
	} else {
		err = json.Unmarshal(tmp, dst)
		return
	}
}
