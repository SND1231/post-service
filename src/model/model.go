package model

import (
	"time"
)

type Post struct {
	ID        int32
	Title     string `gorm:"size:30"`
	Content   string `gorm:"size:400"`
	PhotoUrl  string `gorm:"size:400"`
	UserId    int32
	LikeCount int32
	Likes     []Like `gorm:"many2many:post_likes;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Like struct {
	ID     int32
	UserId int32
}
