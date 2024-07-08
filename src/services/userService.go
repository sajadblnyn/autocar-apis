package services

import (
	"github.com/sajadblnyn/autocar-apis/apis/dto"
	"github.com/sajadblnyn/autocar-apis/common"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/data/db"
	"github.com/sajadblnyn/autocar-apis/pkg/logging"
	"gorm.io/gorm"
)

type UserService struct {
	cfg        *config.Config
	database   *gorm.DB
	logger     logging.Logger
	otpService *OtpService
}

func NewUSerService(cfg *config.Config) *UserService {
	db := db.New()

	return &UserService{cfg: config.GetConfig(), database: db.GetDb(), logger: logging.NewLogger(cfg), otpService: NewOtpService(cfg)}
}

func (u *UserService) SendOtp(r dto.GetOtpRequest) error {

	otp := common.GenerateOtp()

	err := u.otpService.SetOtp(r.MobileNumber, otp)

	if err != nil {
		return err
	}
	return nil
}
