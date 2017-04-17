// Code generated by goagen v1.1.0, DO NOT EDIT.
//
// API "todo": Application User Types
//
// Command:
// $ goagen
// --design=github.com/meier-christoph/todo-backend-golang-goa/design
// --out=$(GOPATH)/src/github.com/meier-christoph/todo-backend-golang-goa
// --version=v1.1.0-dirty

package app

// todoItem user type.
type todoItem struct {
	// Whether the task is completed or not
	Completed *bool `form:"completed,omitempty" json:"completed,omitempty" xml:"completed,omitempty"`
	// Order or priority of the task
	Order *int `form:"order,omitempty" json:"order,omitempty" xml:"order,omitempty"`
	// Title of the task
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
}

// Publicize creates TodoItem from todoItem
func (ut *todoItem) Publicize() *TodoItem {
	var pub TodoItem
	if ut.Completed != nil {
		pub.Completed = ut.Completed
	}
	if ut.Order != nil {
		pub.Order = ut.Order
	}
	if ut.Title != nil {
		pub.Title = ut.Title
	}
	return &pub
}

// TodoItem user type.
type TodoItem struct {
	// Whether the task is completed or not
	Completed *bool `form:"completed,omitempty" json:"completed,omitempty" xml:"completed,omitempty"`
	// Order or priority of the task
	Order *int `form:"order,omitempty" json:"order,omitempty" xml:"order,omitempty"`
	// Title of the task
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
}