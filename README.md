# go-zero-wire
Testing using wire dependency injection with go-zero

Repo Structure:

```
tree . 
.
├── app
│   ├── infra
│   │   ├── docker-compose.yml
│   │   ├── go.mod
│   │   └── mysql.go
│   └── service
│       ├── api
│       │   └── rest
│       │       ├── etc
│       │       │   └── service-api.yaml
│       │       ├── internal
│       │       │   ├── config
│       │       │   │   └── config.go
│       │       │   ├── handler
│       │       │   │   ├── createuserhandler.go
│       │       │   │   ├── getuserhandler.go
│       │       │   │   └── routes.go
│       │       │   ├── logic
│       │       │   │   ├── createuserlogic.go
│       │       │   │   └── getuserlogic.go
│       │       │   ├── svc
│       │       │   │   └── servicecontext.go
│       │       │   └── types
│       │       │       └── types.go
│       │       ├── service.go
│       │       ├── wire_gen.go
│       │       └── wire.go
│       ├── databases
│       │   └── mysqldb
│       │       ├── 000001_create_user_table.down.sql
│       │       └── 000001_create_user_table.up.sql
│       ├── definition
│       │   └── service.api
│       ├── go.mod
│       ├── go.sum
│       └── internal
│           └── user
│               ├── domain
│               │   ├── mysqldb
│               │   │   └── user.go
│               │   └── user
│               │       ├── repository.go
│               │       ├── user.go
│               │       └── user_test.go
│               ├── logic
│               │   └── logic.go
│               └── repo
│                   └── mysql_repository.go
├── Makefile
└── README.md
```