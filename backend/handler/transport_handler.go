package handler

import (
	"errors"

	"pelacakan-fruit-transport/model"
	"pelacakan-fruit-transport/repository"
	ws "pelacakan-fruit-transport/websocket"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TransportHandler contains HTTP handlers for transport endpoints.
type TransportHandler struct {
	repo repository.TransportRepository
}

// NewTransportHandler returns a new handler wired to the given repository.
func NewTransportHandler(repo repository.TransportRepository) *TransportHandler {
	return &TransportHandler{repo: repo}
}

// CreateTransport handles POST /api/v1/transports.
func (h *TransportHandler) CreateTransport(c *fiber.Ctx) error {
	var req model.CreateTransportRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.APIResponse{
			Success: false,
			Message: "Invalid request body",
		})
	}

	if req.NomorTruk == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.APIResponse{
			Success: false,
			Message: "nomor_truk is required",
		})
	}
	if req.IDTPH == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.APIResponse{
			Success: false,
			Message: "id_tph is required",
		})
	}
	if req.BeratEstimasi <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(model.APIResponse{
			Success: false,
			Message: "berat_estimasi must be greater than 0",
		})
	}

	transport := &model.Transport{
		NomorTruk:     req.NomorTruk,
		IDTPH:         req.IDTPH,
		BeratEstimasi: req.BeratEstimasi,
		Status:        "DALAM_PERJALANAN",
	}

	if err := h.repo.Create(transport); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.APIResponse{
			Success: false,
			Message: "Failed to create transport record",
		})
	}

	ws.NotifyChange()

	return c.Status(fiber.StatusCreated).JSON(model.APIResponse{
		Success: true,
		Message: "Transport created successfully",
		Data:    transport,
	})
}

// GetAllTransports handles GET /api/v1/transports.
func (h *TransportHandler) GetAllTransports(c *fiber.Ctx) error {
	transports, err := h.repo.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.APIResponse{
			Success: false,
			Message: "Failed to retrieve transport records",
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.APIResponse{
		Success: true,
		Data:    transports,
	})
}

// DeleteTransport handles DELETE /api/v1/transports/:id.
func (h *TransportHandler) DeleteTransport(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.APIResponse{
			Success: false,
			Message: "Invalid transport ID format",
		})
	}

	if err := h.repo.Delete(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(model.APIResponse{
				Success: false,
				Message: "Transport not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(model.APIResponse{
			Success: false,
			Message: "Failed to delete transport record",
		})
	}

	ws.NotifyChange()

	return c.Status(fiber.StatusOK).JSON(model.APIResponse{
		Success: true,
		Message: "Transport deleted successfully",
	})
}

// UpdateStatus handles PATCH /api/v1/transports/:id/status.
func (h *TransportHandler) UpdateStatus(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.APIResponse{
			Success: false,
			Message: "Invalid transport ID format",
		})
	}

	var req model.UpdateStatusRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.APIResponse{
			Success: false,
			Message: "Invalid request body",
		})
	}

	if req.Status != "DALAM_PERJALANAN" && req.Status != "DITERIMA" {
		return c.Status(fiber.StatusBadRequest).JSON(model.APIResponse{
			Success: false,
			Message: "status must be DALAM_PERJALANAN or DITERIMA",
		})
	}

	if err := h.repo.UpdateStatus(id, req.Status); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(model.APIResponse{
				Success: false,
				Message: "Transport not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(model.APIResponse{
			Success: false,
			Message: "Failed to update transport status",
		})
	}

	updated, err := h.repo.FindByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.APIResponse{
			Success: false,
			Message: "Status updated but failed to retrieve updated record",
		})
	}

	ws.NotifyChange()

	return c.Status(fiber.StatusOK).JSON(model.APIResponse{
		Success: true,
		Message: "Transport status updated successfully",
		Data:    updated,
	})
}
