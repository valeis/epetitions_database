package models

import "time"

type UserModel struct {
	ID       uint `gorm:"primaryKey;autoIncrement:true;unique"`
	Email    string `gorm:"not null;unique" json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (UserModel) TableName() string {
	return "users"
}

type PetitionModel struct {
	ID        uint `gorm:"primaryKey;autoIncrement:true;unique"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Category  string `json:"category"`
	Status    uint `json:"status"`
	UserID    uint `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	UserModel UserModel `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (PetitionModel) TableName() string {
	return "petitions"
}

type VoteModel struct {
	ID         uint `gorm:"primaryKey;autoIncrement:true;unique"`
	PetitionID uint `json:"petitionId"`
	UserID     uint `json:"userId"`
	Status     uint `json:"status"`
	OTP        string `json:"otp"`
	UserModel  UserModel `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PetitionModel PetitionModel `gorm:"foreignKey:PetitionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (VoteModel) TableName() string {
	return "votes"
}