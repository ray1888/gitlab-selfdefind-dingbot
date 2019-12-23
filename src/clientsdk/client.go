package clientsdk

//import "net/http"

type DingResponse interface {
	IsSuccess()
	IsFailed()
}

type ClientSender interface {
	SendLinkMsg() DingResponse
	SendTextMsg() DingResponse
}
