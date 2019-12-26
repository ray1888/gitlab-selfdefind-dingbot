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

const (
	PREFIX = "gitlab"
	SUFFIX = "\n"
)

func Init(content string) *TextMsg {
	msg := new(TextMsg)
	msg.MsgType.MsgType = impl.TextType
	atPersonNames := getAtPersonInContent(content)
	if atPersonNames == "all" {
		msg.SetAtAll()
	} else if atPersonNames != "" {
		msg.SetAtMobilesByNameList(atPersonNames)
	}
	msg.MakeContent(content)
	return msg
}

func (t *TextMsg) MakeContent(content string) {
	// Add Prefix is to pass the dingbot sercurity authentication
	t.Text.Content = PREFIX + content + SUFFIX
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
	msgTypeLine := strings.Split(strings.TrimSpace(splitContent[len(splitContent)-1]), ":")
	msgEventType := strings.TrimSpace(msgTypeLine[len(msgTypeLine)-1])
	if msgEventType != "merge_request" {
		return ""
	}
	atUsersLine := strings.Split(splitContent[len(splitContent)-2], ":")
	atUser := strings.TrimSpace(atUsersLine[len(atUsersLine)-1])
	return atUser
}
