package handler

import (
	"errors"

	"pelacakan-fruit-transport/model"
	"pelacakan-fruit-transport/repository"
	ws "pelacakan-fruit-transport/websocket"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// TransportHandler contains HTTP handlers for transport endpoints.
type TransportHandler struct {
	repo repository.TransportRepository
}

// NewTransportHandler returns a new handler wired to the given repository.
func NewTransportHandler(repo repository.TransportRepository) *TransportHandler {
	return &TransportHandler{repo: repo}
}

// --- Response / request helpers ---

// respond writes the standard API envelope.
func respond(c *fiber.Ctx, status int, success bool, message string, data interface{}) error {
	return c.Status(status).JSON(model.APIResponse{
		Success: success,
		Message: message,
		Data:    data,
	})
}

// parseID parses the :id path param as a UUID; on failure it writes a 400 and returns err.
func parseID(c *fiber.Ctx) (uuid.UUID, error) {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return uuid.Nil, respond(c, fiber.StatusBadRequest, false, "Invalid transport ID format", nil)
	}
	return id, nil
}

// parseBody parses the request body into out; on failure it writes a 400 and returns err.
func parseBody(c *fiber.Ctx, out interface{}) error {
	if err := c.BodyParser(out); err != nil {
		return respond(c, fiber.StatusBadRequest, false, "Invalid request body", nil)
	}
	return nil
}

// handleRepoError maps repository errors to HTTP responses (404 for ErrNotFound, else 500).
func handleRepoError(c *fiber.Ctx, err error, notFoundMsg, internalMsg string) error {
	if errors.Is(err, repository.ErrNotFound) {
		return respond(c, fiber.StatusNotFound, false, notFoundMsg, nil)
	}
	return respond(c, fiber.StatusInternalServerError, false, internalMsg, nil)
}

// CreateTransport handles POST /api/v1/transports.
func (h *TransportHandler) CreateTransport(c *fiber.Ctx) error {
	var req model.CreateTransportRequest
	if err := parseBody(c, &req); err != nil {
		return err
	}

	if req.NomorTruk == "" {
		return respond(c, fiber.StatusBadRequest, false, "nomor_truk is required", nil)
	}
	if req.IDTPH == "" {
		return respond(c, fiber.StatusBadRequest, false, "id_tph is required", nil)
	}
	if req.BeratEstimasi <= 0 {
		return respond(c, fiber.StatusBadRequest, false, "berat_estimasi must be greater than 0", nil)
	}

	transport := &model.Transport{
		NomorTruk:     req.NomorTruk,
		IDTPH:         req.IDTPH,
		BeratEstimasi: req.BeratEstimasi,
		Status:        model.StatusDalamPerjalanan,
	}

	if err := h.repo.Create(transport); err != nil {
		return respond(c, fiber.StatusInternalServerError, false, "Failed to create transport record", nil)
	}

	ws.NotifyChange()

	return respond(c, fiber.StatusCreated, true, "Transport created successfully", transport)
}

// GetAllTransports handles GET /api/v1/transports.
func (h *TransportHandler) GetAllTransports(c *fiber.Ctx) error {
	transports, err := h.repo.FindAll()
	if err != nil {
		return respond(c, fiber.StatusInternalServerError, false, "Failed to retrieve transport records", nil)
	}

	return respond(c, fiber.StatusOK, true, "", transports)
}

// DeleteTransport handles DELETE /api/v1/transports/:id.
func (h *TransportHandler) DeleteTransport(c *fiber.Ctx) error {
	id, err := parseID(c)
	if err != nil {
		return err
	}

	if err := h.repo.Delete(id); err != nil {
		return handleRepoError(c, err, "Transport not found", "Failed to delete transport record")
	}

	ws.NotifyChange()

	return respond(c, fiber.StatusOK, true, "Transport deleted successfully", nil)
}

// UpdateStatus handles PATCH /api/v1/transports/:id/status.
func (h *TransportHandler) UpdateStatus(c *fiber.Ctx) error {
	id, err := parseID(c)
	if err != nil {
		return err
	}

	var req model.UpdateStatusRequest
	if err := parseBody(c, &req); err != nil {
		return err
	}

	if req.Status != model.StatusDalamPerjalanan && req.Status != model.StatusDiterima {
		return respond(c, fiber.StatusBadRequest, false,
			"status must be DALAM_PERJALANAN or DITERIMA", nil)
	}

	updated, err := h.repo.UpdateStatus(id, req.Status)
	if err != nil {
		return handleRepoError(c, err, "Transport not found", "Failed to update transport status")
	}

	ws.NotifyChange()

	return respond(c, fiber.StatusOK, true, "Transport status updated successfully", updated)
}
