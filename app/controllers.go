// Code generated by goagen v1.3.1, DO NOT EDIT.
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

// TasksController is the controller interface for the Tasks actions.
type TasksController interface {
	goa.Muxer
	Show(*ShowTasksContext) error
}

// MountTasksController "mounts" a Tasks resource controller on the given service.
func MountTasksController(service *goa.Service, ctrl TasksController) {
	initService(service)
	var h goa.Handler

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
	service.Mux.Handle("GET", "/v1/tasks", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Tasks", "action", "Show", "route", "GET /v1/tasks")
}

// UserController is the controller interface for the User actions.
type UserController interface {
	goa.Muxer
	Create(*CreateUserContext) error
	Show(*ShowUserContext) error
}

// MountUserController "mounts" a User resource controller on the given service.
func MountUserController(service *goa.Service, ctrl UserController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateUserContext(ctx, req, service)
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
	service.Mux.Handle("POST", "/v1/users", ctrl.MuxHandler("create", h, unmarshalCreateUserPayload))
	service.LogInfo("mount", "ctrl", "User", "action", "Create", "route", "POST /v1/users")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowUserContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/v1/users/:userId", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "Show", "route", "GET /v1/users/:userId")
}

// unmarshalCreateUserPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateUserPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
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

// WalletController is the controller interface for the Wallet actions.
type WalletController interface {
	goa.Muxer
	Show(*ShowWalletContext) error
	Update(*UpdateWalletContext) error
}

// MountWalletController "mounts" a Wallet resource controller on the given service.
func MountWalletController(service *goa.Service, ctrl WalletController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowWalletContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/v1/wallets", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Wallet", "action", "Show", "route", "GET /v1/wallets")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateWalletContext(ctx, req, service)
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
	service.Mux.Handle("POST", "/v1/wallets", ctrl.MuxHandler("update", h, unmarshalUpdateWalletPayload))
	service.LogInfo("mount", "ctrl", "Wallet", "action", "Update", "route", "POST /v1/wallets")
}

// unmarshalUpdateWalletPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateWalletPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
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
