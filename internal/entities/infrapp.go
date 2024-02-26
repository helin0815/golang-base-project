package entities

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type InfraApp struct {
	ID    int64  `json:"id" gorm:"type:bigint unsigned;comment:ID"`
	Pid   int64  `json:"pid" gorm:"type:bigint unsigned;comment:Pid"`
	Name  string `json:"name" gorm:"type:varchar(255);not null;uniqueIndex:uniq_name;comment:Name"`
	Intro string `json:"intro" gorm:"type:text;comment:Intro"`

	CreatedAt time.Time             `json:"created_at,omitempty" gorm:"created_at;comment:CreatedAt"`
	UpdatedAt time.Time             `json:"updated_at,omitempty" gorm:"updated_at;comment:UpdatedAt"`
	DeletedAt soft_delete.DeletedAt `json:"-" gorm:"not null;uniqueIndex:uniq_name;comment:DeletedAt"`
}
