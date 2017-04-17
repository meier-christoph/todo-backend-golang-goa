package controllers

import (
	"fmt"
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/meier-christoph/todo-backend-golang-goa/app"
	"net/http"
)

type TodosController struct {
	*goa.Controller
	*gorm.DB
}

type TodoDAO struct {
	ID        int `gorm:"primary_key"`
	Title     string
	Order     int
	Completed bool
}

// overrides default table name used by gorm
func (TodoDAO) TableName() string {
	return "todos"
}

func TableName() string {
	return "todos"
}

func NewTodosController(service *goa.Service, db *gorm.DB) *TodosController {
	return &TodosController{
		Controller: service.NewController("TodosController"),
		DB:         db,
	}
}

func TodoWithUrl(req *http.Request, dao *TodoDAO) *app.Todo {
	scheme := "http"
	if req.TLS != nil {
		scheme = "https"
	}
	t := app.Todo{}
	t.ID = &dao.ID
	t.Title = &dao.Title
	t.Order = &dao.Order
	t.Completed = &dao.Completed
	self := fmt.Sprintf("%v://%v/todos/%d", scheme, req.Host, dao.ID)
	t.URL = &self
	return &t
}

func (c *TodosController) Create(ctx *app.CreateTodosContext) error {
	obj := TodoDAO{}
	payload := ctx.Payload
	if payload.Completed != nil {
		obj.Completed = *payload.Completed
	}
	if payload.Order != nil {
		obj.Order = *payload.Order
	}
	if payload.Title != nil {
		obj.Title = *payload.Title
	}

	err := c.DB.Create(&obj).Error
	if err != nil {
		goa.LogError(ctx, "error creating todos", "error", err.Error())
		return ctx.BadRequest()
	}

	res := TodoWithUrl(ctx.Request, &obj)
	ctx.ResponseData.Header().Set("Location", *res.URL)
	return ctx.Created(res)
}

func (c *TodosController) Delete(ctx *app.DeleteTodosContext) error {
	var obj TodoDAO
	err := c.DB.Delete(&obj, ctx.ID).Error
	if err != nil {
		return ctx.NotFound()
	}
	return ctx.NoContent()
}

func (c *TodosController) DeleteAll(ctx *app.DeleteAllTodosContext) error {
	var obj TodoDAO
	err := c.DB.Delete(&obj).Error
	if err != nil {
		goa.LogError(ctx, "error deleting todos", "error", err.Error())
		return ctx.BadRequest()
	}
	return ctx.NoContent()
}

func (c *TodosController) Read(ctx *app.ReadTodosContext) error {
	var obj TodoDAO
	err := c.DB.Table(TableName()).Where("id = ?", ctx.ID).Find(&obj).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	}
	if err != nil {
		goa.LogError(ctx, "error reading todos", "error", err.Error())
		return ctx.BadRequest()
	}
	return ctx.OK(TodoWithUrl(ctx.Request, &obj))
}

func (c *TodosController) Search(ctx *app.SearchTodosContext) error {
	var res []*app.Todo
	var obj []*TodoDAO
	err := c.DB.Table(TableName()).Find(&obj).Error
	if err != nil {
		goa.LogError(ctx, "error listing todos", "error", err.Error())
		return ctx.BadRequest()
	}

	if obj == nil || len(obj) == 0 {
		return ctx.OK([]*app.Todo{})
	}

	for _, t := range obj {
		res = append(res, TodoWithUrl(ctx.Request, t))
	}

	return ctx.OK(res)
}

func (c *TodosController) Update(ctx *app.UpdateTodosContext) error {
	var obj TodoDAO
	err := c.DB.Table(TableName()).Where("id = ?", ctx.ID).Find(&obj).Error
	if err != nil {
		goa.LogError(ctx, "error reading todos", "error", err.Error())
		return ctx.BadRequest()
	}

	payload := ctx.Payload
	if payload.Completed != nil {
		obj.Completed = *payload.Completed
	}
	if payload.Order != nil {
		obj.Order = *payload.Order
	}
	if payload.Title != nil {
		obj.Title = *payload.Title
	}

	err = c.DB.Save(&obj).Error
	if err != nil {
		goa.LogError(ctx, "error updating todos", "error", err.Error())
		return ctx.BadRequest()
	}

	return ctx.OK(TodoWithUrl(ctx.Request, &obj))
}
