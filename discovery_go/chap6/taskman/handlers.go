package main

import (
	"encoding/json"
	"errors"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/younggi/study_golang/discovery_go/chap6/taskman/task"
	"github.com/younggi/study_golang/discovery_go/chap6/taskman/task/mongodao"
)

// FIXME: m is NOT thread-safe.
//var m = task.NewInMemoryAccessor()
var m = mongodao.New("mongodb://localhost", "taskman", "tasks")

func getTasks(r *http.Request) ([]task.Task, error) {
	var result []task.Task
	if err := r.ParseForm(); err != nil {
		return nil, err
	}
	encodedTasks, ok := r.PostForm["task"]
	if !ok {
		return nil, errors.New("task parameter expected")
	}
	for _, encodedTask := range encodedTasks {
		var t task.Task
		if err := json.Unmarshal([]byte(encodedTask), &t); err != nil {
			return nil, err
		}
		result = append(result, t)
	}
	return result, nil
}

func getAccessorHelper(method string, r *http.Request, t task.Task) (task.ID, task.Task, error) {
	var err error
	switch method {
	case "GET":
		id := task.ID(mux.Vars(r)["id"])
		t, err = m.Get(id)
		return id, t, err
	case "DELETE":
		id := task.ID(mux.Vars(r)["id"])
		err := m.Delete(id)
		return id, task.Task{}, err
	case "PUT":
		id := task.ID(mux.Vars(r)["id"])
		err := m.Put(id, t)
		return id, t, err
	case "POST":
		id, err := m.Post(t)
		return id, t, err
	}
	return "", task.Task{}, nil
}

func writeResponseWrapper(w http.ResponseWriter) func(task.ID, task.Task, error) {
	return func(id task.ID, t task.Task, err error) {
		resp := NewResponse(id, t, err)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(resp.Error.Code)
	}
}

func apiGetAllHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := m.GetAll()
	if err != nil {
		log.Println(err)
	}
	resp := []*Response{}
	for _, t := range tasks {
		resp = append(resp, NewResponse("", t, err))
	}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Println(err)
	}
}

func apiHandler(method string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		writeResponseWrapper(w)(getAccessorHelper(method, r, task.Task{}))
	}
}

func apiMultiHandler(method string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tasks, err := getTasks(r)
		if err != nil {
			log.Println(err)
			return
		}
		for _, t := range tasks {
			writeResponseWrapper(w)(getAccessorHelper(method, r, t))
		}
	}
}

var apiGetHandler = apiHandler("GET")
var apiPutHandler = apiMultiHandler("PUT")
var apiPostHandler = apiMultiHandler("POST")
var apiDeleteHandler = apiHandler("DELETE")

var tmpl = template.Must(template.ParseGlob("html/*.html"))

// HTMLHandler handles web page
func htmlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Println(r.Method, "method is not supported")
		return
	}
	getID := func() (task.ID, error) {
		id := task.ID(r.URL.Path[len(htmlPathPrefix):])
		if id == "" {
			return id, errors.New("htmlHandler: ID is empty")
		}
		return id, nil
	}
	id, err := getID()
	if err != nil {
		log.Println(err)
		return
	}
	t, err := m.Get(id)
	resp := NewResponse(id, t, err)
	err = tmpl.ExecuteTemplate(w, "task.html", resp)
	if err != nil {
		log.Println(err)
		return
	}
}
