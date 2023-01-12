swag-gen:
	swag init -g api/api.go -o api/docs

run:
	go run cmd/main.go
