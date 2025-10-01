package model

type User struct {
	ID 							int 			`gorm:"type:int;primaryKey;autoIncrement"`
	Username				string 		`gorm:"type:varchar(50);not null;uniqueIndex"`
	PasswordHash		string 		`gorm:"type:varchar(255);not null"`

	Feed						[]Feed		`gorm:"foreignKey:Userid;references:ID"`
	UserFollower		[]Follow	`gorm:"foreignKey:FollowerID;references:ID"`
	UserFollowed		[]Follow	`gorm:"foreignKey:FollowedID;references:ID"`
}