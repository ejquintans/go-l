
## Handler
Se encarga de 3 cosas.
- Traducir el Request
- Consumir un servicio de la aplicacion
- Traducir el Response


.
├── cmd
│   └── api
│       ├── handlers
│       │       └── user
│       │           └── create.go
│       ├── main.go
│       └── tmp
├── go.mod
├── go.sum
├── internal
│       ├── domain
│       │       └── user.go
│       └── services
│               └── user
│                   └── create.go
├── LICENSE
└── README.md
