migration-up:
	migrate -path ./migrations/postgres/ -database 'postgres://postgres:0003@localhost:3003/catalog?sslmode=disable' up

swag-gen:
	swag init -g api/api.go -o api/docs

run:
	go run cmd/main.go

test_Id:
	go test -run TestBookGetById -v

test_Insert:
	go test -run TestBookInsert -v  

