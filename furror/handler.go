package furror

import "github.com/gofiber/fiber/v2"

func DefaultHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		e, ok := err.(*furror)
		if !ok {
			return ctx.
				Status(fiber.StatusInternalServerError).
				JSON(New("Unknown Error!", fiber.StatusInternalServerError))
		}
		return ctx.Status(e.Code).JSON(e)
	}
}
