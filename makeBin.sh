# Build project for server linux.

mkdir server/bin
cd server
GOOS=linux GOARCH=amd64 go build -o bin/server main.go
echo "Compilation for Linux amd64, done!"
