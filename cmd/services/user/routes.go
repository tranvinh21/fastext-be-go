package user

import "github.com/gofiber/fiber/v2"

type UserHandler struct {
	store *UserStore
}

func SetupUserRoutes(store *UserStore) *UserHandler {
	return &UserHandler{store: store}
}

func (h *UserHandler) RegisterRoutes(app fiber.Router) {
	userRouter := app.Group("/users")
	userRouter.Get("", h.GetUsers)
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.store.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(users)
}
