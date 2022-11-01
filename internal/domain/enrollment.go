package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Enrollment struct {
	ID        string     `json:"id" gorm:"type:char(36);not null;primary_key;unique_index"`
	UserID    string     `json:"user_id,omitempty" gorm:"type:char(36)"`
	User      *User      `json:"user,omitempty"`
	CourseID  string     `json:"course_id" gorm:"type:char(36);not null"`
	Course    *Course    `json:"course,omitempty"`
	Status    string     `json:"status" gorm:"type:char(2)"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
}

func (c *Enrollment) BeforeCreate(tx *gorm.DB) (err error) {

	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return
}
