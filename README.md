[![go-ubuntu](https://github.com/sinlov/drone-feishu-group-robot/workflows/go-ubuntu/badge.svg?branch=main)](https://github.com/sinlov/drone-feishu-group-robot/actions)
[![GoDoc](https://godoc.org/github.com/sinlov/drone-feishu-group-robot?status.png)](https://godoc.org/github.com/sinlov/drone-feishu-group-robot/)
[![GoReportCard](https://goreportcard.com/badge/github.com/sinlov/drone-feishu-group-robot)](https://goreportcard.com/report/github.com/sinlov/drone-feishu-group-robot)
[![codecov](https://codecov.io/gh/sinlov/drone-feishu-group-robot/branch/main/graph/badge.svg)](https://codecov.io/gh/sinlov/drone-feishu-group-robot)

## for what

- this project used to drone CI

## Pipeline Settings (.drone.yml)

`1.x`

```yaml
steps:

  - name: notification-feishu-group-robot
    image: sinlov/drone-feishu-group-robot:1.0.1-alpine
    pull: if-not-exists
    settings:
      debug: false
      feishu_webhook:
        # https://docs.drone.io/pipeline/environment/syntax/#from-secrets
        from_secret: feishu_group_bot_token
      feishu_secret:
        from_secret: feishu_group_secret_bot
      feishu_msg_title: your-group-message-title # default [Drone CI Notification]
      timeout_second: 10 # default 10
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

# dev

## depends

in go mod project

```bash
# warning use privte git host must set
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
$ go list -v -m -versions github.com/sinlov/drone-feishu-group-robot
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
$ docker run --rm sinlov/drone-feishu-group-robot:1.0.1-alpine -h
```
