
## Handler
Se encarga de 3 cosas.
- Traducir el Request
- Consumir un servicio de la aplicacion
- Traducir el Response


```
go2music
├── cmd
│   └── api
│       ├── handlers
│       │       └── user
│       │           └── create.go
│       │           └── handler.go
│       ├── main.go
│       └── tmp
├── internal
│       ├── domain
│       │       └── user.go
│       ├── ports
│       │       └── user.go
│       ├── repositories
│       │       └── postgres
│       │           ├── user
│       │           │    └── insert.go
│       │           │    └── repository.go
│       │           └── connect_client.go
│       └── services
│               └── user
│                   ├── create.go
│                   └── service.go
├── go.mod
├── go.sum
├── LICENSE
└── README.md
```
