
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

## Desplegar en Podman

### Construir imagen aplicacion
```bash
podman build -t go2music-app .
```

### Creo un pod para la aplicacion.
```bash
podman pod create --name go2music-pod -p 5432:5432 -p 8080:8080
```

### Crear un volumen de Podman
```bash
podman volume create pgdata
```

### Desplegar base postgres en un pod:
```bash
podman run -d --name pgsql \
  --pod go2music-pod \
  -e POSTGRES_DB=go2music \
  -e POSTGRES_USER=pichupostgres \
  -e POSTGRES_PASSWORD=passpostgres \
  -v pgdata:/var/lib/postgresql/data \
  docker.io/library/postgres:16
```

### Desplegar app en el pod
```bash
podman run -d --name app \
  --pod go2music-pod \
  -e DATABASE_URL="postgres://pichupostgres:passpostgres@localhost:5432/go2music?sslmode=disable" \
  go2music-app
```

### Desplegar base postgres (SOLO):
```bash
podman run -d \
  --name goapp-postgres \
  -e POSTGRES_USER=pichupostgres \
  -e POSTGRES_PASSWORD=passpostgres \
  -e POSTGRES_DB=go2music \
  -p 5432:5432 \
  -v goapp_pgdata:/var/lib/postgresql/data \
  docker.io/library/postgres:16
```


#### Verificar conexion de la base:
```bash
podman exec -it goapp-postgres psql -U pichupostgres -d go2music
```

### Ejecutar aplicacion:
```bash
go run ./cmd/api/main.go
```

### Crear tabla users
``` sql
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    role VARCHAR(50) NOT NULL DEFAULT 'user',
    image TEXT,
    created_data TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
```

### Datos prueba:
```bash
localhost:8001/users
```

```json
{
	"name":"pichu1",
	"surname":"Pichuape",
	"password":"Pichupass",
	"email":"mail@mail.com",
	"role":"PichuRole",
	"image":"PichuImage"
}
```

