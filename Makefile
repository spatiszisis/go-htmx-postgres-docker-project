run:
	ps -ef | grep 'go-jwt' | grep -v grep | awk '{print $2}' | xargs -r kill -9 
	make build
	@./go-jwt

build:
	make tailwind-build
	@templ generate view
	goose -dir migrations postgres "host=localhost port=5433 user=gojwt_user password=gojwt_password dbname=gojwt_db sslmode=disable" up
	@go build -o go-jwt ./cmd/go-jwt

test:
	@go test -v ./...

reset:
	goose -dir migrations postgres "host=localhost port=5433 user=gojwt_user password=gojwt_password dbname=gojwt_db sslmode=disable" reset

templ-generate:
	@templ generate view

tailwind-watch:
	npx tailwindcss -i interval/templates/css/styles.css -o public/styles.css --watch

tailwind-build:
	npx tailwindcss -i interval/templates/css/styles.css -o public/styles.css --minify