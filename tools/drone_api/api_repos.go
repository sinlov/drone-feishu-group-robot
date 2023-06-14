package drone_api

import "fmt"

func ApiReposGroup() string {
	return fmt.Sprintf("%s/%s", ApiBase(), "repos")
}

// Steps
// struct
type Steps struct {

	// Id
	//
	Id int `json:"id,omitempty"`

	// StepId
	//
	StepId int `json:"step_id,omitempty"`

	// Number
	//
	Number int `json:"number,omitempty"`

	// Name
	//
	Name string `json:"name,omitempty"`

	// Status
	//
	Status string `json:"status,omitempty"`

	// ExitCode
	//
	ExitCode int `json:"exit_code,omitempty"`

	// Started
	//
	Started int `json:"started,omitempty"`

	// Stopped
	//
	Stopped int `json:"stopped,omitempty"`

	// Version
	//
	Version int `json:"version,omitempty"`
}

// Stages
// struct
type Stages struct {

	// Id
	//
	Id int `json:"id,omitempty"`

	// RepoId
	//
	RepoId int `json:"repo_id,omitempty"`

	// BuildId
	//
	BuildId int `json:"build_id,omitempty"`

	// Number
	//
	Number int `json:"number,omitempty"`

	// Name
	//
	Name string `json:"name,omitempty"`

	// Kind
	//
	Kind string `json:"kind,omitempty"`

	// Type
	//
	Type string `json:"type,omitempty"`

	// Status
	//
	Status string `json:"status,omitempty"`

	// Errignore
	//
	Errignore bool `json:"errignore,omitempty"`

	// ExitCode
	//
	ExitCode int `json:"exit_code,omitempty"`

	// Machine
	//
	Machine string `json:"machine,omitempty"`

	// Os
	//
	Os string `json:"os,omitempty"`

	// Arch
	//
	Arch string `json:"arch,omitempty"`

	// Started
	//
	Started int `json:"started,omitempty"`

	// Stopped
	//
	Stopped int `json:"stopped,omitempty"`

	// Created
	//
	Created int `json:"created,omitempty"`

	// Updated
	//
	Updated int `json:"updated,omitempty"`

	// Version
	//
	Version int `json:"version,omitempty"`

	// OnSuccess
	//
	OnSuccess bool `json:"on_success,omitempty"`

	// OnFailure
	//
	OnFailure bool `json:"on_failure,omitempty"`

	// Steps
	//
	Steps []Steps `json:"steps,omitempty"`
}

// ReposBuild
// struct
type ReposBuild struct {

	// Id
	//
	Id int `json:"id,omitempty"`

	// RepoId
	//
	RepoId int `json:"repo_id,omitempty"`

	// Number
	//
	Number uint64 `json:"number,omitempty"`

	// Status
	//
	Status string `json:"status,omitempty"`

	// Event
	//
	Event string `json:"event,omitempty"`

	// Action
	//
	Action string `json:"action,omitempty"`

	// Link
	//
	Link string `json:"link,omitempty"`

	// Message
	//
	Message string `json:"message,omitempty"`

	// Before
	//
	Before string `json:"before,omitempty"`

	// After
	//
	After string `json:"after,omitempty"`

	// Ref
	//
	Ref string `json:"ref,omitempty"`

	// SourceRepo
	//
	SourceRepo string `json:"source_repo,omitempty"`

	// Source
	//
	Source string `json:"source,omitempty"`

	// Target
	//
	Target string `json:"target,omitempty"`

	// AuthorLogin
	//
	AuthorLogin string `json:"author_login,omitempty"`

	// AuthorName
	//
	AuthorName string `json:"author_name,omitempty"`

	// AuthorEmail
	//
	AuthorEmail string `json:"author_email,omitempty"`

	// AuthorAvatar
	//
	AuthorAvatar string `json:"author_avatar,omitempty"`

	// Sender
	//
	Sender string `json:"sender,omitempty"`

	// Started
	//
	Started int `json:"started,omitempty"`

	// Finished
	//
	Finished int `json:"finished,omitempty"`

	// Created
	//
	Created int `json:"created,omitempty"`

	// Updated
	//
	Updated int `json:"updated,omitempty"`

	// Version
	//
	Version int `json:"version,omitempty"`

	// Stages
	//
	Stages []Stages `json:"stages,omitempty"`
}
