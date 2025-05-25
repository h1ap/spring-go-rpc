package biz

import (
	"context"

	v1 "spring-go-rpc/api/gen/go/system/service/v1"

	"github.com/go-kratos/kratos/v2/log"
)

// UserUsecase 用户业务用例
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// UserRepo 用户仓储接口
type UserRepo interface {
	// GetAllUsers 获取所有用户
	GetAllUsers(ctx context.Context, req *v1.GetAllUsersReqVO) (*v1.GetAllUsersRespVO, error)
	// GetUserById 根据ID获取用户
	GetUserById(ctx context.Context, req *v1.GetUserByIdReqVO) (*v1.GetUserByIdRespVO, error)
	// CreateUser 创建用户
	CreateUser(ctx context.Context, req *v1.CreateUserReqVO) (*v1.CreateUserRespVO, error)
	// UpdateUser 更新用户
	UpdateUser(ctx context.Context, req *v1.UpdateUserReqVO) (*v1.UpdateUserRespVO, error)
	// DeleteUser 删除用户
	DeleteUser(ctx context.Context, req *v1.DeleteUserReqVO) (*v1.DeleteUserRespVO, error)
	// HealthCheck 健康检查
	HealthCheck(ctx context.Context, req *v1.HealthCheckReqVO) (*v1.HealthCheckRespVO, error)
}

// NewUserUsecase 创建用户业务用例
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

// GetAllUsers 获取所有用户
func (uc *UserUsecase) GetAllUsers(ctx context.Context, req *v1.GetAllUsersReqVO) (*v1.GetAllUsersRespVO, error) {
	uc.log.WithContext(ctx).Infof("GetAllUsers req: %+v", req)
	return uc.repo.GetAllUsers(ctx, req)
}

// GetUserById 根据ID获取用户
func (uc *UserUsecase) GetUserById(ctx context.Context, req *v1.GetUserByIdReqVO) (*v1.GetUserByIdRespVO, error) {
	uc.log.WithContext(ctx).Infof("GetUserById req: %+v", req)
	return uc.repo.GetUserById(ctx, req)
}

// CreateUser 创建用户
func (uc *UserUsecase) CreateUser(ctx context.Context, req *v1.CreateUserReqVO) (*v1.CreateUserRespVO, error) {
	uc.log.WithContext(ctx).Infof("CreateUser req: %+v", req)
	return uc.repo.CreateUser(ctx, req)
}

// UpdateUser 更新用户
func (uc *UserUsecase) UpdateUser(ctx context.Context, req *v1.UpdateUserReqVO) (*v1.UpdateUserRespVO, error) {
	uc.log.WithContext(ctx).Infof("UpdateUser req: %+v", req)
	return uc.repo.UpdateUser(ctx, req)
}

// DeleteUser 删除用户
func (uc *UserUsecase) DeleteUser(ctx context.Context, req *v1.DeleteUserReqVO) (*v1.DeleteUserRespVO, error) {
	uc.log.WithContext(ctx).Infof("DeleteUser req: %+v", req)
	return uc.repo.DeleteUser(ctx, req)
}

// HealthCheck 健康检查
func (uc *UserUsecase) HealthCheck(ctx context.Context, req *v1.HealthCheckReqVO) (*v1.HealthCheckRespVO, error) {
	uc.log.WithContext(ctx).Infof("HealthCheck req: %+v", req)
	return uc.repo.HealthCheck(ctx, req)
}
