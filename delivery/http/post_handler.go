package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/haffjjj/myblog-backend/models"

	"github.com/haffjjj/myblog-backend/usecase/post"
)

//PostHandler ...
type PostHandler struct {
	pUsecase post.Usecase
}

//NewPostHandler ...
func NewPostHandler(e *echo.Echo, pU post.Usecase) {
	handler := &PostHandler{pU}
	middleware := &Middleware{}

	p := e.Group("/posts")
	p.Use(middleware.JWTAuth())

	p.GET("/:id", handler.GetByID)

	pG := e.Group("/postsGroups")
	pG.Use(middleware.JWTAuth())

	pG.GET("", handler.GetGroups)
	pG.GET("/tag/:tag", handler.GetGroupsByTag)
}

//GetByID ...
func (p *PostHandler) GetByID(e echo.Context) error {

	idP := e.Param("id")

	post, err := p.pUsecase.GetByID(idP)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return e.JSON(http.StatusOK, post)
}

// GetGroups ...
func (p *PostHandler) GetGroups(e echo.Context) error {

	pagination := models.Pagination{Start: 0, Limit: 100}

	if startP, ok := e.QueryParams()["start"]; ok {
		start, err := strconv.Atoi(startP[0])
		if err != nil {
			fmt.Println(err)
			return e.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
		}
		pagination.Start = start
	}

	if limitP, ok := e.QueryParams()["limit"]; ok {
		limit, err := strconv.Atoi(limitP[0])
		if err != nil {
			return e.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
		}
		pagination.Limit = limit
	}

	postsGroups, err := p.pUsecase.GetGroups(pagination)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}
	return e.JSON(http.StatusOK, postsGroups)
}

//GetGroupsByTag ...
func (p *PostHandler) GetGroupsByTag(e echo.Context) error {

	tag := e.Param("tag")

	pagination := models.Pagination{Start: 0, Limit: 100}

	if startP, ok := e.QueryParams()["start"]; ok {
		start, err := strconv.Atoi(startP[0])
		if err != nil {
			return e.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
		}
		pagination.Start = start
	}

	if limitP, ok := e.QueryParams()["limit"]; ok {
		limit, err := strconv.Atoi(limitP[0])
		if err != nil {
			return e.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
		}
		pagination.Limit = limit
	}

	postsGroups, err := p.pUsecase.GetGroupsByTag(tag, pagination)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}
	return e.JSON(http.StatusOK, postsGroups)
}
