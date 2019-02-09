package controllers

import (
	"github.com/goadesign/goa"
	log "github.com/sirupsen/logrus"
	"github.com/waymobetta/go-coindrop-api/app"
	"github.com/waymobetta/go-coindrop-api/db"
)

// TaskController implements the task resource.
type TaskController struct {
	*goa.Controller
	db *db.DB
}

// NewTaskController creates a task controller.
func NewTaskController(service *goa.Service, db *db.DB) *TaskController {
	return &TaskController{
		Controller: service.NewController("TaskController"),
		db:         db,
	}
}

// Show runs the show action.
func (c *TaskController) Show(ctx *app.ShowTaskContext) error {
	// TaskController_Show: start_implement

	// Put your logic here

	userTask := new(db.UserTask)
	userTask.AuthUserID = ctx.Params.Get("userId")

	_, err := c.db.GetUserTasks(userTask)
	if err != nil {
		log.Errorf("[controller/task] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not get user's tasks from db",
		})
	}
	
	// initialize new variable tasks of Tasks struct
	tasks := new(db.Tasks)

	// get all tasks
	tasks, err = c.db.GetTasks(tasks)
	if err != nil {
		log.Errorf("[controller/task] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not get tasks from db",
		})
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

	log.Printf("[controller/task] returned task for coindrop user: %v\n", userTask.AuthUserID)

	// TODO:
	// need fix to return entire slice of struct objects instead of simply returning a string

	res := &app.Task{
		TaskName: userTasks.Tasks[0].Title,
	}
	return ctx.OK(res)
	// TaskController_Show: end_implement
}
