package internal

import (
	"context"
	"errors"
	"fmt"
	"github.com/meier-christoph/todo-backend-golang-goa/gen/todos"
	"gorm.io/gorm"
)

type TodosServiceImpl struct {
	*gorm.DB
}

func NewTodosServiceImpl(db *gorm.DB) *TodosServiceImpl {
	return &TodosServiceImpl{
		DB: db,
	}
}

type TodoDAO struct {
	ID        int `gorm:"primaryKey"`
	Title     string
	Order     int
	Completed bool
}

func NewTodoWithURL(ctx context.Context, dao *TodoDAO) *todos.Todo {
	url, _ := ctx.Value(ReverseProxyPublicUrl).(string)
	if url == "" {
		url = "http://localhost:8080"
	}
	t := todos.Todo{}
	t.ID = dao.ID
	t.Title = &dao.Title
	t.Order = &dao.Order
	t.Completed = &dao.Completed
	t.URL = fmt.Sprintf("%s/todos/%d", url, dao.ID)
	return &t
}

func (c *TodosServiceImpl) List(ctx context.Context) (res []*todos.Todo, err error) {
	var obj []*TodoDAO
	err = c.DB.Find(&obj).Error
	if err != nil {
		return nil, err
	}

	if obj == nil || len(obj) == 0 {
		return []*todos.Todo{}, nil
	}

	for _, t := range obj {
		res = append(res, NewTodoWithURL(ctx, t))
	}

	return res, nil
}

func (c *TodosServiceImpl) Create(ctx context.Context, input *todos.TodoPayload) (res *todos.Todo, err error) {
	obj := TodoDAO{}
	if input.Completed != nil {
		obj.Completed = *input.Completed
	}
	if input.Order != nil {
		obj.Order = *input.Order
	}
	if input.Title != nil {
		obj.Title = *input.Title
	}

	err = c.DB.Create(&obj).Error
	if err != nil {
		return nil, err
	}

	return NewTodoWithURL(ctx, &obj), nil
}

func (c *TodosServiceImpl) Read(ctx context.Context, input *todos.ReadPayload) (res *todos.Todo, err error) {
	var obj TodoDAO
	err = c.DB.Where("id = ?", input.ID).First(&obj).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, todos.MakeNotFound(err)
		}
		return nil, err
	}
	return NewTodoWithURL(ctx, &obj), nil
}

func (c *TodosServiceImpl) Update(ctx context.Context, input *todos.UpdatePayload) (res *todos.Todo, err error) {
	var obj TodoDAO
	err = c.DB.Where("id = ?", input.ID).First(&obj).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, todos.MakeNotFound(err)
		}
		return nil, err
	}

	if input.Completed != nil {
		obj.Completed = *input.Completed
	}
	if input.Order != nil {
		obj.Order = *input.Order
	}
	if input.Title != nil {
		obj.Title = *input.Title
	}

	err = c.DB.Save(&obj).Error
	if err != nil {
		return nil, err
	}

	return NewTodoWithURL(ctx, &obj), nil
}

func (c *TodosServiceImpl) Delete(ctx context.Context, input *todos.DeletePayload) (err error) {
	var obj TodoDAO
	err = c.DB.Delete(&obj, input.ID).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *TodosServiceImpl) DeleteAll(ctx context.Context) (err error) {
	err = c.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&TodoDAO{}).Error
	if err != nil {
		return err
	}
	return nil
}
