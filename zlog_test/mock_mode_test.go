package zlog_test

import (
	"github.com/sinlov-go/zlog-zap-wrapper/internal/constant"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnittestMockMode(t *testing.T) {
	assert.True(t, constant.EnableUnitTestMockMode())
}
