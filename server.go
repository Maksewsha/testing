package main

import (
	"net/http"

	controller "github.com/maksewsha/testingGo/controller"
)

type Server struct {
	host   string
	port   string
	server *http.Server
}

func (s *Server) addHandlers() *http.ServeMux {
	serveMux := &http.ServeMux{}
	controllers := controller.NewControllers()

	serveMux.HandleFunc("/", controllers.PageController.Home)
	serveMux.HandleFunc("/user", controllers.UserController.AddUser)

	return serveMux
}

func (s *Server) start() {
	serveMux := s.addHandlers()
	s.server = &http.Server{
		Handler: serveMux,
	}
	s.server.ListenAndServe()
}
