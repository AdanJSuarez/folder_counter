# Folder Counter

## Introduction
This app is able to find folders and files and order them by size.

Instructions:

The app should list the folders and files in a source folder. The list should be ordered by size, and it should return the size and the last modification date of each element. Also, a count of the files and the total size should be provided.
- the app should provide a CLI to choose the folder and output the result.
- the app should provide an API + UI to choose the folder and show the result.
- the app should be production-ready. (no need to do everything,  you could list what left to be prod-ready)
- the app should run on a unix system (i.e. ubuntu)

## How to run server
This software needs Golang installed, if you don't have it, please install it following the instruction found here:

`https://golang.org/`

Once you have it, you could run it typing the following:

`go run server/main.go`

Or if you prefer you can compile the file and then to run it:

* Compile for linux amd64 (create `bin` folder if it doesn't exist): `./makeBin.sh`

* Run CLI: `./bin/server cli`

* Run API: `./bin/server api`

## How to run unit tests on server
This software is far from being well tested but I included only some cases because the time short that I have.

To run all test:

`go test ./...`

## How to run UI
To run UI from project folder, you need to install all dependencies so first type:
`npm install --prefix ui`

When all the dependecies are installed you need to type:
`npm start --prefix ui`

## How to run unit test
You have to first install dependencies from the previous step, and then type:
`npm test --prefix ui`
