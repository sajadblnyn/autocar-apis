package services

import (
	"context"

	"github.com/sajadblnyn/autocar-apis/apis/dto"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/data/models"
)

type ColorService struct {
	base *BaseService[models.Color, dto.CreateColorRequest, dto.UpdateColorRequest, dto.ColorResponse]
}

func NewColorService(cfg *config.Config) *ColorService {

	return &ColorService{base: (NewBaseService[models.Color, dto.CreateColorRequest, dto.UpdateColorRequest, dto.ColorResponse](cfg, []preload{}))}
}

func (s *ColorService) CreateColor(c context.Context, r *dto.CreateColorRequest) (*dto.ColorResponse, error) {
	return s.base.Create(c, r)
}

func (s *ColorService) UpdateColor(c context.Context, id int, r *dto.UpdateColorRequest) (*dto.ColorResponse, error) {
	return s.base.Update(c, id, r)
}

func (s *ColorService) DeleteColor(c context.Context, id int) error {

	return s.base.Delete(c, id)

}

func (s *ColorService) GetColor(c context.Context, id int) (*dto.ColorResponse, error) {
	return s.base.GetById(c, id)
}

func (s *ColorService) GetColorByFilter(c context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.ColorResponse], error) {
	return s.base.GetByFilter(c, req)
}
