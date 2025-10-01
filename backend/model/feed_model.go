package model

import "time"

type Feed struct {
	ID 					int 			`gorm:"type:int;primaryKey;autoIncrement"`
	Userid 			int 			`gorm:"type:int"`
	Content 		string 		`gorm:"type:varchar(200)"`
	Createdat		time.Time	`gorm:"type:timestamp"`

	User				User			`gorm:"foreignKey:Userid;references:ID"`
}