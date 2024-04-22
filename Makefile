run: build
	@./go-jwt

build:
	goose -dir migrations postgres "host=localhost port=5433 user=gojwt_user password=gojwt_password dbname=gojwt_db sslmode=disable" up
	@go build -o go-jwt ./cmd/go-jwt

test:
	@go test -v ./...

reset:
	goose -dir migrations postgres "host=localhost port=5433 user=gojwt_user password=gojwt_password dbname=gojwt_db sslmode=disable" reset