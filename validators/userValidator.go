package validators

import (
	"github.com/Pratham-Mishra04/interact/helpers"
	"github.com/Pratham-Mishra04/interact/initializers"
	"github.com/Pratham-Mishra04/interact/models"
	"github.com/Pratham-Mishra04/interact/schemas"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UserCreateValidator(c *fiber.Ctx) error {
	var reqBody schemas.UserCreateSchema

	if err := c.BodyParser(&reqBody); err != nil {
		return &fiber.Error{Code: 400, Message: "Invalid Req Body"}
	}

	if err := helpers.Validate[schemas.UserCreateSchema](reqBody); err != nil {
		return err
	}

	if reqBody.Password != reqBody.ConfirmPassword {
		return &fiber.Error{Code: 400, Message: "Passwords do not match."}
	}

	// add phone number validator

	var user models.User
	initializers.DB.First(&user, "email = ?", reqBody.Email)

	if user.ID != uuid.Nil {
		return &fiber.Error{Code: 400, Message: "User with this Email ID already exists"}
	}

	initializers.DB.First(&user, "username = ?", reqBody.Username)

	if user.ID != uuid.Nil {
		return &fiber.Error{Code: 400, Message: "User with this Username already exists"}
	}

	return c.Next()
}
