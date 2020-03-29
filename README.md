# receipts-api
Study project using Golang + fasthttp + MongoDB driver.

****

### Project

receipts-api is an rest api to store receipt links along with some notes. 

The purpouse is to keep Golang codding skills and improve and/or test some concepts.

### Tech Stack

* [Golang](https://golang.org/)
* [Mongo Go Driver](https://github.com/mongodb/mongo-go-driver)
* [FastHttp](https://github.com/valyala/fasthttp)

### Running

You can run this api project navigating to cmd folder and running build and run. When executing run command you can pass a third argument as dev, qa or prod indicating the environment you want to execute. When no arg is passing it will assume dev.


##### Navigate to cmd folder
``
cd cmd
``

##### build
``
go build
``

##### run without argument will assume dev envirnment
``
go run main.go
``

##### run with one of those arguments on list
``
go run  main.go [dev, qa, prod]
``

You can check the environment configurations on fodler
configs/files/mongodb and configs/files/server  by its indication on config.mongodb.<env>.json. Exampels:
* config.mongodb.dev.json
* config.server.dev.json

### Architecture

This project is organized in 3 layers.
1. Controller (api)

contains API level executions, such as capturing parameters, parsing body to models and responsing api executions setting http status and messages.

2. Use Case 

Executes the business rules, such as validations.

3. Repository

Execute commands on MongoDB

#### Structure overview

![structure](./_documentation/structure-overview.png)

#### Folders

* api : Register controllers

* cmd : Main application

* configs : Configuration files and procedures
* handlers : Prepare Server 
* internals : Business side information
* _documentation : support for readme files

### Future developments 

* Adding docker 
* Adding GRPC
    * For a log project

#### Good links
* [Project Layout](https://github.com/golang-standards/project-layout)


