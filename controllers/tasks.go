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

// markAssigned marks a user's task as assigned
func markAssigned(c *TasksController, ctx *app.UpdateTasksContext) error {
	// TasksController_Assign: start_implement

	// Put your logic here

	// initialize new copy of TaskUser struct in variable taskUser
	taskUser := new(db.TaskUser)
	taskUser.AuthUserID = ctx.Payload.CognitoAuthUserID
	taskUser.Title = ctx.Payload.TaskName

	// initialize new copy of UserTask struct in variable userTask
	userTask := new(db.UserTask)

	// mark task assigned and pass AuthUserID to userTask struct
	userTask.Assigned = taskUser.Title
	userTask.AuthUserID = taskUser.AuthUserID

	_, err := c.db.MarkUserTaskAssigned(userTask)
	if err != nil {
		log.Errorf("[controller/tasks] %v", err)
		return ctx.BadRequest(&app.StandardError{
			Code:    400,
			Message: "could not assign task to user in db",
		})
	}

	log.Printf("[controller/tasks] assigned task for coindrop user: %v\n", userTask.AuthUserID)

	return nil
	// TasksController_Assign: end_implement
}

// markComplete marks a user's task as complete
func markCompleted(c *TasksController, ctx *app.UpdateTasksContext) error {
	// TasksController_Complete: start_implement

	// Put your logic here

	// initialize new copy of TaskUser struct in variable taskUser
	taskUser := new(db.TaskUser)
	taskUser.AuthUserID = ctx.Payload.CognitoAuthUserID
	taskUser.Title = ctx.Payload.TaskName

	// initialize new copy of UserTask struct in variable userTask
	userTask := new(db.UserTask)

	// mark task complete and pass AuthUserID to userTask struct
	userTask.Completed = taskUser.Title
	userTask.AuthUserID = taskUser.AuthUserID

	_, err := c.db.MarkUserTaskCompleted(userTask)
	if err != nil {
		log.Errorf("[controller/tasks] %v", err)
		return ctx.BadRequest(&app.StandardError{
			Code:    400,
			Message: "could not mark task complete in db",
		})
	}

	log.Printf("[controller/tasks] marked task complete for coindrop user: %v\n", userTask.AuthUserID)

	return nil
	// TasksController_Complete: end_implement
}

// Update runs the update action.
func (c *TasksController) Update(ctx *app.UpdateTasksContext) error {
	// TasksController_Update: start_implement

	// Put your logic here

	switch {
	case ctx.Payload.TaskState == "complete":
		err := markCompleted(c, ctx)
		if err != nil {
			return err
		}
	case ctx.Payload.TaskState == "assign":
		err := markAssigned(c, ctx)
		if err != nil {
			return err
		}
	case ctx.Payload.TaskState != "assign" && ctx.Payload.TaskState != "complete":
		log.Errorf("[controller/tasks] no task state update provided in payload")
		return ctx.BadRequest(&app.StandardError{
			Code:    400,
			Message: "no task state update provided in payload",
		})
	}

	return ctx.OK(nil)
	// TasksController_Update: end_implement
}

// Show runs the show action.
func (c *TasksController) Show(ctx *app.ShowTasksContext) error {
	// TasksController_Show: start_implement

	// Put your logic here

	userTask := new(db.UserTask)
	userTask.AuthUserID = ctx.Params.Get("userId")

	_, err := c.db.GetUserTasks(userTask)
	if err != nil {
		log.Errorf("[controller/tasks] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not get user's tasks from db",
		})
	}

	// initialize new variable tasks of Tasks struct
	tasks := new(db.Tasks)

	// get all tasks
	_, err = c.db.GetTasks(tasks)
	if err != nil {
		log.Errorf("[controller/tasks] %v", err)
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

	log.Printf("[controller/tasks] returned tasks for coindrop user: %v\n", userTask.AuthUserID)

	res := &app.Tasks{
		TaskList: userTasks.Tasks,
	}
	return ctx.OK(res)
	// TasksController_Show: end_implement
}