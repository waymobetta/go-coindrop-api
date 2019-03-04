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

	taskUser := new(db.TaskUser)
	taskUser.AuthUserID = ctx.Params.Get("userId")

	tasks, err := c.db.GetUserTasks(taskUser)
	if err != nil {
		log.Errorf("[controller/tasks] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not get user's tasks from db",
		})
	}

	var t app.TaskCollection

	for _, task := range tasks {
		t = append(t, &app.Task{
			ID:              task.ID,
			Title:           task.Title,
			Type:            task.Type,
			Author:          task.Author,
			Description:     task.Description,
			Token:           task.Token,
			TokenAllocation: task.TokenAllocation,
			Badge: &app.Badge{
				ID:          task.BadgeData.ID,
				Name:        task.BadgeData.Name,
				Description: task.BadgeData.Description,
				Recipients:  task.BadgeData.Recipients,
			},
		})
	}

	log.Printf("[controller/tasks] returned tasks for coindrop user: %v\n", taskUser.AuthUserID)

	res := &app.Tasks{
		Tasks: t,
	}
	return ctx.OK(res)
	// TasksController_Show: end_implement
}

// Create runs the create action.
func (c *TasksController) Create(ctx *app.CreateTasksContext) error {
	// TasksController_Create: start_implement

	// Put your logic here

	userTask := &db.UserTask2{
		UserID: ctx.Payload.UserID,
		TaskID: ctx.Payload.TaskID,
	}
	_, err := c.db.AddUserTask(userTask)
	if err != nil {
		log.Errorf("[controller/tasks] failed to create task; %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    400,
			Message: "Error creating task",
		})
	}

	return ctx.OK(nil)
	// TasksController_Create: end_implement
}

// Update runs the update action.
func (c *TasksController) Update(ctx *app.UpdateTasksContext) error {
	// TasksController_Update: start_implement

	// Put your logic here

	err := markCompleted(c, ctx)
	if err != nil {
		log.Errorf("[controller/tasks] failed to update task; %v", err)
		return ctx.BadRequest(&app.StandardError{
			Code:    400,
			Message: "error updating task",
		})
	}

	return ctx.OK(nil)
	// TasksController_Update: end_implement
}

// markComplete marks a user's task as complete
func markCompleted(c *TasksController, ctx *app.UpdateTasksContext) error {
	// TasksController_Complete: start_implement

	// Put your logic here

	// initialize new copy of TaskUser struct in variable taskUser
	cognitoUserID := ctx.Value("cognitoUserID").(string)
	taskUser := new(db.TaskUser2)
	taskUser.UserID = cognitoUserID
	taskUser.TaskID = ctx.TaskID

	// initialize new copy of UserTask struct in variable userTask
	userTask := new(db.UserTask2)

	// mark task complete and pass AuthUserID to userTask struct
	userTask.Completed = ctx.Payload.Completed
	userTask.UserID = taskUser.UserID

	_, err := c.db.MarkUserTaskCompleted(userTask)
	if err != nil {
		log.Errorf("[controller/tasks] %v", err)
		return ctx.BadRequest(&app.StandardError{
			Code:    400,
			Message: "could not mark task complete in db",
		})
	}

	log.Printf("[controller/tasks] marked task complete for coindrop user: %v\n", userTask.UserID)

	return nil
	// TasksController_Complete: end_implement
}
