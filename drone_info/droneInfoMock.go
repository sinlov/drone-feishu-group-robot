package drone_info

import (
	"fmt"
	"github.com/sinlov/drone-feishu-group-robot/tools"
	"log"
	"net/url"
	"os"
	"strconv"
)

const (
	mockEnvDroneRepo              = "drone-file-browser-plugin"
	mockEnvDroneCommitAuthorEmail = "sinlovgmppt@gmail.com"
	mockEnvDroneRepoOwner         = "sinlov"
	mockEnvDroneCommitBranch      = "main"
	mockEnvDroneRemoteUrlBase     = "https://github.com"
	mockEnvDroneUrlBase           = "https://drone.xxx.com"
	mockEnvDroneRepoScm           = "git"

	mockEnvDroneStageStarted  uint64 = 1674531206
	mockEnvDroneStageFinished uint64 = 1674532106
	mockEnvDroneBuildStarted  uint64 = 1674531206
	mockEnvDroneBuildFinished uint64 = 1674532206
	mockEnvDroneBuildNumber   uint64 = 1
	mockEnvDroneStageMachine  string = "CI-machine"
	mockEnvDroneStageOs       string = "linux"
	mockEnvDroneStageArch     string = "amd64"
	mockEnvDroneStageVariant  string = ""
	mockEnvDroneStageType     string = "docker"
	mockEnvDroneStageKind     string = "pipeline"
	mockEnvDroneStageName     string = "build"

	mockEnvDroneBuildEvent         = "push"
	mockEnvDroneBuildStatusSuccess = "success"
	mockEnvDroneBuildStatusFailure = "failure"

	mockEnvDroneCommitMessage = "mock message commit\nmore line\nand more line\r\n"
	mockEnvDroneCommitSha     = "68e3d62dd69f06077a243a1db1460109377add64"

	mockEnvFailedStages = "build,test"
	mockEnvFailedSteps  = "backend,frontend"
)

func MockDroneInfo(status string) *Drone {

	if status == "" {
		status = mockEnvDroneBuildStatusSuccess
	}
	commitMessage := mockEnvDroneCommitMessage
	commitMessage = tools.Str2LineRaw(commitMessage)

	owner := mockEnvDroneRepoOwner
	email := mockEnvDroneCommitAuthorEmail
	repoName := mockEnvDroneRepo
	repoUrl := fmt.Sprintf("%s/%s/%s", mockEnvDroneRemoteUrlBase, owner, repoName)
	repoHttpUrl := fmt.Sprintf("%s/%s/%s.git", mockEnvDroneRemoteUrlBase, owner, repoName)
	repoSshUrl := fmt.Sprintf("git@%s:%s/%s.git", owner, repoName, "github.com")
	repoHost := ""
	repoHostName := ""
	parse, err := url.Parse(repoHttpUrl)
	if err == nil {
		repoHost = parse.Host
		repoHostName = parse.Hostname()
	}
	commitSHA := mockEnvDroneCommitSha
	branch := mockEnvDroneCommitBranch
	droneBaseUrl := mockEnvDroneUrlBase
	buildNumber := mockEnvDroneBuildNumber

	var drone = Drone{
		//  repo info
		Repo: Repo{
			ShortName: repoName,
			GroupName: mockEnvDroneRepoOwner,
			FullName:  fmt.Sprintf("%s/%s", owner, repoName),
			OwnerName: mockEnvDroneRepoOwner,
			Scm:       mockEnvDroneRepoScm,
			RemoteURL: repoUrl,
			HttpUrl:   repoHttpUrl,
			SshUrl:    repoSshUrl,
			Host:      repoHost,
			HostName:  repoHostName,
		},
		//  build info
		Build: Build{
			Status:       status,
			Number:       buildNumber,
			Tag:          "",
			Link:         fmt.Sprintf("%s/%s/%s/%d", droneBaseUrl, owner, repoName, buildNumber),
			Event:        mockEnvDroneBuildEvent,
			StartAt:      mockEnvDroneBuildStarted,
			FinishedAt:   mockEnvDroneBuildFinished,
			PR:           "",
			DeployTo:     "",
			FailedStages: mockEnvFailedStages,
			FailedSteps:  mockEnvFailedSteps,
		},
		Commit: Commit{
			Link:    fmt.Sprintf("%s/commit/%s", repoHttpUrl, commitSHA),
			Branch:  branch,
			Message: commitMessage,
			Sha:     commitSHA,
			Ref:     fmt.Sprintf("refs/heads/%s", branch),
			Author: CommitAuthor{
				Username: owner,
				Email:    email,
				Name:     owner,
				Avatar:   "",
			},
		},
		Stage: Stage{
			StartedAt:  mockEnvDroneStageStarted,
			FinishedAt: mockEnvDroneStageFinished,
			Machine:    mockEnvDroneStageMachine,
			Os:         mockEnvDroneStageOs,
			Arch:       mockEnvDroneStageArch,
			Variant:    mockEnvDroneStageVariant,
			Type:       mockEnvDroneStageType,
			Kind:       mockEnvDroneStageKind,
			Name:       mockEnvDroneStageName,
		},
	}

	return &drone
}

