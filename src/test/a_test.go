package test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestHttpSent(t *testing.T) {
	fmt.Println("Hello, playground")
	websites := []string{"http://www.baidu.com", "http://www.163.com", "http://www.google.com"}
	client := new(http.Client)
	for _, web := range websites {
		req, err := http.NewRequest("GET", web, strings.NewReader(""))
		if err != nil {
			fmt.Println("make request error")
		}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println("send request to endpoint error")
		}
		fmt.Println(res)
	}
}
