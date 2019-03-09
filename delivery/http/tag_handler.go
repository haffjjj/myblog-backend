package http

import (
	"net/http"

	"github.com/haffjjj/myblog-backend/models"
	"github.com/haffjjj/myblog-backend/usecase/tag"
	"github.com/labstack/echo"
)

//TagHandler ...
type TagHandler struct {
	tUsecase tag.Usecase
}

//NewTagHandler ...
func NewTagHandler(e *echo.Echo, tU tag.Usecase) {
	handler := &TagHandler{tU}
	// middleware := &Middleware{}

	t := e.Group("/tags")
	// t.Use(middleware.JWTAuth())

	t.GET("", handler.Get)
}

//Get ...
func (t *TagHandler) Get(e echo.Context) error {
	tags, err := t.tUsecase.Get()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}
	// fmt.Println(tags)
	return e.JSON(http.StatusOK, tags)
}
