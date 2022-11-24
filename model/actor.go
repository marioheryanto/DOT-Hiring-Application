package model

type Actor struct {
	Id      int    `json:"-" gorm:"primary_key"`
	Name    string `json:"name"`
	Gender  string `json:"gender"`
	MovieID uint   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
