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

In the server folder:

* Run CLI `go run main.go cli`
* Run API `go run main.go api` or just `go run main.go`

Or if you prefer you can compile the file and then to run it:

* Compile for linux amd64 (create `bin` folder if it doesn't exist): `./makeBin.sh`

* Run CLI: `./bin/server cli`
* Run API: `./bin/server api`

### How to run unit tests on server

To run all test:

In the server folder: `go test ./...`

### How to run UI
To run UI, you need to install all dependencies so first type:
`npm install --prefix ui`

When all the dependencies are installed you need to type:
`npm start --prefix ui`

### How to run unit test
You have to first install dependencies from the previous step, and then type:
`npm test --prefix ui`

### How to run with docker
You can run each service independently but you can run both with `docker-compose up`. It is a good practice to build them first with `docker-compose build` and everytime that there is a change. Usually `docker-compose up` builds them for you the first time but if you change anything this last instruction doesn't re-build the images.


## Get production ready

I can see two ways to get this software production ready. One is SaaS where we deploy the UI and the API(server) where a user connects through the UI, and the second one is On-Prem where the user can download it and run it on his computer… it could be both too!

For the second case we have different options, one simple solution can be to include in our UI a button to download the binary for different platforms like Windows, Linux and Mac. I didn’t include this button in the UI because it wasn’t in the specification. We can store the software already compiled in a S3 bucket and when the user clicks on that button we serve that file from the S3.

In my humble opinion, this software is not very useful as SaaS because it (the server) is going to read only the folders that it has access to, if we deploy it in Docker containers, it will access only folders into its container and not the user folders. It makes more sense to use it locally as on-prem.

At this moment, before to deploy it, this software -API(server) and UI need more work as follow:

* Write more unit tests (server and ui). It has some tests to show a bit how it works but it is far from being enough. We should run the test with coverage so depende on the minimum coverage set by the company we know when is 'enough'. We should not think that 100% coverage is a bug free software whatsoever!
* In UI would be great to use a library like bootstrap for React to get a prettier look without needing to work much on html and css.
* Both (API and UI) need to set some kind of security system (authentication-authorization-encryption): Oauth2, jwt, challenge, https, etc... The final security solution depends on the result of our risk analysis and then how strong our security needs to be set.

### Deploy this software:

There are many different ways to deploy this software like running on Docker on the cloud, running directly in a cloud virtual machine, the same but in our servers, etc. I am going to explain one case for simplicity that is going to be running on the cloud on Docker.

1.- A good idea is to have a domain like `https://www.myfilecounter.ca` It is not strictly necessary though. If we decide to go with one, we can have something like `https://www.myfilecounter.ca` for UI and `https://api.myfilecounter.ca` for API. We should hadler secure connections so we need to set a certificate to make it `https` that allows us encrypted communication between UI and the API. Cloud services make this process usually easy.

2.- We need to decide the architecture, for example, I know this software is going to be a complete success :-) and thousands of users are going to connect to it so we need to set a
Load balancer to distribute the workload among different instances of this software.
A container orchestration service like AWS-Elastic Container Services or Kubernetes to handle those containers and scale them up (increasing demand) or down (decreasing demand).
We need a container registry where we deploy the Docker images like AWS-Elastic Container Registry, from where our orchestration solution is going to pick the Docker images.
We need S3 buckets for binaries and if we decide to go with something like AWS Amplify.

In this matter, to define and create the services included in the architecture, we can create them in our cloud provider or use a solution like Terraform that helps us to handle it.

I recommend deploying the UI independently and, instead of running in a Docker container, deploy it in a service like AWS Amplify where it sets the code as it was static in a S3 bucket -even though it is not- and distributes from there instead of having a server/service running. It is much cheaper and easier to scale.

3.- To deploy the software, if we decide to automate the process, we need to set our pipeline to handle CI/CD. This pipeline is going to get the code from the repo -Github in this case- and set them in our Docker registry and store the binaries on S3. To do so every time it is triggered, it will run all tests that we have set like unit test, integration test -we don’t have integration test because there is only one service-, etc. If everything went well, it generates the deployable artifacts -in our case Docker image and binaries- and sends it to AWS Elastic Container Registry and S3 respectively. Our orchestration services -like AWS ECS- will check periodically if there is a new Docker image version in our container registry and if so, it will create new service instances with the new image and when the new services are running, it will turn off the old services. We usually set a specific branch for deployment -in old days ‘master’, nowadays ‘main’- so everytime we merge on master/main the pipeline will automatically trigger the whole process. For this pipeline we can use different tools like, for example, Github Actions, Codeship, etc.

4.- In our cloud provider we need to set some analysis tools like CloudWatch where we can see how our servers perform, peak of demand, etc. Another interesting tool is a log system that is going to catch all the messages generated by our software that helps us to debug, analyze, etc.







