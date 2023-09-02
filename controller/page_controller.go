package controller

import "net/http"

type PageController struct {
}

func (*PageController) Home(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Welcome to home page!"))
}
