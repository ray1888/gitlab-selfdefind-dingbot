package client

import (
	dingding "github.com/ray1888/self-defined-dingbot/src/clientsdk/client/impl/dingding"
	"testing"
)

func TestClientSendTextMsg(t *testing.T) {
	// TODO to send
	type Test struct {
		Name   string
		Input  string
		Output string
	}
	tests := []Test{
		{Name: "test Send Text Message", Input: "gitlab\n" +
			"John Smith pushed to branch master at repository Diaspora\n" +
			"commit: da1560886d4f094c3e6c9ef40349f7d38b5d27d7\n" +
			"commit-messages:\n" +
			"b6568db1b: Update Catalan translation to e38cb41.\n" +
			"da1560886: fixed readme\n" +
			"项目：Diaspora",
			Output: ""},
		{Name: "test MergeRequest Message",
			Input:
			"gitlab\n"+
				"Administrator opened the merge request\n" +
				"合入方向: ms-viewport -> master\n" +
				"last_commit_msg:\n" +
				"da1560886d: fixed readme\n" +
				"状态: unchecked\n" +
				"标题: MS-Viewport\n" +
				"项目：Awesome Project\n" +
				"创建时间：2013-12-03T17:15:43Z\n" +
				"更新时间：2013-12-03T17:15:43Z\n" +
				"链接: http://example.com/diaspora/merge_requests/1\n"+
				"代码审阅人：@User1\n",
			Output: ""},
	}
	client := dingding.Init()
	for _, test := range tests {
		client.SendTextMsg(test.Input)
	}

}

