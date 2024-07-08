package services

import (
	"fmt"
	"time"

	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/constants"
	"github.com/sajadblnyn/autocar-apis/data/cache"
	"github.com/sajadblnyn/autocar-apis/pkg/logging"
	"github.com/sajadblnyn/autocar-apis/pkg/service_errors"
)

type OtpService struct {
	cfg          *config.Config
	cacheService cache.CacheService
	logger       logging.Logger
}

type OtpDto struct {
	Value string
	Used  bool
}

func NewOtpService(cfg *config.Config) *OtpService {
	return &OtpService{cfg: cfg, cacheService: cache.New(), logger: logging.NewLogger(cfg)}
}

func (o *OtpService) SetOtp(mobileNumber string, otp string) error {
	key := fmt.Sprintf("%s:%s", constants.CacheOtpDefaultKey, mobileNumber)

	v := OtpDto{}
	err := o.cacheService.Get(key, &v)
	if err == nil && v.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpUsed}
	} else if err == nil && v.Value != "" {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpExists}
	}

	v.Value = otp
	v.Used = false
	err = o.cacheService.Set(key, &v, time.Second*o.cfg.Otp.ExpireTime)

	return err
}

func (o *OtpService) ValidateOtp(mobileNumber string, otp string) error {
	key := fmt.Sprintf("%s:%s", constants.CacheOtpDefaultKey, mobileNumber)

	v := OtpDto{}
	err := o.cacheService.Get(key, &v)
	if err != nil {
		return err
	} else if v.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpUsed}
	} else if v.Value != otp {
		return &service_errors.ServiceError{EndUserMessage: service_errors.NotValidOtp}
	}

	v.Used = true
	o.cacheService.Set(key, &v, time.Second*o.cfg.Otp.ExpireTime)
	return nil
}
