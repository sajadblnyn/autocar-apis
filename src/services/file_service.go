package services

import (
	"context"

	"github.com/sajadblnyn/autocar-apis/apis/dto"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/data/models"
)

type FileService struct {
	base *BaseService[models.File, dto.CreateFileRequest, dto.UpdateFileRequest, dto.GetFileResponse]
}

func NewFileService(cfg *config.Config) *FileService {

	return &FileService{base: (NewBaseService[models.File, dto.CreateFileRequest, dto.UpdateFileRequest, dto.GetFileResponse](cfg, []preload{}))}
}

func (s *FileService) CreateFile(c context.Context, r *dto.CreateFileRequest) (*dto.GetFileResponse, error) {
	return s.base.Create(c, r)
}

func (s *FileService) UpdateFile(c context.Context, id int, r *dto.UpdateFileRequest) (*dto.GetFileResponse, error) {
	return s.base.Update(c, id, r)
}

func (s *FileService) DeleteFile(c context.Context, id int) error {

	return s.base.Delete(c, id)

}

func (s *FileService) GetFile(c context.Context, id int) (*dto.GetFileResponse, error) {
	return s.base.GetById(c, id)
}

func (s *FileService) GetFileByFilter(c context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.GetFileResponse], error) {
	return s.base.GetByFilter(c, req)
}
