package controller

import (
	"fmt"
	"net/http"
)

func IndexPage(response http.Response, request *http.Request) {
	fmt.Println("indexpage")
}
