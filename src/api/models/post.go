package models

import "gorm.io/gorm"

type Post struct {
	UserId     int
	Title      string
	Comment    string
	Lng        float64
	Lat        float64
	Address    string
	ImagePath  string
	PublicFlag bool
	gorm.Model
}
