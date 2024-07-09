package model

import (
	"database/sql"
	"time"
)

type Movie struct {
	ID 				int 						`gorm:"primary_key;column:id"`
	CreatedAt time.Time 			`gorm:"column:created_at"`
	UpdatedAt time.Time 			`gorm:"column:updated_at"`
	Title 		string 					`gorm:"column:title"`
	Overview 	sql.NullString 	`gorm:"column:overview"`
}
