package drone_info

const (
	EnvDroneCommitAuthor       = "DRONE_COMMIT_AUTHOR"
	EnvDroneCommitAuthorAvatar = "DRONE_COMMIT_AUTHOR_AVATAR"
	EnvDroneCommitAuthorEmail  = "DRONE_COMMIT_AUTHOR_EMAIL"
	EnvDroneCommitBranch       = "DRONE_COMMIT_BRANCH"
	EnvDroneCommitLink         = "DRONE_COMMIT_LINK"
	EnvDroneCommitMessage      = "DRONE_COMMIT_MESSAGE"
	EnvDroneCommitSha          = "DRONE_COMMIT_SHA"
	EnvDroneCommitRef          = "DRONE_COMMIT_REF"
	EnvDroneRepo               = "DRONE_REPO"
	EnvDroneRepoName           = "DRONE_REPO_NAME"
	EnvDroneRepoNamespace      = "DRONE_REPO_NAMESPACE"
	EnvDroneRemoteUrl          = "DRONE_REMOTE_URL"
	EnvDroneRepoOwner          = "DRONE_REPO_OWNER"
	EnvDroneStageStarted       = "DRONE_STAGE_STARTED"
	EnvDroneStageFinished      = "DRONE_STAGE_FINISHED"
	EnvDroneBuildStatus        = "DRONE_BUILD_STATUS"
	EnvDroneBuildNumber        = "DRONE_BUILD_NUMBER"
	EnvDroneBuildLink          = "DRONE_BUILD_LINK"
	EnvDroneBuildEvent         = "DRONE_BUILD_EVENT"
	EnvDroneBuildStarted       = "DRONE_BUILD_STARTED"
	EnvDroneBuildFinished      = "DRONE_BUILD_FINISHED"
	EnvDroneTag                = "DRONE_TAG"
	EnvDronePR                 = "DRONE_PULL_REQUEST"
	EnvDroneDeployTo           = "DRONE_DEPLOY_TO"
)

type (
	// Repo repo base info
	Repo struct {
		ShortName string //  short name
		GroupName string //  group name
		FullName  string //  repository full name
		OwnerName string //  repo owner
		RemoteURL string //  repo remote url
	}

	// Build info
	Build struct {
		Status     string //  providers the current build status
		Number     uint64 //  providers the current build number
		Tag        string //  providers the current build tag
		Link       string //  providers the current build link
		Event      string //  trigger event
		StartAt    uint64 //  build start at ( unix timestamp )
		FinishedAt uint64 //  build finish at ( unix timestamp )
		PR         string //  build pull request
		DeployTo   string //  build deploy to
	}

	// Commit info
	Commit struct {
		Branch  string //  providers the branch for the current commit
		Link    string //  providers the http link to the current commit in the remote source code management system(e.g.GitHub)
		Message string //  providers the commit message for the current build
		Sha     string //  providers the commit sha for the current build
		Ref     string //  commit ref
		Author  CommitAuthor
	}

	// Stage drone stage env
	Stage struct {
		StartedAt  uint64
		FinishedAt uint64
	}

	// CommitAuthor commit author info
	CommitAuthor struct {
		Avatar   string //  providers the author avatar for the current commit
		Email    string //  providers the author email for the current commit
		Name     string //  providers the author name for the current commit
		Username string //  the author username for the current commit
	}

	// Drone drone info
	Drone struct {
		Repo   Repo
		Build  Build
		Commit Commit
		Stage  Stage
	}
)
