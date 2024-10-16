package service

import (
	"context"
	"encoding/json"
	"kenalbatik-be/internal/domain"
	"kenalbatik-be/internal/infra/helper"
	quizRepo "kenalbatik-be/internal/quiz/repository"
	userRepo "kenalbatik-be/internal/user/repository"
	userAnswerRepo "kenalbatik-be/internal/userAnswer/repository"
	"time"

	"github.com/google/uuid"
)

type QuizService interface {
	GetQuizzes(ctx context.Context, userId uuid.UUID) ([]domain.Quiz, error)
	CheckAnswer(ctx context.Context, userId uuid.UUID, userAnswer domain.AnswerRequest) (domain.AnswerResponse, error)
}

type quizService struct {
	userRepo       userRepo.Userepository
	quizRepo       quizRepo.QuizRepository
	userAnswerRepo userAnswerRepo.UserAnswerRepositoryInterface
}

func NewQuizService(userRepo userRepo.Userepository, quizRepo quizRepo.QuizRepository, userAnswerRepo userAnswerRepo.UserAnswerRepositoryInterface) QuizService {
	return &quizService{
		userRepo:       userRepo,
		quizRepo:       quizRepo,
		userAnswerRepo: userAnswerRepo,
	}
}

func (s *quizService) GetQuizzes(ctx context.Context, userId uuid.UUID) ([]domain.Quiz, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user domain.User
	err := s.userRepo.FindUser(ctx, &user, domain.UserParam{ID: userId})
	if err != nil {
		return nil, err
	}

	var quizzes []domain.Quiz

	switch user.TierID {
	case string(domain.TIER1):
		err = s.quizRepo.GetQuiz(ctx, &quizzes, "DIFFICULTY = ?", []interface{}{domain.EASY})
	case string(domain.TIER2):
		err = s.quizRepo.GetQuiz(ctx, &quizzes, "DIFFICULTY = ? OR DIFFICULTY = ?", []interface{}{domain.EASY, domain.MEDIUM})
	case string(domain.TIER3):
		err = s.quizRepo.GetQuiz(ctx, &quizzes, "DIFFICULTY = ?", []interface{}{domain.MEDIUM})
	case string(domain.TIER4):
		err = s.quizRepo.GetQuiz(ctx, &quizzes, "DIFFICULTY = ? OR DIFFICULTY = ?", []interface{}{domain.MEDIUM, domain.HARD})
	case string(domain.TIER5):
		err = s.quizRepo.GetQuiz(ctx, &quizzes, "DIFFICULTY = ?", []interface{}{domain.HARD})
	}
	if err != nil {
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, domain.ErrTimeout
	default:
		return quizzes, nil
	}
}

func (s *quizService) CheckAnswer(ctx context.Context, userId uuid.UUID, userAnswer domain.AnswerRequest) (domain.AnswerResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user domain.User
	err := s.userRepo.FindUser(ctx, &user, domain.UserParam{ID: userId})
	if err != nil {
		return domain.AnswerResponse{}, err
	}

	var correctAnswer int

	for i, quizId := range userAnswer.QuizID {
		var quiz domain.Quiz
		err := s.quizRepo.GetQuizByID(ctx, &quiz, quizId)
		if err != nil {
			return domain.AnswerResponse{}, err
		}

		if quiz.Answer == userAnswer.UserAnswer[i] {
			correctAnswer++

			switch quiz.Difficulty {
			case domain.EASY:
				user.Experience += 3
			case domain.MEDIUM:
				user.Experience += 5
			case domain.HARD:
				user.Experience += 7
			}
		}
	}

	if user.Experience >= 100 {
		user.TierID = string(domain.TIER5)
	} else if user.Experience >= 75 && user.Experience < 100 {
		user.TierID = string(domain.TIER4)
	} else if user.Experience >= 50 && user.Experience < 70 {
		user.TierID = string(domain.TIER3)
	} else if user.Experience >= 25 && user.Experience < 50 {
		user.TierID = string(domain.TIER2)
	} else {
		user.TierID = string(domain.TIER1)
	}

	quizIdJson, err := json.Marshal(userAnswer.QuizID)
	if err != nil {
		return domain.AnswerResponse{}, err
	}

	userAnswerJson, err := json.Marshal(userAnswer.UserAnswer)
	if err != nil {
		return domain.AnswerResponse{}, err
	}

	userAnswerData := domain.UserAnswer{
		UserID:             userId,
		QuizID:             quizIdJson,
		UserAnswer:         userAnswerJson,
		TotalCorrectAnswer: correctAnswer,
	}

	err = s.userAnswerRepo.Create(ctx, &userAnswerData)
	if err != nil {
		return domain.AnswerResponse{}, err
	}

	err = s.userRepo.UpdateUser(ctx, user, domain.UserParam{ID: userId})
	if err != nil {
		return domain.AnswerResponse{}, err
	}

	var updatedUser domain.User
	err = s.userRepo.FindUser(ctx, &updatedUser, domain.UserParam{ID: userId})
	if err != nil {
		return domain.AnswerResponse{}, err
	}

	var totalCorrenctAnswer int
	for _, u := range updatedUser.UserQuiz {
		totalCorrenctAnswer += u.TotalCorrectAnswer
	}

	res := domain.AnswerResponse{
		Username:             updatedUser.Username,
		UserExperience:       updatedUser.Experience,
		UserTier:             domain.UserTier(updatedUser.TierID),
		TierPhotoLink:        updatedUser.Tier.TierPhotoLink,
		ExpToNextTier:        helper.GetNextExp(domain.UserTier(updatedUser.TierID)),
		CurrentCorrectAnswer: correctAnswer,
		TotalQuiz:            len(updatedUser.UserQuiz) * 5,
		TotalCorrectAnswer:   totalCorrenctAnswer,
	}

	select {
	case <-ctx.Done():
		return domain.AnswerResponse{}, domain.ErrTimeout
	default:
		return res, nil
	}
}
