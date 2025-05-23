// Code generated by goa v3.20.1, DO NOT EDIT.
//
// todos HTTP server types
//
// Command:
// $ goa gen github.com/meier-christoph/todo-backend-golang-goa/design

package server

import (
	todos "github.com/meier-christoph/todo-backend-golang-goa/gen/todos"
	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "todos" service "create" endpoint HTTP
// request body.
type CreateRequestBody struct {
	// Title of the task
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Order or priority of the task
	Order *int `form:"order,omitempty" json:"order,omitempty" xml:"order,omitempty"`
	// Whether the task is completed or not
	Completed *bool `form:"completed,omitempty" json:"completed,omitempty" xml:"completed,omitempty"`
}

// UpdateRequestBody is the type of the "todos" service "update" endpoint HTTP
// request body.
type UpdateRequestBody struct {
	// Title of the task
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Order or priority of the task
	Order *int `form:"order,omitempty" json:"order,omitempty" xml:"order,omitempty"`
	// Whether the task is completed or not
	Completed *bool `form:"completed,omitempty" json:"completed,omitempty" xml:"completed,omitempty"`
}

// ListResponseBody is the type of the "todos" service "list" endpoint HTTP
// response body.
type ListResponseBody []*TodoResponse

// CreateResponseBody is the type of the "todos" service "create" endpoint HTTP
// response body.
type CreateResponseBody struct {
	// Unique ID
	ID  int    `form:"id" json:"id" xml:"id"`
	URL string `form:"url" json:"url" xml:"url"`
	// Title of the task
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Order or priority of the task
	Order *int `form:"order,omitempty" json:"order,omitempty" xml:"order,omitempty"`
	// Whether the task is completed or not
	Completed *bool `form:"completed,omitempty" json:"completed,omitempty" xml:"completed,omitempty"`
}

// ReadResponseBody is the type of the "todos" service "read" endpoint HTTP
// response body.
type ReadResponseBody struct {
	// Unique ID
	ID  int    `form:"id" json:"id" xml:"id"`
	URL string `form:"url" json:"url" xml:"url"`
	// Title of the task
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Order or priority of the task
	Order *int `form:"order,omitempty" json:"order,omitempty" xml:"order,omitempty"`
	// Whether the task is completed or not
	Completed *bool `form:"completed,omitempty" json:"completed,omitempty" xml:"completed,omitempty"`
}

// UpdateResponseBody is the type of the "todos" service "update" endpoint HTTP
// response body.
type UpdateResponseBody struct {
	// Unique ID
	ID  int    `form:"id" json:"id" xml:"id"`
	URL string `form:"url" json:"url" xml:"url"`
	// Title of the task
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Order or priority of the task
	Order *int `form:"order,omitempty" json:"order,omitempty" xml:"order,omitempty"`
	// Whether the task is completed or not
	Completed *bool `form:"completed,omitempty" json:"completed,omitempty" xml:"completed,omitempty"`
}

// CreateBadRequestResponseBody is the type of the "todos" service "create"
// endpoint HTTP response body for the "BadRequest" error.
type CreateBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ReadNotFoundResponseBody is the type of the "todos" service "read" endpoint
// HTTP response body for the "NotFound" error.
type ReadNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateNotFoundResponseBody is the type of the "todos" service "update"
// endpoint HTTP response body for the "NotFound" error.
type UpdateNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateBadRequestResponseBody is the type of the "todos" service "update"
// endpoint HTTP response body for the "BadRequest" error.
type UpdateBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// DeleteNotFoundResponseBody is the type of the "todos" service "delete"
// endpoint HTTP response body for the "NotFound" error.
type DeleteNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// TodoResponse is used to define fields on response body types.
type TodoResponse struct {
	// Unique ID
	ID  int    `form:"id" json:"id" xml:"id"`
	URL string `form:"url" json:"url" xml:"url"`
	// Title of the task
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Order or priority of the task
	Order *int `form:"order,omitempty" json:"order,omitempty" xml:"order,omitempty"`
	// Whether the task is completed or not
	Completed *bool `form:"completed,omitempty" json:"completed,omitempty" xml:"completed,omitempty"`
}

// NewListResponseBody builds the HTTP response body from the result of the
// "list" endpoint of the "todos" service.
func NewListResponseBody(res []*todos.Todo) ListResponseBody {
	body := make([]*TodoResponse, len(res))
	for i, val := range res {
		body[i] = marshalTodosTodoToTodoResponse(val)
	}
	return body
}

// NewCreateResponseBody builds the HTTP response body from the result of the
// "create" endpoint of the "todos" service.
func NewCreateResponseBody(res *todos.Todo) *CreateResponseBody {
	body := &CreateResponseBody{
		ID:        res.ID,
		URL:       res.URL,
		Title:     res.Title,
		Order:     res.Order,
		Completed: res.Completed,
	}
	return body
}

// NewReadResponseBody builds the HTTP response body from the result of the
// "read" endpoint of the "todos" service.
func NewReadResponseBody(res *todos.Todo) *ReadResponseBody {
	body := &ReadResponseBody{
		ID:        res.ID,
		URL:       res.URL,
		Title:     res.Title,
		Order:     res.Order,
		Completed: res.Completed,
	}
	return body
}

// NewUpdateResponseBody builds the HTTP response body from the result of the
// "update" endpoint of the "todos" service.
func NewUpdateResponseBody(res *todos.Todo) *UpdateResponseBody {
	body := &UpdateResponseBody{
		ID:        res.ID,
		URL:       res.URL,
		Title:     res.Title,
		Order:     res.Order,
		Completed: res.Completed,
	}
	return body
}

// NewCreateBadRequestResponseBody builds the HTTP response body from the
// result of the "create" endpoint of the "todos" service.
func NewCreateBadRequestResponseBody(res *goa.ServiceError) *CreateBadRequestResponseBody {
	body := &CreateBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewReadNotFoundResponseBody builds the HTTP response body from the result of
// the "read" endpoint of the "todos" service.
func NewReadNotFoundResponseBody(res *goa.ServiceError) *ReadNotFoundResponseBody {
	body := &ReadNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateNotFoundResponseBody builds the HTTP response body from the result
// of the "update" endpoint of the "todos" service.
func NewUpdateNotFoundResponseBody(res *goa.ServiceError) *UpdateNotFoundResponseBody {
	body := &UpdateNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateBadRequestResponseBody builds the HTTP response body from the
// result of the "update" endpoint of the "todos" service.
func NewUpdateBadRequestResponseBody(res *goa.ServiceError) *UpdateBadRequestResponseBody {
	body := &UpdateBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewDeleteNotFoundResponseBody builds the HTTP response body from the result
// of the "delete" endpoint of the "todos" service.
func NewDeleteNotFoundResponseBody(res *goa.ServiceError) *DeleteNotFoundResponseBody {
	body := &DeleteNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateTodoPayload builds a todos service create endpoint payload.
func NewCreateTodoPayload(body *CreateRequestBody) *todos.TodoPayload {
	v := &todos.TodoPayload{
		Title:     body.Title,
		Order:     body.Order,
		Completed: body.Completed,
	}

	return v
}

// NewReadPayload builds a todos service read endpoint payload.
func NewReadPayload(id string) *todos.ReadPayload {
	v := &todos.ReadPayload{}
	v.ID = id

	return v
}

// NewUpdatePayload builds a todos service update endpoint payload.
func NewUpdatePayload(body *UpdateRequestBody, id string) *todos.UpdatePayload {
	v := &todos.UpdatePayload{
		Title:     body.Title,
		Order:     body.Order,
		Completed: body.Completed,
	}
	v.ID = id

	return v
}

// NewDeletePayload builds a todos service delete endpoint payload.
func NewDeletePayload(id string) *todos.DeletePayload {
	v := &todos.DeletePayload{}
	v.ID = id

	return v
}
