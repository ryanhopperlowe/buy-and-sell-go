package model

import (
	"database/sql"
	"time"
)

type Identifier = uint

type Model struct {
	ID        Identifier   `gorm:"primarykey" json:"id"`
	CreatedAt time.Time    `json:"-"`
	UpdatedAt time.Time    `json:"-"`
	DeletedAt sql.NullTime `gorm:"index" json:"-"`
}
