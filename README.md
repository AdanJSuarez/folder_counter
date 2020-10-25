# Folder Counter

## Introduction
This app is able to find folders and files and order them by size.

Instructions:

The app should list the folders and files in a source folder. The list should be ordered by size, and it should return the size and the last modification date of each element. Also, a count of the files and the total size should be provided.
- the app should provide a CLI to choose the folder and output the result.
- the app should provide an API + UI to choose the folder and show the result.
- the app should be production-ready. (no need to do everything,  you could list what left to be prod-ready)
- the app should run on a unix system (i.e. ubuntu)

### How to run server
This software needs Golang installed, if you don't have it, please install it following the instruction found here:

`https://golang.org/`

Once you have it, you could run it typing the following:

Go to server folder: `cd server`

* Run CLI `go run main.go cli`
* Run API `go run main.go api` or just `go run main.go`

Or if you prefer you can compile the file and then to run it:

* Compile for linux amd64 (create `bin` folder if it doesn't exist): `./makeBin.sh`

* Run CLI: `./bin/server cli`
* Run API: `./bin/server api`

### How to run unit tests on server
This software is far from being well tested but I included only some cases because the time short that I have.

To run all test:

In server folder: `go test ./...`

### How to run UI
To run UI, you need to install all dependencies so first type:
`npm install --prefix ui`

When all the dependecies are installed you need to type:
`npm start --prefix ui`

### How to run unit test
You have to first install dependencies from the previous step, and then type:
`npm test --prefix ui`

### How to run with docker
You can run each service independetly but you can run both with `docker-compose up`. It is a good practice build them first with `docker-compose build` and everytime that there is a change. Usually `docker-compose up` build them for you the first time but if you change anything this last instruction doesn't re-build the images.


## Get production ready

This software is so simple and not very useful for deploy it on the cloud. It makes more sense for using locally in a computer.
For both cases this software need the following:
* Write more unit tests (server and ui). It has some tests to show a bit how it works but it is far from been enough. We should
run the test with coverage so depende on the minimum coverage set in the company it helps to know when is 'enough'. We should not think that 100% coverage is a bug free software!
* In UI would be great to use a librery like bootstrap for React to get a prettier looking with almost zero efford.
* It needs to set some kind of authentication and authorization mechanism: Oauth2, jwt, challenge, etc... it depends of the result of our risk analisys and how strong our security needs to be set.

### Deploy this software:

1.- A good idea is to have a domain like `https://www.myfilecounter.ca` It is not strictly necesary though. If we decide
to go with one, we can have something like `https://www.myfilecounter.ca` for UI and `https://api.myfilecounter.ca` for API. We should hadler secure connections we need to set a certificate to make it `https` that allows us a encrypted communication between UI and the API. Cloud services make this process usually easy.

2.- We need to decide the architecture, for example I know this software is going to be a complete success :-) and thousands of
users are going to connect so we need to set something like: load balancer to distribute the work load among different instances of this software. A container orchestration service like AWS-Elastic Container Services or Kubernetes to handle those containers and scale them up (incrising demand) of down (decrese demand). We need a container registry where we deploy de Docker images like AWS-Elastic Container Registry, from where our orchestration solution is going to pick the Docker images. For this example we don't need database or broker message. If this software were more complex we can set some parts async for improving performance.

In this matter to define and create the architecture we can create in our cloud provider or use a solution like Terraform that
helps us to handle it.



