package service

import (
	"Inkspire/dao"
	"Inkspire/dto"
	"Inkspire/models"
	"Inkspire/utils"
	"fmt"
)

type UserService struct {
	userDAO *dao.UserDAO
}

// NewUserService 创建 UserService 实例
func NewUserService(userDAO *dao.UserDAO) *UserService {
	return &UserService{userDAO: userDAO}
}

// RegisterUser 注册新用户（包括密码加密）
func (service *UserService) RegisterUser(req *dto.RegisterRequest) error {
	// 检查用户名或邮箱是否已存在
	if exists, _ := service.userDAO.IsUsernameExists(req.Username); exists {
		return fmt.Errorf("用户名已存在")
	}
	if exists, _ := service.userDAO.IsEmailExists(req.Email); exists {
		return fmt.Errorf("邮箱已存在")
	}

	// 加密密码
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return fmt.Errorf("密码加密失败: %v", err)
	}

	// 创建用户
	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		Role:     "user", // 默认角色
	}
	return service.userDAO.CreateUser(user)
}

// AuthenticateUser 验证用户登录
func (service *UserService) AuthenticateUser(username, plainPassword string) (bool, error) {
	// 根据用户名查询用户
	user, err := service.userDAO.GetUserByUsername(username)
	if err != nil {
		return false, fmt.Errorf("用户不存在或用户名错误")
	}

	// 校验密码
	if !utils.CheckPassword(user.Password, plainPassword) {
		return false, fmt.Errorf("密码错误")
	}

	return true, nil
}

func (service *UserService) GetUserResponse(id uint) (*dto.UserResponse, error) {
	user, err := service.userDAO.GetUserByID(id)
	if err != nil {
		return nil, fmt.Errorf("用户不存在或用户名错误")
	}
	return &dto.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}, nil
}
