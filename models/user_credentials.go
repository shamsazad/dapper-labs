package models

type UserCredential struct {
	Email          string `json:"email" gorm:"primaryKey"`
	HashedPassword string `gorm:"size:200" json:"hashed_password"`
}
