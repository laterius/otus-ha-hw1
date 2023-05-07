package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laterius/service_architecture_hw3/app/internal/service"
	"net/http"
)

func NewPostProfile(r service.UserRememberReader, pu service.UserPartialUpdater) *postProfileHandler {
	return &postProfileHandler{
		reader:         r,
		partialUpdater: pu,
	}
}

type postProfileHandler struct {
	reader         service.UserRememberReader
	partialUpdater service.UserPartialUpdater
}

func (h *postProfileHandler) Handle() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var u service.UserPartialUpdate
		err := ctx.BodyParser(&u)
		if err != nil {
			return fail(ctx, err)
		}

		rememberToken := ctx.Cookies("remember_token")
		if rememberToken == "" {
			return ctx.SendStatus(http.StatusUnauthorized)
		}

		currentUser, err := h.reader.ByRemember(rememberToken)
		if err != nil {
			return fail(ctx, err)
		}

		updatedUser, err := h.partialUpdater.PartialUpdate(currentUser.Id, u.ToDomain())
		if err != nil {
			return fail(ctx, err)
		}

		return ctx.Render("profile", fiber.Map{
			"FirstName": updatedUser.FirstName,
			"LastName":  updatedUser.LastName,
			"Username":  updatedUser.Username,
			"Phone":     updatedUser.Phone,
			"Email":     updatedUser.Email,
			"Age":       updatedUser.Age,
			"Gender":    updatedUser.Gender,
			"Hobby":     updatedUser.Hobby,
			"City":      updatedUser.City,
		})
	}
}
