package drone_api_test

import (
	"bytes"
	"encoding/json"
	"github.com/sinlov/drone-feishu-group-robot/tools/drone_api"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestDroneApiClient_FetchBuildInfo(t *testing.T) {

	if envCheck(t) {
		return
	}

	// mock DroneApiClient_FetchBuildInfo

	droneApiClient := drone_api.NewDroneApiClient(
		envDroneServerUrl,
		envPluginDroneSystemAdminToken,
		defTimeoutSecond, envDebug)
	err := droneApiClient.FetchBuildInfo(envDroneRepoOwner, envDroneRepoName, 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("~> mock DroneApiClient_FetchBuildInfo")
	// do DroneApiClient_FetchBuildInfo
	t.Logf("Last Record Build Success Number: %d", droneApiClient.LastRecordBuildSuccessNumber())
	t.Logf("Last Record Build Fail Number: %d", droneApiClient.LastRecordBuildFailNumber())
	t.Logf("~> do DroneApiClient_FetchBuildInfo")

	lastRecordRepoBuildSuccess := droneApiClient.LastRecordRepoBuildSuccess()
	lastSuccessInfo, errMarshal := json.Marshal(lastRecordRepoBuildSuccess)
	if errMarshal != nil {
		t.Fatalf("lastSuccesBuild info err: %v", errMarshal)
	}
	var str bytes.Buffer
	err = json.Indent(&str, []byte(lastSuccessInfo), "", "    ")
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("checkIgnoreLastSuccessByAdminToken lastSuccesBuild info\n%v\n", str.String())
	// verify DroneApiClient_FetchBuildInfo
	assert.Less(t, uint64(1), droneApiClient.LastRecordBuildSuccessNumber())
}
