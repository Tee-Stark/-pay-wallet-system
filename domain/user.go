package domain

type User struct {
	ID    string `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"not null;unique"`
}
