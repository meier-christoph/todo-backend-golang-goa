# TODO-Backend with GOA

Minimal viable product passing the test suite for todo-backends with golang and goa.

Deployed on heroku at the link below:


## How to build this project

- install golang and setup $GOPATH environment variable
- checkout this repo somewhere into your $GOPATH
- install glide for vendor management
- use glide to initialize vendor dependencies for this project
  - `glide install`
- make goagen cli tools from vendor
  - `cd ./vendor/github.com/goadesign/goa/goagen`
  - `go build`
  - `cd -`
- optional: move goagen to somewhere on the $PATH
  - `mv ./vendor/github.com/goadesign/goa/goagen/goagen $GOPATH/bin`
- generate code using goagen cli tools
  - `./vendor/github.com/goadesign/goa/goagen/goagen app -d github.com/meier-christoph/todo-backend-golang-goa/design`
  - `./vendor/github.com/goadesign/goa/goagen/goagen main -d github.com/meier-christoph/todo-backend-golang-goa/design`
- alternative: run go generate on main.go (goagen must be on $PATH) 
  - `go generate`

## All in one
`go generate && go build -i && ./todo-backend-golang-goa`

## Run with PostgresSQL (or MySQL)

`export DATABASE_DIALECT=postgres`

`export DATABASE_URL="host=192.168.99.100 user=todos dbname=todos sslmode=disable password=s3cr3t"`

Should also work with MySQL but I did not test it :P

Note: if those variables are not set it will run with sqlite (relative 'todo.db' file)
