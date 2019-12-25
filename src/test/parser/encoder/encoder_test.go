package encoder

import (
	"testing"
	"time"

	"github.com/ray1888/self-defined-dingbot/src/messages/middlemsg"
	"github.com/ray1888/self-defined-dingbot/src/parser/encode/impl"
)

type Test struct {
	TestName string
	Input    middlemsg.Body
	Output   string
}

func TestEncodeTextMsg(t *testing.T) {
	tests := []Test{
		{TestName: "Encode Text Push Event Success",
			Input: middlemsg.Body{
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
			Output: "John Smith pushed to branch master at repository Diaspora\n" +
				"commit: da1560886d4f094c3e6c9ef40349f7d38b5d27d7\n" +
				"commit-messages:\n" +
				"b6568db1b: Update Catalan translation to e38cb41.\n" +
				"da1560886: fixed readme\n" +
				"项目：Diaspora",
		},
	}
	for _, test := range tests {
		result, err := impl.EncodeTextMsg(test.Input)
		if err != nil {
			t.Fatalf("TestName %s encode msg meet error +%v", test.TestName, err)
		}
		if result != test.Output {
			t.Fatalf("TestName %s encode msg is not expected", test.TestName)
		}
	}

}

func TestEncodeLinkMsg(t *testing.T) {
	tests := []Test{
		{TestName: "Encode MergeRequest Event Success",
			Input: middlemsg.Body{
				EventType:    "merge_requset",
				Username:     "Administrator",
				Source:       "ms-viewport",
				Target:       "master",
				Project:      "Awesome Project",
				CommitNumber: "da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
				State:        "opened",
				MergeStatus:  "unchecked",
				CreatedAt:    time.Date(2013, 12, 03, 17, 15, 43, 0, time.UTC),
				UpdatedAt:    time.Date(2013, 12, 03, 17, 15, 43, 0, time.UTC),
				Assignee:     "User1",
			},
			Output: "Administrator opened the merge request\n" +
				"合入方向: ms-viewport -> master\n" +
				"last_commit_msg:\n" +
				"da1560886d: fixed readme\n" +
				"状态: unchecked\n" +
				"标题: MS-Viewport\n" +
				"项目：Awesome Project\n" +
				"代码审阅人：@User1\n" +
				"创建时间：2013-12-03T17:15:43Z\n" +
				"更新时间：2013-12-03T17:15:43Z\n" +
				"链接: ",
		},
	}
	for _, test := range tests {
		result, err := impl.EncodeLinkMsg(test.Input)
		if err != nil {
			t.Fatalf("TestName %s encode msg meet error +%v", test.TestName, err)
		}
		if result != test.Output {
			t.Fatalf("TestName %s encode msg is not expected", test.TestName)
		}
	}

}

func TestEncodeActionCardMsg(t *testing.T) {
	tests := []Test{
		{TestName: "Decode Push Event Success",
			Input: middlemsg.Body{
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
			Output: "John Smith pushed to branch master at repository Diaspora\ncommit: da1560886d4f094c3e6c9ef40349f7d38b5d27d7 \ncommit-message: {{ .CommitMsg}} \n项目：Diaspora",
		},
	}
	for _, test := range tests {
		result, err := impl.EncodeActionCardMsg(test.Input)
		if err != nil {
			t.Fatalf("TestName %s encode msg meet error +%v", test.TestName, err)
		}
		if result != test.Output {
			t.Fatalf("TestName %s encode msg is not expected", test.TestName)
		}
	}

}
