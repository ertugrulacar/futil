package fujwt

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func ToFiberHandler(signinMethod string, signinKey interface{}, errorHandler fiber.ErrorHandler, successHandler fiber.Handler) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: signinMethod,
			Key:    signinKey,
		},
		ErrorHandler:   errorHandler,
		SuccessHandler: successHandler,
	})
}
