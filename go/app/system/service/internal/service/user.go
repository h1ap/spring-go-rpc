package service

import (
	"context"

	v1 "spring-go-rpc/api/gen/go/system/service/v1"
	"spring-go-rpc/app/system/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// UserService 用户服务
type UserService struct {
	v1.UnimplementedUserServiceServer

	uc  *biz.UserUsecase
	log *log.Helper
}

// NewUserService 创建用户服务
func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

// GetAllUsers 获取所有用户
func (s *UserService) GetAllUsers(ctx context.Context, req *v1.GetAllUsersReqVO) (*v1.GetAllUsersRespVO, error) {
	s.log.WithContext(ctx).Infof("收到获取所有用户请求: %+v", req)
	return s.uc.GetAllUsers(ctx, req)
}

// GetUserById 根据ID获取用户
func (s *UserService) GetUserById(ctx context.Context, req *v1.GetUserByIdReqVO) (*v1.GetUserByIdRespVO, error) {
	s.log.WithContext(ctx).Infof("收到根据ID获取用户请求: %+v", req)
	return s.uc.GetUserById(ctx, req)
}

// CreateUser 创建用户
func (s *UserService) CreateUser(ctx context.Context, req *v1.CreateUserReqVO) (*v1.CreateUserRespVO, error) {
	s.log.WithContext(ctx).Infof("收到创建用户请求: %+v", req)
	return s.uc.CreateUser(ctx, req)
}

// UpdateUser 更新用户
func (s *UserService) UpdateUser(ctx context.Context, req *v1.UpdateUserReqVO) (*v1.UpdateUserRespVO, error) {
	s.log.WithContext(ctx).Infof("收到更新用户请求: %+v", req)
	return s.uc.UpdateUser(ctx, req)
}

// DeleteUser 删除用户
func (s *UserService) DeleteUser(ctx context.Context, req *v1.DeleteUserReqVO) (*v1.DeleteUserRespVO, error) {
	s.log.WithContext(ctx).Infof("收到删除用户请求: %+v", req)
	return s.uc.DeleteUser(ctx, req)
}

// HealthCheck 健康检查
func (s *UserService) HealthCheck(ctx context.Context, req *v1.HealthCheckReqVO) (*v1.HealthCheckRespVO, error) {
	s.log.WithContext(ctx).Infof("收到健康检查请求: %+v", req)
	return s.uc.HealthCheck(ctx, req)
}
