package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = API("todo", func() {
	Title("Todo")
	Description("Todo-Backend Implementation with GOLANG & GOA")

	Server("main", func() {
		Host("localhost", func() {
			URI("http://localhost:8080")
		})
	})

	cors.Origin("*", func() {
		cors.Headers("*")
		cors.Methods("*")
	})
})

var _ = Service("probes", func() {
	Method("healthy", func() {
		NoSecurity()
		HTTP(func() {
			GET("/healthy")
			Response(StatusOK)
		})
	})
})

var _ = Service("todos", func() {
	HTTP(func() {
		Path("/todos")
	})

	Error("BadRequest", func() {
		Description("Invalid user input.")
	})
	Error("NotFound", func() {
		Description("Todo not found.")
	})

	Method("list", func() {
		Result(ArrayOf(Todo))
		HTTP(func() {
			GET("/")
			Response(StatusOK)
		})
	})

	Method("create", func() {
		Payload(TodoPayload)
		Result(Todo)
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
			Response("BadRequest", StatusBadRequest)
		})
	})

	Method("read", func() {
		Payload(func() {
			Attribute("id", String)
			Required("id")
		})
		Result(Todo)
		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
		})
	})

	Method("update", func() {
		Payload(func() {
			Extend(TodoPayload)
			Attribute("id", String)
			Required("id")
		})
		Result(Todo)
		HTTP(func() {
			PATCH("/{id}")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
			Response("BadRequest", StatusBadRequest)
		})
	})

	Method("delete", func() {
		Payload(func() {
			Attribute("id", String)
			Required("id")
		})
		HTTP(func() {
			DELETE("/{id}")
			Response(StatusNoContent)
			Response("NotFound", StatusNotFound)
		})
	})

	Method("deleteAll", func() {
		HTTP(func() {
			DELETE("/")
			Response(StatusNoContent)
		})
	})
})

var TodoPayload = Type("TodoPayload", func() {
	Attribute("title", String, "Title of the task")
	Attribute("order", Int, "Order or priority of the task")
	Attribute("completed", Boolean, "Whether the task is completed or not")
})

var Todo = Type("Todo", func() {
	Extend(TodoPayload)
	Attribute("id", Int, "Unique ID")
	Attribute("url", String)
	Required("id", "url")
})
