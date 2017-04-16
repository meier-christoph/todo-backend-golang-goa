# TODO-Backend with GOA

## How to build this project

- install golang and setup $GOPATH environment variable
- checkout this repo into your $GOPATH
- install glide for vendor management
- run `glide install` to initialize vendor dependencies for this project
- make goagen cli tools frm vendor
  - `cd ./vendor/github.com/goadesign/goa/goagen`
  - `go build`
  - `cd -`
- generate code using goagen cli tools
  - `./vendor/github.com/goadesign/goa/goagen/goagen app -d github.com/meier-christoph/todo-backend-golang-goa/design`
  - `./vendor/github.com/goadesign/goa/goagen/goagen main -d github.com/meier-christoph/todo-backend-golang-goa/design`

## All in one
`go generate && go build && ./todo-backend-golang-goa`
