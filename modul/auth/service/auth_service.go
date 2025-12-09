package auth_service

import (
	"errors"
	"ijro-nazorat/helper"
	auth_dto "ijro-nazorat/modul/auth/dto"
	user_model "ijro-nazorat/modul/user/model"

	"gorm.io/gorm"
)

type AuthService interface {
	Login(login auth_dto.Login) (auth_dto.LoginResponse, error)
	RefreshToken(refreshToken auth_dto.RefreshToken) (auth_dto.LoginResponse, error)
}
type authService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) AuthService {
	return &authService{db: db}
}

func (service *authService) Login(login auth_dto.Login) (auth_dto.LoginResponse, error) {
	var model user_model.User
	{
		if err := service.db.Where("email = ?", login.Email).First(&model).Error; err != nil {
			return auth_dto.LoginResponse{}, errors.New("user not found")
		}

		// if err := bcrypt.CompareHashAndPassword([]byte(model.Password), []byte(login.Password)); err != nil {
		// 	return auth_dto.LoginResponse{}, errors.New("invalid password")
		// }
	}

	token, err := helper.GenerateJWT(model.ID, model.CountryID, model.Name, model.Email, model.Role)
	{
		if err != nil {
			return auth_dto.LoginResponse{}, err
		}
	}

	return auth_dto.LoginResponse{Token: token}, nil
}

func (service *authService) RefreshToken(refreshToken auth_dto.RefreshToken) (auth_dto.LoginResponse, error) {
	claims, err := helper.ParseJWT(refreshToken.Token)
	{
		if err != nil {
			return auth_dto.LoginResponse{}, errors.New("invalid token")
		}
	}

	token, err := helper.GenerateJWT(claims.UserID, claims.CountryId, claims.Name, claims.Email, claims.Role)
	{
		if err != nil {
			return auth_dto.LoginResponse{}, err
		}
	}

	return auth_dto.LoginResponse{Token: token}, nil
}
