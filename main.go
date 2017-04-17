//go:generate goagen app -d github.com/meier-christoph/todo-backend-golang-goa/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/meier-christoph/todo-backend-golang-goa/app"
	"github.com/meier-christoph/todo-backend-golang-goa/controllers"
	"log"
	"os"
)

func main() {
	dialect := os.Getenv("DATABASE_DIALECT")
	if dialect == "" {
		dialect = "sqlite3"
	}
	source := os.Getenv("DATABASE_URL")
	if source == "" {
		source = "todo.db"
	}

	var db *gorm.DB
	db, err := gorm.Open(dialect, source)
	if err != nil {
		log.Fatal("Failed to connect to db")
	}
	defer db.Close()
	db.AutoMigrate(controllers.TodoDAO{})

	service := goa.New("todo-backend-golang-goa")

	// middleware(s)
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// controller(s)
	c := controllers.NewTodosController(service, db)
	app.MountTodosController(service, c)

	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
