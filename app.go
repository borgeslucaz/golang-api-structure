package main

import (
	"net/http"
    "github.com/go-pg/pg"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/borgeslucaz/golang-api-structure/user"
	"github.com/borgeslucaz/golang-api-structure/auth"
	"fmt"
)

type dbLogger struct { }
func (d dbLogger) BeforeQuery(q *pg.QueryEvent) {}
func (d dbLogger) AfterQuery(q *pg.QueryEvent) {
	fmt.Println(q.FormattedQuery())
}

func main() {
	// Postgres Start
	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "golang",
		Password: "golang",
		Database: "golang",
	})
	db.AddQueryHook(dbLogger{})

	defer db.Close()

	// Start repos
	userRepo := user.StartRepository(db)
	user.UserService = user.NewService(userRepo)

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	user.Routes(e)
	auth.Routes(e)

	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
