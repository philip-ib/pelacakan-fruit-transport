package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Status constants for Transport.
const (
	StatusDalamPerjalanan = "DALAM_PERJALANAN"
	StatusDiterima        = "DITERIMA"
)

// Transport represents a fruit transport record from TPH to PKS.
type Transport struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	NomorTruk     string    `gorm:"type:varchar(100);not null" json:"nomor_truk"`
	IDTPH         string    `gorm:"type:varchar(50);not null" json:"id_tph"`
	BeratEstimasi float64   `gorm:"type:decimal(10,2);not null" json:"berat_estimasi"`
	Status        string    `gorm:"type:varchar(50)" json:"status"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// TableName overrides the default table name to "transports".
func (Transport) TableName() string {
	return "transports"
}

// BeforeCreate hook ensures the transport has a UUID and a status before being persisted.
func (t *Transport) BeforeCreate(tx *gorm.DB) error {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	if t.Status == "" {
		t.Status = StatusDalamPerjalanan
	}
	return nil
}

// --- Request / Response DTOs ---

// CreateTransportRequest is the expected JSON body for creating a new transport.
type CreateTransportRequest struct {
	NomorTruk     string  `json:"nomor_truk"`
	IDTPH         string  `json:"id_tph"`
	BeratEstimasi float64 `json:"berat_estimasi"`
}

// UpdateStatusRequest is the expected JSON body for updating a transport status.
type UpdateStatusRequest struct {
	Status string `json:"status"`
}

// APIResponse is a standard JSON response envelope.
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
