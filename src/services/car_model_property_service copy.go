package services

import (
	"context"

	"github.com/sajadblnyn/autocar-apis/apis/dto"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/data/models"
)

type CarModelPropertyService struct {
	base *BaseService[models.CarModelProperty, dto.CreateCarModelPropertyRequest, dto.UpdateCarModelPropertyRequest, dto.CarModelPropertyResponse]
}

func NewCarModelPropertyService(cfg *config.Config) *CarModelPropertyService {

	return &CarModelPropertyService{base: (NewBaseService[models.CarModelProperty, dto.CreateCarModelPropertyRequest, dto.UpdateCarModelPropertyRequest, dto.CarModelPropertyResponse](cfg, []preload{
		{string: "Property"}}))}
}

func (s *CarModelPropertyService) CreateCarModelProperty(c context.Context, r *dto.CreateCarModelPropertyRequest) (*dto.CarModelPropertyResponse, error) {
	return s.base.Create(c, r)
}

func (s *CarModelPropertyService) UpdateCarModelProperty(c context.Context, id int, r *dto.UpdateCarModelPropertyRequest) (*dto.CarModelPropertyResponse, error) {
	return s.base.Update(c, id, r)
}

func (s *CarModelPropertyService) DeleteCarModelProperty(c context.Context, id int) error {

	return s.base.Delete(c, id)

}

func (s *CarModelPropertyService) GetCarModelProperty(c context.Context, id int) (*dto.CarModelPropertyResponse, error) {
	return s.base.GetById(c, id)
}

func (s *CarModelPropertyService) GetCarModelPropertyByFilter(c context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelPropertyResponse], error) {
	return s.base.GetByFilter(c, req)
}
