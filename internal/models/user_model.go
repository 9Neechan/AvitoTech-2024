package models

/*import (
	"time"

	"github.com/9Neechan/AvitoTech-2024/internal/database"

	"github.com/google/uuid"
)*/

/*type User struct {
	ID        uuid.UUID `json:"id"`
	IsAdmin   bool      `json:"is_admin"`
	CreatedAt time.Time `json:"updated_at"`
	UpdatedAt time.Time `json:"name"`
}*/

type User struct {
	IsAdmin   bool     
}

/*func DatabaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		IsAdmin:   user.IsAdmin,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}*/
