package middlemsg

import (
	"time"

	"github.com/ray1888/self-defined-dingbot/src/messages/codeplatform/gitlab"
)

type Commit struct {
	Number string
	Info   string
}

type Body struct {
	EventType        string
	Username         string
	Source           string
	Target           string
	Project          string
	TotalCommitCount int
	// CommitNumber 只显示最新的一个commit号码
	CommitNumber string
	// 详细的commit数量
	Commits     []Commit
	CommitsText string
	State       string
	MergeStatus string
	Assignee    gitlab.User
	CreatedAt   time.Time
	UpdatedAt   time.Time
	// 打开的链接
	Link string
	// MergeRequest Title
	Title string
}

func (b Body) Equal(body Body) bool {
	if (b.EventType == body.EventType) &&
		(b.CommitNumber == body.CommitNumber) &&
		(b.Source == body.Source) {
		return true
	} else {
		return false
	}
}

func (b Body) MergeEqual(body Body) bool {
	if (b.EventType == body.EventType) &&
		(b.CommitNumber == body.CommitNumber) &&
		(b.Source == body.Source) &&
		(b.CreatedAt == body.CreatedAt) &&
		(b.UpdatedAt == b.UpdatedAt) &&
		(b.Assignee.Username == body.Assignee.Username) &&
		(b.Assignee.Name == body.Assignee.Name) &&
		(b.Link == body.Link) {
		return true
	} else {
		return false
	}
}
