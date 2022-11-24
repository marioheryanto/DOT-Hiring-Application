package model

type Movie struct {
	Id     int     `json:"ID" gorm:"primary_key"`
	Name   string  `json:"name" gorm:"unique;not null;type:varchar(100);default:null"`
	Genre  string  `json:"genre" gorm:"not null;type:varchar(10);default:null"`
	Year   string  `json:"year" gorm:"not null;type:varchar(4);default:null"`
	Actors []Actor `json:"actors" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
