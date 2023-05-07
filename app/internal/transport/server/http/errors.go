package http

import (
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
	"github.com/laterius/service_architecture_hw3/app/internal/domain"
	"net/http"
)

type Error interface {
	HttpCode() int
}

type errMethodNotAllowed struct {
	error
}

func (e errMethodNotAllowed) HttpCode() int {
	return fiber.StatusMethodNotAllowed
}
func (e errMethodNotAllowed) Error() string {
	return "method not allowed"
}

var (
	ErrMethodNotAllowed = errMethodNotAllowed{}
)

func json(c *fiber.Ctx, data interface{}) error {
	return c.Status(http.StatusOK).JSON(data)
}

func created(c *fiber.Ctx, data interface{}, entityId int64) error {
	origPath := c.OriginalURL()
	if origPath[:len(origPath)-1] != "/" {
		origPath += "/"
	}
	c.Set(fiber.HeaderLocation, fmt.Sprintf("%s%d", origPath, entityId))
	return c.Status(http.StatusCreated).JSON(data)
}

func fail(c *fiber.Ctx, err error) error {
	code := codeByErr(err)
	return c.Status(code).JSON(fiber.Map{
		"code":    code,
		"message": err.Error(),
	})
}

func codeByErr(err error) int {
	if _, ok := err.(Error); ok {
		return err.(Error).HttpCode()
	}

	if _, ok := err.(validation.Errors); ok {
		return fiber.StatusBadRequest
	}

	if errors.Is(err, domain.ErrInvalidUserId) || errors.Is(err, domain.ErrInvalidPassHash) {
		return fiber.StatusBadRequest
	}
	if errors.Is(err, domain.ErrUserNotFound) {
		return fiber.StatusNotFound
	}

	return fiber.StatusInternalServerError
}

func DefaultResponse(ctx *fiber.Ctx) error {
	return fail(ctx, ErrMethodNotAllowed)
}
