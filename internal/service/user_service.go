package service

import (
	"context"
	"errors"
	"sponge/internal/dao"
	"time"

	"sponge/internal/model"
	"sponge/pkg/global"
	"sponge/pkg/utils"
)

type UserService struct {
	userDao *dao.UserDao // 明确依赖 DAO
}

// NewUserService 构造函数，Wire 会自动注入 dao
func NewUserService(dao *dao.UserDao) *UserService {
	return &UserService{userDao: dao}
}

// Register 注册逻辑
func (s *UserService) Register(ctx context.Context, username, password string) error {
	// 检查用户是否已存在
	existingUser, err := s.userDao.GetUserByUsername(ctx, username)
	if err == nil && existingUser != nil {
		return errors.New("用户已存在！")
	}

	// 加密密码
	hashPwd, _ := utils.CryptPassword(password)

	user := model.User{
		UserName: username,
		Password: hashPwd,
		NickName: "默认用户",
	}

	// 插入数据库 (ID 会由 BeforeCreate 自动生成)
	return s.userDao.CreateUser(ctx, &user)
}

// Login 登录逻辑
// 返回：Token, UserID, UserName, Error
func (s *UserService) Login(ctx context.Context, username, password string) (string, string, string, error) {
	// 查库
	user, err := s.userDao.GetUserByUsername(ctx, username)
	if err != nil {
		return "", "", "", errors.New("用户或密码错误")
	}
	// 校验密码
	if !utils.CheckPassword(password, user.Password) {
		return "", "", "", errors.New("用户或密码错误")
	}
	// 生成 Token
	token, err := utils.GenerateToken(user.ID, user.UserName)
	if err != nil {
		return "", "", "", err
	}
	// 存入 Redis (实现单点登录或黑名单机制，同时用于自动续期)
	redisKey := "login:token:" + user.ID
	err = global.Redis.Set(ctx, redisKey, token, 24*time.Hour)
	if err != nil {
		return "", "", "", err
	}
	return token, user.ID, user.UserName, nil
}

// GetUserProfile 获取用户信息
func (s *UserService) GetUserProfile(ctx context.Context, userID string) (*model.User, error) {
	return s.userDao.GetUserByID(ctx, userID)
}
