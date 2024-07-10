package services

import (
	"fmt"

	"github.com/sajadblnyn/autocar-apis/apis/dto"
	"github.com/sajadblnyn/autocar-apis/common"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/constants"
	"github.com/sajadblnyn/autocar-apis/data/db"
	"github.com/sajadblnyn/autocar-apis/data/models"
	"github.com/sajadblnyn/autocar-apis/pkg/logging"
	"github.com/sajadblnyn/autocar-apis/pkg/service_errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	cfg          *config.Config
	database     *gorm.DB
	logger       logging.Logger
	otpService   *OtpService
	tokenService *TokenService
}

func NewUSerService(cfg *config.Config) *UserService {
	db := db.New()

	return &UserService{cfg: config.GetConfig(), database: db.GetDb(), logger: logging.NewLogger(cfg), otpService: NewOtpService(cfg), tokenService: NewTokenService(cfg)}
}

func (u *UserService) SendOtp(r dto.GetOtpRequest) error {

	otp := common.GenerateOtp()

	err := u.otpService.SetOtp(r.MobileNumber, otp)

	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) RegisterByUsername(r *dto.RegisterUserByUsernameRequest) error {
	pass, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		u.logger.Error(logging.General, logging.HashPassword, err.Error(), nil)
		return err
	}
	user := models.User{Username: r.Username, FirstName: r.FirstName, LastName: r.LastName, Email: r.Email, Password: string(pass)}

	var exists bool

	exists, err = u.existsByEmail(user.Email)

	if err != nil {
		u.logger.Error(logging.Database, logging.Select, err.Error(), nil)

		return err
	}
	if exists {
		return &service_errors.ServiceError{EndUserMessage: service_errors.EmailExists}
	}

	exists, err = u.existsByUsername(user.Username)

	if err != nil {
		u.logger.Error(logging.Database, logging.Select, err.Error(), nil)

		return err
	}
	if exists {
		return &service_errors.ServiceError{EndUserMessage: service_errors.UsernameExists}
	}

	tx := u.database.Begin()

	err = tx.Model(&models.User{}).Create(&user).Error
	if err != nil {
		u.logger.Error(logging.Database, logging.Insert, err.Error(), nil)

		u.logger.Error(logging.Database, logging.Rollback, err.Error(), nil)

		tx.Rollback()
		return err
	}

	roleId, err := u.getDefaultRole()
	if err != nil {
		u.logger.Error(logging.Database, logging.Select, err.Error(), nil)

		u.logger.Error(logging.Database, logging.Rollback, err.Error(), nil)

		tx.Rollback()
		return err
	}
	err = tx.Model(&models.UserRole{}).Create(&models.UserRole{UserId: user.Id, RoleId: roleId}).Error
	if err != nil {
		u.logger.Error(logging.Database, logging.Insert, err.Error(), nil)
		u.logger.Error(logging.Database, logging.Rollback, err.Error(), nil)

		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil

}

func (u *UserService) LoginOrRegisterByMobile(r *dto.RegisterLoginByMobileRequest) (*dto.TokenDetail, error) {
	var err error
	err = u.otpService.ValidateOtp(r.MobileNumber, r.Otp)
	if err != nil {
		return nil, err
	}
	exists, err := u.existsByMobileNumber(r.MobileNumber)

	if err != nil {
		return nil, err
	}
	if exists {
		return u.directLoginByUsername(r.MobileNumber)
	}

	user := models.User{MobileNumber: r.MobileNumber, Username: r.MobileNumber}
	pass := common.GeneratePassword()
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hash)
	tx := u.database.Begin()
	err = tx.Create(&user).Error
	if err != nil {
		u.logger.Error(logging.Database, logging.Insert, err.Error(), nil)

		u.logger.Error(logging.Database, logging.Rollback, err.Error(), nil)

		tx.Rollback()
		return nil, err
	}

	roleId, err := u.getDefaultRole()
	if err != nil {
		u.logger.Error(logging.Database, logging.Select, err.Error(), nil)

		u.logger.Error(logging.Database, logging.Rollback, err.Error(), nil)

		tx.Rollback()
		return nil, err
	}
	err = tx.Model(&models.UserRole{}).Create(&models.UserRole{UserId: user.Id, RoleId: roleId}).Error
	if err != nil {
		u.logger.Error(logging.Database, logging.Insert, err.Error(), nil)
		u.logger.Error(logging.Database, logging.Rollback, err.Error(), nil)

		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return u.directLoginByUsername(r.MobileNumber)

}

