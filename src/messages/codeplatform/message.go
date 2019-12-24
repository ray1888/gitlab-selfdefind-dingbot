package codeplatform

var MergeRequest = "merge_request"
var Push = "push"
var Tag = "tag"

type GitlabBody interface {
	GetReqType() string
}
