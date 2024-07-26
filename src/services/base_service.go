package services

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"reflect"
	"strings"
	"time"

	"github.com/sajadblnyn/autocar-apis/apis/dto"
	"github.com/sajadblnyn/autocar-apis/common"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/constants"
	"github.com/sajadblnyn/autocar-apis/data/db"
	"github.com/sajadblnyn/autocar-apis/data/models"
	"github.com/sajadblnyn/autocar-apis/pkg/logging"
	"github.com/sajadblnyn/autocar-apis/pkg/metrics"
	"github.com/sajadblnyn/autocar-apis/pkg/service_errors"
	"gorm.io/gorm"
)

type preload struct{ string }
type BaseService[T any, Tc any, Tu any, Tr any] struct {
	db       *gorm.DB
	logger   logging.Logger
	Preloads []preload
}

func NewBaseService[T any, Tc any, Tu any, Tr any](cfg *config.Config, preloads []preload) *BaseService[T, Tc, Tu, Tr] {
	db := db.New()
	return &BaseService[T, Tc, Tu, Tr]{
		db:       db.GetDb(),
		logger:   logging.NewLogger(cfg),
		Preloads: preloads,
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
		metrics.DbCall.WithLabelValues(reflect.TypeOf(*model).String(), "Create", "Failed").Inc()

		return nil, err
	}
	tx.Commit()
	metrics.DbCall.WithLabelValues(reflect.TypeOf(*model).String(), "Create", "Success").Inc()
	bm, _ := common.ConvertType[models.BaseModel](model)
	return s.GetById(ctx, bm.Id)

}

func (s *BaseService[T, Tc, Tu, Tr]) Update(ctx context.Context, id int, r *Tu) (*Tr, error) {
	updateMap, err := common.ConvertType[map[string]interface{}](r)

	snakeMap := map[string]interface{}{}
	for k, v := range *updateMap {
		snakeMap[common.ToSnakeCase(k)] = v
	}
	if err != nil {
		return nil, err
	}
	snakeMap["updated_at"] = sql.NullTime{Valid: true, Time: time.Now().UTC()}
	snakeMap["updated_by"] = &sql.NullInt64{Valid: true, Int64: int64(ctx.Value(constants.UserIdKey).(float64))}

	tx := s.db.WithContext(ctx).Begin()

	model := new(T)
	err = tx.Model(model).Where("id=? and deleted_by is null", id).Updates(snakeMap).Error

	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Database, logging.Update, err.Error(), nil)
		metrics.DbCall.WithLabelValues(reflect.TypeOf(*model).String(), "Update", "Failed").Inc()

		return nil, err
	}
	tx.Commit()
	metrics.DbCall.WithLabelValues(reflect.TypeOf(*model).String(), "Update", "Success").Inc()

	return s.GetById(ctx, id)
}

func (s *BaseService[T, Tc, Tu, Tr]) Delete(ctx context.Context, id int) error {
	updateMap := map[string]interface{}{}

	if ctx.Value(constants.UserIdKey) == "" {
		return &service_errors.ServiceError{EndUserMessage: service_errors.PermissionDenied}
	}
	(updateMap)["deleted_at"] = sql.NullTime{Valid: true, Time: time.Now().UTC()}
	(updateMap)["deleted_by"] = &sql.NullInt64{Valid: true, Int64: int64(ctx.Value(constants.UserIdKey).(float64))}

	model := new(T)
	tx := s.db.WithContext(ctx).Begin()
	err := tx.Model(model).Where("id=? and deleted_by is null", id).Updates(updateMap).Error
	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Database, logging.Delete, err.Error(), nil)
		metrics.DbCall.WithLabelValues(reflect.TypeOf(*model).String(), "Delete", "Failed").Inc()

		return err
	}
	tx.Commit()
	metrics.DbCall.WithLabelValues(reflect.TypeOf(*model).String(), "Delete", "Success").Inc()

	return nil
}

