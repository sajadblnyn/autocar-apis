package services

import (
	"context"

	"github.com/sajadblnyn/autocar-apis/apis/dto"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/data/models"
)

type CarModelColorService struct {
	base *BaseService[models.CarModelColor, dto.CreateCarModelColorRequest, dto.UpdateCarModelColorRequest, dto.CarModelColorResponse]
}

func NewCarModelColorService(cfg *config.Config) *CarModelColorService {

	return &CarModelColorService{base: (NewBaseService[models.CarModelColor, dto.CreateCarModelColorRequest, dto.UpdateCarModelColorRequest, dto.CarModelColorResponse](cfg, []preload{
		{string: "Color"}}))}
}

func (s *CarModelColorService) CreateCarModelColor(c context.Context, r *dto.CreateCarModelColorRequest) (*dto.CarModelColorResponse, error) {
	return s.base.Create(c, r)
}

func (s *CarModelColorService) UpdateCarModelColor(c context.Context, id int, r *dto.UpdateCarModelColorRequest) (*dto.CarModelColorResponse, error) {
	return s.base.Update(c, id, r)
}

func (s *CarModelColorService) DeleteCarModelColor(c context.Context, id int) error {

	return s.base.Delete(c, id)

}

func (s *CarModelColorService) GetCarModelColor(c context.Context, id int) (*dto.CarModelColorResponse, error) {
	return s.base.GetById(c, id)
}

func (s *CarModelColorService) GetCarModelColorByFilter(c context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelColorResponse], error) {
	return s.base.GetByFilter(c, req)
}
