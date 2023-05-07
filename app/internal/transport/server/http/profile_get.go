package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/laterius/service_architecture_hw3/app/internal/domain"
	"github.com/laterius/service_architecture_hw3/app/internal/service"
	"net/http"
)

func NewGetProfile(r service.UserReader, r2 service.UserRememberReader) *getProfileHandler {
	return &getProfileHandler{
		readerRemember: r2,
		readerUser:     r,
	}
}

type getProfileHandler struct {
	readerRemember service.UserRememberReader
	readerUser     service.UserReader
}

func (h *getProfileHandler) Handle() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userId, err := ctx.ParamsInt(UserIdFieldName, 0)
		if err != nil {
			return fail(ctx, err)
		}

		rememberToken := ctx.Cookies("remember_token")
		if rememberToken == "" {
			return ctx.SendStatus(http.StatusUnauthorized)
		}

		currentUser, err := h.readerRemember.ByRemember(rememberToken)
		if err != nil {
			return fail(ctx, err)
		}

		user, err := h.readerUser.Get(domain.UserId(userId))
		if err != nil {
			return fail(ctx, err)
		}

		fmt.Printf("user.Remember = %s, rememberToken = %s, user.RememberHash = %s, user.Id = %s, currentUser.Id = %s",
			user.Remember, rememberToken, user.RememberHash, user.Id, currentUser.Id)

		if user.Id == currentUser.Id {
			return ctx.Render("profile", fiber.Map{
				"FirstName": user.FirstName,
				"LastName":  user.LastName,
				"Username":  user.Username,
				"Phone":     user.Phone,
				"Email":     user.Email,
				"Age":       user.Age,
				"Gender":    user.Gender,
				"Hobby":     user.Hobby,
				"City":      user.City,
				"Token":     user.RememberHash,
			})
		}

		return ctx.SendStatus(http.StatusForbidden)
	}
}
