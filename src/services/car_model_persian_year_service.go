package services

import (
	"context"

	"github.com/sajadblnyn/autocar-apis/apis/dto"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/data/models"
)

type CarModelPersianYearService struct {
	base *BaseService[models.CarModelYear, dto.CreateCarModelPersianYearRequest, dto.UpdateCarModelPersianYearRequest, dto.CarModelPersianYearResponse]
}

func NewCarModelPersianYearService(cfg *config.Config) *CarModelPersianYearService {

	return &CarModelPersianYearService{base: (NewBaseService[models.CarModelYear, dto.CreateCarModelPersianYearRequest, dto.UpdateCarModelPersianYearRequest, dto.CarModelPersianYearResponse](cfg, []preload{
		{string: "PersianYear"}, {string: "CarModelPriceHistories"}}))}
}

func (s *CarModelPersianYearService) CreateCarModelPersianYear(c context.Context, r *dto.CreateCarModelPersianYearRequest) (*dto.CarModelPersianYearResponse, error) {
	return s.base.Create(c, r)
}

func (s *CarModelPersianYearService) UpdateCarModelPersianYear(c context.Context, id int, r *dto.UpdateCarModelPersianYearRequest) (*dto.CarModelPersianYearResponse, error) {
	return s.base.Update(c, id, r)
}

func (s *CarModelPersianYearService) DeleteCarModelPersianYear(c context.Context, id int) error {

	return s.base.Delete(c, id)

}

func (s *CarModelPersianYearService) GetCarModelPersianYear(c context.Context, id int) (*dto.CarModelPersianYearResponse, error) {
	return s.base.GetById(c, id)
}

func (s *CarModelPersianYearService) GetCarModelPersianYearByFilter(c context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelPersianYearResponse], error) {
	return s.base.GetByFilter(c, req)
}
