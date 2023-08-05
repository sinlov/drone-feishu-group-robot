package cli

import (
	"fmt"
	"github.com/sinlov/drone-feishu-group-robot/feishu_plugin"
	"github.com/sinlov/drone-info-tools/drone_log"
	"github.com/sinlov/drone-info-tools/drone_urfave_cli_v2"
	"github.com/sinlov/drone-info-tools/pkgJson"
	"github.com/urfave/cli/v2"
)

var dronePlugin *feishu_plugin.FeishuPlugin

// GlobalBeforeAction
// do command Action before flag global.
func GlobalBeforeAction(c *cli.Context) error {
	isDebug := feishu_plugin.IsBuildDebugOpen(c)
	if isDebug {
		drone_log.OpenDebug()
	}

	// bind droneInfo
	drone := drone_urfave_cli_v2.UrfaveCliBindDroneInfo(c)

	cliVersion := pkgJson.GetPackageJsonVersionGoStyle()
	drone_log.Debugf("cli version is %s\n", cliVersion)
	drone_log.Debugf("load droneInfo finish at link: %v\n", drone.Build.Link)

	p := feishu_plugin.BindFlag(c, cliVersion, pkgJson.GetPackageJsonName(), drone)

	dronePlugin = &p
	drone_log.Infof("=> start  run: %s, version: %s\n", dronePlugin.Name, dronePlugin.Version)
	return nil
}

// GlobalAction
// do cli Action before flag.
func GlobalAction(c *cli.Context) error {

	if dronePlugin == nil {
		panic(fmt.Errorf("must success run GlobalBeforeAction then run GlobalAction"))
	}
	err := dronePlugin.Exec()

	if err != nil {
		return err
	}

	return nil
}

// GlobalAfterAction
//
//	do command Action after flag global.
//
//nolint:golint,unused
func GlobalAfterAction(c *cli.Context) error {
	drone_log.Infof("=> finish run: %s, version: %s\n", dronePlugin.Name, dronePlugin.Version)
	return nil
}
