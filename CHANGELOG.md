# Changelog

All notable changes to this project will be documented in this file. See [convention-change-log](https://github.com/convention-change/convention-change-log) for commit guidelines.

## [1.14.0](https://github.com/sinlov/drone-feishu-group-robot/compare/1.13.0...v1.14.0) (2023-08-05)

### ✨ Features

* change format of show build link ([0139bf47](https://github.com/sinlov/drone-feishu-group-robot/commit/0139bf474ff4053a3684a2da8d760312f2e13e07))

### 👷‍ Build System

* change go version to 1.19.10 to build ([8259f6eb](https://github.com/sinlov/drone-feishu-group-robot/commit/8259f6ebb5ae37562ffabd719d88e42e045da4da))

## [1.13.0](https://github.com/sinlov/drone-feishu-group-robot/compare/1.12.0...v1.13.0) (2023-08-05)

### ✨ Features

* add title build Tag ([e73233c1](https://github.com/sinlov/drone-feishu-group-robot/commit/e73233c1778f162fb062e173f3935738e9c35c02))

* update feature doc ([20489e96](https://github.com/sinlov/drone-feishu-group-robot/commit/20489e96d8ce18fdbafcf213ebb4b787fa2c91bd))

### ♻ Refactor

* use new build system for fast build ([bac6a404](https://github.com/sinlov/drone-feishu-group-robot/commit/bac6a40437792b6ca06eed303e9cc06cf04300d1))

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

* let checkIgnoreLastSuccessByAdminToken not pass when now drone build status not success ([1b27c6a](https://github.com/sinlov/drone-feishu-group-robot/commit/1b27c6a7f80ad2dc488360a627067fce82e9e3b0))

## [1.9.0](https://github.com/sinlov/drone-feishu-group-robot/compare/v1.8.0...v1.9.0) (2023-06-14)

### Features

* add drone_system_admin_token and feishu_ignore_last_success_by_admin_token_distance support ([0af41fc](https://github.com/sinlov/drone-feishu-group-robot/commit/0af41fcdf8dea1ae1cac61857ea736b2657fd52c))

* let checkIgnoreLastSuccessByAdminToken debug show more info ([5ce5047](https://github.com/sinlov/drone-feishu-group-robot/commit/5ce5047c0ed02010204fa7e75562495e8b42f9e9))

* let drone_system_admin_token or ENV: PLUGIN_DRONE_SYSTEM_ADMIN_TOKEN fail ([1d3de85](https://github.com/sinlov/drone-feishu-group-robot/commit/1d3de85a3795bde99a5b919526e3ef82f17160ab))

## [1.8.0](https://github.com/sinlov/drone-feishu-group-robot/compare/v1.7.0...v1.8.0) (2023-06-14)

### Features

* add feishu_ignore_last_success_by_badges ([a36c4ec](https://github.com/sinlov/drone-feishu-group-robot/commit/a36c4ec1f25eed6101a0965d9689777d44bfa7f0))

## [1.7.0](https://github.com/sinlov/drone-feishu-group-robot/compare/v1.6.0...v1.7.0) (2023-03-15)

### Features

* let detail show start time and finish time ([bc46a56](https://github.com/sinlov/drone-feishu-group-robot/commit/bc46a5691c8c627d5614e8d7819c2a710f4c679d))

## [1.6.0](https://github.com/sinlov/drone-feishu-group-robot/compare/v1.5.0...v1.6.0) (2023-02-11)

### Features

* fix Drone.Commit.Link compare not support, when tags Link get error ([c55854a](https://github.com/sinlov/drone-feishu-group-robot/commit/c55854acf74887699b754ff51ef8859937665087))

* update github.com/sinlov/drone-info-tools v1.7.0 and card render by Drone Tag ([6ee3f3c](https://github.com/sinlov/drone-feishu-group-robot/commit/6ee3f3cf01459e68aa97ece2358aed0973d2ff96))

## [1.6.0](https://github.com/sinlov/drone-feishu-group-robot/compare/v1.5.0...v1.6.0) (2023-02-11)

### Features

* fix Drone.Commit.Link compare not support, when tags Link get error ([c55854a](https://github.com/sinlov/drone-feishu-group-robot/commit/c55854acf74887699b754ff51ef8859937665087))

* update github.com/sinlov/drone-info-tools v1.7.0 and card render by Drone Tag ([6ee3f3c](https://github.com/sinlov/drone-feishu-group-robot/commit/6ee3f3cf01459e68aa97ece2358aed0973d2ff96))

## [1.5.0](https://github.com/sinlov/drone-feishu-group-robot/compare/v1.4.0...v1.5.0) (2023-02-04)

### Features

* let flag bind and maintain at package feishu_plugin ([af68743](https://github.com/sinlov/drone-feishu-group-robot/commit/af687439627de513eb54750241252aa0de0d8b8c))

## [1.4.0](https://github.com/sinlov/drone-feishu-group-robot/compare/v1.3.1...v1.4.0) (2023-02-03)

### Features

* embed package.json to config cli ([911a269](https://github.com/sinlov/drone-feishu-group-robot/commit/911a26938ce2e81aae62e90d59523e9bb5e5e232))

### [1.3.1](https://github.com/sinlov/drone-feishu-group-robot/compare/v1.3.0...v1.3.1) (2023-02-03)

### Bug Fixes

* add feishu_oss_info_send_result and fix template render init func ([affd62f](https://github.com/sinlov/drone-feishu-group-robot/commit/affd62f18aae34fb7d4b6ea3c7715de043847f1c))

* fix not set oss will show drone build error ([9427ed8](https://github.com/sinlov/drone-feishu-group-robot/commit/9427ed8b45a4f67df5da87cee8caa72763538b7b))

## [1.3.0](https://github.com/sinlov/drone-feishu-group-robot/compare/v1.2.0...v1.3.0) (2023-02-03)

### Features

* change to use github.com/sinlov/drone-info-tools v1.2.0 and remove useless code ([b00f07e](https://github.com/sinlov/drone-feishu-group-robot/commit/b00f07e93d2f484a0bbac666185ca2af6f9ec465))

## 1.2.0 (2023-02-02)

* feat: add feishu_plugin can support OSS info to show ([11ab740](https://github.com/sinlov/drone-feishu-group-robot/commit/11ab740))

* feat: add more info to show drone at template ([52c05cd](https://github.com/sinlov/drone-feishu-group-robot/commit/52c05cd))

* feat: add print plugin by name ([7118042](https://github.com/sinlov/drone-feishu-group-robot/commit/7118042))

* feat: add show cli version at debug ([ed8f3c5](https://github.com/sinlov/drone-feishu-group-robot/commit/ed8f3c5))

* docs: update docker hub status to show ([735f598](https://github.com/sinlov/drone-feishu-group-robot/commit/735f598))

## <small>1.0.2 (2023-01-24)</small>

* fix: fix env not use settings feishu_webhook and feishu_secret, and update document ([ec76622](https://github.com/sinlov/drone-feishu-group-robot/commit/ec76622))

* feat: change settings of set ([3d508c4](https://github.com/sinlov/drone-feishu-group-robot/commit/3d508c4))

## <small>1.0.1 (2023-01-24)</small>

* feat: add build tag and codecov config ([92a4e2c](https://github.com/sinlov/drone-feishu-group-robot/commit/92a4e2c))

* feat: add CODECOV_TOKEN ([e28a777](https://github.com/sinlov/drone-feishu-group-robot/commit/e28a777))

* feat: add docker-image-latest for build test ([beb7048](https://github.com/sinlov/drone-feishu-group-robot/commit/beb7048))

* feat: change to build dockerx at action ([27de633](https://github.com/sinlov/drone-feishu-group-robot/commit/27de633))

* feat: change to codecov/codecov-action@v3.1.1 ([37bc78d](https://github.com/sinlov/drone-feishu-group-robot/commit/37bc78d))

* feat: mark version to v1.0.1 for building ([5843891](https://github.com/sinlov/drone-feishu-group-robot/commit/5843891))

* first commit ([99f59ee](https://github.com/sinlov/drone-feishu-group-robot/commit/99f59ee))
