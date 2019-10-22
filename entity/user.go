package entity

// User is user models property
type User struct {
	ID        int    `gorm:"primary_key"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}
