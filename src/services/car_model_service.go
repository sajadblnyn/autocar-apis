package services

import (
	"context"

	"github.com/sajadblnyn/autocar-apis/apis/dto"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/data/models"
)

type CarModelService struct {
	base *BaseService[models.CarModel, dto.CreateCarModelRequest, dto.UpdateCarModelRequest, dto.CarModelResponse]
}

func NewCarModelService(cfg *config.Config) *CarModelService {

	return &CarModelService{base: (NewBaseService[models.CarModel, dto.CreateCarModelRequest, dto.UpdateCarModelRequest, dto.CarModelResponse](cfg, []preload{
		{string: "CarType"}, {string: "Gearbox"},
		{string: "CarModelColors.Color"}, {string: "CarModelYears.PersianYear"}, {string: "CarModelYears.CarModelPriceHistories"},
		{string: "CarModelProperties.Property.Category"}, {string: "CarModelImages.Image"},
		{string: "Company.Country"}, {string: "CarModelComments.User"}}))}
}

func (s *CarModelService) CreateCarModel(c context.Context, r *dto.CreateCarModelRequest) (*dto.CarModelResponse, error) {
	return s.base.Create(c, r)
}

func (s *CarModelService) UpdateCarModel(c context.Context, id int, r *dto.UpdateCarModelRequest) (*dto.CarModelResponse, error) {
	return s.base.Update(c, id, r)
}

func (s *CarModelService) DeleteCarModel(c context.Context, id int) error {

	return s.base.Delete(c, id)

}

func (s *CarModelService) GetCarModel(c context.Context, id int) (*dto.CarModelResponse, error) {
	return s.base.GetById(c, id)
}

func (s *CarModelService) GetCarModelByFilter(c context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelResponse], error) {
	return s.base.GetByFilter(c, req)
}
