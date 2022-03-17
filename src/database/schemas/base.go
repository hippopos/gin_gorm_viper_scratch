package schemas

import "time"

// gorm.Model definition
type Model struct {
	Id        uint      `gorm:"primary_key" json:"id"`
	Code      string    `gorm:"size:100;unique"  json:"code"`
	Name      string    `gorm:"size:100;"  json:"name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
