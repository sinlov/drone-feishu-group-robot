# Features

### v1.13.+

- use new github action build workflow
- show `tag info` at drone tag build at `message card title`
- add Trigger info at each `Stage details info`

![img.png](https://github.com/sinlov/drone-feishu-group-robot/raw/main/features/release-v1.13.x/release-v1.13.0-Trigger-info.png)

## [1.12.0](https://github.com/sinlov/drone-feishu-group-robot/compare/1.11.0...v1.12.0) (2023-08-04)

### ✨ Features

* support env DRONE_BUILD_DEBUG for debug of this tools as drone admin ([cf4df62d](https://github.com/sinlov/drone-feishu-group-robot/commit/cf4df62ddd2a9d66a4a4e4cff73781beb3f455bd))

### 👷‍ Build System

* change to go 1.18 and 1.18.10 to build ([f8c04d34](https://github.com/sinlov/drone-feishu-group-robot/commit/f8c04d341be92feea0bb0d869e7f8a81197e2642))

## [1.11.0](https://github.com/sinlov/drone-feishu-group-robot/compare/v1.10.0...v1.11.0) (2023-06-14)

### Features

* let feishu_ignore_last_success_by_admin_token_distance not pass tag to notifty ([ec1555b](https://github.com/sinlov/drone-feishu-group-robot/commit/ec1555bce120d9d24c910076ae1d38d83d45fc9b))

## [1.10.0](https://github.com/sinlov/drone-feishu-group-robot/compare/v1.9.0...v1.10.0) (2023-06-14)

### Features

* let checkIgnoreLastSuccessByAdminToken not pass when now drone build status not
  success ([1b27c6a](https://github.com/sinlov/drone-feishu-group-robot/commit/1b27c6a7f80ad2dc488360a627067fce82e9e3b0))

## [1.8.0](https://github.com/sinlov/drone-feishu-group-robot/compare/v1.7.0...v1.8.0) (2023-06-14)

### Features

* add
  feishu_ignore_last_success_by_badges ([a36c4ec](https://github.com/sinlov/drone-feishu-group-robot/commit/a36c4ec1f25eed6101a0965d9689777d44bfa7f0))

## [1.7.0](https://github.com/sinlov/drone-feishu-group-robot/compare/v1.6.0...v1.7.0) (2023-03-15)

### Features

* let detail show start time and finish
  time ([bc46a56](https://github.com/sinlov/drone-feishu-group-robot/commit/bc46a5691c8c627d5614e8d7819c2a710f4c679d))

![img.png](release-v1.7.x/release-v1.7.0-detail.png)

## [1.6.0](https://github.com/sinlov/drone-feishu-group-robot/compare/v1.5.0...v1.6.0) (2023-02-11)

### Features

* fix Drone.Commit.Link compare not support, when tags Link get
  error ([c55854a](https://github.com/sinlov/drone-feishu-group-robot/commit/c55854acf74887699b754ff51ef8859937665087))
* update github.com/sinlov/drone-info-tools v1.7.0 and card render by Drone
  Tag ([6ee3f3c](https://github.com/sinlov/drone-feishu-group-robot/commit/6ee3f3cf01459e68aa97ece2358aed0973d2ff96))

![img.png](release-v1.6.x/release-v1.6-tag-render.png)

## [1.5.0](https://github.com/sinlov/drone-feishu-group-robot/compare/v1.4.0...v1.5.0) (2023-02-04)

### Features

* let flag bind and maintain at package
  feishu_plugin ([af68743](https://github.com/sinlov/drone-feishu-group-robot/commit/af687439627de513eb54750241252aa0de0d8b8c))

## [1.4.0](https://github.com/sinlov/drone-feishu-group-robot/compare/v1.3.1...v1.4.0) (2023-02-03)

### Features

* embed package.json to config
  cli ([911a269](https://github.com/sinlov/drone-feishu-group-robot/commit/911a26938ce2e81aae62e90d59523e9bb5e5e232))

### [1.3.1](https://github.com/sinlov/drone-feishu-group-robot/compare/v1.3.0...v1.3.1) (2023-02-03)

- failure notification by build success but oss status error

![img.png](release-v1.3/img-v1.3.1-fail.png)

### Bug Fixes

* add feishu_oss_info_send_result and fix template render init
  func ([affd62f](https://github.com/sinlov/drone-feishu-group-robot/commit/affd62f18aae34fb7d4b6ea3c7715de043847f1c))
* fix not set oss will show drone build
  error ([9427ed8](https://github.com/sinlov/drone-feishu-group-robot/commit/9427ed8b45a4f67df5da87cee8caa72763538b7b))

## [1.3.0](https://github.com/sinlov/drone-feishu-group-robot/compare/v1.2.0...v1.3.0) (2023-02-03)

### Features

* change to use github.com/sinlov/drone-info-tools v1.2.0 and remove useless
  code ([b00f07e](https://github.com/sinlov/drone-feishu-group-robot/commit/b00f07e93d2f484a0bbac666185ca2af6f9ec465))

### v1.2.0

- failure notification

![img.png](release-v1.2.0/img-v1.2.0-failure.png)

- success notification

![img.png](release-v1.2.0/img-v1.2.0-success.png)

### v1.1.0

- success

![img.png](release-v1.1.0/img-v1.1.0-success.png)

- failure

![img.png](release-v1.1.0/img-v1.1.0-failure.png)