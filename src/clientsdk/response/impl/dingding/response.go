package dingding

import (
	"net/http"
)

type Response http.Response

func (r Response) IsSuccess() bool {
	if r.StatusCode != 200 {
		return false
	}
	return true
}

func (r Response) Failed() {

}
