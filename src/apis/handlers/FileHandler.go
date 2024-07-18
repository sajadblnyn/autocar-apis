package handlers

import (
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sajadblnyn/autocar-apis/apis/dto"
	"github.com/sajadblnyn/autocar-apis/apis/helper"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/pkg/logging"
	"github.com/sajadblnyn/autocar-apis/services"
)

var logger logging.Logger

type FileHandler struct {
	service *services.FileService
}

func NewFileHandler(cfg *config.Config) *FileHandler {
	logger = logging.NewLogger(cfg)
	return &FileHandler{service: services.NewFileService(cfg)}
}

func (h *FileHandler) CreateFile(c *gin.Context) {
	req := dto.UploadFileRequest{}
	err := c.ShouldBind(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationErrors(nil, false, int(helper.ValidationError), err))
		return
	}
	createFileReq := dto.CreateFileRequest{}
	createFileReq.Description = req.Description
	createFileReq.Directory = "files"
	createFileReq.MediaType = req.File.Header.Get("Content-Type")
	createFileReq.Name, err = saveFile(req.File, createFileReq.Directory)

	if err != nil {
		logger.Error(logging.Io, logging.UploadFile, err.Error(), nil)
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err))
		return
	}

	res, err := h.service.CreateFile(c, &createFileReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err))
		return

	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, int(helper.Success)))

}

func saveFile(file *multipart.FileHeader, dir string) (string, error) {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", err
	}

	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	ct := file.Header.Get("Content-Type")
	mimeType, _, err := mime.ParseMediaType(ct)
	if err != nil {
		return "", err
	}

	extension, err := mime.ExtensionsByType(mimeType)
	if err != nil {
		return "", err
	}

	fname := uuid.New().String()

	sf, err := os.Create(dir + "/" + fname + extension[0])

	if err != nil {
		return "", err
	}
	defer sf.Close()

	_, err = io.Copy(sf, f)

	if err != nil {
		return "", err
	}

	return fname + extension[0], nil

}

func (h *FileHandler) UpdateFile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	req := dto.UpdateFileRequest{}
	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationErrors(nil, false, int(helper.ValidationError), err))
		return
	}
	res, err := h.service.UpdateFile(c, id, &req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err))
		return

	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, int(helper.Success)))

}

func (h *FileHandler) DeleteFile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	f, err := h.service.GetFile(c, id)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err))
		return
	}
	err = os.Remove(f.Directory + "/" + f.Name)
	if err != nil {
		logger.Error(logging.Io, logging.RemoveFile, err.Error(), nil)
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err))
		return
	}
	err = h.service.DeleteFile(c, id)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err))
		return

	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(nil, true, int(helper.Success)))
}

func (h *FileHandler) GetFile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	res, err := h.service.GetFile(c, id)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err))
		return

	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, int(helper.Success)))
}

func (h *FileHandler) GetFilesByFilter(c *gin.Context) {
	req := dto.PaginationInputWithFilter{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationErrors(nil, false, int(helper.ValidationError), err))
		return
	}

	res, err := h.service.GetFileByFilter(c, &req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, int(helper.Success)))
}
