package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	httpHandlers *HTTPHandlers
}

func NewHTTPServer(httphandler *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		httpHandlers: httphandler,
	}
}

func (s *HTTPServer) StartServer() error {
	router := mux.NewRouter()

	router.Path("/tasks").Methods("POST").HandlerFunc(s.httpHandlers.HandleCreateTask)
	router.Path("/tasks/{title}").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetTask)
	router.Path("/tasks").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetAllTasks)
	router.Path("/tasks").Methods("GET").Queries("completed", "true").HandlerFunc(s.httpHandlers.HandleGetUncompletedTasks)
	router.Path("/tasks{title}").Methods("PATCH").HandlerFunc(s.httpHandlers.HandleMakeTaskCompleted)
	router.Path("/tasks{title}").Methods("DELETE").HandlerFunc(s.httpHandlers.handleDeleteTask)


	return http.ListenAndServe(":8081", router)
	} 