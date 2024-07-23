package services

import (
	"context"

	"github.com/sajadblnyn/autocar-apis/apis/dto"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/data/models"
)

type CarModelPriceHistoryService struct {
	base *BaseService[models.CarModelPriceHistory, dto.CreateCarModelPriceHistoryRequest, dto.UpdateCarModelPriceHistoryRequest, dto.CarModelPriceHistoryResponse]
}

func NewCarModelPriceHistoryService(cfg *config.Config) *CarModelPriceHistoryService {

	return &CarModelPriceHistoryService{base: (NewBaseService[models.CarModelPriceHistory, dto.CreateCarModelPriceHistoryRequest, dto.UpdateCarModelPriceHistoryRequest, dto.CarModelPriceHistoryResponse](cfg, []preload{}))}
}

func (s *CarModelPriceHistoryService) CreateCarModelPriceHistory(c context.Context, r *dto.CreateCarModelPriceHistoryRequest) (*dto.CarModelPriceHistoryResponse, error) {
	return s.base.Create(c, r)
}

func (s *CarModelPriceHistoryService) UpdateCarModelPriceHistory(c context.Context, id int, r *dto.UpdateCarModelPriceHistoryRequest) (*dto.CarModelPriceHistoryResponse, error) {
	return s.base.Update(c, id, r)
}

func (s *CarModelPriceHistoryService) DeleteCarModelPriceHistory(c context.Context, id int) error {

	return s.base.Delete(c, id)

}

func (s *CarModelPriceHistoryService) GetCarModelPriceHistory(c context.Context, id int) (*dto.CarModelPriceHistoryResponse, error) {
	return s.base.GetById(c, id)
}

func (s *CarModelPriceHistoryService) GetCarModelPriceHistoryByFilter(c context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelPriceHistoryResponse], error) {
	return s.base.GetByFilter(c, req)
}
