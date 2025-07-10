package handlers

import (
	"github.com/lareii/siker.im/internal/models"
	"github.com/lareii/siker.im/internal/services"
	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

type URLHandler struct {
	service *services.URLService
	logger  *zap.Logger
}

func NewURLHandler(service *services.URLService, logger *zap.Logger) *URLHandler {
	return &URLHandler{
		service: service,
		logger:  logger,
	}
}

func (h *URLHandler) CreateURL(c fiber.Ctx) error {
	var req models.CreateURLRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.TargetURL == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Target URL is required",
		})
	}

	url, err := h.service.CreateURL(c.Context(), &req)
	if err != nil {
		h.logger.Error("Failed to create short URL", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(url)
}

func (h *URLHandler) GetURL(c fiber.Ctx) error {
	param := c.Params("param")
	if param == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "URL slug or ID is required",
		})
	}

	id, err := bson.ObjectIDFromHex(param)
	var url *models.URLResponse
	if err != nil {
		url, err = h.service.GetURLBySlug(c.Context(), param)
	} else {
		url, err = h.service.GetURLByID(c.Context(), id)
	}

	if err != nil {
		h.logger.Error("Failed to get URL", zap.Error(err))
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(url)
}

func (h *URLHandler) DeleteURL(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "URL ID is required",
		})
	}

	if err := h.service.DeleteURL(c.Context(), id); err != nil {
		h.logger.Error("Failed to delete URL", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete URL",
		})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}

func (h *URLHandler) RedirectURL(c fiber.Ctx) error {
	slug := c.Params("slug")
	if slug == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "URL slug is required",
		})
	}

	url, err := h.service.GetTargetURL(c.Context(), slug)
	if err != nil {
		h.logger.Error("Failed to get URL", zap.Error(err))
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "URL not found",
		})
	}

	return c.Redirect().To(url)
}
