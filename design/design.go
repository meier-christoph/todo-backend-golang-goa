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
	BasePath("/todos")
	DefaultMedia(TodoMedia)
	Origin("*", func() {
		Methods("OPTION", "GET", "POST", "PATCH", "DELETE")
		Headers("Content-Type")
	})
	Action("search", func() {
		Routing(GET("/"))
		Response(OK, ArrayOf(TodoMedia))
		Response(BadRequest)
	})
	Action("create", func() {
		Routing(POST("/"))
		Payload(TodoItem)
		Response("Created", func() {
			Status(201)
			Headers(func() {
				Header("Location", String, "Resource location", func() {
					Pattern("/todos/[0-9]+")
				})
			})
			Media(TodoMedia)
		})
		Response(BadRequest)
	})
	Action("deleteAll", func() {
		Routing(DELETE("/"))
		Response(NoContent)
		Response(BadRequest)
	})
	Action("read", func() {
		Routing(GET("/:id"))
		Params(func() {
			Param("id", Integer, "Unique ID")
		})
		Response(OK, TodoMedia)
		Response(NotFound)
		Response(BadRequest)
	})
	Action("update", func() {
		Routing(PATCH("/:id"))
		Params(func() {
			Param("id", Integer, "Unique ID")
		})
		Payload(TodoItem)
		Response(OK, TodoMedia)
		Response(NotFound)
		Response(BadRequest)
	})
	Action("delete", func() {
		Routing(DELETE("/:id"))
		Params(func() {
			Param("id", Integer, "Unique ID")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest)
	})
})

var TodoItem = Type("TodoItem", func() {
	Attribute("title", String, "Title of the task")
	Attribute("order", Integer, "Order or priority of the task")
	Attribute("completed", Boolean, "Whether the task is completed or not")
})

var TodoMedia = MediaType("application/vnd.todo+json", func() {
	Reference(TodoItem)
	Attributes(func() {
		Attribute("id", Integer, "Unique ID")
		Attribute("title")
		Attribute("order")
		Attribute("completed")
		Attribute("url", String, "Link to the task")
	})
	View("default", func() {
		Attribute("id")
		Attribute("title")
		Attribute("order")
		Attribute("completed")
		Attribute("url")
	})
})
