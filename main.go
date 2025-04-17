//go:generate goa gen github.com/meier-christoph/todo-backend-golang-goa/design

package main

import (
	"github.com/glebarez/sqlite"
	genhttpprobes "github.com/meier-christoph/todo-backend-golang-goa/gen/http/probes/server"
	genhttptodos "github.com/meier-christoph/todo-backend-golang-goa/gen/http/todos/server"
	genprobes "github.com/meier-christoph/todo-backend-golang-goa/gen/probes"
	gentodos "github.com/meier-christoph/todo-backend-golang-goa/gen/todos"
	"github.com/meier-christoph/todo-backend-golang-goa/internal"
	"github.com/spf13/cobra"
	goadebug "goa.design/clue/debug"
	goalog "goa.design/clue/log"
	goahttp "goa.design/goa/v3/http"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	cmd := NewRootCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "todo-backend",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			var dial gorm.Dialector
			uri := os.Getenv("DATABASE_URI")
			if uri != "" {
				dial = postgres.Open(uri)
			} else {
				dial = sqlite.Open("todos.db")
			}

			db, err := gorm.Open(dial, &gorm.Config{})
			if err != nil {
				return err
			}

			err = db.AutoMigrate(internal.TodoDAO{})
			if err != nil {
				return err
			}

			mux := goahttp.NewMuxer()
			requestDecoder := goahttp.RequestDecoder
			responseEncoder := goahttp.ResponseEncoder

			// todos
			todoSrv := &internal.TodosService{DB: db}
			todosEndpoints := gentodos.NewEndpoints(todoSrv)
			todosEndpoints.Use(goadebug.LogPayloads())
			todosEndpoints.Use(goalog.Endpoint)
			todosServer := genhttptodos.New(todosEndpoints, mux, requestDecoder, responseEncoder, nil, nil)
			genhttptodos.Mount(mux, todosServer)

			// probes
			probesSrv := &internal.ProbesService{}
			probesEndpoints := genprobes.NewEndpoints(probesSrv)
			probesServer := genhttpprobes.New(probesEndpoints, mux, requestDecoder, responseEncoder, nil, nil)
			genhttpprobes.Mount(mux, probesServer)

			var handler http.Handler = mux
			handler = internal.ReverseProxy()(handler)

			server := &http.Server{Addr: ":8080", Handler: handler}
			log.Println("Starting service on :8080")
			return server.ListenAndServe()
		},
	}

	return cmd
}
