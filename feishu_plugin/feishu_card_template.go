package feishu_plugin

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"github.com/sinlov/drone-info-tools/template"
	tools "github.com/sinlov/drone-info-tools/tools/str_tools"
)

// defaultCardTemplate
// use FeishuPlugin and feishu_message.FeishuRobotMsgTemplate
const defaultCardTemplate string = `{
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
        "content": "{{#failure Drone.Build.Status}}[Failure]{{/failure}}{{ Drone.Repo.FullName }}"
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
      {
        "tag": "markdown",
        "content": "📝 Commit by {{ Drone.Commit.Author.Username }} on **{{ Drone.Commit.Branch }}**\nCommitCode: {{ Drone.Commit.Sha }}"
      },
      {
        "tag": "markdown",
        "content": "{{#success Drone.Build.Status }}✅{{/success}}{{#failure Drone.Build.Status}}❌{{/failure}} Build [#{{ Drone.Build.Number }}]({{ Drone.Build.Link }}) {{ Drone.Build.Status }}{{#failure Drone.Build.Status}}\n failedStages: {{Drone.Build.FailedStages}}\n failedSteps: {{Drone.Build.FailedSteps}} {{/failure}}"
      },
      {
        "tag": "markdown",
        "content": "**Commit:**\n\n{{ Drone.Commit.Message }}"
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
        "content": "**Stage**\nName: {{ Drone.Stage.Name }}\nMachine: {{ Drone.Stage.Machine }}\nOS: {{ Drone.Stage.Os }}\nArch: {{ Drone.Stage.Arch }}\nType: {{ Drone.Stage.Type }}\nKind: {{ Drone.Stage.Kind }}"
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

	renderPlugin.Drone.Commit.Message = tools.Str2LineRaw(renderPlugin.Drone.Commit.Message)

	message, err := template.RenderTrim(tpl, &renderPlugin)
	if err != nil {
		return "", err
	}
	return message, nil
}

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
