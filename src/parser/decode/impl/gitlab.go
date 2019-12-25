package impl

import (
	//"log"
	"strings"

	"github.com/ray1888/self-defined-dingbot/src/messages/codeplatform"
	"github.com/ray1888/self-defined-dingbot/src/messages/codeplatform/gitlab"
	"github.com/ray1888/self-defined-dingbot/src/messages/middlemsg"
)

type GitlabDecoder struct {
	JsonMsg      string
	InnerChannel chan (middlemsg.Body)
}

//func (gd *GitlabDecoder) Decode(b codeplatform.GitlabBody) {
//	switch b.(type) {
//	case gitlab.MergeRequestBody:
//		gd.InnerChannel <- MRParse(b)
//	case gitlab.PushBody:
//		gd.InnerChannel <- PushParse(b)
//	default:
//		log.Fatal("")
//	}
//}

func PushParse(body gitlab.PushBody) middlemsg.Body {
	msg := middlemsg.Body{}
	msg.EventType = codeplatform.Push
	msg.TotalCommitCount = body.TotalCommitCount
	msg.Project = body.Project.Name
	msg.Username = body.User.Username
	sourceBranch := strings.Split(body.Ref, "/")
	msg.Source = sourceBranch[len(sourceBranch)-1]
	length := len(body.Commits)
	msg.CommitNumber = body.Commits[length-1].Id
	msg.UpdatedAt = body.Commits[length-1].TimeStamp
	msg.CreatedAt = body.Commits[0].TimeStamp

	return msg
}

func MRParse(body gitlab.MergeRequestBody) middlemsg.Body {
	msg := middlemsg.Body{}
	msg.EventType = codeplatform.MergeRequest
	msg.Project = body.Project.Name
	//msg. = body.User.Name
	msg.Username = body.User.Username
	msg.Source = body.ObjectAttributes.SourceBranch
	msg.Target = body.ObjectAttributes.TargetBranch
	msg.CommitNumber = body.ObjectAttributes.LastCommit.Id
	msg.CreatedAt = body.ObjectAttributes.CreatedAt
	msg.UpdatedAt = body.ObjectAttributes.Updatedat
	msg.State = body.State
	msg.MergeStatus = body.MergeStatus
	msg.Assignee = body.ObjectAttributes.Assingee
	msg.Link = body.ObjectAttributes.Link
	return msg
}

//func TagParse(body gitlab.MergeRequestBody) middlemsg.Body {
//	msg := middlemsg.Body{}
//	msg.EventType = codeplatform.MergeRequest
//	msg.Project = body.Project.Name
//	msg.Username = body.User.Username
//	msg.Source = body.ObjectAttributes.SourceBranch
//	msg.Target = body.ObjectAttributes.TargetBranch
//	msg.CommitNumber = body.ObjectAttributes.LastCommit.Id
//	msg.CreatedAt = body.ObjectAttributes.CreatedAt
//	msg.UpdatedAt = body.ObjectAttributes.Updatedat
//	msg.State = body.State
//	msg.MergeStatus = body.MergeStatus
//	msg.AssigneeId = body.AssingeeId
//	return msg
//}
