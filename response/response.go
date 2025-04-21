package response

import "github.com/gofiber/fiber/v2"

func ErrorResponse(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"message": message,
		"error":   true,
	})
}

func SuccessResponse(c *fiber.Ctx, status int, message string, data map[string]interface{}) error {
	response := map[string]interface{}{
		"message": message,
		"error":   false,
	}
	for k, v := range data {
		response[k] = v
	}

	return c.Status(status).JSON(response)
}
