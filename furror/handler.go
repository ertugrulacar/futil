package furror

import (
	"github.com/gofiber/fiber/v2"
)

func DefaultHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		switch e := err.(type) {
		case *furror:
			return ctx.Status(e.Code).JSON(e)
		case *fiber.Error:
			return ctx.Status(e.Code).JSON(e)
		default:
			return ctx.Status(fiber.StatusInternalServerError).
				JSON(New("Unknown Error!", fiber.StatusInternalServerError))
		}
	}
}
