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

// TasksGet handles queries to return all stored tasks
func TasksGet(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	// initialize new variable tasks of Tasks struct
	tasks := new(db.Tasks)

	// return slice of structs of all task listings
	_, err := db.GetTasks(tasks)
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

	fmt.Printf("Successfully returned information for %d tasks\n\n", len(tasks.Tasks))
}

// TaskAdd adds a single task listing to db
func TaskAdd(w http.ResponseWriter, r *http.Request) {
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

	// add user listing to db
	_, err = db.AddTask(task)
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

	fmt.Printf("Successfully added task: %s\n\n", task.Title)
}

// UserTasksGet returns all the associated task information for the user, including assigned and completed tasks
func UserTasksGet(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})

	// initialize new copy of UserTask struct in variable userTask
	userTask := new(db.UserTask)

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

	// unmarshal bytes into task struct
	err = json.Unmarshal(body, userTask)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// get list of user's assigned and completed tasks
	_, err = db.GetUserTasks(userTask)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// initialize new variable tasks of Tasks struct
	tasks := new(db.Tasks)

	// get all tasks
	tasks, err = db.GetTasks(tasks)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	userTasks := new(db.Tasks)

	// TODO:
	// refactor to eliminate for loops if possible
	for task := range tasks.Tasks {
		for assignedTask := range userTask.ListData.AssignedTasks {
			if tasks.Tasks[task].Title == userTask.ListData.AssignedTasks[assignedTask] {
				tasks.Tasks[task].IsAssigned = true
				for completedTask := range userTask.ListData.CompletedTasks {
					if tasks.Tasks[task].Title == userTask.ListData.CompletedTasks[completedTask] {
						tasks.Tasks[task].IsCompleted = true
					}
				}
				userTasks.Tasks = append(userTasks.Tasks, tasks.Tasks[task])
			}
		}
	}

	response = utils.Message(true, userTasks)
	w.WriteHeader(http.StatusCreated)
	utils.Respond(w, response)

	fmt.Printf("Successfully returned task data for user: %s\n\n", userTask.AuthUserID)
}

// UserTaskAdd adds a single user listing to the db
func UserTaskAdd(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})

	// initialize new copy of UserTask struct in variable userTask
	userTask := new(db.UserTask)

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

	// unmarshal bytes into userTask struct
	err = json.Unmarshal(body, userTask)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// add user task listing in db
	// NOTE: userTask.TaskStatus.Assigned + "".Completed will both be empty maps
	_, err = db.AddUserTask(userTask)
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

	fmt.Printf("Successfully added task listing for user: %s\n\n", userTask.AuthUserID)
}

// UserTaskComplete adds a completed task to the existing list of completed tasks
func UserTaskComplete(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})

	// initialize new copy of TaskUser struct in variable taskUser
	taskUser := new(db.TaskUser)

	// initialize new copy of UserTask struct in variable userTask
	userTask := new(db.UserTask)

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

	// unmarshal bytes into task struct
	err = json.Unmarshal(body, taskUser)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	userTask.Completed = taskUser.Title
	userTask.AuthUserID = taskUser.AuthUserID

	// update user listing in db
	_, err = db.MarkUserTaskCompleted(userTask)
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

	fmt.Printf("Successfully added completed task: %s for user: %s\n\n", taskUser.Title, userTask.AuthUserID)
}

// UserTaskAssign adds an assigned task to a user's existing list of assigned tasks
func UserTaskAssign(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})

	// initialize new copy of TaskUser struct in variable taskUser
	taskUser := new(db.TaskUser)

	// initialize new copy of UserTask struct in variable userTask
	userTask := new(db.UserTask)

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

	// unmarshal bytes into task struct
	err = json.Unmarshal(body, taskUser)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	userTask.Assigned = taskUser.Title
	userTask.AuthUserID = taskUser.AuthUserID

	// update user listing in db
	_, err = db.MarkUserTaskAssigned(userTask)
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

	fmt.Printf("Successfully assigned task: %s to user: %s\n\n", taskUser.Title, userTask.AuthUserID)
}
