package main

import (
	"encoding/json"
	"errors"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/younggi/study_golang/discovery_go/chap6/taskman/task"
)

// FIXME: m is NOT thread-safe.
var m = task.NewInMemoryAccessor()

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

func apiGetHandler(w http.ResponseWriter, r *http.Request) {
	id := task.ID(mux.Vars(r)["id"])
	t, err := m.Get(id)
	err = json.NewEncoder(w).Encode(Response{
		ID:    id,
		Task:  t,
		Error: ResponseError{err},
	})
	if err != nil {
		log.Println(err)
	}
}

func apiPutHandler(w http.ResponseWriter, r *http.Request) {
	id := task.ID(mux.Vars(r)["id"])
	tasks, err := getTasks(r)
	if err != nil {
		log.Println(err)
		return
	}
	for _, t := range tasks {
		err = m.Put(id, t)
		err = json.NewEncoder(w).Encode(Response{
			ID:    id,
			Task:  t,
			Error: ResponseError{err},
		})
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func apiPostHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := getTasks(r)
	if err != nil {
		log.Println(err)
		return
	}
	for _, t := range tasks {
		id, err := m.Post(t)
		err = json.NewEncoder(w).Encode(Response{
			ID:    id,
			Task:  t,
			Error: ResponseError{err},
		})
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func apiDeleteHandler(w http.ResponseWriter, r *http.Request) {
	id := task.ID(mux.Vars(r)["id"])
	err := m.Delete(id)
	err = json.NewEncoder(w).Encode(Response{
		ID:    id,
		Error: ResponseError{err},
	})
	if err != nil {
		log.Println(err)
		return
	}
}

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
	err = tmpl.ExecuteTemplate(w, "task.html", &Response{
		ID:    id,
		Task:  t,
		Error: ResponseError{err},
	})
	if err != nil {
		log.Println(err)
		return
	}
}
