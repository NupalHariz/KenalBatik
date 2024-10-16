package domain

import "github.com/google/uuid"

type Difficulty string

const (
	EASY   Difficulty = "easy"
	MEDIUM Difficulty = "medium"
	HARD   Difficulty = "hard"
)

type Quiz struct {
	ID         int          `json:"id" gorm:"primaryKey;autoIncrement"`
	Question   string       `json:"question"`
	Answer     string       `json:"answer"`
	OptionA    string       `json:"option_a"`
	OptionB    string       `json:"option_b"`
	OptionC    string       `json:"option_c"`
	OptionD    string       `json:"option_d"`
	Difficulty Difficulty   `json:"difficulty"`
	Image_Link string       `json:"image_link"`
	UserQuiz   []UserAnswer `json:"user_quiz" gorm:"foreignKey:QuizID"`
}

type AnswerRequest struct {
	QuizID     []int     `json:"quiz_id"`
	UserId     uuid.UUID `json:"-"`
	UserAnswer []string  `json:"user_answer"`
}

type AnswerResponse struct {
	UserExperience       int      `json:"user_experience"`
	Username             string   `json:"username"`
	UserTier             UserTier `json:"user_tier"`
	TierPhotoLink        string   `json:"tier_photo_link"`
	ExpToNextTier        int      `json:"exp_to_next_tier"`
	CurrentCorrectAnswer int      `json:"current_correct_answer"`
	TotalQuiz            int      `json:"total_quiz"`
	TotalCorrectAnswer   int      `json:"total_correct_answer"`
}

type QuizParams struct {
	ID int `json:"id"`
}
