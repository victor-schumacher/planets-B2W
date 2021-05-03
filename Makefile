build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd
test:
	go test ./...
run:
	go run ./cmd/main.go