func MockDroneInfoEnvFull(debug bool) {
	setEnvBool("PLUGIN_DEBUG", debug)

	owner := mockEnvDroneRepoOwner
	email := mockEnvDroneCommitAuthorEmail
	repoName := mockEnvDroneRepo
	repoUrl := fmt.Sprintf("%s/%s/%s", mockEnvDroneRemoteUrlBase, owner, repoName)
	repoHttpUrl := fmt.Sprintf("%s/%s/%s.git", mockEnvDroneRemoteUrlBase, owner, repoName)
	repoSshUrl := fmt.Sprintf("git@%s:%s/%s.git", owner, repoName, "github.com")

	commitSHA := mockEnvDroneCommitSha
	branch := mockEnvDroneCommitBranch
	droneBaseUrl := mockEnvDroneUrlBase
	buildNumber := mockEnvDroneBuildNumber

	setEnvStr(EnvDroneRepo, fmt.Sprintf("%s/%s", owner, repoName))
	setEnvStr(EnvDroneRepoName, repoName)
	setEnvStr(EnvDroneRepoNamespace, owner)
	setEnvStr(EnvDroneRemoteUrl, repoUrl)
	setEnvStr(EnvDroneRepoOwner, owner)
	setEnvStr(EnvDroneGitHttpUrl, repoHttpUrl)
	setEnvStr(EnvDroneGitSshUrl, repoSshUrl)

	setEnvStr(EnvDroneBuildStatus, mockEnvDroneBuildStatusSuccess)
	setEnvU64(EnvDroneBuildNumber, mockEnvDroneBuildNumber)
	setEnvStr(EnvDroneTag, "")
	setEnvStr(EnvDroneBuildLink, fmt.Sprintf("%s/%s/%s/%d", droneBaseUrl, owner, repoName, buildNumber))
	setEnvStr(EnvDroneBuildEvent, mockEnvDroneBuildEvent)
	setEnvU64(EnvDroneBuildStarted, mockEnvDroneBuildStarted)
	setEnvU64(EnvDroneBuildFinished, mockEnvDroneBuildFinished)
	setEnvStr(EnvDroneFailedStages, "")
	setEnvStr(EnvDroneFailedSteps, "")

	setEnvStr(EnvDroneCommitAuthor, owner)
	setEnvStr(EnvDroneCommitAuthorName, owner)
	setEnvStr(EnvDroneCommitAuthorAvatar, "")
	setEnvStr(EnvDroneCommitAuthorEmail, email)
	setEnvStr(EnvDroneCommitLink, fmt.Sprintf("%s/commit/%s", repoUrl, commitSHA))
	setEnvStr(EnvDroneCommitBranch, branch)
	setEnvStr(EnvDroneCommitMessage, mockEnvDroneCommitMessage)
	setEnvStr(EnvDroneCommitSha, commitSHA)
	setEnvStr(EnvDroneCommitRef, fmt.Sprintf("refs/heads/%s", branch))

	setEnvU64(EnvDroneStageStarted, mockEnvDroneStageStarted)
	setEnvU64(EnvDroneStageFinished, mockEnvDroneStageFinished)
	setEnvStr(EnvDroneStageMachine, mockEnvDroneStageMachine)
	setEnvStr(EnvDroneStageOs, mockEnvDroneStageOs)
	setEnvStr(EnvDroneStageArch, mockEnvDroneStageArch)
	setEnvStr(EnvDroneStageVariant, mockEnvDroneStageVariant)
	setEnvStr(EnvDroneStageType, mockEnvDroneStageType)
	setEnvStr(EnvDroneStageKind, mockEnvDroneStageKind)
	setEnvStr(EnvDroneStageName, mockEnvDroneStageName)

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
		log.Fatalf("set env key [%v] string err: %v", key, err)
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
		log.Fatalf("set env key [%v] bool err: %v", key, err)
	}
}

func setEnvU64(key string, val uint64) {
	err := os.Setenv(key, strconv.FormatUint(val, 10))
	if err != nil {
		log.Fatalf("set env key [%v] uint64 err: %v", key, err)
	}
}
