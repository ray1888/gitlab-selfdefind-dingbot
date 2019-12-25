package middlemsg

import "time"

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
	Assignee    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
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
		(b.UpdatedAt == b.UpdatedAt) {
		return true
	} else {
		return false
	}
}
