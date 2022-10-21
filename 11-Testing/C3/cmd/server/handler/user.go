package handler

import (
	"C3/internal/users"
	"C3/pkg/web"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type request struct {
	Name   string `json:"name" binding:"required"`
	Email  string `binding:"required"`
	Age    *int   `binding:"required"`
	Height *int   `binding:"required"`
	Active *bool  `binding:"required"`
	Date   string `binding:"required"`
}

type patchRequest struct {
	Name   string
	Email  string
	Age    *int
	Height *int
	Active *bool
	Date   string
}

type UserController struct {
	service users.Service
}

func NewUserController(u users.Service) *UserController {
	return &UserController{
		service: u,
	}
}

// ListUsers godoc
// @Summary List users
// @Tags Users
// @Description get users
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /users [get]
func (c *UserController) GetAll(ctx *gin.Context) {
	token := ctx.Request.Header.Get("token")
	if token != os.Getenv("TOKEN") {
		ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
		return
	}

	users, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		return
	}
	ctx.JSON(200, web.NewResponse(200, users, ""))
}

func (c *UserController) Store(ctx *gin.Context) {
	token := ctx.Request.Header.Get("token")
	if token != os.Getenv("TOKEN") {
		ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
		return
	}
	var req request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		return
	}
	user, err := c.service.Store(req.Name, req.Email, *req.Age, *req.Height, *req.Active, req.Date)
	if err != nil {
		ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		return
	}
	ctx.JSON(200, web.NewResponse(200, user, ""))
}

func (c *UserController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "ID Not Found"))
		return
	}

	var req request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}

	user, err := c.service.Update(id, req.Name, req.Email, *req.Age, *req.Height, *req.Active, req.Date)
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
		return
	}
	ctx.JSON(200, web.NewResponse(200, user, ""))
}

func (c *UserController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "ID Not Found"))
		return
	}

	user, err := c.service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
		return
	}

	ctx.JSON(200, web.NewResponse(200, user, ""))
}

func (c *UserController) Patch(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "ID Not Found"))
		return
	}

	var req patchRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}

	user, err := c.service.Patch(id, req.Name, req.Email, req.Age, req.Height, req.Active, req.Date)
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
		return
	}
	ctx.JSON(200, web.NewResponse(200, user, ""))
}
