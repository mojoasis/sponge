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
func (s *UserService) Register(username, password string) error {
	var count int64
	global.DB.Model(&model.User{}).Where("user_name = ?", username).Count(&count)
	if count > 0 {
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
	return s.userDao.CreateUser(context.Background(), &user)
}

// Login 登录逻辑
// 返回：Token, UserInfo, Error
func (s *UserService) Login(ctx context.Context, username, password string) (string, *model.User, error) {
	// 查库
	user, err := s.userDao.GetUserByUsername(ctx, username)
	if err != nil {
		return "", nil, errors.New("用户或密码错误")
	}
	// 校验密码
	if !utils.CheckPassword(password, user.Password) {
		return "", nil, errors.New("用户或密码错误")
	}
	// 生成 Token
	token, err := utils.GenerateToken(user.ID, user.UserName)
	if err != nil {
		return "", nil, err
	}
	// 存入 Redis (实现单点登录或黑名单机制，同时用于自动续期)
	redisKey := "login:token:" + user.ID
	err = global.Redis.SetObject(ctx, redisKey, token, 24*time.Hour)
	return token, user, err
}
