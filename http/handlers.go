package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"rest-api/todo"
	"time"

	"github.com/gorilla/mux"
)

type HTTPHandlers struct {
	todoList *todo.List
}

func newHTTPHandlers(todoList *todo.List) *HTTPHandlers {
	return &HTTPHandlers{
		todoList: todoList,
	}
}

/* pattern: /tasks
method: POST
info: json body request
succeed:
	-status code: 201 created
 	-response body: json represent created task

failed:
	-status code: 500, 501, etc.
	-response body: json err message + time
*/

func (h *HTTPHandlers) HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	var taskDTO TaskDTO
	if err := json.NewDecoder(r.Body).Decode(&taskDTO); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToJSON(), http.StatusBadRequest)
	}

	if err := taskDTO.ValidateForCreate(); err != nil {
		errDto := ErrorDTO {
			Message: err.Error(),
			Time: time.Now(),
		}

		http.Error(w, errDto.ToJSON(), http.StatusBadRequest)
	}

	todoTask := todo.NewTask(taskDTO.Title, taskDTO.Description)
	if err := h.todoList.AddTask(todoTask); err  != nil {
		errDto := ErrorDTO {
			Message: err.Error(),
			Time: time.Now(),
		}
		if errors.Is(err, todo.ErrTaskAlradyExists) {
			http.Error(w, errDto.ToJSON(), http.StatusConflict)
		} else {
			http.Error(w, errDto.ToJSON(), http.StatusInternalServerError)
		}

		return
	}

	b, err := json.MarshalIndent(todoTask, "", "	")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		log.Println(err)
		return
	}

}

/*
	pattern: /tasks/{title}
	method: GET
	info: pattern
	suceed:
		-status code 200(!OK)
		-response body: json body response(found task)
	failed:
		-status code: 404, 400, 500...
		-response body: json err message + time
*/

func (h *HTTPHandlers) HandleGetTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	h.todoList
}

/*
	pattern: /tasks
	method: GET
	info: -
	succeed:
		-status code 200(!OK)
		-response body: json body response(found task)
	failed:
		-status code: 400, 500...
		-response body: json err message + time
*/

func (h *HTTPHandlers) HandleGetAllTasks(w http.ResponseWriter, r *http.Request) {

}

/*
	pattern: /tasks?completed=true
	method: GET
	info: query-parametres
	succeed:
		-status code 200(!OK)
		-response body: json body response(found task)
	failed:
		-status code: 400, 500...
		-response body: json err message + time
*/

func (h *HTTPHandlers) HandleGetUncompletedTasks(w http.ResponseWriter, r *http.Request) {

}

/*
pattern: /tasks/{title}
method: PATCH
info: pattern + json body request
succeed:

	-status code 200(!OK)
	-response body: changed task

failed:

	-status code: 404, 400, 500...
	-response body: json err message + time
*/
func (h *HTTPHandlers) HandleMakeTaskCompleted(w http.ResponseWriter, r *http.Request) {

}

/*
pattern: /tasks/{title}
method: DELETE
info: pattern
succeed:

	-status code 204(NO-CONTENT)
	-response body: -

failed:

	-status code: 400, 500...
	-response body: json err message + time
*/
func (h *HTTPHandlers) handleDeleteTask(w http.ResponseWriter, r *http.Request) {

}
