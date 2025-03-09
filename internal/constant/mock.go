package constant

import "github.com/sinlov-go/unittest-kit/env_kit"

const (
	// EnvMockModeEnable mock mode for unit test
	//	open mode at unit test init
	// UnitTestMockModeOpen
	//	then use
	// EnableUnitTestMockMode
	//	to switch unit test case
	EnvMockModeEnable = "ENV_MOCK_MODE_ENABLE"
)

// EnableUnitTestMockMode
// check now in unit test mode by env: EnvMockModeEnable
func EnableUnitTestMockMode() bool {
	return env_kit.FetchOsEnvBool(EnvMockModeEnable, false)
}

// UnitTestMockModeOpen
// open unit test mock mode by set env: EnvMockModeEnable
func UnitTestMockModeOpen() {
	env_kit.SetEnvBool(EnvMockModeEnable, true)
}
