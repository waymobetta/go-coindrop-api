// Code generated by goagen v1.4.1, DO NOT EDIT.
//
// API "coindrop": CLI Commands
//
// Command:
// $ goagen
// --design=github.com/waymobetta/go-coindrop-api/design
// --out=$(GOPATH)/src/github.com/waymobetta/go-coindrop-api
// --version=v1.3.1

package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	uuid "github.com/goadesign/goa/uuid"
	"github.com/spf13/cobra"
	"github.com/waymobetta/go-coindrop-api/client"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type (
	// ShowHealthcheckCommand is the command line data structure for the show action of healthcheck
	ShowHealthcheckCommand struct {
		PrettyPrint bool
	}

	// CreateQuizzesCommand is the command line data structure for the create action of quizzes
	CreateQuizzesCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// ListQuizzesCommand is the command line data structure for the list action of quizzes
	ListQuizzesCommand struct {
		PrettyPrint bool
	}

	// ShowQuizzesCommand is the command line data structure for the show action of quizzes
	ShowQuizzesCommand struct {
		// Quiz ID
		QuizID      string
		PrettyPrint bool
	}

	// CreateRedditCommand is the command line data structure for the create action of reddit
	CreateRedditCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// ShowRedditCommand is the command line data structure for the show action of reddit
	ShowRedditCommand struct {
		// User ID
		UserID      string
		PrettyPrint bool
	}

	// UpdateRedditharvestCommand is the command line data structure for the update action of redditharvest
	UpdateRedditharvestCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// ShowResultsCommand is the command line data structure for the show action of results
	ShowResultsCommand struct {
		// User ID
		UserID      string
		PrettyPrint bool
	}

	// CreateTasksCommand is the command line data structure for the create action of tasks
	CreateTasksCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// ListTasksCommand is the command line data structure for the list action of tasks
	ListTasksCommand struct {
		// User ID
		UserID      string
		PrettyPrint bool
	}

	// ShowTasksCommand is the command line data structure for the show action of tasks
	ShowTasksCommand struct {
		// Task ID
		TaskID      string
		PrettyPrint bool
	}

	// UpdateTasksCommand is the command line data structure for the update action of tasks
	UpdateTasksCommand struct {
		Payload     string
		ContentType string
		// Task ID
		TaskID      string
		PrettyPrint bool
	}

	// CreateUsersCommand is the command line data structure for the create action of users
	CreateUsersCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// ListUsersCommand is the command line data structure for the list action of users
	ListUsersCommand struct {
		// Cognito Auth User ID
		CognitoAuthUserID string
		PrettyPrint       bool
	}

	// ShowUsersCommand is the command line data structure for the show action of users
	ShowUsersCommand struct {
		// User ID
		UserID      string
		PrettyPrint bool
	}

	// ShowVerifyredditCommand is the command line data structure for the show action of verifyreddit
	ShowVerifyredditCommand struct {
		// User ID
		UserID      string
		PrettyPrint bool
	}

	// UpdateVerifyredditCommand is the command line data structure for the update action of verifyreddit
	UpdateVerifyredditCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// ShowWalletsCommand is the command line data structure for the show action of wallets
	ShowWalletsCommand struct {
		// User ID
		UserID      string
		PrettyPrint bool
	}

	// UpdateWalletsCommand is the command line data structure for the update action of wallets
	UpdateWalletsCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}
)

