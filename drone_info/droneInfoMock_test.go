package drone_info

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMockDroneInfoFull(t *testing.T) {
	// mock MockDroneInfoEnvFull
	MockDroneInfoEnvFull(true)
	t.Logf("~> mock MockDroneInfoEnvFull")
	// do MockDroneInfoEnvFull
	t.Logf("~> do MockDroneInfoEnvFull")
	MockEnvDebugPrint()
	// verify MockDroneInfoEnvFull

	assert.Equal(t, mockEnvDroneRepoOwner, os.Getenv(EnvDroneRepoOwner))
}

func TestMockDroneInfo(t *testing.T) {
	// mock MockDroneInfo

	droneInfo := MockDroneInfo(mockEnvDroneBuildStatusSuccess)
	t.Logf("~> mock MockDroneInfo")
	// do MockDroneInfo
	t.Logf("~> do MockDroneInfo")
	// verify MockDroneInfo
	assert.Equal(t, droneInfo.Repo.OwnerName, mockEnvDroneRepoOwner)
	assert.Equal(t, droneInfo.Repo.GroupName, mockEnvDroneRepoOwner)
	assert.Equal(t, droneInfo.Repo.ShortName, mockEnvDroneRepo)
	assert.Equal(t, droneInfo.Build.Status, mockEnvDroneBuildStatusSuccess)

	droneInfoFail := MockDroneInfo(mockEnvDroneBuildStatusFailure)
	assert.Equal(t, droneInfoFail.Build.Status, mockEnvDroneBuildStatusFailure)
}
