# Go Boilerplate
An API boilerplate written in Golang with chi-route and Gorm

## Table of Contents
- [Motivation](#motivation)
- [Configuration Manage](#configuration-manage)
  - [ENV Manage](#env-manage)
  - [Server Configuration](#server-configuration)
  - [Database Configuration](#database-configuration)
  - [PgAdmin](#pg-admin)
- [Installation](#installation)
  - [Local Setup Instruction](#local-setup-instruction)
  - [Develop Application in Docker with Live Reload](#develop-application-in-docker-with-live-reload)
- [Middlewares](#middlewares)
- [Boilerplate Structure](#boilerplate-structure)
- [Code Examples](#examples)
- [Let's Build an API](#lets-build-an-api)
- [Deployment](#deployment)
  - [Container Development Build](#container-development-build)
  - [Container Production Build and Up](#container-production-build-and-up)
- [Useful Commands](#useful-commands)
- [ENV YAML Configure](#env-yaml-configure)
- [Use Packages](#use-packages)

### Motivation
Write restful API with fast development and developer friendly

### Configuration Manage
#### ENV Manage

- Default ENV Configuration Manage from `.env`. sample file `.env.example`
```text
# Server Configuration
SECRET=h9wt*pasj6796j##w(w8=xaje8tpi6h*r&hzgrz065u&ed+k2)
DEBUG=True # `False` in Production
ALLOWED_HOSTS=0.0.0.0
SERVER_HOST=0.0.0.0
SERVER_PORT=8000

# Database Configuration
MASTER_DB_NAME=test_pg_go
MASTER_DB_USER=mamun
MASTER_DB_PASSWORD=123
MASTER_DB_HOST=postgres_db
MASTER_DB_PORT=5432
MASTER_DB_LOG_MODE=True # `False` in Production
MASTER_SSL_MODE=disable

REPLICA_DB_NAME=test_pg_go
REPLICA_DB_USER=mamun
REPLICA_DB_PASSWORD=123
REPLICA_DB_HOST=localhost
REPLICA_DB_PORT=5432
REPLICA_DB_LOG_MODE=True # `False` in Production
REPLICA_SSL_MODE=disable
```
- Server `DEBUG` set `False` in Production
- Database Logger `MASTER_DB_LOG_MODE` and `REPLICA_DB_LOG_MODE`  set `False` in production
- If ENV Manage from YAML file add a config.yml file and configuration [db.go](pkg/config/db.go) and [server.go](pkg/config/server.go). See More [ENV YAML Configure](#env-yaml-configure)

#### Server Configuration
- Use [chi](https://github.com/go-chi/chi) Route

#### Database Configuration
- Use [GORM](https://github.com/go-gorm/gorm) as an ORM
- Use database `MASTER_DB_HOST` value set as `localhost` for local development, and use `postgres_db` for docker development 
#### PG Admin
- Check  PG Admin on [http://0.0.0.0:5050/browser/](http://0.0.0.0:5050/browser/)
- Login with Credential Email `admin@admin.com` Password `root`
- Connect Database Host as `postgres_db`, DB Username and Password as per `.env` set
- Note: if not configure `.env`, default Username `mamun` and password `123`

### Installation
#### Local Setup Instruction
Follow these steps:
- Copy [.env.example](.env.example) as `.env` and configure necessary values
- To add all dependencies for a package in your module `go get .` in the current directory
- Locally run `go run main.go` or `go build main.go` and run `./main`
- Check Application health available on [0.0.0.0:8000/health](http://0.0.0.0:8000/health)

#### Develop Application in Docker with Live Reload
Follow these steps:
- Make sure install the latest version of docker and docker-compose
- Docker Installation for your desire OS https://docs.docker.com/engine/install/ubuntu/
- Docker Composer Installation https://docs.docker.com/compose/install/
- Run and Develop `make dev`
- Check Application health available on [0.0.0.0:8000/health](http://0.0.0.0:8000/health)

### Boilerplate Structure
<<pre><font color="#2A7BDE"><b>.</b></font>
├── <font color="#2A7BDE"><b>controllers</b></font>
│   └── example_controller.go
├── go.mod
├── go.sum
├── LICENSE
├── main.go
├── Makefile
├── <font color="#2A7BDE"><b>migrations</b></font>
│   └── migrations.go
├── <font color="#2A7BDE"><b>models</b></font>
│   ├── base_model.go
│   └── example_model.go
├── <font color="#2A7BDE"><b>pkg</b></font>
│   ├── <font color="#2A7BDE"><b>config</b></font>
│   │   ├── config.go
│   │   ├── db.go
│   │   └── server.go
│   ├── <font color="#2A7BDE"><b>database</b></font>
│   │   └── database.go
│   ├── <font color="#2A7BDE"><b>helpers</b></font>
│   │   ├── <font color="#2A7BDE"><b>pagination</b></font>
│   │   │   └── pagination.go
│   │   ├── response.go
│   │   └── search.go
│   └── <font color="#2A7BDE"><b>logger</b></font>
│       └── logger.go
├── README.md
├── <font color="#2A7BDE"><b>repository</b></font>
│   └── example_repo.go
├── <font color="#2A7BDE"><b>routers</b></font>
     ├── examples.go
     ├── index.go 
     └── router.go</pre>

### Let's Build an API

1. [models](models) folder add a new file name `example_model.go`

```go
package models

import (
	"time"
)

type Example struct {
	Id        int        `json:"id"`
	Data      string     `json:"data" binding:"required"`
	CreatedAt *time.Time `json:"created_at,string,omitempty"`
	UpdatedAt *time.Time `json:"updated_at_at,string,omitempty"`
}
// TableName is Database Table Name of this model
func (e *Example) TableName() string {
	return "examples"
}
```
2. Add Model to [migration](migrations/migrations.go)
```go
package migrations

import (
  "chi-boilerplate/models"
  "chi-boilerplate/pkg/database"
)

func Migrate() {
  var migrationModels = []interface{}{&models.Example{}}
  err := database.GetDB().AutoMigrate(migrationModels...)
  if err != nil {
    return
  }
}
```
3. [controller](controllers) folder add a file `example_controller.go`
- Create API Endpoint 
- Use any syntax of GORM after `base.DB`, this is wrapper of `*gorm.DB`

```go
package controllers

import (
  "chi-boilerplate/models"
  "encoding/json"
  "net/http"
)

func CreateExample(w http.ResponseWriter, request *http.Request) {
  example := new(models.Example)
  err := json.NewDecoder(request.Body).Decode(&example)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }
  models.SaveExample(&example)
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(&example)
}

func GetData(w http.ResponseWriter, request *http.Request) {
  var example []models.Example
  models.GetAll(&example)
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(&example)
}

```
4. [routers](routers) folder add a file `example.go`
```go
package routers

import (
  "chi-boilerplate/controllers"
  "github.com/go-chi/chi/v5"
)

func ExamplesRoutes(router *chi.Mux) {
  router.Group(func(r chi.Router) {
    r.Post("/test/", controllers.CreateExample)
    r.Get("/test/", controllers.GetData)
  })
}
```
5. Finally, register routes to [index.go](routers/index.go)
```go
package routers

import (
  "github.com/go-chi/chi/v5"
  "net/http"
)

//RegisterRoutes add all routing list here automatically get main router
func RegisterRoutes(router *chi.Mux) {
  router.Get("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("\"live\": \"ok\""))
  })
  //Add All route
  ExamplesRoutes(router)
}
```
- Congratulation, your new endpoint `0.0.0.0:8000/v1/example/`

### Deployment
#### Container Development Build
- Run `make build`

#### Container Production Build and Up
- Run `make production`

#### ENV Yaml Configure
```yaml
database:
  driver: "postgres"
  dbname: "test_pg_go"
  username: "mamun"
  password: "123"
  host: "postgres_db" # use `localhost` for local development
  port: "5432"
  ssl_mode: disable
  log_mode: false

server:
  host: "0.0.0.0"
  port: "8000"
  secret: "secret"
  allow_hosts: "localhost"
  debug: false #use `false` in production
  request:
    timeout: 100
```
- [Server Config](pkg/config/server.go)
```go
func ServerConfig() string {
viper.SetDefault("server.host", "0.0.0.0")
viper.SetDefault("server.port", "8000")
appServer := fmt.Sprintf("%s:%s", viper.GetString("server.host"), viper.GetString("server.port"))
return appServer
}
```
- [DB Config](pkg/config/db.go)
```go
func DbConfiguration() string {
	
dbname := viper.GetString("database.dbname")
username := viper.GetString("database.username")
password := viper.GetString("database.password")
host := viper.GetString("database.host")
port := viper.GetString("database.port")
sslMode := viper.GetString("database.ssl_mode")

dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
host, username, password, dbname, port, sslMode)
return dsn
}
```

### Useful Commands

- `make dev`: make dev for development work
- `make build`: make build container
- `make production`: docker production build and up
- `clean`: clean for all clear docker images

### Use Packages
- [Viper](https://github.com/spf13/viper) - Go configuration with fangs.
- [Gorm](https://github.com/go-gorm/gorm) - The fantastic ORM library for Golang
- [Logger](https://github.com/sirupsen/logrus) - Structured, pluggable logging for Go.
- [Air](https://github.com/cosmtrek/air) - Live reload for Go apps (Docker Development)

