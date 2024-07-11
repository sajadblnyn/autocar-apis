package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/sajadblnyn/autocar-apis/common"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/constants"
	"github.com/sajadblnyn/autocar-apis/data/db"
	"github.com/sajadblnyn/autocar-apis/pkg/logging"
	"github.com/sajadblnyn/autocar-apis/pkg/service_errors"
	"gorm.io/gorm"
)

type BaseService[T any, Tc any, Tu any, Tr any] struct {
	db     *gorm.DB
	logger logging.Logger
}

func NewBaseService[T any, Tc any, Tu any, Tr any](cfg *config.Config) *BaseService[T, Tc, Tu, Tr] {
	db := db.New()
	return &BaseService[T, Tc, Tu, Tr]{
		db:     db.GetDb(),
		logger: logging.NewLogger(cfg),
	}

}

func (s *BaseService[T, Tc, Tu, Tr]) Create(ctx context.Context, r *Tc) (*Tr, error) {
	model, err := common.ConvertType[T](r)

	if err != nil {
		return nil, err
	}

	tx := s.db.WithContext(ctx).Begin()

	err = tx.Create(model).Error

	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Database, logging.Insert, err.Error(), nil)
		return nil, err
	}
	tx.Commit()
	return common.ConvertType[Tr](model)

}

func (s *BaseService[T, Tc, Tu, Tr]) Update(ctx context.Context, id int, r *Tu) (*Tr, error) {
	updateMap, err := common.ConvertType[map[string]interface{}](r)

	if err != nil {
		return nil, err
	}
	(*updateMap)["updated_at"] = sql.NullTime{Valid: true, Time: time.Now().UTC()}
	(*updateMap)["updated_by"] = &sql.NullInt64{Valid: true, Int64: int64(ctx.Value(constants.UserIdKey).(float64))}

	tx := s.db.WithContext(ctx).Begin()

	err = tx.Model(new(T)).Where("id=? and deleted_by is null", id).Updates(*updateMap).Error

	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Database, logging.Update, err.Error(), nil)

		return nil, err
	}
	tx.Commit()
	return s.GetById(ctx, id)
}

func (s *BaseService[T, Tc, Tu, Tr]) Delete(ctx context.Context, id int) error {
	updateMap := map[string]interface{}{}

	if ctx.Value(constants.UserIdKey) == "" {
		return &service_errors.ServiceError{EndUserMessage: service_errors.PermissionDenied}
	}
	(updateMap)["deleted_at"] = sql.NullTime{Valid: true, Time: time.Now().UTC()}
	(updateMap)["deleted_by"] = &sql.NullInt64{Valid: true, Int64: int64(ctx.Value(constants.UserIdKey).(float64))}

	tx := s.db.WithContext(ctx).Begin()
	err := tx.Model(new(T)).Where("id=? and deleted_by is null", id).Updates(updateMap).Error
	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Database, logging.Delete, err.Error(), nil)
		return err
	}
	tx.Commit()
	return nil
}

func (s *BaseService[T, Tc, Tu, Tr]) GetById(ctx context.Context, id int) (*Tr, error) {
	model := new(T)

	err := s.db.Model(model).Where("id=? and deleted_by is null", id).First(&model).Error
	if err != nil {
		s.logger.Error(logging.Database, logging.Select, err.Error(), nil)

		return nil, err
	}

	return common.ConvertType[Tr](model)
}
