package models

import (
	"gorm.io/gorm"
	"time"
)

// 栏目
type Category struct {
	ID          uint           `json:"id" gorm:"type:int(11);primaryKey;not null;autoIncrement"`
	Title       string         `json:"title" gorm:"type:varchar(255);not null;index"`
	Description string         `json:"description" gorm:"size:255"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime:milli"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (receiver Category) TableName() string {
	return "blog_category"
}
