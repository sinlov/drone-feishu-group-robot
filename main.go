package main

import (
	"fmt"
	"github.com/sinlov/drone-feishu-group-robot/drone_info"
	"github.com/sinlov/drone-feishu-group-robot/feishu_plugin"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

// Version of cli
var Version = "v1.0.1"

func action(c *cli.Context) error {

	isDebug := c.Bool("config.debug")

	drone := bindDroneInfo(c)

	if isDebug {
		log.Printf("load droneInfo finish at link: %v\n", drone.Build.Link)
	}
	config := feishu_plugin.Config{
		Debug:             c.Bool("config.debug"),
		TimeoutSecond:     c.Int("config.timeout_second"),
		Webhook:           c.String("config.webhook"),
		Secret:            c.String("config.secret"),
		MsgType:           c.String("config.msg_type"),
		Title:             c.String("config.msg_title"),
		PoweredByImageKey: c.String("config.msg_powered_by_image_key"),
		PoweredByImageAlt: c.String("config.msg_powered_by_image_alt"),
	}

	p := feishu_plugin.Plugin{
		Drone:  *drone,
		Config: config,
	}
	err := p.Exec()

	if err != nil {
		log.Fatalf("err: %v", err)
		return err
	}

	return nil
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
	app.Flags = []cli.Flag{
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
			Name:    "config.webhook,webhook",
			Usage:   "feishu webhook for send message",
			EnvVars: []string{"PLUGIN_FEISHU_WEBHOOK"},
		},
		&cli.StringFlag{
			Name:    "config.secret,secret",
			Usage:   "feishu secret",
			EnvVars: []string{"PLUGIN_FEISHU_SECRET"},
		},
		&cli.StringFlag{
			Name:    "config.msg_type,msg_type",
			Usage:   "feishu message type",
			EnvVars: []string{"PLUGIN_FEISHU_MSG_TYPE"},
		},
		&cli.StringFlag{
			Name:    "config.msg_title,msg_title",
			Usage:   "feishu message title",
			EnvVars: []string{"PLUGIN_FEISHU_MSG_TITLE"},
		},
		&cli.StringFlag{
			Name:    "config.msg_powered_by_image_key,msg_powered_by_image_key",
			Usage:   "feishu message powered by image key",
			EnvVars: []string{"PLUGIN_FEISHU_MSG_POWERED_BY_IMAGE_KEY"},
		},
		&cli.StringFlag{
			Name:    "config.msg_powered_by_image_alt,msg_powered_by_image_alt",
			Usage:   "feishu message powered by image alt",
			EnvVars: []string{"PLUGIN_FEISHU_MSG_POWERED_BY_IMAGE_ALT"},
		},
		// plugin end
		// droneInfo start
		&cli.StringFlag{
			Name:    "commit.author.username",
			Usage:   "providers the author username for the current commit",
			EnvVars: []string{drone_info.EnvDroneCommitAuthor},
		},
		&cli.StringFlag{
			Name:    "commit.author.avatar",
			Usage:   "providers the author avatar url for the current commit",
			EnvVars: []string{drone_info.EnvDroneCommitAuthorAvatar},
		},
		&cli.StringFlag{
			Name:    "commit.author.email",
			Usage:   "providers the author email for the current commit",
			EnvVars: []string{drone_info.EnvDroneCommitAuthorEmail},
		},
		&cli.StringFlag{
			Name:    "commit.author.name",
			Usage:   "providers the author name for the current commit",
			EnvVars: []string{drone_info.EnvDroneCommitAuthor},
		},
		&cli.StringFlag{
			Name:    "commit.branch",
			Usage:   "providers the branch for the current build",
			EnvVars: []string{drone_info.EnvDroneCommitBranch},
			Value:   "master",
		},
		&cli.StringFlag{
			Name:    "commit.link",
			Usage:   "providers the http link to the current commit in the remote source code management system(e.g.GitHub)",
			EnvVars: []string{drone_info.EnvDroneCommitLink},
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
		&cli.StringFlag{
			Name:    "repo.full.name",
			Usage:   "providers the full name of the repository",
			EnvVars: []string{drone_info.EnvDroneRepo},
		},
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
			Name:    "repo.remote.url",
			Usage:   "providers the remote url of the repository",
			EnvVars: []string{drone_info.EnvDroneRemoteUrl},
		},
		&cli.StringFlag{
			Name:    "repo.owner",
			Usage:   "providers the owner of the repository",
			EnvVars: []string{drone_info.EnvDroneRepoOwner},
		},
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
			Name:    "build.tag",
			Usage:   "build tag",
			EnvVars: []string{drone_info.EnvDroneTag},
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
		// droneInfo end
	}

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

func bindDroneInfo(c *cli.Context) *drone_info.Drone {
	var drone = drone_info.Drone{
		//  repo info
		Repo: drone_info.Repo{
			ShortName: c.String("repo.name"),
			GroupName: c.String("repo.group"),
			OwnerName: c.String("repo.owner"),
			RemoteURL: c.String("repo.remote.url"),
			FullName:  c.String("repo.full.name"),
		},
		//  build info
		Build: drone_info.Build{
			Status:     c.String("build.status"),
			Number:     c.Uint64("build.number"),
			Tag:        c.String("build.tag"),
			Link:       c.String("build.link"),
			Event:      c.String("build.event"),
			StartAt:    c.Uint64("build.started"),
			FinishedAt: c.Uint64("build.finished"),
			PR:         c.String("pull.request"),
			DeployTo:   c.String("deploy.to"),
		},
		Commit: drone_info.Commit{
			Sha:     c.String("commit.sha"),
			Branch:  c.String("commit.branch"),
			Message: c.String("commit.message"),
			Link:    c.String("commit.link"),
			Author: drone_info.CommitAuthor{
				Avatar:   c.String("commit.author.avatar"),
				Email:    c.String("commit.author.email"),
				Name:     c.String("commit.author.name"),
				Username: c.String("commit.author.username"),
			},
		},
		Stage: drone_info.Stage{
			StartedAt:  c.Uint64("stage.started"),
			FinishedAt: c.Uint64("stage.finished"),
		},
	}
	return &drone
}
