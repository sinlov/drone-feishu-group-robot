package feishu_plugin_test

import (
	"github.com/sinlov/drone-info-tools/drone_info"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnvKeys(t *testing.T) {
	// mock EnvKeys
	const keyEnvs = "ENV_KEYS"
	t.Logf("~> mock EnvKeys")
	drone_info.MockDroneInfoEnvFull(false)
	// do EnvKeys
	t.Logf("~> do EnvKeys")
	assert.False(t, fetchOsEnvBool(drone_info.EnvKeyPluginDebug, false))
	assert.Equal(t, 1, fetchOsEnvInt(drone_info.EnvDroneBuildNumber, 0))
	envArray := fetchOsEnvArray(keyEnvs)
	assert.Nil(t, envArray)

	setEnvStr(t, keyEnvs, "foo, bar,My ")
	envArray = fetchOsEnvArray(keyEnvs)

	// verify EnvKeys
	assert.NotNil(t, envArray)
	assert.Equal(t, "foo", envArray[0])
	assert.Equal(t, "bar", envArray[1])
	assert.Equal(t, "My", envArray[2])
}
