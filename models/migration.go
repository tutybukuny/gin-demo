package models

import "time"

type Migration struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"size:100;unique"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func (Migration) TableName() string {
	return "migration"
}