func (u *UserService) LoginByUsername(r *dto.LoginByUsernameRequest) (*dto.TokenDetail, error) {

	user := models.User{}
	var err error
	err = u.database.Model(&models.User{}).Where("username=?", r.Username).Preload("UserRoles", func(t *gorm.DB) *gorm.DB {
		return t.Preload("Role")
	}).Find(&user).Error

	if err != nil {
		u.logger.Error(logging.Database, logging.Select, err.Error(), nil)
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Password))
	if err != nil {
		return nil, err
	}
	tDto := tokenDto{UserId: user.Id, FirstName: user.FirstName,
		LastName: user.LastName, MobileNumber: user.MobileNumber,
		Email: user.Email, Username: user.Email}

	for _, v := range *user.UserRoles {
		tDto.Roles = append(tDto.Roles, v.Role.Name)
	}

	token, err := u.tokenService.GenerateToken(&tDto)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (u *UserService) directLoginByUsername(username string) (*dto.TokenDetail, error) {

	defer func() {
		err := recover()
		if err != nil {
			u.logger.Error(logging.General, logging.RecoverError, fmt.Sprintf("%v", err), nil)
		}
	}()
	var user models.User
	var err error
	err = u.getUserWithRolesByUsername(username, &user)

	if err != nil {
		u.logger.Error(logging.Database, logging.Select, err.Error(), nil)
		return nil, err
	}
	tDto := tokenDto{UserId: user.Id, FirstName: user.FirstName,
		LastName: user.LastName, MobileNumber: user.MobileNumber,
		Email: user.Email, Username: user.Email}

	for _, v := range *user.UserRoles {
		tDto.Roles = append(tDto.Roles, v.Role.Name)
	}

	token, err := u.tokenService.GenerateToken(&tDto)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (u *UserService) getUserWithRolesByUsername(username string, user *models.User) (err error) {

	err = u.database.Model(&models.User{}).Where("username=?", username).Preload("UserRoles", func(t *gorm.DB) *gorm.DB {
		return t.Preload("Role")
	}).Find(user).Error

	if err != nil {
		u.logger.Error(logging.Database, logging.Select, err.Error(), nil)
		return err
	}
	return nil
}

func (u *UserService) existsByEmail(e string) (bool, error) {
	var count int64

	er := u.database.Model(&models.User{}).Where("email=?", e).Count(&count).Error
	if er != nil {
		u.logger.Error(logging.Database, logging.Select, er.Error(), nil)
		return false, er
	}
	return count > 0, nil
}

func (u *UserService) existsByUsername(username string) (bool, error) {
	var count int64
	er := u.database.Model(&models.User{}).Where("username=?", username).Count(&count).Error
	if er != nil {
		u.logger.Error(logging.Database, logging.Select, er.Error(), nil)

		return false, er
	}
	return count > 0, nil
}

func (u *UserService) existsByMobileNumber(mobileNumber string) (bool, error) {
	var count int64
	er := u.database.Model(&models.User{}).Where("mobile_number=?", mobileNumber).Count(&count).Error
	if er != nil {
		u.logger.Error(logging.Database, logging.Select, er.Error(), nil)

		return false, er
	}
	return count > 0, nil
}

func (u *UserService) getDefaultRole() (roleId int, err error) {
	er := u.database.Model(&models.Role{}).Select("id").Where("name=?", constants.DefaultRoleName).First(&roleId).Error
	if er != nil {
		u.logger.Error(logging.Database, logging.Select, er.Error(), nil)

		return 0, er
	}
	return roleId, nil
}
