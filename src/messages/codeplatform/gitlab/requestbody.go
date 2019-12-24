package gitlab

import "time"

type User struct {
	Name               string `json:"name"`
	UserNameUnderScore string `json:"user_name"`
	Username           string `json:"username"`
	Avatar_url         string `json:"avatar_url"`
}

type Project struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	WebUrl    string `json:"web_url"`
	NameSpace string `json:"namespace"`
	Url       string `json:"url"`
}

type Repository struct {
	Name        string `json:"name"`
	Url         string `json:"url"`
	Description string `json:"description"`
	HomePage    string `json:"homepage"`
}

type CommitAuthor struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Commit struct {
	Id           string    `json:"id"`
	Message      string    `json:"message"`
	TimeStamp    time.Time `json:"timestamp"`
	CommitAuthor `json:"author"`
}

type Commits []Commit

type MergeObject struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	WebUrl      string `json:"web_url"`
	Avatar_url  string `json:"avatar_url"`
	NameSpace   string `json:"namespace"`
}

type ObjectAttributes struct {
	Id           int         `json:"id"`
	TargetBranch string      `json:"target_branch"`
	SourceBranch string      `json:"source_branch"`
	AuthorId     int         `json:"author_id"`
	AssingeeId   int         `json:"assignee_id"`
	Title        string      `json:"title"`
	CreatedAt    time.Time   `json:"created_at"`
	Updatedat    time.Time   `json:"updated_at"`
	State        string      `json:"state"`
	MergeStatus  string      `json:"merge_status"`
	Description  string      `json:"description"`
	Source       MergeObject `json:"source"`
	Target       MergeObject `json:"target"`
	LastCommit   Commit      `json:"last_commit"`
}

type BaseBody struct {
	ObjectKind string `json:"object_kind"`
	User       `json:"user"`
	Project    `json:"project"`
	Repository `json:"repository"`
	Ref        string `json:"ref"`
}

// Gitlab Push Type Event data is same to its Tag type event
type PushBody struct {
	BaseBody
	Commits          `json:"commits"`
	TotalCommitCount int `json:"total_commits_count"`
}

func (pb PushBody) GetReqType() string {
	return pb.BaseBody.ObjectKind
}

type TagBody PushBody

type MergeRequestBody struct {
	BaseBody
	User             `json:"user"`
	ObjectAttributes `json:"object_attributes"`
}

func (mb MergeRequestBody) GetReqType() string {
	return mb.BaseBody.ObjectKind
}
