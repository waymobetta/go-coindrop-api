package controllers

import (
	"github.com/goadesign/goa"
	log "github.com/sirupsen/logrus"
	"github.com/waymobetta/go-coindrop-api/app"
	"github.com/waymobetta/go-coindrop-api/db"
)

// TasksController implements the tasks resource.
type TasksController struct {
	*goa.Controller
	db *db.DB
}

// NewTasksController creates a tasks controller.
func NewTasksController(service *goa.Service, db *db.DB) *TasksController {
	return &TasksController{
		Controller: service.NewController("TasksController"),
		db:         db,
	}
}

// Show runs the show action.
func (c *TasksController) Show(ctx *app.ShowTasksContext) error {
	// TasksController_Show: start_implement

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

	res := &app.Tasks{
		TaskList: userTasks.Tasks,
	}
	return ctx.OK(res)
	// TasksController_Show: end_implement
}
