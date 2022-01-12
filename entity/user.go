package entity

type User struct {
	ID        string `gorm:"primary_key" json:"id"`
	Firstname string `gorm:"type:varchar(255);NOT NULL" json:"firstname" binding:"required"`
	Lastname  string `gorm:"type:varchar(255)" json:"lastname"`
}
