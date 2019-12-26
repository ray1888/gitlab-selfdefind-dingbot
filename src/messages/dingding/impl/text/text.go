package text

import (
	"github.com/ray1888/self-defined-dingbot/src/messages/dingding/impl"
	"strings"
)

type Text struct {
	Content string `json:"content"`
}

type At struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

type TextMsg struct {
	impl.MsgType
	Text `json:"text"`
	At   `json:"at"`
}

func Init(content string) *TextMsg {
	msg := new(TextMsg)
	msg.MsgType.MsgType = impl.TextType
	msg.Text.Content = content
	atPersonNames := getAtPersonInContent(content)
	if atPersonNames == "all" {
		msg.SetAtAll()
	} else {
		msg.SetAtMobilesByNameList(atPersonNames)
	}
	return msg
}

func (t *TextMsg) SetAtAll() {
	t.At.IsAtAll = true
}

func (t *TextMsg) SetAtMobilesByNameList(name string) {
	// This is limit by gitlab can't assign assignee more than two person
	t.At.AtMobiles = []string{name}
}

func getAtPersonInContent(content string) string {
	if content == "" {
		return content
	}
	splitContent := strings.Split(content, "\n")
	// TODO use const instead
	if splitContent[len(splitContent)-1] != "merge_request" {
		return ""
	}
	atUsersLine := strings.Split(splitContent[len(splitContent)-2], ":")
	atUser := strings.TrimSpace(atUsersLine[len(atUsersLine)-1])
	return atUser
}
