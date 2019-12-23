package dingding

import (
	"net/http"

	"github.com/ray1888/self-defined-dingbot/src/clientsdk/response"
	"github.com/ray1888/self-defined-dingbot/src/clientsdk/response/impl/dingding"
	"github.com/ray1888/self-defined-dingbot/src/messages/dingding/impl/link"
	"github.com/ray1888/self-defined-dingbot/src/messages/dingding/impl/text"
)

type DingClient struct {
	http.Client
	Endpoint string
}

func init() *DingClient {
	client := new(DingClient)
	client.Endpoint = "https://oapi.dingtalk.com/robot/send?access_token=xxxxx"
	return client
}

func (c DingClient) SendLinkMsg(msg *) response.DingResponse {
	req := http.Request{URL: c.Endpoint, Body:*(link.Init())}
	req, err := c.Do(&req)
	if err != nil{

	}
	return dingding.Response{}
}

func (c DingClient) SendTextMsg() response.DingResponse {
	req := http.Request{URL: c.Endpoint, Body:*(text.Init())}
	req, err := c.Do(&req)
	if err != nil{

	}
	return dingding.Response{}
}
