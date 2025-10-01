package model

import "time"

type Follow struct {
	ID         				int       `gorm:"primaryKey;autoIncrement" json:"id"`
	FollowerID 				int       `gorm:"not null" json:"follower_id"`
	FollowedID 				int       `gorm:"not null" json:"followed_id"`
	CreatedAt  				time.Time `gorm:"autoCreateTime" json:"created_at"`

	UserFollower			User			`gorm:"foreignKey:FollowerID;references:ID"`
	UserFollowed			User			`gorm:"foreignKey:FollowedID;references:ID"`
}