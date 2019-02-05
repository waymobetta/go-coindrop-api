package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/waymobetta/go-coindrop-api/db"
	"github.com/waymobetta/go-coindrop-api/utils"
)

// TODO:
// create 2 table schemas:
// 1. to store actions (linked to tasks table)
// 2. to store results of action tasks

// ActionGet handles queries to return a specific action task
func (h *Handlers) ActionGet(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	// initialize new variable tasks of Tasks struct
	tasks := new(db.Tasks)

	// return slice of structs of all task listings
	_, err := h.db.GetTasks(tasks)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// marshall into JSON
	_, err = json.Marshal(tasks)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	response = utils.Message(true, tasks)
	w.WriteHeader(http.StatusOK)
	utils.Respond(w, response)

	fmt.Printf("[db] successfully returned information for %d tasks\n\n", len(tasks.Tasks))
}

// ActionAdd adds a single action task listing to db
func (h *Handlers) ActionAdd(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	// initialize new variable user of User struct
	task := new(db.Task)

	// add limit for large payload protection
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}
	defer r.Body.Close()

	// unmarshal bytes into user struct
	err = json.Unmarshal(body, task)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// add task action listing to db
	_, err = h.db.AddTask(task)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	response = utils.Message(true, "success")
	w.WriteHeader(http.StatusCreated)
	utils.Respond(w, response)

	fmt.Printf("[db] successfully added task: %s\n\n", task.Title)
}
