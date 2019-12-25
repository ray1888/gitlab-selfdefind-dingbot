package dingding

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ray1888/self-defined-dingbot/src/messages/dingding/impl/link"

	//"fmt"
	"github.com/ray1888/self-defined-dingbot/src/messages/dingding/impl/text"
	"net/http"
	//"strings"

	"github.com/ray1888/self-defined-dingbot/src/clientsdk/response"
	//"github.com/ray1888/self-defined-dingbot/src/clientsdk/response/impl/dingding"
	//"github.com/ray1888/self-defined-dingbot/src/messages/dingding/impl/text"
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

func (c DingClient) SendLinkMsg(msg string) response.DingResponse {
	dingLinkMsg := link.Init(msg)
	data, err := json.Marshal(dingLinkMsg)
	req, err := http.NewRequest("POST", c.Endpoint, bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	res, err := c.Do(req)
	if err != nil {

	}
	fmt.Println(res)
	var Res response.DingResponse
	return Res
}

func (c DingClient) SendTextMsg(msg string) response.DingResponse {
	dingTextMsg := text.Init(msg)
	data, err := json.Marshal(*dingTextMsg)
	req, err := http.NewRequest("POST", c.Endpoint, bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	res, err := c.Do(req)
	if err != nil {

	}
	fmt.Println(res)
	var Res response.DingResponse
	return Res
}
