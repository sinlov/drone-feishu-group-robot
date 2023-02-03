package main

import (
	_ "embed"
	"fmt"
	"github.com/sinlov/drone-feishu-group-robot/feishu_plugin"
	"github.com/sinlov/drone-info-tools/drone_urfave_cli_v2"
	"github.com/sinlov/drone-info-tools/pkgJson"
	"github.com/sinlov/drone-info-tools/template"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

const (
	Name = "drone-feishu-group-robot"
)

//go:embed package.json
var packageJson string

func action(c *cli.Context) error {

	isDebug := c.Bool("config.debug")

	drone := drone_urfave_cli_v2.UrfaveCliBindDroneInfo(c)

	cliVersion := pkgJson.GetPackageJsonVersionGoStyle()

	if isDebug {
		log.Printf("debug: cli version is %s", cliVersion)
		log.Printf("debug: load droneInfo finish at link: %v\n", drone.Build.Link)
	}
	p := feishu_plugin.BindFlag(c, cliVersion, Name, drone)
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
	app := cli.NewApp()
	app.Version = pkgJson.GetPackageJsonVersionGoStyle()
	app.Name = "Drone feishu Message FeishuPlugin"
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
