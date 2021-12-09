package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thsanan/idolist/models"
	"github.com/thsanan/idolist/services"
)

type actdressHandler struct {
	actService services.ActressService
}

func NewActdressHandler(actService services.ActressService) actdressHandler {
	return actdressHandler{actService}
}

func (actHandler actdressHandler) ActAddHandler(c *fiber.Ctx) error {
	actRequest := models.Actdress{}
	c.BodyParser(&actRequest)
	act, err := actHandler.actService.InsertAct(actRequest)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": act,
	})
}