// RegisterCommands registers the resource action CLI commands.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "create",
		Short: `create action`,
	}
	tmp1 := new(CreateQuizzesCommand)
	sub = &cobra.Command{
		Use:   `quizzes ["/v1/quizzes"]`,
		Short: ``,
		Long: `

Payload example:

{
   "title": "Aut qui nisi dolore ipsa."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp1.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp2 := new(CreateRedditCommand)
	sub = &cobra.Command{
		Use:   `reddit ["/v1/social/reddit/userid"]`,
		Short: ``,
		Long: `

Payload example:

{
   "userId": "9302608f-f6a4-4004-b088-63e5fb43cc26",
   "username": "Ea porro sunt facilis alias."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp2.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp3 := new(CreateTasksCommand)
	sub = &cobra.Command{
		Use:   `tasks ["/v1/tasks"]`,
		Short: ``,
		Long: `

Payload example:

{
   "taskId": "Natus neque pariatur reprehenderit officiis eos.",
   "userId": "Sit amet non voluptatem id."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp3.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp4 := new(CreateUsersCommand)
	sub = &cobra.Command{
		Use:   `users ["/v1/users"]`,
		Short: ``,
		Long: `

Payload example:

{
   "cognitoAuthUserId": "Et molestias sed possimus enim et."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp4.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "list",
		Short: `list action`,
	}
	tmp5 := new(ListQuizzesCommand)
	sub = &cobra.Command{
		Use:   `quizzes ["/v1/quizzes"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp5.Run(c, args) },
	}
	tmp5.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp5.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp6 := new(ListTasksCommand)
	sub = &cobra.Command{
		Use:   `tasks ["/v1/tasks"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp6.Run(c, args) },
	}
	tmp6.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp6.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp7 := new(ListUsersCommand)
	sub = &cobra.Command{
		Use:   `users ["/v1/users"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp7.Run(c, args) },
	}
	tmp7.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp7.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "show",
		Short: `show action`,
	}
	tmp8 := new(ShowHealthcheckCommand)
	sub = &cobra.Command{
		Use:   `healthcheck ["/v1/health"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp8.Run(c, args) },
	}
	tmp8.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp8.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp9 := new(ShowQuizzesCommand)
	sub = &cobra.Command{
		Use:   `quizzes ["/v1/quizzes/QUIZID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp9.Run(c, args) },
	}
	tmp9.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp9.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp10 := new(ShowRedditCommand)
	sub = &cobra.Command{
		Use:   `reddit ["/v1/social/reddit/userid"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp10.Run(c, args) },
	}
	tmp10.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp10.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp11 := new(ShowResultsCommand)
	sub = &cobra.Command{
		Use:   `results ["/v1/quiz/results"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp11.Run(c, args) },
	}
	tmp11.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp11.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp12 := new(ShowTasksCommand)
	sub = &cobra.Command{
		Use:   `tasks ["/v1/tasks/TASKID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp12.Run(c, args) },
	}
	tmp12.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp12.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp13 := new(ShowUsersCommand)
	sub = &cobra.Command{
		Use:   `users ["/v1/users/USERID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp13.Run(c, args) },
	}
	tmp13.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp13.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp14 := new(ShowVerifyredditCommand)
	sub = &cobra.Command{
		Use:   `verifyreddit ["/v1/social/reddit/userid/verify"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp14.Run(c, args) },
	}
	tmp14.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp14.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp15 := new(ShowWalletsCommand)
	sub = &cobra.Command{
		Use:   `wallets ["/v1/wallets"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp15.Run(c, args) },
	}
	tmp15.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp15.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "update",
		Short: `update action`,
	}
	tmp16 := new(UpdateRedditharvestCommand)
	sub = &cobra.Command{
		Use:   `redditharvest ["/v1/social/reddit/harvest"]`,
		Short: ``,
		Long: `

Payload example:

{
   "userId": "9302608f-f6a4-4004-b088-63e5fb43cc26",
   "username": "Qui rerum totam quia amet reiciendis velit."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp16.Run(c, args) },
	}
	tmp16.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp16.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp17 := new(UpdateTasksCommand)
	sub = &cobra.Command{
		Use:   `tasks ["/v1/tasks/TASKID"]`,
		Short: ``,
		Long: `

Payload example:

{
   "completed": true
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp17.Run(c, args) },
	}
	tmp17.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp17.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp18 := new(UpdateVerifyredditCommand)
	sub = &cobra.Command{
		Use:   `verifyreddit ["/v1/social/reddit/userid/verify"]`,
		Short: ``,
		Long: `

Payload example:

{
   "postedVerificationCode": "Non quia aliquam voluptatem quibusdam ut.",
   "userId": "9302608f-f6a4-4004-b088-63e5fb43cc26"
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp18.Run(c, args) },
	}
	tmp18.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp18.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp19 := new(UpdateWalletsCommand)
	sub = &cobra.Command{
		Use:   `wallets ["/v1/wallets"]`,
		Short: ``,
		Long: `

Payload example:

{
   "walletAddress": "0x845fdD93Cca3aE9e380d5556818e6d0b902B977c"
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp19.Run(c, args) },
	}
	tmp19.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp19.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
}

func intFlagVal(name string, parsed int) *int {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func float64FlagVal(name string, parsed float64) *float64 {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func boolFlagVal(name string, parsed bool) *bool {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func stringFlagVal(name string, parsed string) *string {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func hasFlag(name string) bool {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--"+name) {
			return true
		}
	}
	return false
}

func jsonVal(val string) (*interface{}, error) {
	var t interface{}
	err := json.Unmarshal([]byte(val), &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func jsonArray(ins []string) ([]interface{}, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []interface{}
	for _, id := range ins {
		val, err := jsonVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)
	}
	return vals, nil
}

func timeVal(val string) (*time.Time, error) {
	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func timeArray(ins []string) ([]time.Time, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []time.Time
	for _, id := range ins {
		val, err := timeVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func uuidVal(val string) (*uuid.UUID, error) {
	t, err := uuid.FromString(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func uuidArray(ins []string) ([]uuid.UUID, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []uuid.UUID
	for _, id := range ins {
		val, err := uuidVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func float64Val(val string) (*float64, error) {
	t, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func float64Array(ins []string) ([]float64, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []float64
	for _, id := range ins {
		val, err := float64Val(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func boolVal(val string) (*bool, error) {
	t, err := strconv.ParseBool(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func boolArray(ins []string) ([]bool, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []bool
	for _, id := range ins {
		val, err := boolVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

// Run makes the HTTP request corresponding to the ShowHealthcheckCommand command.
func (cmd *ShowHealthcheckCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/health"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowHealthcheck(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowHealthcheckCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the CreateQuizzesCommand command.
func (cmd *CreateQuizzesCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/quizzes"
	}
	var payload client.QuizPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateQuizzes(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateQuizzesCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the ListQuizzesCommand command.
func (cmd *ListQuizzesCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/quizzes"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListQuizzes(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListQuizzesCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the ShowQuizzesCommand command.
func (cmd *ShowQuizzesCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/quizzes/%v", url.QueryEscape(cmd.QuizID))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowQuizzes(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowQuizzesCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var quizID string
	cc.Flags().StringVar(&cmd.QuizID, "quizId", quizID, `Quiz ID`)
}

// Run makes the HTTP request corresponding to the CreateRedditCommand command.
func (cmd *CreateRedditCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/social/reddit/userid"
	}
	var payload client.CreateUserPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateReddit(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateRedditCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the ShowRedditCommand command.
func (cmd *ShowRedditCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/social/reddit/userid"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowReddit(ctx, path, stringFlagVal("userId", cmd.UserID))
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowRedditCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var userID string
	cc.Flags().StringVar(&cmd.UserID, "userId", userID, `User ID`)
}

// Run makes the HTTP request corresponding to the UpdateRedditharvestCommand command.
func (cmd *UpdateRedditharvestCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/social/reddit/harvest"
	}
	var payload client.UpdateUserPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdateRedditharvest(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateRedditharvestCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the ShowResultsCommand command.
func (cmd *ShowResultsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/quiz/results"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowResults(ctx, path, stringFlagVal("userId", cmd.UserID))
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowResultsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var userID string
	cc.Flags().StringVar(&cmd.UserID, "userId", userID, `User ID`)
}

// Run makes the HTTP request corresponding to the CreateTasksCommand command.
func (cmd *CreateTasksCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/tasks"
	}
	var payload client.CreateTaskPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateTasks(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateTasksCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the ListTasksCommand command.
func (cmd *ListTasksCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/tasks"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListTasks(ctx, path, stringFlagVal("userId", cmd.UserID))
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListTasksCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var userID string
	cc.Flags().StringVar(&cmd.UserID, "userId", userID, `User ID`)
}

// Run makes the HTTP request corresponding to the ShowTasksCommand command.
func (cmd *ShowTasksCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/tasks/%v", url.QueryEscape(cmd.TaskID))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowTasks(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowTasksCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var taskID string
	cc.Flags().StringVar(&cmd.TaskID, "taskId", taskID, `Task ID`)
}

// Run makes the HTTP request corresponding to the UpdateTasksCommand command.
func (cmd *UpdateTasksCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/tasks/%v", url.QueryEscape(cmd.TaskID))
	}
	var payload client.TaskPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdateTasks(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateTasksCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var taskID string
	cc.Flags().StringVar(&cmd.TaskID, "taskId", taskID, `Task ID`)
}

// Run makes the HTTP request corresponding to the CreateUsersCommand command.
func (cmd *CreateUsersCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/users"
	}
	var payload client.UserPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateUsers(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateUsersCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the ListUsersCommand command.
func (cmd *ListUsersCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/users"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListUsers(ctx, path, stringFlagVal("cognitoAuthUserId", cmd.CognitoAuthUserID))
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListUsersCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var cognitoAuthUserID string
	cc.Flags().StringVar(&cmd.CognitoAuthUserID, "cognitoAuthUserId", cognitoAuthUserID, `Cognito Auth User ID`)
}

// Run makes the HTTP request corresponding to the ShowUsersCommand command.
func (cmd *ShowUsersCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/users/%v", url.QueryEscape(cmd.UserID))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowUsers(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowUsersCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var userID string
	cc.Flags().StringVar(&cmd.UserID, "userId", userID, `User ID`)
}

// Run makes the HTTP request corresponding to the ShowVerifyredditCommand command.
func (cmd *ShowVerifyredditCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/social/reddit/userid/verify"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowVerifyreddit(ctx, path, stringFlagVal("userId", cmd.UserID))
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowVerifyredditCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var userID string
	cc.Flags().StringVar(&cmd.UserID, "userId", userID, `User ID`)
}

// Run makes the HTTP request corresponding to the UpdateVerifyredditCommand command.
func (cmd *UpdateVerifyredditCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/social/reddit/userid/verify"
	}
	var payload client.VerificationPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdateVerifyreddit(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateVerifyredditCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the ShowWalletsCommand command.
func (cmd *ShowWalletsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/wallets"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowWallets(ctx, path, stringFlagVal("userId", cmd.UserID))
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowWalletsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var userID string
	cc.Flags().StringVar(&cmd.UserID, "userId", userID, `User ID`)
}

// Run makes the HTTP request corresponding to the UpdateWalletsCommand command.
func (cmd *UpdateWalletsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/wallets"
	}
	var payload client.WalletPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdateWallets(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateWalletsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}
