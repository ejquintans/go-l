# Etapa 1: build (opcional, solo si querés construir dentro del contenedor)
FROM golang:1.24 AS builder

WORKDIR /app

# Copiamos los archivos
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compilamos el binario
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/api

# Etapa 2: imagen final mínima
FROM alpine:latest

# Instalar ca-certificates para conexiones HTTPS/SSL (útil si se conecta a PostgreSQL con TLS)
# RUN apk add --no-cache ca-certificates

# Crear directorio de la app
WORKDIR /app

# Copiar binario desde la etapa de build
COPY --from=builder /app/app .

# Variables de entorno por defecto (pueden sobrescribirse en tiempo de ejecución)
ENV PORT=8080
ENV DATABASE_URL=postgres://pichupostgres:passpostgres@localhost:5432/go2music?sslmode=disable

# Exponer el puerto (opcional, para documentación)
EXPOSE ${PORT}

# Ejecutar la app
ENTRYPOINT ["/app/app"]
