// +heroku install bitbucket.org/liamstask/goose/cmd/goose ./...

module app

go 1.14

require (
	github.com/flosch/pongo2 v0.0.0-20200529170236-5abacdfa4915
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/jmoiron/sqlx v1.2.0
	github.com/labstack/echo/v4 v4.9.0
	github.com/leodido/go-urn v1.2.0 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0
)
