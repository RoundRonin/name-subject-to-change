package dto

import "time"

type UserRequest struct {
	Name  string `json:"name" binding:"required" example:"Ronin"`
	Email string `json:"email" binding:"required,email" example:"ronin@example.com"`
}

type UserResponse struct {
	ID        string    `json:"id" example:"r9n-123abc"`
	Name      string    `json:"name" example:"Ronin"`
	Email     string    `json:"email" example:"ronin@example.com"`
	CreatedAt time.Time `json:"created_at" example:"2025-10-21T12:00:00Z"`
}
