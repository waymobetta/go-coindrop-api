package controllers

import (
	"github.com/goadesign/goa"
	log "github.com/sirupsen/logrus"
	"github.com/waymobetta/go-coindrop-api/app"
	"github.com/waymobetta/go-coindrop-api/db"
	"github.com/waymobetta/go-coindrop-api/types"
)

// TasksController implements the tasks resource.
type TasksController struct {
	*goa.Controller
	db *db.DB
}

// NewTasksController creates a tasks controller.
func NewTasksController(service *goa.Service, dbs *db.DB) *TasksController {
	return &TasksController{
		Controller: service.NewController("TasksController"),
		db:         dbs,
	}
}

// Show runs the show action.
func (c *TasksController) Show(ctx *app.ShowTasksContext) error {
	// TasksController_Show: start_implement

	// Put your logic here
	userID := ctx.Params.Get("userId")
	// Note: if query string `userId` is empty,
	// then use user ID from auth token
	if userID == "" {
		userID = ctx.Value("authUserID").(string)
	}
	taskUser := new(types.TaskUser)
	taskUser.UserID = userID
	taskUser.TaskID = ctx.TaskID

	task, err := c.db.GetUserTask(taskUser)
	if err != nil {
		log.Errorf("[controller/tasks] failed to get user task: %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not get user task",
		})
	}

	log.Printf("[controller/tasks] returned tasks for coindrop user: %v\n", taskUser.UserID)

	return ctx.OK(&app.Task{
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
		LogoURL: task.LogoURL,
	})
	// TasksController_Show: end_implement
}

// List runs the list action.
func (c *TasksController) List(ctx *app.ListTasksContext) error {
	// TasksController_List: start_implement

	// Put your logic here

	userID := ctx.Params.Get("userId")
	// Note: if query string `userId` is empty,
	// then use user ID from auth token
	if userID == "" {
		userID = ctx.Value("authUserID").(string)
	}

	taskUser := new(types.TaskUser)
	taskUser.UserID = userID

	tasks, err := c.db.GetUserTasks(taskUser)
	if err != nil {
		log.Errorf("[controller/tasks] failed to get user tasks: %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not get user's tasks",
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
			LogoURL: task.LogoURL,
		})
	}

	log.Printf("[controller/tasks] returned tasks for coindrop user: %v\n", taskUser.UserID)

	return ctx.OK(&app.Tasks{
		Tasks: t,
	})

	// TasksController_List: end_implement
}

// Create runs the create action.
func (c *TasksController) Create(ctx *app.CreateTasksContext) error {
	// TasksController_Create: start_implement

	// Put your logic here

	userTask := &types.UserTask{
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

	// initialize new copy of TaskUser struct in variable taskUser
	userID := ctx.Value("authUserID").(string)
	taskUser := new(types.TaskUser2)
	taskUser.UserID = userID
	taskUser.TaskID = ctx.TaskID

	// initialize new copy of UserTask struct in variable userTask
	userTask := new(types.UserTask)

	// mark task complete and pass UserID to userTask struct
	userTask.Completed = ctx.Payload.Completed
	userTask.UserID = taskUser.UserID
	userTask.TaskID = taskUser.TaskID

	_, err := c.db.MarkUserTaskCompleted(userTask)
	if err != nil {
		log.Errorf("[controller/tasks] error updating task, %v", err)
		return ctx.BadRequest(&app.StandardError{
			Code:    400,
			Message: "could not mark task complete in db",
		})
	}

	log.Printf("[controller/tasks] marked task complete for coindrop user: %v\n", userTask.UserID)

	return ctx.OK(nil)
	// TasksController_Update: end_implement
}
