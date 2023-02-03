[![go-ubuntu](https://github.com/sinlov/drone-feishu-group-robot/workflows/go-ubuntu/badge.svg?branch=main)](https://github.com/sinlov/drone-feishu-group-robot/actions)
[![GoDoc](https://godoc.org/github.com/sinlov/drone-feishu-group-robot?status.png)](https://godoc.org/github.com/sinlov/drone-feishu-group-robot/)
[![GoReportCard](https://goreportcard.com/badge/github.com/sinlov/drone-feishu-group-robot)](https://goreportcard.com/report/github.com/sinlov/drone-feishu-group-robot)
[![codecov](https://codecov.io/gh/sinlov/drone-feishu-group-robot/branch/main/graph/badge.svg)](https://codecov.io/gh/sinlov/drone-feishu-group-robot)
![docker version semver](https://img.shields.io/docker/v/sinlov/drone-feishu-group-robot?sort=semver)
[![docker image size](https://img.shields.io/docker/image-size/sinlov/drone-feishu-group-robot)](https://hub.docker.com/r/sinlov/drone-feishu-group-robot)
[![docker pulls](https://img.shields.io/docker/pulls/sinlov/drone-feishu-group-robot)](https://hub.docker.com/r/sinlov/drone-feishu-group-robot/tags?page=1&ordering=last_updated)

## for what

- this project used to drone CI use [Custom bot guide](https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN?lang=en-US)
- support drone-docker `platform=linux/arm,linux/arm64,linux/amd64`
- support drone-exec

## Features

- more see [features/README.md](features/README.md)

### v1.2.0

- failure notification

![img.png](features/release-v1.2.0/img-v1.2.0-failure.png)

- success notification

![img.png](features/release-v1.2.0/img-v1.2.0-success.png)

## before use

- Configure webhook like `https://open.feishu.cn/open-apis/bot/v2/hook/{web_hook}` end `{web_hook}`
- `{web_hook}` must settings at `settings.feishu_webhook`or `PLUGIN_FEISHU_WEBHOOK`
- if set `Custom keywords` you can change `settings.feishu_msg_title` or `PLUGIN_FEISHU_MSG_TITLE`
- or set `Signature validation` by `settings.feishu_secret` or `PLUGIN_FEISHU_SECRET`

## Pipeline Settings (.drone.yml)

- `1.x` docker

```yaml
steps:
  - name: notification-feishu-group-robot
    image: sinlov/drone-feishu-group-robot:1.3.0-alpine
    pull: if-not-exists
    settings:
      debug: false
#      ntp_target: "pool.ntp.org" # if not set will not sync
      timeout_second: 10 # default 10
      feishu_webhook:
        # https://docs.drone.io/pipeline/environment/syntax/#from-secrets
        from_secret: feishu_group_bot_token
      feishu_secret:
        from_secret: feishu_group_secret_bot
      # let notification card change more info see https://open.feishu.cn/document/ukTMukTMukTM/uAjNwUjLwYDM14CM2ATN
      feishu_msg_title: "Drone CI Notification" # default [Drone CI Notification]
      feishu_enable_forward: true
      feishu_oss_host: "https://xxx.com" # OSS host for show oss info, if empty will not show oss info
      feishu_oss_info_user: "admin" # OSS user for show at card
      feishu_oss_info_path: "dist/foo/bar" # OSS path for show at card
      feishu_oss_resource_url: "https://xxx.com/s/xxx" # OSS resource url
      feishu_oss_page_url: "https://xxx.com/p/xxx" # OSS page url
      feishu_oss_page_passwd: "abc_xyz" # OSS password at page url, will hide PLUGIN_FEISHU_OSS_RESOURCE_URL when PAGE_PASSWD not empty 
    when:
      event: # https://docs.drone.io/pipeline/exec/syntax/conditions/#by-event
        - promote
        - rollback
        - push
        - pull_request
        - tag
      status: # only support failure/success,  both open will send anything
        - failure
        - success
```

- `1.x` drone-exec only support env
- download by [https://github.com/sinlov/drone-feishu-group-robot/releases](https://github.com/sinlov/drone-feishu-group-robot/releases) to get platform binary, then has local path
- binary path like `C:\Drone\drone-runner-exec\plugins\drone-feishu-group-robot.exe` can be drone run env like `EXEC_DRONE_FEISHU_GROUP_ROBOT_FULL_PATH`
- env:EXEC_DRONE_FEISHU_GROUP_ROBOT_FULL_PATH can set at file which define as [DRONE_RUNNER_ENVFILE](https://docs.drone.io/runner/exec/configuration/reference/drone-runner-envfile/) to support each platform to send feishu message

```yaml
steps:
  - name: notification-feishu-group-robot-exec # must has env EXEC_DRONE_FEISHU_GROUP_ROBOT_FULL_PATH and exec tools
    environment:
      PLUGIN_DEBUG: false
      # PLUGIN_NTP_TARGET: "pool.ntp.org" # if not set will not sync
      PLUGIN_TIMEOUT_SECOND: 10 # default 10
      PLUGIN_FEISHU_WEBHOOK:
        from_secret: feishu_group_bot_token
      PLUGIN_FEISHU_SECRET:
        from_secret: feishu_group_secret_bot
      # let notification card change more info see https://open.feishu.cn/document/ukTMukTMukTM/uAjNwUjLwYDM14CM2ATN
      PLUGIN_FEISHU_MSG_TITLE: "Drone CI Notification" # default [Drone CI Notification]
      PLUGIN_FEISHU_ENABLE_FORWARD: true
      PLUGIN_FEISHU_OSS_HOST: "https://xxx.com" # OSS host for show oss info, if empty will not show oss info
      PLUGIN_FEISHU_OSS_INFO_USER: "admin" # OSS user for show at card
      PLUGIN_FEISHU_OSS_INFO_PATH: "dist/foo/bar" # OSS path for show at card 
      PLUGIN_FEISHU_OSS_RESOURCE_URL: "https://xxx.com/s/xxx" # OSS resource url
      PLUGIN_FEISHU_OSS_PAGE_URL: "https://xxx.com/p/xxx" # OSS page url
      PLUGIN_FEISHU_OSS_PAGE_PASSWD: "abc_xyz" # OSS password at page url, will hide PLUGIN_FEISHU_OSS_RESOURCE_URL when PAGE_PASSWD not empty
    commands:
      - ${EXEC_DRONE_FEISHU_GROUP_ROBOT_FULL_PATH} `
        ""
    when:
      event: # https://docs.drone.io/pipeline/exec/syntax/conditions/#by-event
        - promote
        - rollback
        - push
        - pull_request
        - tag
      status: # only support failure/success,  both open will send anything
        - failure
        - success
```

### custom settings

- `settings.debug` or `PLUGIN_DEBUG` can open plugin debug mode
- `settings.timeout_second` or `PLUGIN_TIMEOUT_SECOND` can set send message timeout
- `settings.ntp_target` or `PLUGIN_NTP_TARGET` set ntp server to sync time for `Signature validation` by error code 19021
- `settings.feishu_msg_title` or `PLUGIN_FEISHU_MSG_TITLE` can change message card title
- `settings.feishu_enable_forward` or `PLUGIN_FEISHU_ENABLE_FORWARD` can change message share way
- `settings.feishu_msg_powered_by_image_key` or `PLUGIN_FEISHU_MSG_POWERED_BY_IMAGE_KEY` can change card img by feishu-image-key
- `settings.feishu_msg_powered_by_image_alt` or `PLUGIN_FEISHU_MSG_POWERED_BY_IMAGE_ALT` can change card img alt tag name

# dev

## depends

in go mod project

```bash
# warning use private git host must set
# global set for once
# add private git host like github.com to evn GOPRIVATE
$ go env -w GOPRIVATE='github.com'
# use ssh proxy
# set ssh-key to use ssh as http
$ git config --global url."git@github.com:".insteadOf "http://github.com/"
# or use PRIVATE-TOKEN
# set PRIVATE-TOKEN as gitlab or gitea
$ git config --global http.extraheader "PRIVATE-TOKEN: {PRIVATE-TOKEN}"
# set this rep to download ssh as https use PRIVATE-TOKEN
$ git config --global url."ssh://github.com/".insteadOf "https://github.com/"

# before above global settings
# test version info
$ git ls-remote -q http://github.com/sinlov/drone-feishu-group-robot.git

# test depends see full version
$ go list -mod=readonly -v -m -versions github.com/sinlov/drone-feishu-group-robot
# or use last version add go.mod by script
$ echo "go mod edit -require=$(go list -m -versions github.com/sinlov/drone-feishu-group-robot | awk '{print $1 "@" $NF}')"
$ echo "go mod vendor"
```

## evn

- golang sdk 1.17+

```bash
make init dep
```

- test code

add env then test

```bash
export PLUGIN_FEISHU_WEBHOOK= \
  export PLUGIN_FEISHU_SECRET=
```

```bash
make test
```

- see help

```bash
make dev
```

update main.go file set env then and run

```bash
export PLUGIN_FEISHU_WEBHOOK= \
  export PLUGIN_FEISHU_SECRET= \
  export DRONE_REPO=sinlov/drone-feishu-group-robot \
  export DRONE_REPO_NAME=drone-feishu-group-robot \
  export DRONE_REPO_NAMESPACE=sinlov \
  export DRONE_REMOTE_URL=https://github.com/sinlov/drone-feishu-group-robot \
  export DRONE_REPO_OWNER=sinlov \
  export DRONE_COMMIT_AUTHOR=sinlov \
  export DRONE_COMMIT_AUTHOR_AVATAR=  \
  export DRONE_COMMIT_AUTHOR_EMAIL=sinlovgmppt@gmail.com \
  export DRONE_COMMIT_BRANCH=main \
  export DRONE_COMMIT_LINK=https://github.com/sinlov/drone-feishu-group-robot/commit/68e3d62dd69f06077a243a1db1460109377add64 \
  export DRONE_COMMIT_SHA=68e3d62dd69f06077a243a1db1460109377add64 \
  export DRONE_COMMIT_REF=refs/heads/main \
  export DRONE_COMMIT_MESSAGE="mock message commit" \
  export DRONE_STAGE_STARTED=1674531206 \
  export DRONE_STAGE_FINISHED=1674532106 \
  export DRONE_BUILD_STATUS=success \
  export DRONE_BUILD_NUMBER=1 \
  export DRONE_BUILD_LINK=https://drone.xxx.com/sinlov/drone-feishu-group-robot/1 \
  export DRONE_BUILD_EVENT=push \
  export DRONE_BUILD_STARTED=1674531206 \
  export DRONE_BUILD_FINISHED=1674532206
```

- then run

```bash
make run
```

## docker

```bash
# then test build as test/Dockerfile
$ make dockerTestRestartLatest
# if run error
# like this error
# err: missing feishu webhook, please set feishu webhook
#  fix env settings then test

# see run docker fast
$ make dockerTestRunLatest

# clean test build
$ make dockerTestPruneLatest

# see how to use
$ docker run --rm sinlov/drone-feishu-group-robot:latest -h
# or version
$ docker run --rm sinlov/drone-feishu-group-robot:1.1.0-alpine -h
```
