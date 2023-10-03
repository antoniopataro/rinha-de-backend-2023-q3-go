package people

import (
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type createRequest struct {
	Birthdate string   `json:"nascimento"`
	Name      *string  `json:"nome"`
	Nickname  *string  `json:"apelido"`
	Stack     []string `json:"stack"`
}

func (controller *Controller) Create(c *fiber.Ctx) error {
	dto := createRequest{}

	if err := c.BodyParser(&dto); err != nil {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "bad request"})
	}

	if err := func() error {
		if dto.Name == nil {
			return errors.New("missing name")
		}

		if len(*dto.Name) > 100 {
			return errors.New("invalid name")
		}

		if dto.Nickname == nil {
			return errors.New("missing nickname")
		}

		if len(*dto.Nickname) > 32 {
			return errors.New("invalid nickname")
		}

		if _, err := time.Parse("2006-01-02", dto.Birthdate); err != nil {
			return errors.New("invalid birthdate")
		}

		for _, stack := range dto.Stack {
			if len(stack) > 32 {
				return fmt.Errorf("invalid stack: %s", stack)
			}
		}

		return nil
	}(); err != nil {
		return c.
			Status(fiber.StatusUnprocessableEntity).
			JSON(fiber.Map{"error": err.Error()})
	}

	person, err := controller.repository.Create(
		dto.Birthdate,
		*dto.Name,
		*dto.Nickname,
		dto.Stack,
	)

	if err != nil {
		return c.
			Status(fiber.StatusUnprocessableEntity).
			JSON(fiber.Map{"error": err.Error()})
	}

	c.Set(fiber.HeaderLocation, fmt.Sprintf("/pessoas/%s", person.ID))

	return c.
		Status(fiber.StatusCreated).
		JSON(person)
}

func (controller *Controller) Count(c *fiber.Ctx) error {
	count, err := controller.repository.Count()

	if err != nil {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	return c.
		Status(fiber.StatusOK).
		JSON(fiber.Map{"count": count})
}

type findRequest struct {
	ID string `binding:"required" uri:"id"`
}

func (controller *Controller) Find(c *fiber.Ctx) error {
	dto := findRequest{
		ID: c.Params("id"),
	}

	person, err := controller.repository.Find(
		dto.ID,
	)

	if err != nil {
		return c.
			Status(fiber.StatusNotFound).
			JSON(fiber.Map{"error": err.Error()})
	}

	return c.
		Status(fiber.StatusOK).
		JSON(person)
}

type searchRequest struct {
	t string `query:"t"`
}

func (controller *Controller) Search(c *fiber.Ctx) error {
	dto := searchRequest{
		t: c.Query("t"),
	}

	if len(dto.t) == 0 {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "missing query"})
	}

	people, err := controller.repository.Search(dto.t)

	if err != nil {
		return c.
			Status(fiber.StatusNotFound).
			JSON(fiber.Map{"error": err.Error()})
	}

	return c.
		Status(fiber.StatusOK).
		JSON(people)
}

type Controller struct {
	repository *Repository
}

func MakeController(repository *Repository) *Controller {
	return &Controller{
		repository: repository,
	}
}
