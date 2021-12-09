package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/jmoiron/sqlx"
	"github.com/thsanan/idolist/handlers"
	"github.com/thsanan/idolist/repositories"
	"github.com/thsanan/idolist/services"
)

func main() {

	db, err := sqlx.Open("mysql", "sandland:IntelliP24.X@tcp(203.159.94.79:3306)/idolist")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	engine := html.NewFileSystem(http.Dir("./views"), ".html")
	engine.Reload(true)
	engine.Debug(true)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./views")

	/* app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "*",
		AllowHeaders: "*",
	})) */
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("page_mov_form", fiber.Map{
			"alert": "info",
		}, "index")
	})
	repoAct := repositories.NewActDb(db)
	serviceAct := services.NewActressEvo(repoAct)
	handlerAct := handlers.NewActdressHandler(serviceAct)

	app.Post("/actdress", handlerAct.ActAddHandler)

	app.Listen(":3000")
}
