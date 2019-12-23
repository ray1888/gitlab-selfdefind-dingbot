package codeplatform

var MergeRequest = "merge_Request"
var Push = "push"
var Tag = "tag"

type GitlabBody interface {
	GetReqType() string
}
