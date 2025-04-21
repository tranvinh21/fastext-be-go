package middleware

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func init() {
	validate.RegisterValidation("passwd", func(fl validator.FieldLevel) bool {
		password := strings.TrimSpace(fl.Field().String())

		hasLetter := regexp.MustCompile(`[A-Za-z]`).MatchString(password)
		hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)

		return hasLetter && hasNumber
	})
}
func ValidateBody[T any]() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body T

		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid JSON format",
			})
		}

		if err := validate.Struct(body); err != nil {
			var messages []string
			for _, e := range err.(validator.ValidationErrors) {
				field := e.Field()
				tag := e.Tag()

				// Custom error cho tag passwd
				if tag == "passwd" {
					messages = append(messages, fmt.Sprintf("%s must contain at least one letter and one number", field))
					continue
				}

				messages = append(messages, fmt.Sprintf("Field '%s' failed on the '%s' rule", field, tag))
			}

			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"errors": messages,
			})
		}

		// Gán vào context để controller dùng
		c.Locals("body", &body)
		return c.Next()
	}
}
