package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/thsanan/idolist/models"
	"github.com/thsanan/idolist/repositories"
	"github.com/thsanan/idolist/services"
)

func main() {

	db, err := sqlx.Open("mysql", "sandland:IntelliP24.X@tcp(203.159.94.79:3306)/idolist")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	app := fiber.New()

	repoAct := repositories.NewActDb(db)
	serviceAct := services.NewActressEvo(repoAct)

	app.Post("/actdress", func(c *fiber.Ctx) error {
		actRequest := models.Actdress{}
		c.BodyParser(&actRequest)
		act, err := serviceAct.InsertAct(actRequest)
		if err != nil {
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"message": act,
		})

	})

	app.Listen(":3000")
}
