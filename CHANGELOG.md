# Changelog

All notable changes to this project will be documented in this file. See [convention-change-log](https://github.com/convention-change/convention-change-log) for commit guidelines.

## [1.1.0](https://github.com/sinlov-go/zlog-zap-wrapper/compare/1.0.2...v1.1.0) (2025-03-11)

### ✨ Features

* open DestructorInit to unit test ([15321fe7](https://github.com/sinlov-go/zlog-zap-wrapper/commit/15321fe7ecfc1f5f126e586254f479485cc29b51))

## [1.0.2](https://github.com/sinlov-go/zlog-zap-wrapper/compare/1.0.1...v1.0.2) (2025-03-10)

### 👷‍ Build System

* update codecov action and rename config file ([3b54471c](https://github.com/sinlov-go/zlog-zap-wrapper/commit/3b54471c3ec70bbffd1ff440d3ec83c9181879ef))

## [1.0.1](https://github.com/sinlov-go/zlog-zap-wrapper/compare/1.0.0...v1.0.1) (2025-03-10)

### 📝 Documentation

* update golang-codecov workflow to use codecov-action v5 ([c297c36b](https://github.com/sinlov-go/zlog-zap-wrapper/commit/c297c36beda2dbc91e66ff4a43dd13f256b01220))

## 1.0.0 (2025-03-10)

### 🐛 Bug Fixes

* update file permissions before removal ([c00fd98b](https://github.com/sinlov-go/zlog-zap-wrapper/commit/c00fd98b3fb77079764b614d34f0049296038aa1))

* use `path.Join` instead of `filepath.Join` ([f94db263](https://github.com/sinlov-go/zlog-zap-wrapper/commit/f94db263e9cd410a334772112a5422e9f6d35672))

### ✨ Features

* add config.yml and update example_test.go ([90fd05a4](https://github.com/sinlov-go/zlog-zap-wrapper/commit/90fd05a4f783c03d0539d47d737a4b90af66aca9))

* implement log flavors and improve logger initialization ([50a6dab0](https://github.com/sinlov-go/zlog-zap-wrapper/commit/50a6dab0ca229f33f30562ccd420d32674bbb91c))

* add export functionality and enhance help documentation ([d71b7139](https://github.com/sinlov-go/zlog-zap-wrapper/commit/d71b71397a65a6ea2bc7adf0872c51958e4cd933))

* add new logging functionality ([876ecabd](https://github.com/sinlov-go/zlog-zap-wrapper/commit/876ecabd428b5d0462ea75d1297968a0ee0d3705))

### 📝 Documentation

* add example init and update config options ([80f4c4d0](https://github.com/sinlov-go/zlog-zap-wrapper/commit/80f4c4d0f07145ed11c331f45c435aa2892fbdf5))

* add examples for zLogInitLogger and zLogInitWithFlavors ([253b49e5](https://github.com/sinlov-go/zlog-zap-wrapper/commit/253b49e52bf349cbdf5166e598973b2dad8a085f))

### ♻ Refactor

* improve log pruning functionality ([070b034f](https://github.com/sinlov-go/zlog-zap-wrapper/commit/070b034f96c001afa227077b75a69ab8c5f2dfaf))

* replace filepath with path for log file handling ([8634b744](https://github.com/sinlov-go/zlog-zap-wrapper/commit/8634b744bc68d2f19719e3aaec918f99f452fe96))

* improve log file path construction ([cfb03251](https://github.com/sinlov-go/zlog-zap-wrapper/commit/cfb032511ed7bbd2fd2002a09935183a105cd5f9))

### 👷‍ Build System

* update Go version to 1.21 ([f4f17c95](https://github.com/sinlov-go/zlog-zap-wrapper/commit/f4f17c95ea059a1251f977b3db71fcf976cdfc0d))

* add configurable test and linting options ([595001ea](https://github.com/sinlov-go/zlog-zap-wrapper/commit/595001eab20f66179ce0e66fe608de0ee4f89bd9))
