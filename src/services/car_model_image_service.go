package services

import (
	"context"

	"github.com/sajadblnyn/autocar-apis/apis/dto"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/data/models"
)

type CarModelImageService struct {
	base *BaseService[models.CarModelImage, dto.CreateCarModelImageRequest, dto.UpdateCarModelImageRequest, dto.CarModelImageResponse]
}

func NewCarModelImageService(cfg *config.Config) *CarModelImageService {

	return &CarModelImageService{base: (NewBaseService[models.CarModelImage, dto.CreateCarModelImageRequest, dto.UpdateCarModelImageRequest, dto.CarModelImageResponse](cfg, []preload{
		{string: "Image"}}))}
}

func (s *CarModelImageService) CreateCarModelImage(c context.Context, r *dto.CreateCarModelImageRequest) (*dto.CarModelImageResponse, error) {
	return s.base.Create(c, r)
}

func (s *CarModelImageService) UpdateCarModelImage(c context.Context, id int, r *dto.UpdateCarModelImageRequest) (*dto.CarModelImageResponse, error) {
	return s.base.Update(c, id, r)
}

func (s *CarModelImageService) DeleteCarModelImage(c context.Context, id int) error {

	return s.base.Delete(c, id)

}

func (s *CarModelImageService) GetCarModelImage(c context.Context, id int) (*dto.CarModelImageResponse, error) {
	return s.base.GetById(c, id)
}

func (s *CarModelImageService) GetCarModelImageByFilter(c context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelImageResponse], error) {
	return s.base.GetByFilter(c, req)
}