func (s *BaseService[T, Tc, Tu, Tr]) GetById(ctx context.Context, id int) (*Tr, error) {
	model := new(T)

	databse := Preload(s.db, s.Preloads)

	err := databse.Model(model).Where("id=? and deleted_by is null", id).First(&model).Error
	if err != nil {
		s.logger.Error(logging.Database, logging.Select, err.Error(), nil)
		metrics.DbCall.WithLabelValues(reflect.TypeOf(*model).String(), "GetById", "Failed").Inc()

		return nil, err
	}
	metrics.DbCall.WithLabelValues(reflect.TypeOf(*model).String(), "GetById", "Success").Inc()

	return common.ConvertType[Tr](model)
}
func (s *BaseService[T, Tc, Tu, Tr]) GetByFilter(c context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[Tr], error) {
	model := new(T)

	res, err := Paginate[T, Tr](req, s.Preloads, s.db)
	if err != nil {
		metrics.DbCall.WithLabelValues(reflect.TypeOf(*model).String(), "GetByFilter", "Failed").Inc()
		return nil, err
	}
	metrics.DbCall.WithLabelValues(reflect.TypeOf(*model).String(), "GetByFilter", "Success").Inc()
	return res, nil
}
func GetFilterQuery[T any](filter *dto.DynamicFilter) string {

	if filter.Filter == nil {
		return ""
	}

	model := new(T)
	modelType := reflect.TypeOf(*model)

	filters := []string{}
	filters = append(filters, "deleted_by is null")

	for k, v := range filter.Filter {
		fld, ok := modelType.FieldByName(k)
		if !ok {
			continue
		}
		fld.Name = common.ToSnakeCase(fld.Name)
		switch v.Type {
		case "contains":
			filters = append(filters, fmt.Sprintf("%s ILIKE '%%%s%%'", fld.Name, v.From))
		case "notContains":
			filters = append(filters, fmt.Sprintf("%s ILIKE '%%%s%%'", fld.Name, v.From))
		case "equals":
			filters = append(filters, fmt.Sprintf("%s = '%s'", fld.Name, v.From))
		case "greaterThan":
			filters = append(filters, fmt.Sprintf("%s > %s", fld.Name, v.From))
		case "lessThan":
			filters = append(filters, fmt.Sprintf("%s < %s", fld.Name, v.From))
		case "startsWith":
			filters = append(filters, fmt.Sprintf("%s ILIKE '%s%%'", fld.Name, v.From))
		case "endsWith":
			filters = append(filters, fmt.Sprintf("%s ILIKE '%%%s'", fld.Name, v.From))
		case "greaterThanOrEqual":
			filters = append(filters, fmt.Sprintf("%s >= %s", fld.Name, v.From))
		case "lessThanOrEqual":
			filters = append(filters, fmt.Sprintf("%s <= %s", fld.Name, v.From))
		case "inRange":
			if fld.Type.Kind() == reflect.String {
				filters = append(filters, fmt.Sprintf("%s >= '%s'", fld.Name, v.From))
				filters = append(filters, fmt.Sprintf("%s <= '%s'", fld.Name, v.To))

			} else {
				filters = append(filters, fmt.Sprintf("%s >= %s", fld.Name, v.From))
				filters = append(filters, fmt.Sprintf("%s <= %s", fld.Name, v.To))
			}
		}
	}

	return strings.Join(filters, " AND ")
}

func GetSortQuery[T any](filter *dto.DynamicFilter) string {

	if filter.Sort == nil {
		return ""
	}

	model := new(T)
	modelType := reflect.TypeOf(*model)

	sorts := []string{}
	for _, v := range *filter.Sort {
		fld, ok := modelType.FieldByName(v.ColId)
		if !ok || (v.Sort != "asc" && v.Sort != "desc") {
			continue
		}
		fld.Name = common.ToSnakeCase(fld.Name)
		sorts = append(sorts, fmt.Sprintf("%s %s", fld.Name, v.Sort))
	}

	return strings.Join(sorts, ", ")
}

func Paginate[T any, Tr any](p *dto.PaginationInputWithFilter, preloads []preload, database *gorm.DB) (*dto.PagedList[Tr], error) {

	model := new(T)
	var items *[]T
	var rItems *[]Tr

	var totalRows int64

	query := GetFilterQuery[T](&p.DynamicFilter)
	sort := GetSortQuery[T](&p.DynamicFilter)

	database = Preload(database, preloads)

	database.Model(model).Where(query).Count(&totalRows)

	err := database.Model(model).
		Where(query).
		Limit(p.GetPageSize()).Offset(p.GetOffset()).Order(sort).Find(&items).Error
	if err != nil {
		return nil, err
	}

	rItems, err = common.ConvertType[[]Tr](items)

	if err != nil {
		return nil, err
	}

	return MakePagedList(rItems, totalRows, p.GetPageSize(), p.GetPageNumber()), nil

}

func MakePagedList[T any](items *[]T, totalRowsCount int64, pageSize int, pageNumber int) *dto.PagedList[T] {
	pagedList := &dto.PagedList[T]{
		PageNumber: pageNumber,
		TotalRows:  totalRowsCount,
		TotalPages: int(math.Ceil(float64(totalRowsCount) / float64(pageSize))),
		Items:      items,
	}
	pagedList.HasPreviousPage = pageNumber > 1
	pagedList.HasNextPage = pageNumber < pagedList.TotalPages

	return pagedList
}

func Preload(database *gorm.DB, preloads []preload) *gorm.DB {
	for _, v := range preloads {
		database = database.Preload(v.string)
	}
	return database
}
