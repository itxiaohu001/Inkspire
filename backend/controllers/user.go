package controllers

import (
	"Inkspire/dto"
	"Inkspire/errors"
	"Inkspire/service"
	"Inkspire/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

// Login 用户登录
func (uc *UserController) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errors.ErrInvalidRequestFormat)
		return
	}
	isValid, err := uc.userService.AuthenticateUser(req.UserName, req.Password)
	if err != nil {
		c.JSON(http.StatusOK, errors.ErrUserAuthFailed)
		return
	}
	if !isValid {
		c.JSON(http.StatusOK, errors.ErrUserAuthFailed)
		return
	}

	token, err := utils.GenerateToken(req.UserName, utils.DefaultExpTime)
	if err != nil {
		c.JSON(http.StatusOK, errors.ErrTokenGeneration)
		return
	}

	c.JSON(http.StatusOK, &dto.LoginResponse{
		Basic: dto.SuccessBasic,
		Token: token,
	})
}

// Register 用户注册
func (uc *UserController) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errors.ErrInvalidRequestFormat)
		return
	}
	if err := uc.userService.RegisterUser(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}

// GetUser 获取用户信息
func (uc *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")
	userID, _ := strconv.ParseUint(id, 10, 64)
	response, err := uc.userService.GetUserResponse(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}
	c.JSON(http.StatusOK, response)
}
