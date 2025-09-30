package restful

import (
	"github.com/faujiahmat/zentra-order-service/src/common/errors"
	"github.com/gofiber/fiber/v2"
)

func HandleResponseError(c *fiber.Ctx, err *errors.Response) error {
	return c.Status(int(err.HttpCode)).JSON(fiber.Map{
		"errors": err.Message,
	})
}
