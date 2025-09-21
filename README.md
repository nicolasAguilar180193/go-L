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