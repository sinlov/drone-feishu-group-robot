package feishu_plugin

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"github.com/sinlov/drone-info-tools/template"
	tools "github.com/sinlov/drone-info-tools/tools/str_tools"
	"strings"
)

// DefaultCardTemplate
// use FeishuPlugin and feishu_message.FeishuRobotMsgTemplate
const DefaultCardTemplate string = `{
  "timestamp": {{ FeishuRobotMsgTemplate.Timestamp }},
  "sign": "{{ FeishuRobotMsgTemplate.Sign }}",
  "msg_type": "interactive",
  "card": {
    "config": {
      "enable_forward": {{ FeishuRobotMsgTemplate.CtxTemp.CardTemp.EnableForward }}
    },
    "header": {
      "template": "{{#success Drone.Build.Status }}blue{{/success}}{{#failure Drone.Build.Status}}red{{/failure}}",
      "title": {
        "tag": "plain_text",
        "content": "{{#failure Drone.Build.Status}}[Failure]{{/failure}}{{ Drone.Repo.FullName }} {{#success Config.CardOss.InfoTagResult }}Tag: {{ Drone.Build.Tag }}{{/success}}"
      }
    },
    "elements": [
      {
        "tag": "markdown",
        "content": "**{{ FeishuRobotMsgTemplate.CtxTemp.CardTemp.CardTitle }}**"
      },
      {
        "tag": "hr"
      },
{{#success Config.CardOss.InfoTagResult }}
      {
        "tag": "markdown",
        "content": "📝 **Drone Tag:** {{ Drone.Build.Tag }}\nCommitCode: {{ Drone.Commit.Sha }}"
      },
{{/success}}
{{#failure Config.CardOss.InfoTagResult }}
      {
        "tag": "markdown",
        "content": "📝 Commit by {{ Drone.Commit.Author.Username }} on **{{ Drone.Build.Branch }}**\nCommitCode: {{ Drone.Commit.Sha }}"
      },
{{/failure}}
      {
        "tag": "markdown",
        "content": "{{#success Drone.Build.Status }}✅{{/success}}{{#failure Drone.Build.Status}}❌{{/failure}} Build [#{{ Drone.Build.Number }}]({{ Drone.Build.Link }}) {{ Drone.Build.Status }}{{#failure Drone.Build.Status}}\n failedStages: {{Drone.Build.FailedStages}}\n failedSteps: {{Drone.Build.FailedSteps}} {{/failure}}"
      },
      {
        "tag": "markdown",
        "content": "**Commit:**\n{{ Drone.Commit.Message }}"
      },
      {
        "tag": "markdown",
        "content": "[See Commit Details]({{ Drone.Commit.Link }}) | [See Build Details]({{ Drone.Build.Link }})"
      },
      {
        "tag": "hr"
      },
{{#success Config.RenderOssCard }}
{{#success Config.CardOss.InfoSendResult }}
      {
        "tag": "markdown",
        "content": "[OSS {{ Config.CardOss.InfoUser }} ]({{ Config.CardOss.Host }})\nPath: {{ Config.CardOss.InfoPath }}\nPage: [{{ Config.CardOss.PageUrl }}]({{ Config.CardOss.PageUrl }}){{#failure Config.CardOss.RenderResourceUrl }}\nPassword: {{ Config.CardOss.PagePasswd }}\n{{/failure}}{{#success Config.CardOss.RenderResourceUrl }}\nDownload: [click me]({{ Config.CardOss.ResourceUrl }})\n{{/success}}"
      },
{{/success}}
{{#failure Config.CardOss.InfoSendResult }}
      {
        "tag": "markdown",
        "content": "[OSS {{ Config.CardOss.InfoUser }} ]({{ Config.CardOss.Host }}) send error, please check at [build Details]({{ Drone.Build.Link }})"
      },
{{/failure}}
      {
        "tag": "hr"
      },
{{/success}}
      {
        "tag": "markdown",
        "content": "**Started:** {{ Drone.Stage.StartedTime }}\n**Finished:** {{ Drone.Stage.FinishedTime }}\n**Stage details info**\nName: {{ Drone.Stage.Name }}\nTrigger: {{ Drone.Build.Trigger }}\nMachine: {{ Drone.Stage.Machine }}\nOS: {{ Drone.Stage.Os }}\nArch: {{ Drone.Stage.Arch }}\nType: {{ Drone.Stage.Type }}\nKind: {{ Drone.Stage.Kind }}"
      },
      {
        "tag": "hr"
      },
      {
        "tag": "note",
        "elements": [
          {
            "tag": "plain_text",
            "content": "From drone {{ Drone.DroneSystem.Version }}@{{Drone.DroneSystem.HostName}}. Powered By"
          },
          {
            "tag": "img",
            "img_key": "{{ FeishuRobotMsgTemplate.CtxTemp.CardTemp.LogoImgKey }}",
            "alt": {
              "tag": "plain_text",
              "content": "{{ FeishuRobotMsgTemplate.CtxTemp.CardTemp.LogoAltStr }}"
            }
          }
        ]
      }
    ]
  }
}`

