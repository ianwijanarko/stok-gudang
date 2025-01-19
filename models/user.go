package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Nama     string `gorm:"size:50"`
	Password string `gorm:"size:100"`
	Token    string `gorm:"type:text"`
}
