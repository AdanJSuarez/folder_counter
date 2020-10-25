# Build project for server linux.

mkdir server/bin
GOOS=linux GOARCH=amd64 go build -o server/bin/server server/main.go
echo "Compilation for Linux amd64, done!"
