package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sajadblnyn/autocar-apis/apis/helper"
)

type TestHandler struct {
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (t *TestHandler) BindQuery(c *gin.Context) {
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(struct {
		Id string `json:"id"`
	}{
		Id: c.Query("id"),
	}, true, 1))
	return
}

func (t *TestHandler) BindQueryArray(c *gin.Context) {

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"Ids": c.QueryArray("id"),
	}, true, 1))
	return
}

func (t *TestHandler) BindHeader1(c *gin.Context) {
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"UserId": c.GetHeader("UserId"),
	}, true, 1))
	return
}

func (t *TestHandler) BindHeader2(c *gin.Context) {
	h := struct {
		UserId  int
		Browser string
	}{}
	c.ShouldBindHeader(&h)
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(h, true, 1))
	return
}

func (t *TestHandler) BindJsonBody(c *gin.Context) {
	b := struct {
		Name   string `json:"name" binding:"required,alpha,min=3,max=8"`
		Family string `json:"family" binding:"required,alpha,min=3,max=8"`
		Mobile string `json:"mobile" binding:"required,numeric,iran-mobile-validator"`
	}{}
	err := c.ShouldBindJSON(&b)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationErrors(nil, false, -1, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(b, true, 1))
	return
}

func (t *TestHandler) BindForm(c *gin.Context) {
	b := struct {
		Name   string
		Family string
	}{}
	c.ShouldBind(&b)
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(b, true, 1))
	return
}

func (t *TestHandler) BindFormFile(c *gin.Context) {
	f, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.GenerateBaseResponseWithError(nil, false, -1, err))
	}
	err = c.SaveUploadedFile(f, "file.txt")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"file": f.Filename,
	}, true, 1))
	return
}

func (t *TestHandler) BindUri(c *gin.Context) {
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"id": c.Param("id"),
	}, true, 1))
	return
}
