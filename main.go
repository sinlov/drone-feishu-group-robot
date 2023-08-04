package main

import (
	_ "embed"
	"fmt"
	"github.com/sinlov/drone-feishu-group-robot/feishu_plugin"
	"github.com/sinlov/drone-info-tools/drone_log"
	"github.com/sinlov/drone-info-tools/drone_urfave_cli_v2"
	"github.com/sinlov/drone-info-tools/pkgJson"
	"github.com/sinlov/drone-info-tools/template"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

const (
	CopyrightStartYear = "2023"
)

//go:embed package.json
var packageJson string

func action(c *cli.Context) error {

	isDebug := c.Bool(feishu_plugin.NamePluginDebug)
	if isDebug {
		drone_log.OpenDebug()
	}

	drone := drone_urfave_cli_v2.UrfaveCliBindDroneInfo(c)

	cliVersion := pkgJson.GetPackageJsonVersionGoStyle()

	if isDebug {
		log.Printf("debug: cli version is %s", cliVersion)
		log.Printf("debug: load droneInfo finish at link: %v\n", drone.Build.Link)
	}
	p := feishu_plugin.BindFlag(c, cliVersion, pkgJson.GetPackageJsonName(), drone)
	err := p.Exec()

	if err != nil {
		log.Fatalf("err: %v", err)
		return err
	}
	return nil
}

func main() {
	pkgJson.InitPkgJsonContent(packageJson)
	template.RegisterSettings(template.DefaultFunctions)
	name := pkgJson.GetPackageJsonName()

	app := cli.NewApp()
	app.Version = pkgJson.GetPackageJsonVersionGoStyle()
	app.Name = "Drone feishu Message FeishuPlugin"
	app.Usage = "Sending message to feishu group by robot using WebHook"
	app.Name = name
	app.Usage = pkgJson.GetPackageJsonDescription()
	year := time.Now().Year()
	jsonAuthor := pkgJson.GetPackageJsonAuthor()
	app.Copyright = fmt.Sprintf("© %s-%d %s by: %s, run on %s %s",
		CopyrightStartYear, year, jsonAuthor.Name, runtime.Version(), runtime.GOOS, runtime.GOARCH)
	author := &cli.Author{
		Name:  jsonAuthor.Name,
		Email: jsonAuthor.Email,
	}
	app.Authors = []*cli.Author{
		author,
	}

	app.Action = action
	flags := drone_urfave_cli_v2.UrfaveCliAppendCliFlag(drone_urfave_cli_v2.DroneInfoUrfaveCliFlag(), feishu_plugin.Flag())
	flags = drone_urfave_cli_v2.UrfaveCliAppendCliFlag(flags, feishu_plugin.HideFlag())
	flags = drone_urfave_cli_v2.UrfaveCliAppendCliFlag(flags, feishu_plugin.CommonFlag())
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
