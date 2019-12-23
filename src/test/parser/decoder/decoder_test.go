package decoder

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	"github.com/ray1888/self-defined-dingbot/src/messages/codeplatform/gitlab"
	"github.com/ray1888/self-defined-dingbot/src/parser/decode/impl"
)

type Test struct {
	TestName string
	TestFile string
	Output   string
}

func TestGitlabDecoderPushEvent(t *testing.T) {
	decoder := new(impl.GitlabDecoder)

	tests := []Test{
		{TestName: "Decode Push Event Success", TestFile: "./fixture/push.json", Output: ""},
	}
	for _, test := range tests {
		file, err := os.Open(test.TestFile)
		if err != nil {
			t.Fatalf("Test %s Failed, can't open Fixture file", test.TestName)
		}
		data := make([]byte, 1000000)
		_, err = file.Read(data)
		PushMsg := new(gitlab.PushBody)
		if err := json.NewDecoder(bytes.NewReader(data)).Decode(&PushMsg); err != nil {
			t.Fatalf("Test %s Failed, can't parse json file to json struct", test.TestName)
		}
		decoder.Decode(PushMsg)
		if decoder.Msg != test.Output {
			t.Fatalf("Test %s, Parse Data error.", test.TestName)
		}
	}

}

func TestGitlabDecoderTagEvent(t *testing.T) {
	decoder := new(impl.GitlabDecoder)

	tests := []Test{
		{TestName: "Decode Tag Event Success", TestFile: "./fixture/tag.json", Output: ""},
	}
	for _, test := range tests {
		file, err := os.Open(test.TestFile)
		if err != nil {
			t.Fatalf("Test %s Failed, can't open Fixture file", test.TestName)
		}
		data := make([]byte, 1000000)
		_, err = file.Read(data)
		PushMsg := new(gitlab.PushBody)
		if err := json.NewDecoder(bytes.NewReader(data)).Decode(&PushMsg); err != nil {
			t.Fatalf("Test %s Failed, can't parse json file to json struct", test.TestName)
		}
		if decoder.Msg != test.Output {
			t.Fatalf("Test %s, Parse Data error.", test.TestName)
		}
	}

}

func TestGitlabDecoderMergeRequestEvent(t *testing.T) {
	decoder := new(impl.GitlabDecoder)

	tests := []Test{
		{TestName: "Decode MergeRequest Event Success", TestFile: "./fixture/mergerequest.json", Output: ""},
	}
	for _, test := range tests {
		file, err := os.Open(test.TestFile)
		if err != nil {
			t.Fatalf("Test %s Failed, can't open Fixture file", test.TestName)
		}
		data := make([]byte, 1000000)
		_, err = file.Read(data)
		MergeReqMsg := new(gitlab.MergeRequestBody)
		if err := json.NewDecoder(bytes.NewReader(data)).Decode(&MergeReqMsg); err != nil {
			t.Fatalf("Test %s Failed, can't parse json file to json struct", test.TestName)
		}
		if decoder.Msg != test.Output {
			t.Fatalf("Test %s, Parse Data error.", test.TestName)
		}
	}
}
