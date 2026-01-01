package service

import (
	"context"
	"errors"
	"fmt"
	"sponge/internal/consts"
	"sponge/internal/dao"
	"sponge/internal/model/dto"
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
func (s *UserService) Register(ctx context.Context, req *dto.RegisterReq) error {
	// 检查用户是否已存在
	existingUser, err := s.userDao.GetUserByUsername(ctx, req.UserName)
	if err == nil && existingUser != nil {
		return errors.New("用户已存在！")
	}
	// dto -> model
	user := model.User{}
	if err := utils.ToModel(req, &user); err != nil {
		return err
	}
	// 加密密码
	user.Password, _ = utils.CryptPassword(req.Password)
	// 插入数据库 (ID 会由 BeforeCreate 自动生成)
	return s.userDao.CreateUser(ctx, &user)
}

// Login 登录逻辑
// 返回：Token, UserID, UserName, Error
func (s *UserService) Login(ctx context.Context, username, password string) (string, int64, string, error) {
	// 查库
	user, err := s.userDao.GetUserByUsername(ctx, username)
	if err != nil {
		return "", 0, "", errors.New("用户或密码错误")
	}

	// 检查用户状态
	if user.Status != model.UserStatusNormal {
		return "", 0, "", errors.New("用户已被禁用或注销")
	}

	// 校验密码
	if !utils.CheckPassword(password, user.Password) {
		return "", 0, "", errors.New("用户或密码错误")
	}

	// 生成 Token
	token, err := utils.GenerateToken(user.ID, user.UserName)
	if err != nil {
		return "", 0, "", err
	}

	// 存入 Redis (实现单点登录或黑名单机制，同时用于自动续期)
	redisKey := fmt.Sprintf("%s%d", consts.RedisKeyLoginToken, user.ID)
	err = global.Redis.Set(ctx, redisKey, token, 24*time.Hour)
	if err != nil {
		return "", 0, "", err
	}

	return token, user.ID, user.UserName, nil
}

// GetUserProfile 获取用户信息
func (s *UserService) GetUserProfile(ctx context.Context, userID int64) (*model.User, error) {
	return s.userDao.GetUserByID(ctx, userID)
}

// GetUsersByIDs 批量获取用户信息（性能优化）
func (s *UserService) GetUsersByIDs(ctx context.Context, userIDs []int64) ([]*model.User, error) {
	return s.userDao.GetUsersByIDs(ctx, userIDs)
}

// UpdateUserProfile 更新用户信息逻辑
func (s *UserService) UpdateUserProfile(ctx context.Context, id int64, req *dto.UpdateUserProfileReq) (*model.User, error) {
	// 先查出旧数据
	user, err := s.userDao.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	// 将 req 中非 nil 的字段覆盖到 user 对象上
	if err := utils.CopyToModelForUpdate(req, user); err != nil {
		return nil, err
	}
	// 将对象交给 DAO 保存
	if err := s.userDao.UpdateUserInfo(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

// UpdatePassword 更新密码逻辑
func (s *UserService) UpdatePassword(ctx context.Context, uid int64, req *dto.UpdatePasswordReq) error {
	// 获取用户信息
	user, err := s.userDao.GetUserByID(ctx, uid)
	if err != nil {
		return errors.New("用户不存在")
	}
	// 校验旧密码
	if !utils.CheckPassword(user.Password, req.OldPassword) {
		return errors.New("原密码错误")
	}
	// 加密新密码
	hashedPassword, _ := utils.CryptPassword(req.NewPassword)
	// 更新
	return s.userDao.UpdateColumn(ctx, uid, "password", hashedPassword)
}

// UpdatePhone 更新手机号逻辑
func (s *UserService) UpdatePhone(ctx context.Context, uid int64, req *dto.UpdatePhoneReq) error {
	// 校验验证码 (假设从 Redis 中读取)
	// cacheCode, _ := global.Redis.Get(ctx, "code:"+req.NewPhone)
	// if cacheCode != req.Code { return errors.New("验证码错误") }
	// 检查新手机号是否已被占用
	exists, _ := s.userDao.CheckPhoneExists(ctx, req.NewPhone)
	if exists {
		return errors.New("该手机号已被绑定")
	}
	// 更新
	return s.userDao.UpdateColumn(ctx, uid, "phone", req.NewPhone)
}
