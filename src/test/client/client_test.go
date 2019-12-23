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
		{Name: "test Send Text Message", Input: "", Output: ""},
	}
	client := dingding.DingClient{}
	for _, test := range tests {
		req := client.SendLinkMsg()
	}

}

func TestClientSendLinkMsg(t *testing.T) {
	// TODO to send
	type Test struct {
		Name   string
		Input  string
		Output string
	}
	tests := []Test{
		{Name: "test Send Link Message", Input: "", Output: ""},
	}
	client := dingding.DingClient{}
	for _, test := range tests {
		client.SendLinkMsg()
	}

}
