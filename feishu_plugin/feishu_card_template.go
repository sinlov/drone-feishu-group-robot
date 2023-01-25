package feishu_plugin

import "github.com/sinlov/drone-feishu-group-robot/template"

// defaultCardTemplate
// use Plugin and feishu_message.FeishuRobotMsgTemplate
const defaultCardTemplate string = `
{
    "timestamp": {{ FeishuRobotMsgTemplate.Timestamp }},
    "sign": "{{ FeishuRobotMsgTemplate.Sign }}",
    "msg_type": "interactive",
    "card": {
        "config": {
            "wide_screen_mode": true,
            "enable_forward":true
        },
        "header": {
            "template": "{{#success Drone.Build.Status }}blue{{/success}}{{#failure Drone.Build.Status}}red{{/failure}}",
            "title": {
                "tag": "plain_text",
                "content": "{{ FeishuRobotMsgTemplate.CtxTemp.CardTemp.CardTitle }}"
            }
        },
        "elements": [
            {
                "tag": "markdown",
                "content": "**{{ Drone.Repo.FullName }}**"
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
                "content": "{{#success Drone.Build.Status }}✅{{/success}}{{#failure Drone.Build.Status}}❌{{/failure}} Build [#{{ Drone.Build.Number }}]({{ Drone.Build.Link }}) {{ Drone.Build.Status }}."
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
            {
                "tag": "note",
                "elements": [
                    {
                        "tag": "plain_text",
                        "content": "Powered By"
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
}
`

func RenderFeishuCard(tpl string, p *Plugin) (string, error) {
	message, err := template.RenderTrim(tpl, &p)
	if err != nil {
		return "", err
	}
	return message, nil
}
