package main

import (
	"log"
	"os"

	"app/handler"
	"app/repository"

	_ "github.com/go-sql-driver/mysql" // Using MySQL driver
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/go-playground/validator.v9"
)

var db *sqlx.DB
var e = createMux()

func main() {
	// connectDB() の戻り値をグローバル変数に格納
	db = connectDB()
	repository.SetDB(db)

	e.GET("/", handler.ArticleIndex)
	e.GET("/new", handler.ArticleNew)
	e.GET("/:id", handler.ArticleShow)
	e.GET("/:id/edit", handler.ArticleEdit)
	e.POST("/", handler.ArticleCreate)

	// Webサーバーをポート番号 8080 で起動する
	e.Logger.Fatal(e.Start(":8080"))
}

func connectDB() *sqlx.DB {
	dsn := os.Getenv("DSN")
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		e.Logger.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		e.Logger.Fatal(err)
	}
	log.Println("db connection succeeded")
	return db
}

func createMux() *echo.Echo {
	// アプリケーションインスタンスを生成
	e := echo.New()

	// アプリケーションに各種ミドルウェアを設定
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(middleware.CSRF())

	e.Static("/css", "src/css")
	e.Static("/js", "src/js")

	e.Validator = &CustomValidator{validator: validator.New()}

	// アプリケーションインスタンスを返却
	return e
}

// CustomValidator ...
type CustomValidator struct {
	validator *validator.Validate
}

// Validate ...
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
