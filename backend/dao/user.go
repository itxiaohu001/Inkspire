package dao

import (
	"Inkspire/models"

	"gorm.io/gorm"
)

type UserDAO struct {
	db *gorm.DB
}

// NewUserDAO 创建 UserDAO 实例
func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{db: db}
}

// CreateUser 创建新用户
func (dao *UserDAO) CreateUser(user *models.User) error {
	return dao.db.Create(user).Error
}

// GetUserByID 根据 ID 获取用户
func (dao *UserDAO) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := dao.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByUsername 根据用户名获取用户
func (dao *UserDAO) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := dao.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByEmail 根据邮箱获取用户
func (dao *UserDAO) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := dao.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetAllUsers 获取所有用户
func (dao *UserDAO) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	if err := dao.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUsersByRole 根据角色获取用户
func (dao *UserDAO) GetUsersByRole(role string) ([]*models.User, error) {
	var users []*models.User
	if err := dao.db.Where("role = ?", role).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// UpdateUser 更新用户信息
func (dao *UserDAO) UpdateUser(user *models.User) error {
	return dao.db.Save(user).Error
}

// UpdateUserPassword 更新用户密码
func (dao *UserDAO) UpdateUserPassword(id uint, newPassword string) error {
	return dao.db.Model(&models.User{}).Where("id = ?", id).Update("password", newPassword).Error
}

// UpdateUserRole 更新用户角色
func (dao *UserDAO) UpdateUserRole(id uint, newRole string) error {
	return dao.db.Model(&models.User{}).Where("id = ?", id).Update("role", newRole).Error
}

// DeleteUserByID 根据 ID 删除用户
func (dao *UserDAO) DeleteUserByID(id uint) error {
	return dao.db.Delete(&models.User{}, id).Error
}

// DeleteUserByUsername 根据用户名删除用户
func (dao *UserDAO) DeleteUserByUsername(username string) error {
	return dao.db.Where("username = ?", username).Delete(&models.User{}).Error
}

// IsUsernameExists 检查用户名是否存在
func (dao *UserDAO) IsUsernameExists(username string) (bool, error) {
	var count int64
	if err := dao.db.Model(&models.User{}).Where("username = ?", username).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// IsEmailExists 检查邮箱是否存在
func (dao *UserDAO) IsEmailExists(email string) (bool, error) {
	var count int64
	if err := dao.db.Model(&models.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// GetUsersWithPagination 分页获取用户
func (dao *UserDAO) GetUsersWithPagination(offset, limit int) ([]*models.User, error) {
	var users []*models.User
	if err := dao.db.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
