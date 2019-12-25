package impl

import (
	"bytes"
	"errors"
	"strings"
	"text/template"

	"github.com/ray1888/self-defined-dingbot/src/messages/codeplatform"
	"github.com/ray1888/self-defined-dingbot/src/messages/middlemsg"
	selfDefindTemplate "github.com/ray1888/self-defined-dingbot/src/template"
)

type Encoder struct {
	InnerChannel  chan (middlemsg.Body)
	OutputChannel chan (string)
}

func Init(inChan chan (middlemsg.Body), outChan chan (string)) *Encoder {
	encoder := new(Encoder)
	encoder.InnerChannel = inChan
	encoder.OutputChannel = outChan
	return encoder
}

func (e *Encoder) encode() {
	for msg := range e.InnerChannel {
		switch msg.EventType {
		case codeplatform.MergeRequest:
			// Default use Link msg

		case codeplatform.Push:
			// Default use Text Msg
		case codeplatform.Tag:
			// Default use ActionCard Msg
		}
	}
}

func loadTemplate(templateName string) (*template.Template, error) {
	switch templateName {
	case "text":
		return selfDefindTemplate.TextTemplate, nil
	case "link":
		return selfDefindTemplate.LinkMergeRequestTemplate, nil
	case "actioncard":
		return selfDefindTemplate.LinkMergeRequestTemplate, nil
	default:
		return nil, errors.New("not support name")
	}
}

func EncodeLinkMsg(body middlemsg.Body) (string, error) {
	tpl, err := loadTemplate("link")
	if err != nil {
		return "", err
	}
	var content bytes.Buffer
	err = tpl.Execute(&content, body)
	contentString := content.String()
	if err != nil {
		return "", err
	}
	return contentString, nil
}

func EncodeTextMsg(body middlemsg.Body) (string, error) {
	tpl, err := loadTemplate("text")
	if err != nil {
		return "", err
	}
	body.CommitsText = parseCommitsToText(body.Commits)
	var content bytes.Buffer
	err = tpl.Execute(&content, body)
	contentString := content.String()
	if err != nil {
		return "", err
	}
	return contentString, nil
}

func EncodeActionCardMsg(body middlemsg.Body) (string, error) {
	tpl, err := loadTemplate("text")
	if err != nil {
		return "", err
	}
	var content bytes.Buffer
	err = tpl.Execute(&content, body)
	contentString := content.String()
	if err != nil {
		return "", err
	}
	return contentString, nil
}

func parseCommit(commit middlemsg.Commit) string {
	return commit.Number[:9] + ": " + commit.Info + "\n"
}

func parseCommitsToText(commits []middlemsg.Commit) string {
	result := ""
	for _, commit := range commits {
		workCommit := commit
		result += parseCommit(workCommit)
	}
	result = strings.TrimSpace(result)
	return result
}
