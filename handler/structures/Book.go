package structures

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Book struct {
	ID          uint       `gorm:"not null;primary_key"`
	IDBook      string     `gorm:"not null;uniqueIndex" json:"uid" form:"uid" valid:"required~Unique ID Book Required"`
	Title       string     `gorm:"null" json:"Title" form:"Title" valid:"required~Title Book Required"`
	Author      string     `gorm:"null" json:"Author"`
	Description string     `gorm:"null" json:"Desc"`
	CreatedAt   *time.Time `gorm:"autoCreateTime"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

func (t_Book *Book) BeforeCreate(t_DB *gorm.DB) (err error) {
	_, ErrParam := govalidator.ValidateStruct(t_Book)

	if ErrParam != nil {
		err = ErrParam
		return
	}

	return nil
}

func (t_Book *Book) BeforeUpdate(t_DB *gorm.DB) (err error) {
	_, ErrParam := govalidator.ValidateStruct(t_Book)

	if ErrParam != nil {
		err = ErrParam
		return
	}

	return nil
}
