package middlewares

import (
	"crypto/rsa"
	"net/http"
	"nowgoal/pkg/appresponse"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func AuthMiddleware(publicKey *rsa.PublicKey) fiber.Handler {
	return jwtware.New(jwtware.Config{
		AuthScheme:    "Bearer",
		SigningMethod: "RS256",
		SigningKey:    publicKey,
		ErrorHandler:  jwtError,
	})

}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(appresponse.ResponseBody{
			Status: appresponse.Status{
				Code:    http.StatusBadRequest,
				Message: "Sorry, Not responding because of incorrect syntax",
			},
		})
	}
	return c.Status(fiber.StatusUnauthorized).JSON(appresponse.ResponseBody{
		Status: appresponse.Status{
			Code:    http.StatusUnauthorized,
			Message: "Sorry, We are not able to process your request. Please try again",
		},
	})
}
