package entities

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type Project struct {
	ID      int64  `json:"id" gorm:"type:bigint unsigned;comment:ID"`
	Name    string `json:"name" gorm:"type:varchar(255);not null;uniqueIndex:uniq_name;comment:Name"`
	Intro   string `json:"intro" gorm:"type:text;comment:Intro"`
	Status  int    `json:"status" gorm:"type:tinyint unsigned;comment:Status"`
	OwnerBy string `json:"owner_by" gorm:"type:varchar(255);comment:OwnerBy"`

	CreatedAt time.Time             `json:"created_at,omitempty" gorm:"created_at;comment:CreatedAt"`
	UpdatedAt time.Time             `json:"updated_at,omitempty" gorm:"updated_at;comment:UpdatedAt"`
	DeletedAt soft_delete.DeletedAt `json:"-" gorm:"not null;uniqueIndex:uniq_name;comment:DeletedAt"`
}
