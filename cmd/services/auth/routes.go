package auth

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/tranvinh21/fastext-be-go/config"
	"github.com/tranvinh21/fastext-be-go/db/schema"
	"github.com/tranvinh21/fastext-be-go/middleware"
	"github.com/tranvinh21/fastext-be-go/request"
	"github.com/tranvinh21/fastext-be-go/response"
	"github.com/tranvinh21/fastext-be-go/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthHandler struct {
	store *AuthStore
}

func SetupAuthRoutes(store *AuthStore) *AuthHandler {
	return &AuthHandler{store: store}
}

func (h *AuthHandler) RegisterRoutes(app fiber.Router) {
	authRouter := app.Group("/auth")
	authRouter.Post("/signup", middleware.ValidateBody[request.SignupRequest](), h.Signup)
	authRouter.Post("/signin", middleware.ValidateBody[request.SigninRequest](), h.Signin)
	authRouter.Post("/signout", h.Signout)
	authRouter.Post("/refresh-token", h.RefreshToken)

}

func (h *AuthHandler) checkUserExists(c *fiber.Ctx, field string, value string) error {
	var user *schema.User
	var err error

	switch field {
	case "email":
		user, err = h.store.GetUserByEmail(value)
	case "name":
		user, err = h.store.GetUserByUsername(value)
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return errors.New("error checking user existence")
	}

	if user != nil {
		message := "Email already exists"
		if field == "name" {
			message = "name already exists"
		}
		return errors.New(message)
	}

	return nil
}

func (h *AuthHandler) Signup(c *fiber.Ctx) error {
	body := c.Locals("body").(*request.SignupRequest)

	// Check if user exists
	if err := h.checkUserExists(c, "email", body.Email); err != nil {
		return response.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}
	if err := h.checkUserExists(c, "name", body.Name); err != nil {
		return response.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return response.ErrorResponse(c, fiber.StatusBadRequest, "Error hashing password")
	}

	// Create user
	user := schema.User{
		Email:    body.Email,
		Password: string(hashedPassword),
		Name:     body.Name,
	}

	if err := h.store.CreateUser(&user); err != nil {
		return response.ErrorResponse(c, fiber.StatusBadRequest, "User not created")
	}

	return response.SuccessResponse(c, fiber.StatusCreated, "User created successfully", nil)
}

func (h *AuthHandler) Signin(c *fiber.Ctx) error {
	body := c.Locals("body").(*request.SigninRequest)
	user, err := h.store.GetUserByEmail(body.Email)
	if err != nil {
		return response.ErrorResponse(c, fiber.StatusBadRequest, "User not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return response.ErrorResponse(c, fiber.StatusBadRequest, "Invalid password")
	}

	accessToken, err := utils.GenerateToken(jwt.MapClaims{
		"userId": user.ID,
	}, config.Envs.JWT.ACCESS_TOKEN_SECRET)
	if err != nil {
		return response.ErrorResponse(c, fiber.StatusBadRequest, "Error generating access token")
	}

	refreshToken, err := utils.GenerateToken(jwt.MapClaims{
		"userId": user.ID,
	}, config.Envs.JWT.REFRESH_TOKEN_SECRET)
	if err != nil {
		return response.ErrorResponse(c, fiber.StatusBadRequest, "Error generating refresh token")
	}

	cookie := fiber.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "lax",
	}
	c.Cookie(&cookie)

	return response.SuccessResponse(c, fiber.StatusOK, "User signed in successfully", fiber.Map{
		"accessToken": accessToken,
	})
}

func (h *AuthHandler) Signout(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
