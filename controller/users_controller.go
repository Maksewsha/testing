package controller

import "net/http"

type UserController struct {
}

func (*UserController) AddUser(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("User has been added!"))
}
