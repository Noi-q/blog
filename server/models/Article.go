package models

import (
	"time"
)

// 文章
type Article struct {
	ID         uint      `json:"id" gorm:"type:int(11);primaryKey;not null;autoIncrement;column:id"`
	Title      string    `json:"title" gorm:"varchar(255)"`
	Content    string    `json:"content" gorm:"varchar(255)"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime:milli"`
	Category   Category  `json:"category"`
	CategoryId uint      `json:"category_id"`
}

func (receiver Article) TableName() string {
	return "blog_article"
}
