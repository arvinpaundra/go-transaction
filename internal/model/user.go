package model

// Define the table, here we declare struct name as User, so the table will be created as `users` table, this was define in the Gorm docs
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name" gorm:"varchar;not_null"`
	Email string `json:"email" gorm:"varchar;not_null;unique"`
	Common
}
