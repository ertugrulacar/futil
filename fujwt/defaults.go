package fujwt

import (
	"github.com/ErtugrulAcar/futil/fuconf"
	"github.com/ErtugrulAcar/futil/furror"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type defaults struct {
}

func Defaults() *defaults {
	return &defaults{}
}

func (d *defaults) ErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		return furror.New("Unauthorized Request!", fiber.StatusUnauthorized)
	}
}

// SuccessHandler returns a default Fiber JWT Success Handler
// jwtKeys parameter can be passed or configured from fuconf.Conf().AddJwtKeys()
func (d *defaults) SuccessHandler(jwtKeys ...string) fiber.Handler {
	var keys []string
	if jwtKeys != nil {
		keys = jwtKeys
	} else {
		keys = fuconf.Conf().GetJwtKeys()
	}

	return func(ctx *fiber.Ctx) error {
		// Set is_authorized as true
		ctx.Locals(fuconf.Conf().GetIsAuthorizedQueryKey(), true)
		// Get the payload of the auth-token
		payload := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
		// Set the jwt payload keys
		for _, key := range keys {
			ctx.Locals(key, payload[key])
		}
		return ctx.Next()
	}
}
