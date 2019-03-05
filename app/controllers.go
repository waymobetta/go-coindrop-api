// Code generated by goagen v1.4.1, DO NOT EDIT.
//
// API "coindrop": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/waymobetta/go-coindrop-api/design
// --out=$(GOPATH)/src/github.com/waymobetta/go-coindrop-api
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// HealthcheckController is the controller interface for the Healthcheck actions.
type HealthcheckController interface {
	goa.Muxer
	Show(*ShowHealthcheckContext) error
}

// MountHealthcheckController "mounts" a Healthcheck resource controller on the given service.
func MountHealthcheckController(service *goa.Service, ctrl HealthcheckController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/v1/health", ctrl.MuxHandler("preflight", handleHealthcheckOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowHealthcheckContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleHealthcheckOrigin(h)
	service.Mux.Handle("GET", "/v1/health", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Healthcheck", "action", "Show", "route", "GET /v1/health")
}

// handleHealthcheckOrigin applies the CORS response headers corresponding to the origin.
func handleHealthcheckOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "OPTIONS, HEAD, POST, GET, UPDATE, DELETE, PATCH")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// QuizzesController is the controller interface for the Quizzes actions.
type QuizzesController interface {
	goa.Muxer
	Create(*CreateQuizzesContext) error
	List(*ListQuizzesContext) error
	Show(*ShowQuizzesContext) error
}

