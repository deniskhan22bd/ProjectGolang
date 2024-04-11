up:
	@echo 'Running up migrations...'
	migrate -path='./pkg/migrations' -database='postgres://postgres:123@localhost/golang_project?sslmode=disable' up

down:
	@echo 'Running down migrations...'
	migrate -path='./pkg/migrations' -database='postgres://postgres:123@localhost/golang_project?sslmode=disable' down
	@echo y

run:
	@echo "Running run server..."
	make down
	make up
	go run ./cmd


migration:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./pkg/migrations ${name}

clear:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
