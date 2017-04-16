//go:generate goagen app -d github.com/meier-christoph/todo-backend-golang-goa/design

package main

import (
	"context"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/meier-christoph/todo-backend-golang-goa/app"
	"net/http"
)

func main() {
	// Create service
	service := goa.New("todo")

	// Mount middleware
	// service.Use(CustomOriginMiddleware())
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "todos" controller
	c := NewTodosController(service)
	app.MountTodosController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}

// example on how to define a custom middleware e.g. setup origin headers manually
func CustomOriginMiddleware() goa.Middleware {
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			rw.Header().Set("Access-Control-Allow-Origin", "*")
			rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			rw.Header().Set("Access-Control-Allow-Methods", "OPTION, GET, POST, PATCH, DELETE")
			return h(ctx, rw, req)
		}
	}
}