// MountQuizzesController "mounts" a Quizzes resource controller on the given service.
func MountQuizzesController(service *goa.Service, ctrl QuizzesController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/v1/quizzes", ctrl.MuxHandler("preflight", handleQuizzesOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/v1/quizzes/:quizId", ctrl.MuxHandler("preflight", handleQuizzesOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateQuizzesContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*QuizPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleSecurity("JWTAuth", h)
	h = handleQuizzesOrigin(h)
	service.Mux.Handle("POST", "/v1/quizzes", ctrl.MuxHandler("create", h, unmarshalCreateQuizzesPayload))
	service.LogInfo("mount", "ctrl", "Quizzes", "action", "Create", "route", "POST /v1/quizzes", "security", "JWTAuth")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListQuizzesContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleSecurity("JWTAuth", h)
	h = handleQuizzesOrigin(h)
	service.Mux.Handle("GET", "/v1/quizzes", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Quizzes", "action", "List", "route", "GET /v1/quizzes", "security", "JWTAuth")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowQuizzesContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleSecurity("JWTAuth", h)
	h = handleQuizzesOrigin(h)
	service.Mux.Handle("GET", "/v1/quizzes/:quizId", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Quizzes", "action", "Show", "route", "GET /v1/quizzes/:quizId", "security", "JWTAuth")
}

// handleQuizzesOrigin applies the CORS response headers corresponding to the origin.
func handleQuizzesOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "OPTIONS, HEAD, POST, GET, UPDATE, DELETE, PATCH")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalCreateQuizzesPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateQuizzesPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &quizPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// RedditController is the controller interface for the Reddit actions.
type RedditController interface {
	goa.Muxer
	Create(*CreateRedditContext) error
	Show(*ShowRedditContext) error
}

// MountRedditController "mounts" a Reddit resource controller on the given service.
func MountRedditController(service *goa.Service, ctrl RedditController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/v1/social/reddit/userid", ctrl.MuxHandler("preflight", handleRedditOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateRedditContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateUserPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleSecurity("JWTAuth", h)
	h = handleRedditOrigin(h)
	service.Mux.Handle("POST", "/v1/social/reddit/userid", ctrl.MuxHandler("create", h, unmarshalCreateRedditPayload))
	service.LogInfo("mount", "ctrl", "Reddit", "action", "Create", "route", "POST /v1/social/reddit/userid", "security", "JWTAuth")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowRedditContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleSecurity("JWTAuth", h)
	h = handleRedditOrigin(h)
	service.Mux.Handle("GET", "/v1/social/reddit/userid", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Reddit", "action", "Show", "route", "GET /v1/social/reddit/userid", "security", "JWTAuth")
}

// handleRedditOrigin applies the CORS response headers corresponding to the origin.
func handleRedditOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "OPTIONS, HEAD, POST, GET, UPDATE, DELETE, PATCH")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalCreateRedditPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateRedditPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createUserPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// RedditharvestController is the controller interface for the Redditharvest actions.
type RedditharvestController interface {
	goa.Muxer
	Update(*UpdateRedditharvestContext) error
}

// MountRedditharvestController "mounts" a Redditharvest resource controller on the given service.
func MountRedditharvestController(service *goa.Service, ctrl RedditharvestController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/v1/social/reddit/harvest", ctrl.MuxHandler("preflight", handleRedditharvestOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateRedditharvestContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateUserPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	h = handleSecurity("JWTAuth", h)
	h = handleRedditharvestOrigin(h)
	service.Mux.Handle("POST", "/v1/social/reddit/harvest", ctrl.MuxHandler("update", h, unmarshalUpdateRedditharvestPayload))
	service.LogInfo("mount", "ctrl", "Redditharvest", "action", "Update", "route", "POST /v1/social/reddit/harvest", "security", "JWTAuth")
}

// handleRedditharvestOrigin applies the CORS response headers corresponding to the origin.
func handleRedditharvestOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "OPTIONS, HEAD, POST, GET, UPDATE, DELETE, PATCH")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalUpdateRedditharvestPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateRedditharvestPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &updateUserPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// ResultsController is the controller interface for the Results actions.
type ResultsController interface {
	goa.Muxer
	Show(*ShowResultsContext) error
}

// MountResultsController "mounts" a Results resource controller on the given service.
func MountResultsController(service *goa.Service, ctrl ResultsController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/v1/quiz/results", ctrl.MuxHandler("preflight", handleResultsOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowResultsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleSecurity("JWTAuth", h)
	h = handleResultsOrigin(h)
	service.Mux.Handle("GET", "/v1/quiz/results", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Results", "action", "Show", "route", "GET /v1/quiz/results", "security", "JWTAuth")
}

// handleResultsOrigin applies the CORS response headers corresponding to the origin.
func handleResultsOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "OPTIONS, HEAD, POST, GET, UPDATE, DELETE, PATCH")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// TasksController is the controller interface for the Tasks actions.
type TasksController interface {
	goa.Muxer
	Create(*CreateTasksContext) error
	List(*ListTasksContext) error
	Show(*ShowTasksContext) error
	Update(*UpdateTasksContext) error
}

// MountTasksController "mounts" a Tasks resource controller on the given service.
func MountTasksController(service *goa.Service, ctrl TasksController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/v1/tasks", ctrl.MuxHandler("preflight", handleTasksOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/v1/tasks/:taskId", ctrl.MuxHandler("preflight", handleTasksOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateTasksContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateTaskPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleSecurity("JWTAuth", h)
	h = handleTasksOrigin(h)
	service.Mux.Handle("POST", "/v1/tasks", ctrl.MuxHandler("create", h, unmarshalCreateTasksPayload))
	service.LogInfo("mount", "ctrl", "Tasks", "action", "Create", "route", "POST /v1/tasks", "security", "JWTAuth")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListTasksContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleSecurity("JWTAuth", h)
	h = handleTasksOrigin(h)
	service.Mux.Handle("GET", "/v1/tasks", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Tasks", "action", "List", "route", "GET /v1/tasks", "security", "JWTAuth")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowTasksContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleSecurity("JWTAuth", h)
	h = handleTasksOrigin(h)
	service.Mux.Handle("GET", "/v1/tasks/:taskId", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Tasks", "action", "Show", "route", "GET /v1/tasks/:taskId", "security", "JWTAuth")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateTasksContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*TaskPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	h = handleSecurity("JWTAuth", h)
	h = handleTasksOrigin(h)
	service.Mux.Handle("POST", "/v1/tasks/:taskId", ctrl.MuxHandler("update", h, unmarshalUpdateTasksPayload))
	service.LogInfo("mount", "ctrl", "Tasks", "action", "Update", "route", "POST /v1/tasks/:taskId", "security", "JWTAuth")
}

// handleTasksOrigin applies the CORS response headers corresponding to the origin.
func handleTasksOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "OPTIONS, HEAD, POST, GET, UPDATE, DELETE, PATCH")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalCreateTasksPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateTasksPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createTaskPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateTasksPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateTasksPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &taskPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// UsersController is the controller interface for the Users actions.
type UsersController interface {
	goa.Muxer
	Create(*CreateUsersContext) error
	List(*ListUsersContext) error
	Show(*ShowUsersContext) error
}

// MountUsersController "mounts" a Users resource controller on the given service.
func MountUsersController(service *goa.Service, ctrl UsersController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/v1/users", ctrl.MuxHandler("preflight", handleUsersOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/v1/users/:userId", ctrl.MuxHandler("preflight", handleUsersOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateUsersContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UserPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleUsersOrigin(h)
	service.Mux.Handle("POST", "/v1/users", ctrl.MuxHandler("create", h, unmarshalCreateUsersPayload))
	service.LogInfo("mount", "ctrl", "Users", "action", "Create", "route", "POST /v1/users")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListUsersContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleUsersOrigin(h)
	service.Mux.Handle("GET", "/v1/users", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Users", "action", "List", "route", "GET /v1/users")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowUsersContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleUsersOrigin(h)
	service.Mux.Handle("GET", "/v1/users/:userId", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Users", "action", "Show", "route", "GET /v1/users/:userId")
}

// handleUsersOrigin applies the CORS response headers corresponding to the origin.
func handleUsersOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "OPTIONS, HEAD, POST, GET, UPDATE, DELETE, PATCH")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalCreateUsersPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateUsersPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &userPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// VerifyredditController is the controller interface for the Verifyreddit actions.
type VerifyredditController interface {
	goa.Muxer
	Show(*ShowVerifyredditContext) error
	Update(*UpdateVerifyredditContext) error
}

// MountVerifyredditController "mounts" a Verifyreddit resource controller on the given service.
func MountVerifyredditController(service *goa.Service, ctrl VerifyredditController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/v1/social/reddit/userid/verify", ctrl.MuxHandler("preflight", handleVerifyredditOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowVerifyredditContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleSecurity("JWTAuth", h)
	h = handleVerifyredditOrigin(h)
	service.Mux.Handle("GET", "/v1/social/reddit/userid/verify", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Verifyreddit", "action", "Show", "route", "GET /v1/social/reddit/userid/verify", "security", "JWTAuth")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateVerifyredditContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*VerificationPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	h = handleSecurity("JWTAuth", h)
	h = handleVerifyredditOrigin(h)
	service.Mux.Handle("POST", "/v1/social/reddit/userid/verify", ctrl.MuxHandler("update", h, unmarshalUpdateVerifyredditPayload))
	service.LogInfo("mount", "ctrl", "Verifyreddit", "action", "Update", "route", "POST /v1/social/reddit/userid/verify", "security", "JWTAuth")
}

// handleVerifyredditOrigin applies the CORS response headers corresponding to the origin.
func handleVerifyredditOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "OPTIONS, HEAD, POST, GET, UPDATE, DELETE, PATCH")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalUpdateVerifyredditPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateVerifyredditPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &verificationPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// WalletsController is the controller interface for the Wallets actions.
type WalletsController interface {
	goa.Muxer
	Show(*ShowWalletsContext) error
	Update(*UpdateWalletsContext) error
}

// MountWalletsController "mounts" a Wallets resource controller on the given service.
func MountWalletsController(service *goa.Service, ctrl WalletsController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/v1/wallets", ctrl.MuxHandler("preflight", handleWalletsOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowWalletsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleSecurity("JWTAuth", h)
	h = handleWalletsOrigin(h)
	service.Mux.Handle("GET", "/v1/wallets", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Wallets", "action", "Show", "route", "GET /v1/wallets", "security", "JWTAuth")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateWalletsContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*WalletPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	h = handleSecurity("JWTAuth", h)
	h = handleWalletsOrigin(h)
	service.Mux.Handle("POST", "/v1/wallets", ctrl.MuxHandler("update", h, unmarshalUpdateWalletsPayload))
	service.LogInfo("mount", "ctrl", "Wallets", "action", "Update", "route", "POST /v1/wallets", "security", "JWTAuth")
}

// handleWalletsOrigin applies the CORS response headers corresponding to the origin.
func handleWalletsOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "OPTIONS, HEAD, POST, GET, UPDATE, DELETE, PATCH")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalUpdateWalletsPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateWalletsPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &walletPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}
