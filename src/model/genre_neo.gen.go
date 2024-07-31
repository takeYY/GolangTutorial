// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameGenreNeo = "genre_neo"

// GenreNeo mapped from table <genre_neo>
type GenreNeo struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	Code      string    `gorm:"column:code;not null" json:"code"`
	Name      string    `gorm:"column:name;not null" json:"name"`
	MovieID   int32     `gorm:"column:movie_id;not null" json:"movie_id"`
}

// TableName GenreNeo's table name
func (*GenreNeo) TableName() string {
	return TableNameGenreNeo
}