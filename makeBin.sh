# Build project for server linux.

mkdir bin
GOOS=linux GOARCH=amd64 go build -o bin/folderReader src/main.go
echo "Compilation for Linux amd64, done!"
