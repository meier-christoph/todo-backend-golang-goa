# TODO-Backend with GOA

Minimal viable product passing the test suite for todo-backends with golang and goa.

Deployed at the link below:

https://todo-backend-golang-goa.herokuapp.com/todos

### How to build this project

- install golang
- checkout the code
- generate code using goa gen (*)
    - `go generate`
- run the app
    - `go run ./main.go`

(*): Normally you would not commit the generated code but since this is a showcase for goa I decided to include it
anyway. All the generated code is in the gen folder, everything else is part of the actual source.

### Run the todo backend test suite

Locally:

https://www.todobackend.com/specs/index.html?http://localhost:8080/todos

Hosted:

https://www.todobackend.com/specs/index.html?http://localhost:8080/todos

### Run with PostgreSQL

```shell
export DATABASE_URI="postgresql://user:secret@localhost:5432/todos"
```

Note: if this variable is not set it will run with sqlite (relative 'todos.db' file)
