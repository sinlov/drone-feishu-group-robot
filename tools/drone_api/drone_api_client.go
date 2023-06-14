package drone_api

import (
	"fmt"
	"github.com/monaco-io/request"
	"github.com/monaco-io/request/response"
	"log"
	"net/http"
	"strconv"
	"time"
)

type (
	DroneApiClient struct {
		isDebug        bool
		DroneServerUrl string
		TimeoutSecond  uint

		baseHeader map[string]string

		lastRecordBuildFailNumber    uint64
		lastRecordRepoBuildFail      *ReposBuild
		lastRecordBuildSuccessNumber uint64
		lastRecordRepoBuildSuccess   *ReposBuild
	}
)

func NewDroneApiClient(serverUrl, token string, timeoutSecond uint, isDebug bool) DroneApiClient {
	SetApiBase(serverUrl)
	baseHeader := map[string]string{
		"User-Agent": baseUserAgent + " github.com/sinlov/filebrowser-client",
		AuthHeadKey:  fmt.Sprintf("Bearer %s", token),
	}
	return DroneApiClient{
		isDebug:        isDebug,
		DroneServerUrl: serverUrl,
		TimeoutSecond:  timeoutSecond,
		baseHeader:     baseHeader,
	}
}

// FetchBuildInfo
// will fetch build info
// owner drone build owner
// repo drone build repo
// perPage set 0 will change to default 10
// can use LastRecordBuildFailNumber() LastRecordRepoBuildFail() LastRecordBuildSuccessNumber() LastRecordRepoBuildSuccess()
func (d *DroneApiClient) FetchBuildInfo(owner, repo string, perPage uint) error {
	if perPage == 0 {
		perPage = 10
	}
	c := request.Client{
		Timeout: time.Duration(d.TimeoutSecond) * time.Second,
		URL:     fmt.Sprintf("%s/%s/%s/builds/", ApiReposGroup(), owner, repo),
		Method:  request.GET,
		Header:  d.baseHeader,
		Query: map[string]string{
			"page":     "1",
			"per_page": strconv.FormatUint(uint64(perPage), 10),
		},
	}
	var repoBuilds []ReposBuild
	_, err := d.sendRespJson(c, &repoBuilds, "Repos builds")
	if err != nil {
		return err
	}
	if len(repoBuilds) == 0 {
		return nil
	}

	var lastFailNumber uint64
	var lastFailRepoBuild ReposBuild
	var lastSuccessNumber uint64
	var lastSuccessRepoBuild ReposBuild
	for _, repoBuild := range repoBuilds {
		if repoBuild.Status == StatusPending || repoBuild.Status == StatusRunning || repoBuild.Status == StatusBlocked {
			continue
		}

		if repoBuild.Status == StatusError || repoBuild.Status == StatusFailing {
			if lastFailNumber < repoBuild.Number {
				lastFailNumber = repoBuild.Number
				lastFailRepoBuild = repoBuild
			}
		}

		if repoBuild.Status == StatusPassing {
			if lastSuccessNumber < repoBuild.Number {
				lastSuccessNumber = repoBuild.Number
				lastSuccessRepoBuild = repoBuild
			}
		}
	}

	d.lastRecordBuildFailNumber = lastFailNumber
	d.lastRecordRepoBuildFail = &lastFailRepoBuild
	d.lastRecordBuildSuccessNumber = lastSuccessNumber
	d.lastRecordRepoBuildSuccess = &lastSuccessRepoBuild

	return nil
}

// LastRecordBuildFailNumber
// must use before FetchBuildInfo()
func (d *DroneApiClient) LastRecordBuildFailNumber() uint64 {
	return d.lastRecordBuildFailNumber
}

// LastRecordRepoBuildFail
// must use before FetchBuildInfo()
func (d *DroneApiClient) LastRecordRepoBuildFail() *ReposBuild {
	return d.lastRecordRepoBuildFail
}

// LastRecordBuildSuccessNumber
// must use before FetchBuildInfo()
func (d *DroneApiClient) LastRecordBuildSuccessNumber() uint64 {
	return d.lastRecordBuildSuccessNumber
}

// LastRecordRepoBuildSuccess
// must use before FetchBuildInfo()
func (d *DroneApiClient) LastRecordRepoBuildSuccess() *ReposBuild {
	return d.lastRecordRepoBuildSuccess
}

func (d *DroneApiClient) sendRespJson(c request.Client, data interface{}, apiMark string) (*response.Sugar, error) {
	if d.isDebug {
		log.Printf("debug: DroneApiClient sendRespJson url: %s ", c.URL)
		c.PrintCURL()
	}
	send := c.Send()
	if !send.OK() {
		return send, fmt.Errorf("try %v send url [ %v ] fail: %v", apiMark, c.URL, send.Error())
	}
	if send.Code() != http.StatusOK {
		return send, fmt.Errorf("try %v send [ %v ] fail: code [ %v ], msg: %v", apiMark, c.URL, send.Code(), send.String())
	}
	if d.isDebug {
		log.Printf("debug: DroneApiClient sendRespJson try %v succes by code [ %v ], content:\n%s", apiMark, send.Code(), send.String())
	}

	resp := send.ScanJSON(data)
	if !resp.OK() {
		return resp, fmt.Errorf("try DroneApiClient sendRespJson %v ScanJSON fail: %v", apiMark, resp.Error())
	}
	return resp, nil
}
