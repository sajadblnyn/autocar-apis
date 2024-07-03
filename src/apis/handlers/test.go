package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestHandler struct {
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (t *TestHandler) BindQuery(c *gin.Context) {
	c.JSON(http.StatusOK, struct {
		Id string `json:"id"`
	}{
		Id: c.Query("id"),
	})
	return
}

func (t *TestHandler) BindQueryArray(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"Ids": c.QueryArray("id"),
	})
	return
}

func (t *TestHandler) BindHeader1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"UserId": c.GetHeader("UserId"),
	})
	return
}

func (t *TestHandler) BindHeader2(c *gin.Context) {
	h := struct {
		UserId  int
		Browser string
	}{}
	c.ShouldBindHeader(&h)
	c.JSON(http.StatusOK, h)
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Body": b,
	})
	return
}

func (t *TestHandler) BindForm(c *gin.Context) {
	b := struct {
		Name   string
		Family string
	}{}
	c.ShouldBind(&b)
	c.JSON(http.StatusOK, gin.H{
		"Form": b,
	})
	return
}

func (t *TestHandler) BindFormFile(c *gin.Context) {
	f, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	err = c.SaveUploadedFile(f, "file.txt")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"file": f.Filename,
	})
	return
}

func (t *TestHandler) BindUri(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id": c.Param("id"),
	})
	return
}
