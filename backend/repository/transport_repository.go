package repository

import (
	"pelacakan-fruit-transport/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TransportRepository defines the contract for transport data access.
type TransportRepository interface {
	Create(transport *model.Transport) error
	FindAll() ([]model.Transport, error)
	FindByID(id uuid.UUID) (*model.Transport, error)
	UpdateStatus(id uuid.UUID, status string) error
	Delete(id uuid.UUID) error
}

// transportRepository is the GORM-backed implementation of TransportRepository.
type transportRepository struct {
	db *gorm.DB
}

// NewTransportRepository returns a new repository instance.
func NewTransportRepository(db *gorm.DB) TransportRepository {
	return &transportRepository{db: db}
}

// Create inserts a new transport record into the database.
func (r *transportRepository) Create(transport *model.Transport) error {
	return r.db.Create(transport).Error
}

// FindAll returns all transport records ordered by newest first.
func (r *transportRepository) FindAll() ([]model.Transport, error) {
	var transports []model.Transport
	err := r.db.Order("created_at DESC").Find(&transports).Error
	return transports, err
}

// FindByID looks up a single transport by its UUID.
func (r *transportRepository) FindByID(id uuid.UUID) (*model.Transport, error) {
	var transport model.Transport
	err := r.db.First(&transport, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &transport, nil
}

// Delete removes a transport record by its UUID.
func (r *transportRepository) Delete(id uuid.UUID) error {
	result := r.db.Delete(&model.Transport{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// UpdateStatus updates the status field of a transport record.
func (r *transportRepository) UpdateStatus(id uuid.UUID, status string) error {
	result := r.db.Model(&model.Transport{}).Where("id = ?", id).Update("status", status)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
