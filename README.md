⚽ GO-L

# Introduction

Introducing "GO-L", a cutting-edge RESTful API backend, meticulously crafted with Go. This project is not just about managing soccer tournaments and leagues, it's about redefining the way we interact with sports data.

## Init

```bash
# Copy .env
cp .env.example .env

# Install
go mod tidy

# Run
go run cmd/api/main.go
```

## Use mongodb linux

```bash
# Execute shell
mongosh

# show data bases
show dbs

# select go-l db
use go-l

# show collections
show colecttions

# Get documents
db.players.find()
```

## Ejecutar Test unitarios

```bash

# 1. Cobertura básica:
go test -cover ./...

# 2. Generar perfil:
go test -coverprofile=coverage.out ./...

#3. Reporte HTML:
go tool cover -html=coverage.out -o coverage.html
```

#### El reporte HTML te mostrará visualmente:
- Verde: Líneas cubiertas por tests
- Rojo: Líneas no cubiertas por tests  
- Gris: Líneas no ejecutables (comentarios, declaraciones, etc.)