package handlers

import (
	"fmt"
	"nowgoal/internal/core/domain"
	"nowgoal/internal/core/ports"
	"nowgoal/pkg/appresponse"
	"nowgoal/pkg/validator"

	"github.com/gofiber/fiber/v2"
)

type HTTPHandler struct {
	srv       ports.Service
	validator validator.Validator
}

func New(srv ports.Service) *HTTPHandler {
	return &HTTPHandler{
		srv:       srv,
		validator: validator.New(),
	}
}

func (hdl *HTTPHandler) ReadStat(c *fiber.Ctx) error {
	// err = hdl.validator.ValidateStruct(payload)
	// if err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(appresponse.BadRequest)
	// }
	fmt.Print("dssss")
	data, err := hdl.srv.ReadValueFromCSVFile()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(appresponse.InternalServerError)
	}
	fmt.Print("12122121")
	return c.Status(200).JSON(appresponse.ResponseBody{Status: appresponse.Success, Data: data})
}

func (hdl *HTTPHandler) GetStatsPattern1(c *fiber.Ctx) error {
	payload := domain.GetStatsPattern1Request{}
	err := c.BodyParser(&payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(appresponse.BadRequest)
	}
	results, err := hdl.srv.GetStatPattern1(payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(appresponse.InternalServerError)
	}
	return c.Status(fiber.StatusOK).JSON(appresponse.ResponseBody{Status: appresponse.Success, Data: results})
}
