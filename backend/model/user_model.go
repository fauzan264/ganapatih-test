package model

type User struct {
	ID 							int 		`gorm:"type:int;primaryKey;autoIncrement"`
	Username				string 	`gorm:"type:varchar(50);not null;uniqueIndex"`
	PasswordHash		string 	`gorm:"type:varchar(255);not null"`
}