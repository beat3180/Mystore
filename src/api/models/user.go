package models

import "gorm.io/gorm"

// gorm.Modelの中身
// type Model struct {
// 	ID        uint `gorm:"primarykey"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt DeletedAt `gorm:"index"`
// }

type User struct {
	gorm.Model
	NickName string `json:"nick_name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"` // -を指定すると非表示にできる
}
