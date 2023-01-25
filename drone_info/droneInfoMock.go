package drone_info

import (
	"fmt"
	"github.com/sinlov/drone-feishu-group-robot/tools"
	"log"
	"os"
	"strconv"
)

const (
	mockEnvDroneRepo              = "drone-feishu-group-robot"
	mockEnvDroneCommitAuthorEmail = "sinlovgmppt@gmail.com"
	mockEnvDroneRepoOwner         = "sinlov"
	mockEnvDroneCommitBranch      = "main"
	mockEnvDroneRemoteUrlBase     = "https://github.com"
	mockEnvDroneUrlBase           = "https://drone.xxx.com"

	mockEnvDroneStageStarted  uint64 = 1674531206
	mockEnvDroneStageFinished uint64 = 1674532106
	mockEnvDroneBuildStarted  uint64 = 1674531206
	mockEnvDroneBuildFinished uint64 = 1674532206
	mockEnvDroneBuildNumber   uint64 = 1

	mockEnvDroneBuildEvent         = "push"
	mockEnvDroneBuildStatusSuccess = "success"
	mockEnvDroneBuildStatusFailure = "failure"

	mockEnvDroneCommitMessage = "mock message commit\nmore line\nand more line\r\n"
	mockEnvDroneCommitSha     = "68e3d62dd69f06077a243a1db1460109377add64"
)

func MockDroneInfo(status string) *Drone {

	if status == "" {
		status = mockEnvDroneBuildStatusSuccess
	}

	owner := mockEnvDroneRepoOwner
	email := mockEnvDroneCommitAuthorEmail
	repoName := mockEnvDroneRepo
	gitUrl := fmt.Sprintf("%s/%s/%s", mockEnvDroneRemoteUrlBase, owner, repoName)
	commitSHA := mockEnvDroneCommitSha
	branch := mockEnvDroneCommitBranch
	droneBaseUrl := mockEnvDroneUrlBase
	buildNumber := mockEnvDroneBuildNumber
	commitMessage := tools.Str2LineRaw(mockEnvDroneCommitMessage)

	var drone = Drone{
		//  repo info
		Repo: Repo{
			ShortName: repoName,
			GroupName: mockEnvDroneRepoOwner,
			OwnerName: mockEnvDroneRepoOwner,
			RemoteURL: gitUrl,
			FullName:  fmt.Sprintf("%s/%s", owner, repoName),
		},
		//  build info
		Build: Build{
			Status:     status,
			Number:     buildNumber,
			Link:       fmt.Sprintf("%s/%s/%s/%d", droneBaseUrl, owner, repoName, buildNumber),
			Event:      mockEnvDroneBuildEvent,
			StartAt:    mockEnvDroneBuildStarted,
			FinishedAt: mockEnvDroneBuildFinished,
		},
		Commit: Commit{
			Sha:     commitSHA,
			Branch:  branch,
			Message: commitMessage,
			Link:    fmt.Sprintf("%s/commit/%s", gitUrl, commitSHA),
			Author: CommitAuthor{
				Avatar:   "",
				Email:    email,
				Name:     owner,
				Username: owner,
			},
		},
		Stage: Stage{
			StartedAt:  mockEnvDroneStageStarted,
			FinishedAt: mockEnvDroneStageFinished,
		},
	}

	return &drone
}

func MockDroneInfoEnvFull(debug bool) {
	setEnvBool("PLUGIN_DEBUG", debug)

	owner := mockEnvDroneRepoOwner
	email := mockEnvDroneCommitAuthorEmail
	repoName := mockEnvDroneRepo
	gitUrl := fmt.Sprintf("%s/%s/%s", mockEnvDroneRemoteUrlBase, owner, repoName)
	commitSHA := mockEnvDroneCommitSha
	branch := mockEnvDroneCommitBranch
	droneBaseUrl := mockEnvDroneUrlBase

	commitMessage := tools.Str2LineRaw(mockEnvDroneCommitMessage)

	setEnvStr(EnvDroneRepo, fmt.Sprintf("%s/%s", owner, repoName))
	setEnvStr(EnvDroneRepoName, repoName)
	setEnvStr(EnvDroneRepoNamespace, owner)
	setEnvStr(EnvDroneRemoteUrl, gitUrl)
	setEnvStr(EnvDroneRepoOwner, owner)
	setEnvStr(EnvDroneCommitAuthor, owner)
	setEnvStr(EnvDroneCommitAuthorAvatar, "")
	setEnvStr(EnvDroneCommitAuthorEmail, email)
	setEnvStr(EnvDroneCommitBranch, branch)
	setEnvStr(EnvDroneCommitLink, fmt.Sprintf("%s/commit/%s", gitUrl, commitSHA))
	setEnvStr(EnvDroneCommitSha, commitSHA)
	setEnvStr(EnvDroneCommitRef, fmt.Sprintf("refs/heads/%s", branch))
	setEnvStr(EnvDroneCommitMessage, commitMessage)
	setEnvU64(EnvDroneStageStarted, mockEnvDroneStageStarted)
	setEnvU64(EnvDroneStageFinished, mockEnvDroneStageFinished)
	setEnvStr(EnvDroneBuildStatus, mockEnvDroneBuildStatusSuccess)
	setEnvU64(EnvDroneBuildNumber, mockEnvDroneBuildNumber)
	setEnvStr(EnvDroneBuildLink, fmt.Sprintf("%s/%s/%s/%d", droneBaseUrl, owner, repoName, mockEnvDroneBuildNumber))
	setEnvStr(EnvDroneBuildEvent, mockEnvDroneBuildEvent)
	setEnvU64(EnvDroneBuildStarted, mockEnvDroneBuildStarted)
	setEnvU64(EnvDroneBuildFinished, mockEnvDroneBuildFinished)
}

func MockEnvDebugPrint() {
	envDebug, find := os.LookupEnv("PLUGIN_DEBUG")
	if find && envDebug == "true" {
		for _, e := range os.Environ() {
			log.Println(e)
		}
	}
}

func setEnvStr(key string, val string) {
	err := os.Setenv(key, val)
	if err != nil {
		log.Fatalf("set env %v err: %v", key, err)
	}
}

func setEnvBool(key string, val bool) {
	var err error
	if val {
		err = os.Setenv(key, "true")
	} else {
		err = os.Setenv(key, "false")
	}
	if err != nil {
		log.Fatalf("set env %v err: %v", key, err)
	}
}

func setEnvU64(key string, val uint64) {
	err := os.Setenv(key, strconv.FormatUint(val, 10))
	if err != nil {
		log.Fatalf("set env %v err: %v", key, err)
	}
}
