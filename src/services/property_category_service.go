package services

import (
	"context"

	"github.com/sajadblnyn/autocar-apis/apis/dto"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/data/models"
)

type PropertyCategoryService struct {
	base *BaseService[models.PropertyCategory, dto.CreatePropertyCategoryRequest, dto.UpdatePropertyCategoryRequest, dto.GetPropertyCategoryResponse]
}

func NewPropertyCategoryService(cfg *config.Config) *PropertyCategoryService {

	return &PropertyCategoryService{base: (NewBaseService[models.PropertyCategory, dto.CreatePropertyCategoryRequest, dto.UpdatePropertyCategoryRequest, dto.GetPropertyCategoryResponse](cfg, []preload{{string: "Properties"}}))}
}

func (s *PropertyCategoryService) CreatePropertyCategory(c context.Context, r *dto.CreatePropertyCategoryRequest) (*dto.GetPropertyCategoryResponse, error) {
	return s.base.Create(c, r)
}

func (s *PropertyCategoryService) UpdatePropertyCategory(c context.Context, id int, r *dto.UpdatePropertyCategoryRequest) (*dto.GetPropertyCategoryResponse, error) {
	return s.base.Update(c, id, r)
}

func (s *PropertyCategoryService) DeletePropertyCategory(c context.Context, id int) error {

	return s.base.Delete(c, id)

}

func (s *PropertyCategoryService) GetPropertyCategory(c context.Context, id int) (*dto.GetPropertyCategoryResponse, error) {
	return s.base.GetById(c, id)
}

func (s *PropertyCategoryService) GetPropertyCategoryByFilter(c context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.GetPropertyCategoryResponse], error) {
	return s.base.GetByFilter(c, req)
}