func RenderFeishuCard(tpl string, p *FeishuPlugin) (string, error) {
	var renderPlugin FeishuPlugin
	err := deepCopyByPlugin(p, &renderPlugin)
	if err != nil {
		return "", err
	}

	renderPlugin.Drone.Build.Tag = tools.Str2LineRaw(renderPlugin.Drone.Build.Tag)
	renderPlugin.Drone.Commit.Sha = tools.Str2LineRaw(renderPlugin.Drone.Commit.Sha)
	renderPlugin.Drone.Commit.Branch = tools.Str2LineRaw(renderPlugin.Drone.Commit.Branch)
	renderPlugin.Drone.Build.Branch = tools.Str2LineRaw(renderPlugin.Drone.Build.Branch)
	renderPlugin.Drone.Build.Status = tools.Str2LineRaw(renderPlugin.Drone.Build.Status)
	renderPlugin.Drone.Commit.Message = tools.Str2LineRaw(renderPlugin.Drone.Commit.Message)

	renderPlugin.Drone.Commit.Link = tools.Str2LineRaw(renderPlugin.Drone.Commit.Link)
	renderPlugin.Drone.Build.Link = tools.Str2LineRaw(renderPlugin.Drone.Build.Link)

	renderPlugin.Config.CardOss.InfoUser = tools.Str2LineRaw(renderPlugin.Config.CardOss.InfoUser)
	renderPlugin.Config.CardOss.InfoPath = tools.Str2LineRaw(renderPlugin.Config.CardOss.InfoPath)
	renderPlugin.Config.CardOss.PageUrl = tools.Str2LineRaw(renderPlugin.Config.CardOss.PageUrl)
	renderPlugin.Config.CardOss.RenderResourceUrl = tools.Str2LineRaw(renderPlugin.Config.CardOss.RenderResourceUrl)
	renderPlugin.Config.CardOss.PagePasswd = tools.Str2LineRaw(renderPlugin.Config.CardOss.PagePasswd)
	renderPlugin.Config.CardOss.Host = tools.Str2LineRaw(renderPlugin.Config.CardOss.Host)

	renderPlugin.Drone.Stage.StartedTime = tools.Str2LineRaw(renderPlugin.Drone.Stage.StartedTime)
	renderPlugin.Drone.Stage.FinishedTime = tools.Str2LineRaw(renderPlugin.Drone.Stage.FinishedTime)
	renderPlugin.Drone.Stage.Name = tools.Str2LineRaw(renderPlugin.Drone.Stage.Name)
	renderPlugin.Drone.Build.Trigger = tools.Str2LineRaw(renderPlugin.Drone.Build.Trigger)
	renderPlugin.Drone.Stage.Machine = tools.Str2LineRaw(renderPlugin.Drone.Stage.Machine)
	renderPlugin.Drone.Stage.Os = tools.Str2LineRaw(renderPlugin.Drone.Stage.Os)
	renderPlugin.Drone.Stage.Arch = tools.Str2LineRaw(renderPlugin.Drone.Stage.Arch)
	renderPlugin.Drone.Stage.Type = tools.Str2LineRaw(renderPlugin.Drone.Stage.Type)
	renderPlugin.Drone.Stage.Kind = tools.Str2LineRaw(renderPlugin.Drone.Stage.Kind)

	// check out p.Config.CardOss.InfoTagResult
	if renderPlugin.Drone.Build.Tag == "" {
		renderPlugin.Config.CardOss.InfoTagResult = RenderStatusHide
	} else {
		renderPlugin.Config.CardOss.InfoTagResult = RenderStatusShow
		// fix Drone.Commit.Link compare not support, when tags Link get error
		renderPlugin.Drone.Commit.Link = strings.Replace(renderPlugin.Drone.Commit.Link, "compare/0000000000000000000000000000000000000000...", "commit/", -1)
	}

	message, err := template.RenderTrim(tpl, &renderPlugin)
	if err != nil {
		return "", err
	}
	return message, nil
}

// deepCopyByGob deep copy by gob
//
//nolint:golint,unused
func deepCopyByGob(src, dst interface{}) error {
	var buffer bytes.Buffer
	if err := gob.NewEncoder(&buffer).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(&buffer).Decode(dst)
}

func deepCopyByPlugin(src, dst *FeishuPlugin) error {
	if tmp, err := json.Marshal(&src); err != nil {
		return err
	} else {
		err = json.Unmarshal(tmp, dst)
		return err
	}
}
