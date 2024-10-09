package domain

import "github.com/google/uuid"

type UserAnswer struct {
	UserID             uuid.UUID `json:"user_id" gorm:"type:varchar(36)"`
	QuizID             []int     `json:"quiz_id" gorm:"type:json"`
	UserAnswer         []string  `json:"user_answer" gorm:"type:json"`
	TotalCorrectAnswer int       `json:"total_correct_answer"`
}

type UserAnswerResponse struct {
	TotalCorrectAnswer int `json:"total_correct_answer"`
}
