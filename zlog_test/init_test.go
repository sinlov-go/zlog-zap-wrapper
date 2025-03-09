package zlog_test

import (
	"fmt"
	"github.com/sinlov-go/unittest-kit/unittest_file_kit"
	"github.com/sinlov-go/zlog-zap-wrapper/internal/constant"
	"github.com/sinlov-go/zlog-zap-wrapper/zlog"
	"path/filepath"
	"runtime"
)

var (
	// testBaseFolderPath
	//  test base dir will auto get by package init()
	testBaseFolderPath = ""
	testGoldenKit      *unittest_file_kit.TestGoldenKit
)

func init() {
	testBaseFolderPath, _ = getCurrentFolderPath()
	zlog.MockZapLoggerInit()
	mockTestEnv()
	testGoldenKit = unittest_file_kit.NewTestGoldenKit(testBaseFolderPath)
}

func mockTestEnv() {
	// mock mode will open mock mode
	constant.UnitTestMockModeOpen()
}

// test case basic tools start
// getCurrentFolderPath
//
//	can get run path this golang dir
func getCurrentFolderPath() (string, error) {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return "", fmt.Errorf("can not get current file info")
	}
	return filepath.Dir(file), nil
}

// test case basic tools end
