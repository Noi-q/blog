package dto

import (
	"blog-admin/models"
	"time"
)

type CategoryDto struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ToCategoryDto(category models.Category) CategoryDto {
	return CategoryDto{
		ID:          category.ID,
		Title:       category.Title,
		Description: category.Description,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}
}
