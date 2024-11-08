run: 
	go run main.go

build: 
	go build -ldflags="-s -w" -o build-app main.go

run-dev: 
	nodemon --exec go run main.go --signal SIGTERM

db-generate: 
	go run github.com/steebchen/prisma-client-go generate

db-pull: 
	go run github.com/steebchen/prisma-client-go db pull

db-push: 
	go run github.com/steebchen/prisma-client-go db push