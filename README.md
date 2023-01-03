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

#### /projects/title
* `GET` : Get a project
* `PUT` : Update a project
* `DELETE` : Delete a project

#### /projects/title/tasks
* `GET` : Get all tasks
* `POST` : Create a new task

#### /projects/title/tasks/id
* `GET` : Get a specific task
* `PUT` : Update a specific task
* `DELETE` : Delete a specific task

## Challanges Faced
* Resolved circular dependency while breaking code into packages. Introduced interface to resolve cycle dependency issue. Have a look at this [blog](https://quoeamaster.medium.com/golang-gotchas-2-the-curse-of-import-cycle-not-allowed-6abfa3523f57) to better understand the circular dependency issues.

## Future Improvement
* The empty struct for project and tasks handler doesn't look like a good approach. I will try to find the alternative solution for that.

## Feature TODO
- [ ] Make Config file environment specific rather then hard coded values.
- [ ] Introduce Testing.
- [ ] Support Authentication with user for securing the APIs.
- [ ] Introduce CI/CD and security in pipelines.
- [ ] Deploying this api behind proxy server using Nginx