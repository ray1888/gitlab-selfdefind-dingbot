package impl

import (
	"log"

	"github.com/ray1888/self-defined-dingbot/src/messages/codeplatform"
	"github.com/ray1888/self-defined-dingbot/src/messages/codeplatform/gitlab"
)

type GitlabDecoder struct {
	Msg     string
	JsonMsg string
}

func (gd *GitlabDecoder) Decode(b codeplatform.GitlabBody) {

	switch b.(type) {
	case gitlab.MergeRequestBody:
		gd.Msg = MRParse(b)
	case gitlab.PushBody:
		gd.Msg = PushParse(b)
	default:
		log.Fatal("")
	}
}

func PushParse(body codeplatform.GitlabBody) string {
	return ""
}

func MRParse(body codeplatform.GitlabBody) string {
	return ""
}
