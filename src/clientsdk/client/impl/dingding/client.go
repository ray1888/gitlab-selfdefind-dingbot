package dingding

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ray1888/self-defined-dingbot/src/messages/dingding/impl/text"
)

type DingClient struct {
	http.Client
	Endpoint string
}

func Init() *DingClient {
	client := new(DingClient)
	client.Endpoint = "https://oapi.dingtalk.com/robot/send?access_token=f1b1b10623cc173845c5d197ec4aa1fc76bbaa8d64b9aaefd3ef614933646be0"
	return client
}

func (c DingClient) SendTextMsg(msg string) error {
	dingTextMsg := text.Init(msg)
	data, err := json.Marshal(*dingTextMsg)
	req, err := http.NewRequest("POST", c.Endpoint, bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	res, err := c.Do(req)
	if err != nil {
		// TODO log error
		return err
	}
	if res.StatusCode != 200 {
		// TODO log status code & msg
		return errors.New("Send Text Msg to DingBot is meetError")
	}
	return nil
}
