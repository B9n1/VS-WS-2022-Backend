package models

type Todo struct {
	Todo     string `json:"todo" gorm:"primaryKey"`
	Priority int    `json:"priority" gorm:"default:2"`
}
