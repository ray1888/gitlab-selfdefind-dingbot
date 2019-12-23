package client

import (
	"github.com/ray1888/self-defined-dingbot/src/clientsdk/response"
)

type ClientSender interface {
	SendLinkMsg() response.DingResponse
	SendTextMsg() response.DingResponse
}
