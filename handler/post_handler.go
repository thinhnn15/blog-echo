package handler

import (
	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/thinhnn15/blog-echo/model"
	req2 "github.com/thinhnn15/blog-echo/model/req"
	"github.com/thinhnn15/blog-echo/repository"
	"net/http"
	"time"
)

type PostHandler struct {
	PostRepo repository.PostRepo
}

func (p* PostHandler) HandleNewPost(c echo.Context) error {
	req := req2.ReqNewPost{}
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	post := model.Post{
		Title:    req.Title,
		Content:  req.Content,
		CreateAt: time.Time{},
	}
	post, err := p.PostRepo.SavePost(c.Request().Context(), post)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Handling successful",
		Data:       post,
	})
}
