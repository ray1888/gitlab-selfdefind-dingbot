package impl

import (
	"bytes"
	"errors"
	"text/template"

	"github.com/ray1888/self-defined-dingbot/src/messages/codeplatform"
	"github.com/ray1888/self-defined-dingbot/src/messages/middlemsg"
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
	var templateFilePath string
	switch templateName {
	case "text":
		templateFilePath = ""
	case "link":
		templateFilePath = ""
	case "actioncard":
		templateFilePath = ""
	default:
		return nil, errors.New("not support name")
	}
	tpl, err := template.ParseFiles(templateFilePath)
	if err != nil {
		return nil, err
	}
	return tpl, nil
}

func EncodeLinkMsg(body middlemsg.Body) string {
	return ""
}

func EncodeTextMsg(body middlemsg.Body) (string, error) {
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

func EncodeActionCardMsg(body middlemsg.Body) {

}
