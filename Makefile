run:
	go mod tidy
	clear
	go run main.go

migrate_up:
	migrate -path migrations -database postgres://postgres:postgres@localhost:5432/chat?sslmode=disable -verbose up

migrate_down:
	migrate -path migrations -database postgres://postgres:postgres@localhost:5432/chat?sslmode=disable -verbose down

migrate_force:
	migrate -path migrations -database postgres://postgres:postgres@localhost:5432/chat?sslmode=disable -verbose force 1

migrate_file:
	migrate create -ext sql -dir migrations -seq create_table

swag-gen:
	~/go/bin/swag init -g ./internal/delivery/api.go -o ./internal/delivery/docs force 1