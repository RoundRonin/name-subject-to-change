package user

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type User struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
}

func NewUser(name, email string) *User {
	createdAt := time.Now().UTC()
	hash := sha256.Sum256([]byte(name + email + createdAt.Format(time.RFC3339Nano)))
	return &User{
		ID:        hex.EncodeToString(hash[:]),
		Name:      name,
		Email:     email,
		CreatedAt: createdAt,
	}
}
