package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Metadata struct {
	ID        uuid.UUID      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *Metadata) BeforeCreate(tx *gorm.DB) error {
	if u.ID != uuid.Nil {
		return nil
	}
	u.ID = uuid.New()
	return nil
}
