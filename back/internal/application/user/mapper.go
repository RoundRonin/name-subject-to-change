package user

import (
	"github.com/RoundRonin/name-subject-to-change/back/internal/application/user/dto"
	"github.com/RoundRonin/name-subject-to-change/back/internal/domain/user"
)

func ToUserResponse(u *user.User) dto.UserResponse {
	return dto.UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
	}
}
