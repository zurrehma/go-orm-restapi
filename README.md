## GO ORM based REST API
A REST API example using Router(gorilla/mux) and ORM(gorm).

```bash
# Build and Run
cd go-orm-restapi
go build
./go-orm-restapi
# API Endpoint : http://127.0.0.1:3000
```

## API

#### /projects
* `GET` : Get all projects
* `POST` : Create a new project

#### /projects/:title
* `GET` : Get a project
* `PUT` : Update a project
* `DELETE` : Delete a project