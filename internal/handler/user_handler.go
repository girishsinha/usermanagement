package handler

import (
	"fmt"
	"strconv"

	"github.com/girishsinha/user-manage/internal/service"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

// POST /users
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	type request struct {
		Name string `json:"name"`
		Dob  string `json:"dob"`
	}

	var body request
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	fmt.Println(body.Dob)
	user, err := h.svc.RegisterUser(c.Context(), body.Name, body.Dob)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(user)
}

// GET /users
func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	// 1. Get pagination parameters from URL (e.g. /users?limit=10&page=1)
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	page, _ := strconv.Atoi(c.Query("page", "1"))
	offset := (page - 1) * limit

	users, err := h.svc.GetAllUsers(c.Context(), int32(limit), int32(offset))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

// GET /users/:id
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	user, err := h.svc.GetUserByID(c.Context(), id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user)
}

// PUT /users/:id
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	type request struct {
		Name string `json:"name"`
		Dob  string `json:"dob"`
	}
	var body request
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	user, err := h.svc.UpdateUser(c.Context(), id, body.Name, body.Dob)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(user)
}

// DELETE /users/:id
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	err := h.svc.DeleteUser(c.Context(), id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "User deleted successfully"})
}
