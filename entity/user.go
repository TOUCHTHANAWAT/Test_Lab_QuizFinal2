package entity

import "time"

type User struct {
	StudentID string `valid:"required~StudentID is required, matches(^[BMD]\\d{7}$)"`
	FirstName string    `valid:"required~FirstName is required"`
	LastName  string    `valid:"required~LastName is required"`
	Email     string    `valid:"required~Email is required,email~Email is invalid "`
	Phone     string    `valid:"required~Phone is required,stringlength(10|10)"`
	Profile   string    `gorm:"type:longtext"`
	Link      string    `valid:"required~Link is required,url~URL is invalid"`
	Password  string   
	BirthDay  time.Time 

	GenderID uint `valid:"required~GenderID is required"`
	Genders []Gender `gorm:"foreignKey:GenderID"`
}
