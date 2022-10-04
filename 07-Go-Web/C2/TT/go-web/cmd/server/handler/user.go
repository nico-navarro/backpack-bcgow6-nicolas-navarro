package handler

import (
	"go-web/internal/users"

	"github.com/gin-gonic/gin"
)

type request struct {
	Name   string `json:"name" binding:"required"`
	Email  string `binding:"required"`
	Age    int    `binding:"required"`
	Height int    `binding:"required"`
	Active bool   `binding:"required"`
	Date   string `binding:"required"`
}

type UserController struct {
	service users.Service
}

func NewUserController(u users.Service) *UserController {
	return &UserController{
		service: u,
	}
}

func (c *UserController) GetAll(ctx *gin.Context) {
	token := ctx.Request.Header.Get("token")
	if token != "123456" {
		ctx.JSON(401, gin.H{
			"error": "token inválido",
		})
		return
	}

	users, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, users)
}

func (c *UserController) Store(ctx *gin.Context) {
	token := ctx.Request.Header.Get("token")
	if token != "123456" {
		ctx.JSON(401, gin.H{"error": "token inválido"})
		return
	}
	var req request
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := c.service.Store(req.Name, req.Email, req.Age, req.Height, req.Active, req.Date)
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, user)
}
