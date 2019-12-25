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
	}
	client := dingding.Init()
	for _, test := range tests {
		client.SendTextMsg(test.Input)
	}

}

//func TestClientSendLinkMsg(t *testing.T) {
//	// TODO to send
//	type Test struct {
//		Name   string
//		Input  string
//		Output string
//	}
//	tests := []Test{
//		{Name: "test Send Link Message", Input: "", Output: ""},
//	}
//	client := dingding.DingClient{}
//	for _, test := range tests {
//		client.SendLinkMsg()
//	}
//
//}
