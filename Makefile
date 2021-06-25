all: dev

dev:
	go run main.go

test:
	go test ./...

build:
	go build main.go

deploy:
	gcloud app deploy

log:
	gcloud app logs tail -s default
