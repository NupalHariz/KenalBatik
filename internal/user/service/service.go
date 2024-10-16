package service

import (
	"context"
	"kenalbatik-be/internal/domain"
	"kenalbatik-be/internal/infra/gomail"
	"kenalbatik-be/internal/infra/helper"
	"kenalbatik-be/internal/infra/jwt"
	"kenalbatik-be/internal/user/repository"
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(ctx context.Context, userRegister domain.UserRegister) error
	Login(ctx context.Context, userLogin domain.UserLogin) (domain.UserLoginResponse, error)
	Oauth(ctx context.Context, user domain.UserOauth) (string, error)
	ForgotPassword(ctx context.Context, userForgot domain.UserForgotPassword) error
	ResetPassword(ctx context.Context, userReset domain.ResetPassword) error
	GetUserByID(ctx context.Context, userId uuid.UUID) (domain.UserProfile, error)
}

type userService struct {
	userRepo repository.Userepository
	jwt      jwt.JWT
	goMail   gomail.GoMailInterface
}

func NewUserService(userRepo repository.Userepository, jwt jwt.JWT, goMail gomail.GoMailInterface) UserService {
	return &userService{
		userRepo: userRepo,
		jwt:      jwt,
		goMail:   goMail,
	}
}

func (s *userService) RegisterUser(ctx context.Context, userRegister domain.UserRegister) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user domain.User

	//Ga guna udah ada unique constraint di db
	err := s.userRepo.FindUser(ctx, &user, domain.UserParam{Email: userRegister.Email})
	if err != nil {
		if err != domain.ErrRecordNotFound {
			return err
		}
	}

	if user.Email != "" {
		return domain.ErrEmailAlreadyExist
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(userRegister.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user = domain.User{
		ID:         uuid.New(),
		Username:   userRegister.Username,
		Email:      userRegister.Email,
		Password:   string(hashPassword),
		Experience: 0,
		Level:      1,
		TierID:       string(domain.TIER1),
		ForgotPasswordToken: "",
		ForgotPasswordExpired: time.Date(1970, 1, 1, 0, 0, 0, 0, time.Local),
	}

	err = s.userRepo.CreateUser(ctx, user)

	select {
	case <-ctx.Done():
		return domain.ErrTimeout
	default:
		return err
	}
}

func (s *userService) Login(ctx context.Context, userLogin domain.UserLogin) (domain.UserLoginResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user domain.User

	err := s.userRepo.FindUser(ctx, &user, domain.UserParam{Email: userLogin.Email})
	if err != nil {
		if err == domain.ErrRecordNotFound {
			return domain.UserLoginResponse{}, domain.ErrInvalidEmailOrPassword
		}

		return domain.UserLoginResponse{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password))
	if err != nil {
		return domain.UserLoginResponse{}, domain.ErrInvalidEmailOrPassword
	}

	token, err := s.jwt.GenerateToken(user.ID)

	res := domain.UserLoginResponse{
		Token: token,
	}

	select {
	case <-ctx.Done():
		return domain.UserLoginResponse{}, domain.ErrTimeout
	default:
		return res, err
	}
}

func (s *userService) Oauth(ctx context.Context, user domain.UserOauth) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var userDomain domain.User

	err := s.userRepo.FindUser(ctx, &userDomain, domain.UserParam{Email: user.Email})
	if err != nil {
		if err != domain.ErrRecordNotFound {
			return "", err
		}

		newUser := domain.User{
			ID:         uuid.New(),
			Username:   user.Name,
			Email:      user.Email,
			Experience: 0,
			Level:      1,
			TierID:       string(domain.TIER1),
			ForgotPasswordExpired: time.Date(1970, 1, 1, 0, 0, 0, 0, time.Local),
		}

		err = s.userRepo.CreateUser(ctx, newUser)
		if err != nil {
			return "", err
		}
	}

	token, err := s.jwt.GenerateToken(userDomain.ID)

	select {
	case <-ctx.Done():
		return "", domain.ErrTimeout
	default:
		return token, err
	}
}

func (s *userService) ForgotPassword(ctx context.Context, userForgot domain.UserForgotPassword) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user domain.User

	err := s.userRepo.FindUser(ctx, &user, domain.UserParam{Email: userForgot.Email})
	if err != nil {
		return err
	}

	forgotPasswordToken := helper.GenerateRandomStringNumber()
	forgotPasswordTokenExpired := time.Now().Add(time.Hour * 1)

	subject := "Reset Password"
	HTMLbody := "<p>PLease input this code to reset your password: " + forgotPasswordToken + "</p>"

	updateUser := domain.User{
		ForgotPasswordToken: 	forgotPasswordToken,
		ForgotPasswordExpired: 	forgotPasswordTokenExpired,
	}

	err = s.userRepo.UpdateUser(ctx, updateUser, domain.UserParam{Email: user.Email})
	if err != nil {
		return err
	}

	go func() {
		err = s.goMail.SendEmail(subject, HTMLbody, user.Email)
		if err != nil {
			log.Println("Error sending email: ", err)
		}
	}()
	
	
	select {
	case <-ctx.Done():
		return domain.ErrTimeout
	default:
		return err
	}
}

func (s *userService) ResetPassword(ctx context.Context, userReset domain.ResetPassword) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user domain.User

	err := s.userRepo.FindUser(ctx, &user, domain.UserParam{ForgotPasswordToken: userReset.ForgotPasswordToken})
	if err != nil {
		return err
	}

	if user.ForgotPasswordExpired.Before(time.Now()) {
		return domain.ErrTokenExpired
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(userReset.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	updateUser := domain.User{
		Password: string(hashPassword),
		ForgotPasswordExpired: time.Now(),
	}

	err = s.userRepo.UpdateUser(ctx, updateUser, domain.UserParam{ID: user.ID})
	
	select {
	case <-ctx.Done():
		return domain.ErrTimeout
	default:
		return err
	}
}

func (s *userService) GetUserByID(ctx context.Context, userId uuid.UUID) (domain.UserProfile, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user domain.User
	err := s.userRepo.FindUser(ctx, &user, domain.UserParam{ID: userId})
	
	var totalCorrenctAnswer int
	for _, u := range user.UserQuiz {
		totalCorrenctAnswer += u.TotalCorrectAnswer
	}

	totalQuiz := len(user.UserQuiz) * 5

	nextExp := helper.GetNextExp(domain.UserTier(user.TierID))

	userResponse := domain.UserProfile{
		Username: user.Username,
		Email: user.Email,
		Experience: user.Experience,
		Level: user.Level,
		Tier: domain.UserTier(user.TierID),
		TierPhotoLink: user.Tier.TierPhotoLink,
		ExpToNextTier: nextExp,
		TotalQuiz: totalQuiz,
		TotalCorrectAnswer: totalCorrenctAnswer,
	}

	select {
	case <-ctx.Done():
		return userResponse, domain.ErrTimeout
	default:
		return userResponse, err
	}
}
