package domain

import (
	"time"

	"github.com/google/uuid"
)

type UserTier string

const (
	TIER1 UserTier = "Batik Pemula"
	TIER2 UserTier = "Batik Penjelajah"
	TIER3 UserTier = "Batik Satria"
	TIER4 UserTier = "Batik Jawara"
	TIER5 UserTier = "Batik Legenda"
)

type User struct {
	ID                    uuid.UUID    `json:"id" gorm:"type:varchar(36);primaryKey"`
	Username              string       `json:"username"`
	Email                 string       `json:"email" gorm:"uniqueIndex type:varchar(100)"`
	Password              string       `json:"password"`
	Experience            int          `json:"experience"`
	Level                 int          `json:"level"`
	Tier                  UserTier     `json:"tier" gorm:"type:ENUM('Batik Pemula', 'Batik Penjelajah', 'Batik Satria', 'Batik Jawara', 'Batik Legenda')"`
	ForgotPasswordToken   string       `json:"forgot_password_token"`
	ForgotPasswordExpired time.Time    `json:"forgot_password_expired" gorm:"type:datetime"`
	UserQuiz              []UserAnswer `json:"user_quiz" gorm:"foreignKey:UserID"`
}

type UserRegister struct {
	Username        string `json:"username" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}
type UserParam struct {
	ID                  uuid.UUID `json:"id"`
	Username            string    `json:"username"`
	Email               string    `json:"email"`
	ForgotPasswordToken string    `json:"forgot_password_token"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type OauthRedirectLink struct {
	RedirectLink string `json:"redirect_link"`
}

type UserOauth struct {
	Email       string `json:"email" gorm:"primary key"`
	Family_name string `json:"family_name"`
	Given_name  string `json:"given_name"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	Picture     string `json:"picture"`
}

type UserForgotPassword struct {
	Email string `json:"email" binding:"required"`
}

type ResetPassword struct {
	ForgotPasswordToken string `json:"forgot_password_token" binding:"required"`
	Password            string `json:"password" binding:"required"`
	ConfirmPassword     string `json:"confirm_password" binding:"required"`
}

type UserProfile struct {
	Username           string   `json:"username"`
	Email              string   `json:"email"`
	Experience         int      `json:"experience"`
	Level              int      `json:"level"`
	Tier               UserTier `json:"tier"`
	ExpToNextTier      int      `json:"exp_to_next_tier"`
	TotalQuiz          int      `json:"total_quiz"`
	TotalCorrectAnswer int      `json:"total_correct_answer"`
}
