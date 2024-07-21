package services

import (
	"context"

	"github.com/sajadblnyn/autocar-apis/apis/dto"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/data/models"
)

type CarTypeService struct {
	base *BaseService[models.CarType, dto.CreateCarTypeRequest, dto.UpdateCarTypeRequest, dto.CarTypeResponse]
}

func NewCarTypeService(cfg *config.Config) *CarTypeService {

	return &CarTypeService{base: (NewBaseService[models.CarType, dto.CreateCarTypeRequest, dto.UpdateCarTypeRequest, dto.CarTypeResponse](cfg, []preload{}))}
}

func (s *CarTypeService) CreateCarType(c context.Context, r *dto.CreateCarTypeRequest) (*dto.CarTypeResponse, error) {
	return s.base.Create(c, r)
}

func (s *CarTypeService) UpdateCarType(c context.Context, id int, r *dto.UpdateCarTypeRequest) (*dto.CarTypeResponse, error) {
	return s.base.Update(c, id, r)
}

func (s *CarTypeService) DeleteCarType(c context.Context, id int) error {

	return s.base.Delete(c, id)

}

func (s *CarTypeService) GetCarType(c context.Context, id int) (*dto.CarTypeResponse, error) {
	return s.base.GetById(c, id)
}

func (s *CarTypeService) GetCarTypeByFilter(c context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarTypeResponse], error) {
	return s.base.GetByFilter(c, req)
}
