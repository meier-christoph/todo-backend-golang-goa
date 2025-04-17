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

### Run with PostgreSQL
