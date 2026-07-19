package router

import (
	"pelacakan-fruit-transport/handler"
	ws "pelacakan-fruit-transport/websocket"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// SetupRoutes configures CORS middleware and registers all API routes.
func SetupRoutes(app *fiber.App, h *handler.TransportHandler) {
	// CORS middleware — allows Flutter Web access from any origin during development.
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PATCH,DELETE,OPTIONS",
		AllowHeaders: "Origin,Content-Type,Accept,Authorization",
	}))

	// WebSocket endpoint — real-time notifikasi ke dashboard.
	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		ws.DefaultHub.HandleConnection(c)
	}))

	// API v1 route group
	v1 := app.Group("/api/v1")

	// Transport endpoints
	v1.Post("/transports", h.CreateTransport)
	v1.Get("/transports", h.GetAllTransports)
	v1.Patch("/transports/:id/status", h.UpdateStatus)
	v1.Delete("/transports/:id", h.DeleteTransport)
}
