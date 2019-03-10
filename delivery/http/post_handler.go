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
func NewPostHandler(c *echo.Echo, pU post.Usecase) {
	handler := &PostHandler{pU}
	// middleware := &Middleware{}

	g := c.Group("/posts")
	// p.Use(middleware.JWTAuth())

	g.GET("/:id", handler.GetByID)
	g.POST("", handler.Store)

	gG := c.Group("/postsGroups")
	// pG.Use(middleware.JWTAuth())

	gG.GET("", handler.GetGroups)
	gG.GET("/tag/:tag", handler.GetGroupsByTag)
}

//GetByID ...
func (pH *PostHandler) GetByID(c echo.Context) error {
	idP := c.Param("id")

	post, err := pH.pUsecase.GetByID(idP)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, post)
}

//Store ...
func (pH *PostHandler) Store(c echo.Context) error {
	var p models.Post
	c.Bind(&p)

	err := c.Validate(p)

	fmt.Println(err)

	return c.JSON(http.StatusOK, "test")
}

// GetGroups ...
func (pH *PostHandler) GetGroups(c echo.Context) error {

	pagination := models.Pagination{Start: 0, Limit: 100}

	if startP, ok := c.QueryParams()["start"]; ok {
		start, err := strconv.Atoi(startP[0])
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
		}
		pagination.Start = start
	}

	if limitP, ok := c.QueryParams()["limit"]; ok {
		limit, err := strconv.Atoi(limitP[0])
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
		}
		pagination.Limit = limit
	}

	postsGroups, err := pH.pUsecase.GetGroups(pagination)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, postsGroups)
}

//GetGroupsByTag ...
func (pH *PostHandler) GetGroupsByTag(c echo.Context) error {

	tag := c.Param("tag")

	pagination := models.Pagination{Start: 0, Limit: 100}

	if startP, ok := c.QueryParams()["start"]; ok {
		start, err := strconv.Atoi(startP[0])
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
		}
		pagination.Start = start
	}

	if limitP, ok := c.QueryParams()["limit"]; ok {
		limit, err := strconv.Atoi(limitP[0])
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
		}
		pagination.Limit = limit
	}

	postsGroups, err := pH.pUsecase.GetGroupsByTag(tag, pagination)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, postsGroups)
}
