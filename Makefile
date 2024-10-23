build:
	@go build -o bin/dist cmd/main.go

run: build
	@./bin/dist

# Run tests
test:
	@go test ./... -v

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down

# Notes:

# to install a package run
# go get -u github.com/golang-jwt/jwt/v5

# To create a migration file run 

# 1.create this file
# cmd/migrate/main.go

# 2.then create this foalder
# cmd/migrate/migrations 

# 3.Run at root level
# make migration add-order-table     
# make migration <table name>

# 4.files will be created in the cmd/migrate/migrations foalders write sql code

# 5.To apply migrations run
# make migrate-up      

# 6.to delete migrations run
# make migrate-down