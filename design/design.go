package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("todo", func() {
	Description("Todo-Backend Implementation with GOLANG & GOA")
	Scheme("http")
	Host("localhost:8080")
	Consumes("application/json")
	Produces("application/json")
})

var _ = Resource("todos", func() {
	BasePath("/todo")
	DefaultMedia(TodoMedia)
	Origin("*", func() {
		Methods("OPTION", "GET", "POST", "PATCH", "DELETE")
		Headers("Content-Type")
	})
	Action("search", func() {
		Routing(GET("/"))
		Response(OK)
	})
	Action("create", func() {
		Routing(POST("/"))
		Response(NoContent)
	})
	Action("deleteAll", func() {
		Routing(DELETE("/"))
		Response(NoContent)
	})
	Action("read", func() {
		Routing(GET("/:id"))
		Params(func() {
			Param("id", Integer, "Unique ID")
		})
		Response(OK)
		Response(NotFound)
	})
	Action("update", func() {
		Routing(PATCH("/:id"))
		Params(func() {
			Param("id", Integer, "Unique ID")
		})
		Response(NoContent)
		Response(NotFound)
	})
	Action("delete", func() {
		Routing(DELETE("/:id"))
		Params(func() {
			Param("id", Integer, "Unique ID")
		})
		Response(NoContent)
		Response(NotFound)
	})
})

var TodoMedia = MediaType("application/vnd.todo+json", func() {
	Attributes(func() {
		Attribute("id", Integer, "Unique ID")
		Attribute("title", String, "Title of the task")
		Attribute("order", Integer, "Order or priority of the task")
		Attribute("completed", Boolean, "Whether the task is completed or not")
		Required("id", "title", "order", "completed")
	})
	View("default", func() {
		Attribute("id")
		Attribute("title")
		Attribute("order")
		Attribute("completed")
	})
})
