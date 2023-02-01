package main

import (
	"fmt"
	"github.com/sinlov/drone-feishu-group-robot/drone_info"
	"github.com/sinlov/drone-feishu-group-robot/feishu_plugin"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

const (
	// Version of cli
	Version = "v1.1.0"
)

func action(c *cli.Context) error {

	isDebug := c.Bool("config.debug")

	drone := bindDroneInfo(c)

	if isDebug {
		log.Printf("debug: cli version is %s", Version)
		log.Printf("debug: load droneInfo finish at link: %v\n", drone.Build.Link)
	}
	config := feishu_plugin.Config{
		Debug:               c.Bool("config.debug"),
		TimeoutSecond:       c.Int("config.timeout_second"),
		NtpTarget:           c.String("config.ntp_target"),
		Webhook:             c.String("config.webhook"),
		Secret:              c.String("config.secret"),
		FeishuEnableForward: c.Bool("config.feishu_enable_forward"),
		MsgType:             c.String("config.msg_type"),
		Title:               c.String("config.msg_title"),
		PoweredByImageKey:   c.String("config.msg_powered_by_image_key"),
		PoweredByImageAlt:   c.String("config.msg_powered_by_image_alt"),
	}

	if isDebug {
		log.Printf("config.timeout_second: %v", config.TimeoutSecond)
	}

	p := feishu_plugin.Plugin{
		Version: Version,
		Drone:   drone,
		Config:  config,
	}
	err := p.Exec()

	if err != nil {
		log.Fatalf("err: %v", err)
		return err
	}
	return nil
}

// pluginFlag
// set plugin flag at here
func pluginFlag() []cli.Flag {
	return []cli.Flag{
		// plugin start
		&cli.BoolFlag{
			Name:    "config.debug,debug",
			Usage:   "debug mode",
			EnvVars: []string{"PLUGIN_DEBUG"},
		},
		&cli.IntFlag{
			Name:    "config.timeout_second,timeout_second",
			Usage:   "do request timeout setting second",
			EnvVars: []string{"PLUGIN_TIMEOUT_SECOND"},
		},
		&cli.StringFlag{
			Name:    "config.ntp_target,ntp_target",
			Usage:   "ntp target like: pool.ntp.org, time1.google.com,time.pool.aliyun.com, default not use ntpd to sync",
			EnvVars: []string{"PLUGIN_NTP_TARGET"},
		},
		&cli.StringFlag{
			Name:    "config.webhook,feishu_webhook",
			Usage:   "feishu webhook for send message",
			EnvVars: []string{"PLUGIN_FEISHU_WEBHOOK"},
		},
		&cli.StringFlag{
			Name:    "config.secret,feishu_secret",
			Usage:   "feishu secret",
			EnvVars: []string{"PLUGIN_FEISHU_SECRET"},
		},
		&cli.BoolFlag{
			Name:    "config.feishu_enable_forward,feishu_enable_forward",
			Usage:   "feishu message enable forward, default false",
			EnvVars: []string{"PLUGIN_FEISHU_ENABLE_FORWARD"},
		},
		&cli.StringFlag{
			Name:    "config.msg_type,feishu_msg_type",
			Usage:   "feishu message type",
			EnvVars: []string{"PLUGIN_FEISHU_MSG_TYPE"},
		},
		&cli.StringFlag{
			Name:    "config.msg_title,feishu_msg_title",
			Usage:   "feishu message title",
			EnvVars: []string{"PLUGIN_FEISHU_MSG_TITLE"},
		},
		&cli.StringFlag{
			Name:    "config.msg_powered_by_image_key,feishu_msg_powered_by_image_key",
			Usage:   "feishu message powered by image key",
			EnvVars: []string{"PLUGIN_FEISHU_MSG_POWERED_BY_IMAGE_KEY"},
		},
		&cli.StringFlag{
			Name:    "config.msg_powered_by_image_alt,feishu_msg_powered_by_image_alt",
			Usage:   "feishu message powered by image alt",
			EnvVars: []string{"PLUGIN_FEISHU_MSG_POWERED_BY_IMAGE_ALT"},
		},
		// plugin end
	}
}

// droneInfoFlag
// Please do not edit unless you understand drone's environment variables.
func droneInfoFlag() []cli.Flag {
	return []cli.Flag{
		// droneInfo start
		&cli.StringFlag{
			Name:    "repo.name",
			Usage:   "providers the name of the repository",
			EnvVars: []string{drone_info.EnvDroneRepoName},
		},
		&cli.StringFlag{
			Name:    "repo.group",
			Usage:   "providers the group of the repository",
			EnvVars: []string{drone_info.EnvDroneRepoNamespace},
		},
		&cli.StringFlag{
			Name:    "repo.full.name",
			Usage:   "providers the full name of the repository",
			EnvVars: []string{drone_info.EnvDroneRepo},
		},
		&cli.StringFlag{
			Name:    "repo.owner",
			Usage:   "providers the owner of the repository",
			EnvVars: []string{drone_info.EnvDroneRepoOwner},
		},
		&cli.StringFlag{
			Name:    "repo.scm",
			Usage:   "Provides the repository type for the current running build",
			EnvVars: []string{drone_info.EnvDroneRepoScm},
		},
		&cli.StringFlag{
			Name:    "repo.remote.url",
			Usage:   "Provides the git+ssh url that should be used to clone the repository.",
			EnvVars: []string{drone_info.EnvDroneRemoteUrl},
		},
		&cli.StringFlag{
			Name:    "repo.http.url",
			Usage:   "Provides the http url that should be used to clone the repository",
			EnvVars: []string{drone_info.EnvDroneGitHttpUrl},
		},
		&cli.StringFlag{
			Name:    "repo.ssh.url",
			Usage:   "Provides the ssh url that should be used to clone the repository",
			EnvVars: []string{drone_info.EnvDroneGitSshUrl},
		},

		// drone_info.Build
		&cli.StringFlag{
			Name:    "build.workspace",
			Usage:   "drone’s working directory for a pipeline",
			EnvVars: []string{drone_info.EnvDroneBuildWorkSpace},
		},
		&cli.StringFlag{
			Name:    "build.status",
			Usage:   "build status",
			Value:   "success",
			EnvVars: []string{drone_info.EnvDroneBuildStatus},
		},
		&cli.Uint64Flag{
			Name:    "build.number",
			Usage:   "providers the current build number",
			EnvVars: []string{drone_info.EnvDroneBuildNumber},
		},
		&cli.StringFlag{
			Name:    "build.tag",
			Usage:   "build tag",
			EnvVars: []string{drone_info.EnvDroneTag},
		},
		&cli.StringFlag{
			Name:    "build.target_branch",
			Usage:   "This environment variable can be used in conjunction with the source branch variable to get the pull request base and head branch.",
			EnvVars: []string{drone_info.EnvDroneTargetBranch},
		},
		&cli.StringFlag{
			Name:    "build.link",
			Usage:   "build link",
			EnvVars: []string{drone_info.EnvDroneBuildLink},
		},
		&cli.StringFlag{
			Name:    "build.event",
			Usage:   "build event",
			EnvVars: []string{drone_info.EnvDroneBuildEvent},
		},
		&cli.Uint64Flag{
			Name:    "build.started",
			Usage:   "build started",
			EnvVars: []string{drone_info.EnvDroneBuildStarted},
		},
		&cli.Uint64Flag{
			Name:    "build.finished",
			Usage:   "build finished",
			EnvVars: []string{drone_info.EnvDroneBuildFinished},
		},
		&cli.StringFlag{
			Name:    "pull.request",
			Usage:   "pull request",
			EnvVars: []string{drone_info.EnvDronePR},
		},
		&cli.StringFlag{
			Name:    "deploy.to",
			Usage:   "provides the target deployment environment for the running build. This value is only available to promotion and rollback pipelines.",
			EnvVars: []string{drone_info.EnvDroneDeployTo},
		},
		&cli.StringFlag{
			Name:    "failed.stages",
			Usage:   "Provides a comma-separate list of failed pipeline stages for the current running build.",
			EnvVars: []string{drone_info.EnvDroneFailedStages},
		},
		&cli.StringFlag{
			Name:    "failed.steps",
			Usage:   "Provides a comma-separate list of failed pipeline steps",
			EnvVars: []string{drone_info.EnvDroneFailedSteps},
		},

		&cli.StringFlag{
			Name:    "commit.author.username",
			Usage:   "Provides the commit author name for the current running build. Note this is a user-defined value and may be empty or inaccurate",
			EnvVars: []string{drone_info.EnvDroneCommitAuthorName},
		},
		&cli.StringFlag{
			Name:    "commit.author.email",
			Usage:   "Provides the commit email address for the current running build. Note this is a user-defined value and may be empty or inaccurate",
			EnvVars: []string{drone_info.EnvDroneCommitAuthorEmail},
		},
		&cli.StringFlag{
			Name:    "commit.author.name",
			Usage:   "Provides the commit author username for the current running build. This is the username from source control management system (e.g. GitHub username)",
			EnvVars: []string{drone_info.EnvDroneCommitAuthor},
		},
		&cli.StringFlag{
			Name:    "commit.author.avatar",
			Usage:   "Provides the commit author avatar for the current running build. This is the avatar from source control management system (e.g. GitHub)",
			EnvVars: []string{drone_info.EnvDroneCommitAuthorAvatar},
		},
		&cli.StringFlag{
			Name:    "commit.link",
			Usage:   "providers the http link to the current commit in the remote source code management system(e.g.GitHub)",
			EnvVars: []string{drone_info.EnvDroneCommitLink},
		},
		&cli.StringFlag{
			Name:    "commit.branch",
			Usage:   "providers the branch for the current build",
			EnvVars: []string{drone_info.EnvDroneCommitBranch},
			Value:   "master",
		},
		&cli.StringFlag{
			Name:    "commit.message",
			Usage:   "providers the commit message for the current build",
			EnvVars: []string{drone_info.EnvDroneCommitMessage},
		},
		&cli.StringFlag{
			Name:    "commit.sha",
			Usage:   "providers the commit sha for the current build",
			EnvVars: []string{drone_info.EnvDroneCommitSha},
		},
		&cli.StringFlag{
			Name:    "commit.ref",
			Usage:   "providers the commit ref for the current build",
			EnvVars: []string{drone_info.EnvDroneCommitRef},
		},

		// drone_info.Stage
		&cli.Uint64Flag{
			Name:    "stage.started",
			Usage:   "stage started ",
			EnvVars: []string{drone_info.EnvDroneStageStarted},
		},
		&cli.Uint64Flag{
			Name:    "stage.finished",
			Usage:   "stage finished",
			EnvVars: []string{drone_info.EnvDroneStageFinished},
		},
		&cli.StringFlag{
			Name:    "stage.machine",
			Usage:   "stage machine",
			EnvVars: []string{drone_info.EnvDroneStageMachine},
		},
		&cli.StringFlag{
			Name:    "stage.os",
			Usage:   "stage OS",
			EnvVars: []string{drone_info.EnvDroneStageOs},
		},
		&cli.StringFlag{
			Name:    "stage.arch",
			Usage:   "stage arch",
			EnvVars: []string{drone_info.EnvDroneStageArch},
		},
		&cli.StringFlag{
			Name:    "stage.variant",
			Usage:   "stage variant",
			EnvVars: []string{drone_info.EnvDroneStageVariant},
		},
		&cli.StringFlag{
			Name:    "stage.type",
			Usage:   "stage type",
			EnvVars: []string{drone_info.EnvDroneStageType},
		},
		&cli.StringFlag{
			Name:    "stage.kind",
			Usage:   "stage kind",
			EnvVars: []string{drone_info.EnvDroneStageKind},
		},
		&cli.StringFlag{
			Name:    "stage.name",
			Usage:   "stage name",
			EnvVars: []string{drone_info.EnvDroneStageName},
		},

		// drone_info.DroneSystem
		&cli.StringFlag{
			Name:    "drone.system.version",
			Usage:   "Provides the version of the Drone server.",
			EnvVars: []string{drone_info.EnvDroneSystemVersion},
		},
		&cli.StringFlag{
			Name:    "drone.system.host",
			Usage:   "Provides the host used by the Drone server. This can be combined with the protocol to construct to the server url.",
			EnvVars: []string{drone_info.EnvDroneSystemHost},
		},
		&cli.StringFlag{
			Name:    "drone.system.hostname",
			Usage:   "Provides the hostname used by the Drone server. This can be combined with the protocol to construct to the server url.",
			EnvVars: []string{drone_info.EnvDroneSystemHostName},
		},
		&cli.StringFlag{
			Name:    "drone.system.proto",
			Usage:   "Provides the protocol used by the Drone server. This can be combined with the hostname to construct to the server url.",
			EnvVars: []string{drone_info.EnvDroneSystemProto},
		},
		// droneInfo end
	}
}

// bindDroneInfo
// Please do not edit unless you understand drone's environment variables.
func bindDroneInfo(c *cli.Context) drone_info.Drone {
	repoHttpUrl := c.String("repo.http.url")
	repoHost := ""
	repoHostName := ""
	parse, err := url.Parse(repoHttpUrl)
	if err == nil {
		repoHost = parse.Host
		repoHostName = parse.Hostname()
	}
	stageStartT := c.Uint64("stage.started")
	stageStartTime := time.Unix(int64(stageStartT), 0).Format(drone_info.DroneTimeFormatDefault)
	stageFinishedT := c.Uint64("stage.finished")
	stageFinishedTime := time.Unix(int64(stageStartT), 0).Format(drone_info.DroneTimeFormatDefault)
	var drone = drone_info.Drone{
		//  repo info
		Repo: drone_info.Repo{
			ShortName: c.String("repo.name"),
			GroupName: c.String("repo.group"),
			FullName:  c.String("repo.full.name"),
			OwnerName: c.String("repo.owner"),
			Scm:       c.String("repo.scm"),
			RemoteURL: c.String("repo.remote.url"),
			HttpUrl:   repoHttpUrl,
			SshUrl:    c.String("repo.ssh.url"),
			Host:      repoHost,
			HostName:  repoHostName,
		},
		//  drone_info.Build
		Build: drone_info.Build{
			WorkSpace:    c.String("build.workspace"),
			Status:       c.String("build.status"),
			Number:       c.Uint64("build.number"),
			Tag:          c.String("build.tag"),
			TargetBranch: c.String("build.target_branch"),
			Link:         c.String("build.link"),
			Event:        c.String("build.event"),
			StartAt:      c.Uint64("build.started"),
			FinishedAt:   c.Uint64("build.finished"),
			PR:           c.String("pull.request"),
			DeployTo:     c.String("deploy.to"),
			FailedStages: c.String("failed.stages"),
			FailedSteps:  c.String("failed.steps"),
		},
		Commit: drone_info.Commit{
			Link:    c.String("commit.link"),
			Branch:  c.String("commit.branch"),
			Message: c.String("commit.message"),
			Sha:     c.String("commit.sha"),
			Ref:     c.String("commit.ref"),
			Author: drone_info.CommitAuthor{
				Username: c.String("commit.author.username"),
				Email:    c.String("commit.author.email"),
				Name:     c.String("commit.author.name"),
				Avatar:   c.String("commit.author.avatar"),
			},
		},
		Stage: drone_info.Stage{
			StartedAt:    stageStartT,
			StartedTime:  stageStartTime,
			FinishedAt:   stageFinishedT,
			FinishedTime: stageFinishedTime,
			Machine:      c.String("stage.machine"),
			Os:           c.String("stage.os"),
			Arch:         c.String("stage.arch"),
			Variant:      c.String("stage.variant"),
			Type:         c.String("stage.type"),
			Kind:         c.String("stage.kind"),
			Name:         c.String("stage.name"),
		},
		DroneSystem: drone_info.DroneSystem{
			Version:  c.String("drone.system.version"),
			Host:     c.String("drone.system.host"),
			HostName: c.String("drone.system.hostname"),
			Proto:    c.String("drone.system.proto"),
		},
	}
	return drone
}

func main() {
	app := cli.NewApp()
	app.Version = Version
	app.Name = "Drone feishu Message Plugin"
	app.Usage = "Sending message to feishu group by robot using WebHook"
	year := time.Now().Year()
	app.Copyright = fmt.Sprintf("© 2022-%d sinlov", year)
	authorSinlov := &cli.Author{
		Name:  "sinlov",
		Email: "sinlovgmppt@gmail.com",
	}
	app.Authors = []*cli.Author{
		authorSinlov,
	}

	app.Action = action
	flags := appendCliFlag(droneInfoFlag(), pluginFlag())
	app.Flags = flags

	// kubernetes runner patch
	if _, err := os.Stat("/run/drone/env"); err == nil {
		errDotEnv := godotenv.Overload("/run/drone/env")
		if errDotEnv != nil {
			log.Fatalf("load /run/drone/env err: %v", errDotEnv)
		}
	}

	// app run as urfave
	if err := app.Run(os.Args); nil != err {
		log.Println(err)
	}
}

// appendCliFlag
// append cli.Flag
func appendCliFlag(target []cli.Flag, elem []cli.Flag) []cli.Flag {
	if len(elem) == 0 {
		return target
	}

	return append(target, elem...)
}
