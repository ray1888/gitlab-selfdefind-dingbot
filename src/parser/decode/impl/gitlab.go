package impl

import (
	"log"
	"strings"

	"github.com/ray1888/self-defined-dingbot/src/messages/codeplatform"
	"github.com/ray1888/self-defined-dingbot/src/messages/codeplatform/gitlab"
	"github.com/ray1888/self-defined-dingbot/src/messages/middlemsg"
)

type GitlabDecoder struct {
	Msg     middlemsg.Body
	JsonMsg string
}

func (gd *GitlabDecoder) Decode(b codeplatform.GitlabBody) {
	switch b.(type) {
	//case gitlab.MergeRequestBody:
	//	gd.Msg = MRParse(b)
	//case gitlab.PushBody:
	//	gd.Msg = PushParse(gitlab.PushBody(b))
	default:
		log.Fatal("")
	}
}

func PushParse(body gitlab.PushBody) middlemsg.Body {
	msg := middlemsg.Body{}
	msg.EventType = codeplatform.Push
	msg.TotalCommitCount = body.TotalCommitCount
	msg.Project = body.Project.Name
	msg.Username = body.UserName
	sourceBranch := strings.Split(body.Ref, "/")
	msg.Source = sourceBranch[len(sourceBranch)-1]
	length := len(body.Commits)
	msg.CommitNumber = body.Commits[length-1].Id
	msg.UpdatedAt = body.Commits[length-1].TimeStamp
	msg.CreateAt = body.Commits[0].TimeStamp

	return msg
}

//func MRParse(body ) middlemsg.Body {
//	return  middlemsg.Body{}
//}
