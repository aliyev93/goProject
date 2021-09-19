dev:
	go run web/main.go

build:
	CGO_ENABLED=0 GOOS=linux go build -a -o app