package drone_api

import "fmt"

const (
	AuthHeadKey   = "Authorization"
	baseApi       = "api"
	baseUserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	// https://github.com/harness/drone/blob/master/core/status.go#L20

	StatusSkipped  = "skipped"
	StatusBlocked  = "blocked"
	StatusDeclined = "declined"
	StatusWaiting  = "waiting_on_dependencies"
	StatusPending  = "pending"
	StatusRunning  = "running"
	StatusPassing  = "success"
	StatusFailing  = "failure"
	StatusKilled   = "killed"
	StatusError    = "error"
)

var (
	_apiDroneBaseUrl = ""
	_apiBaseUrl      = ""
)

func SetApiBase(baseUrl string) {
	_apiDroneBaseUrl = baseUrl
	_apiBaseUrl = fmt.Sprintf("%s/%s", baseUrl, baseApi)
}

func BaseUrl() string {
	if _apiDroneBaseUrl == "" {
		panic("please use SetApiBase() first")
	}
	return _apiDroneBaseUrl
}

func ApiBase() string {
	if _apiBaseUrl == "" {
		panic("please use SetApiBase() first")
	}
	return _apiBaseUrl
}
