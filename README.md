# github.com/rahul-golang/mongocrud.


This application is desiended in golang with mongodb Of sample Rest APIs

## Project Architecture

        ├── cmd
        │   └── service.go
        ├── database
        │   └── mongo.go
        ├── go.mod
        ├── go.sum
        ├── log
        │   └── logging.go
        ├── logfile.log
        ├── main.go
        ├── pkg
        │   ├── endpoints
        │   │   └── router.go
        │   ├── handler
        │   │   ├── helpers.go
        │   │   └── user_handler.go
        │   ├── models
        │   │   ├── model.go
        │   │   └── user.go
        │   ├── repository
        │   │   └── user_repository.go
        │   └── service
        │       └── user_service.go
        └── README.md