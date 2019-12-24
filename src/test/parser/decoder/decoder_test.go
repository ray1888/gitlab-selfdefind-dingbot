package decoder

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/ray1888/self-defined-dingbot/src/messages/codeplatform/gitlab"
	"github.com/ray1888/self-defined-dingbot/src/messages/middlemsg"
	"github.com/ray1888/self-defined-dingbot/src/parser/decode/impl"
)

type Test struct {
	TestName string
	TestFile string
	Output   middlemsg.Body
}

func TestGitlabDecoderPushEvent(t *testing.T) {
	decoder := new(impl.GitlabDecoder)

	tests := []Test{
		{TestName: "Decode Push Event Success", TestFile: "./fixture/push.json",
			Output: middlemsg.Body{
				EventType:    "push",
				Username:     "John Smith",
				Source:       "master",
				Project:      "Diaspora",
				CommitNumber: "da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
				Commits: []middlemsg.Commit{
					{Number: "b6568db1bc1dcd7f8b4d5a946b0b91f9dacd7327",
						Info: "Update Catalan translation to e38cb41."},
					{Number: "da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
						Info: "fixed readme"},
				},
				TotalCommitCount: 2,
			},
		},
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
		result := impl.PushParse(*PushMsg)
		if test.Output.Equal(result) != true {
			t.Fatalf("Test %s, Parse Data error. Expected is %+v\n Actual is %+v\n",
				test.TestName, test.Output, decoder.Msg)
		}
	}

}

//func TestGitlabDecoderTagEvent(t *testing.T) {
//	decoder := new(impl.GitlabDecoder)
//
//	tests := []Test{
//		{TestName: "Decode Tag Event Success", TestFile: "./fixture/tag.json", Output: ""},
//	}
//	for _, test := range tests {
//		file, err := os.Open(test.TestFile)
//		if err != nil {
//			t.Fatalf("Test %s Failed, can't open Fixture file", test.TestName)
//		}
//		data := make([]byte, 1000000)
//		_, err = file.Read(data)
//		PushMsg := new(gitlab.PushBody)
//		if err := json.NewDecoder(bytes.NewReader(data)).Decode(&PushMsg); err != nil {
//			t.Fatalf("Test %s Failed, can't parse json file to json struct", test.TestName)
//		}
//		if decoder.Msg != test.Output {
//			t.Fatalf("Test %s, Parse Data error.", test.TestName)
//		}
//	}
//
//}

func TestGitlabDecoderMergeRequestEvent(t *testing.T) {
	decoder := new(impl.GitlabDecoder)

	tests := []Test{
		{TestName: "Decode MergeRequest Event Success", TestFile: "./fixture/push.json",
			Output: middlemsg.Body{
				EventType:    "merge_request",
				Username:     "root",
				Source:       "ms-viewport",
				Target:       "master",
				Project:      "Gitlab Test",
				CommitNumber: "da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
				Commits: []middlemsg.Commit{
					{Number: "b6568db1bc1dcd7f8b4d5a946b0b91f9dacd7327",
						Info: "Update Catalan translation to e38cb41."},
					{Number: "da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
						Info: "fixed readme"},
				},
				TotalCommitCount: 2,
				CreateAt:         time.Date(2011, 12, 12, 14, 27, 31, 0, time.UTC),
				UpdatedAt:        time.Date(2013, 12, 03, 17, 15, 43, 0, time.UTC),
			},
		},
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

		if test.Output.Equal(decoder.Msg) {
			t.Fatalf("Test %s, Parse Data error.", test.TestName)
		}
	}
}
