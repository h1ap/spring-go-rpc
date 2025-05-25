package data

import (
	"context"
	"encoding/json"
	"spring-go-rpc/app/system/service/conf"
	"time"

	"github.com/go-kratos/kratos/v2/middleware/logging"
	mmd "github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/http"

	v1 "spring-go-rpc/api/gen/go/system/service/v1"
	"spring-go-rpc/app/system/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	r            registry.Discovery
	log          *log.Helper
	remoteClient v1.UserServiceHTTPClient
}

// NewUserRepo 创建用户仓储
func NewUserRepo(r registry.Discovery, logger log.Logger) biz.UserRepo {

	conn, err := http.NewClient(
		context.Background(),
		http.WithMiddleware(
			recovery.Recovery(),
			logging.Client(logger),
			mmd.Client(),
		),
		http.WithDiscovery(r),
		http.WithEndpoint("discovery:///spring-go-rpc-spring"),
		http.WithBlock(),
		http.WithTimeout(5*time.Second),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	return &userRepo{
		r:            r,
		log:          log.NewHelper(logger),
		remoteClient: v1.NewUserServiceHTTPClient(conn),
	}
}

// GetAllUsers 获取所有用户
func (r *userRepo) GetAllUsers(ctx context.Context, req *v1.GetAllUsersReqVO) (*v1.GetAllUsersRespVO, error) {
	r.log.WithContext(ctx).Infof("调用Spring Boot服务获取所有用户")

	_, c := conf.GetConfig()
	value, _ := c.Value("spring.go.config.spring.user").String()
	r.log.WithContext(ctx).Infof(value)

	// 通过HTTP调用Spring Boot的用户服务，传递空请求
	rsp, err := r.remoteClient.GetAllUsers(ctx, &v1.GetAllUsersReqVO{})
	if err != nil {
		r.log.WithContext(ctx).Errorf("调用Spring Boot服务失败: %v", err)
		return nil, err
	}
	
	var remoteUser v1.UserInfo
	json.Unmarshal([]byte(value), &remoteUser)
	rsp.Data = append(rsp.Data, &remoteUser)

	r.log.WithContext(ctx).Infof("成功获取用户列表，响应码: %d, 消息: %s, 用户数量: %d",
		rsp.Code, rsp.Message, len(rsp.Data))
	return rsp, nil
}

// GetUserById 根据ID获取用户
func (r *userRepo) GetUserById(ctx context.Context, req *v1.GetUserByIdReqVO) (*v1.GetUserByIdRespVO, error) {
	r.log.WithContext(ctx).Infof("调用Spring Boot服务根据ID获取用户, userId: %d", req.Id)

	// 通过HTTP调用Spring Boot的用户服务
	resp, err := r.remoteClient.GetUserById(ctx, req)
	if err != nil {
		r.log.WithContext(ctx).Errorf("调用Spring Boot服务失败: %v", err)
		return nil, err
	}

	r.log.WithContext(ctx).Infof("成功获取用户信息，响应码: %d, 消息: %s, 用户: %+v",
		resp.Code, resp.Message, resp.Data)
	return resp, nil
}

// CreateUser 创建用户
func (r *userRepo) CreateUser(ctx context.Context, req *v1.CreateUserReqVO) (*v1.CreateUserRespVO, error) {
	r.log.WithContext(ctx).Infof("调用Spring Boot服务创建用户, username: %s, email: %s", req.Username, req.Email)

	// 通过HTTP调用Spring Boot的用户服务
	resp, err := r.remoteClient.CreateUser(ctx, req)
	if err != nil {
		r.log.WithContext(ctx).Errorf("调用Spring Boot服务失败: %v", err)
		return nil, err
	}

	r.log.WithContext(ctx).Infof("成功创建用户，响应码: %d, 消息: %s, 用户: %+v",
		resp.Code, resp.Message, resp.Data)
	return resp, nil
}

// UpdateUser 更新用户
func (r *userRepo) UpdateUser(ctx context.Context, req *v1.UpdateUserReqVO) (*v1.UpdateUserRespVO, error) {
	r.log.WithContext(ctx).Infof("调用Spring Boot服务更新用户, userId: %d", req.Id)

	// 通过HTTP调用Spring Boot的用户服务
	resp, err := r.remoteClient.UpdateUser(ctx, req)
	if err != nil {
		r.log.WithContext(ctx).Errorf("调用Spring Boot服务失败: %v", err)
		return nil, err
	}

	r.log.WithContext(ctx).Infof("成功更新用户，响应码: %d, 消息: %s, 用户: %+v",
		resp.Code, resp.Message, resp.Data)
	return resp, nil
}

// DeleteUser 删除用户
func (r *userRepo) DeleteUser(ctx context.Context, req *v1.DeleteUserReqVO) (*v1.DeleteUserRespVO, error) {
	r.log.WithContext(ctx).Infof("调用Spring Boot服务删除用户, userId: %d", req.Id)

	// 通过HTTP调用Spring Boot的用户服务
	resp, err := r.remoteClient.DeleteUser(ctx, req)
	if err != nil {
		r.log.WithContext(ctx).Errorf("调用Spring Boot服务失败: %v", err)
		return nil, err
	}

	r.log.WithContext(ctx).Infof("删除用户操作完成，响应码: %d, 消息: %s", resp.Code, resp.Message)
	return resp, nil
}

// HealthCheck 健康检查
func (r *userRepo) HealthCheck(ctx context.Context, req *v1.HealthCheckReqVO) (*v1.HealthCheckRespVO, error) {
	r.log.WithContext(ctx).Infof("调用Spring Boot服务健康检查")

	// 通过HTTP调用Spring Boot的用户服务
	resp, err := r.remoteClient.HealthCheck(ctx, req)
	if err != nil {
		r.log.WithContext(ctx).Errorf("调用Spring Boot服务失败: %v", err)
		return nil, err
	}

	r.log.WithContext(ctx).Infof("健康检查完成，响应码: %d, 消息: %s, 数据: %s",
		resp.Code, resp.Message, resp.Data)
	return resp, nil
}
