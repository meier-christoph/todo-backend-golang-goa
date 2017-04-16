package main

import (
	"github.com/goadesign/goa"
	"github.com/meier-christoph/todo-backend-golang-goa/app"
)

// TodosController implements the todos resource.
type TodosController struct {
	*goa.Controller
}

// NewTodosController creates a todos controller.
func NewTodosController(service *goa.Service) *TodosController {
	return &TodosController{Controller: service.NewController("TodosController")}
}

// Create runs the create action.
func (c *TodosController) Create(ctx *app.CreateTodosContext) error {
	// TodosController_Create: start_implement

	// Put your logic here

	// TodosController_Create: end_implement
	return nil
}

// Delete runs the delete action.
func (c *TodosController) Delete(ctx *app.DeleteTodosContext) error {
	// TodosController_Delete: start_implement

	// Put your logic here

	// TodosController_Delete: end_implement
	return nil
}

// DeleteAll runs the deleteAll action.
func (c *TodosController) DeleteAll(ctx *app.DeleteAllTodosContext) error {
	// TodosController_DeleteAll: start_implement

	// Put your logic here

	// TodosController_DeleteAll: end_implement
	return nil
}

// Read runs the read action.
func (c *TodosController) Read(ctx *app.ReadTodosContext) error {
	// TodosController_Read: start_implement

	// Put your logic here

	// TodosController_Read: end_implement
	res := &app.Todo{}
	return ctx.OK(res)
}

// Search runs the search action.
func (c *TodosController) Search(ctx *app.SearchTodosContext) error {
	// TodosController_Search: start_implement

	// Put your logic here

	// TodosController_Search: end_implement
	res := &app.Todo{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *TodosController) Update(ctx *app.UpdateTodosContext) error {
	// TodosController_Update: start_implement

	// Put your logic here

	// TodosController_Update: end_implement
	return nil
}
