package tech.cooplink.springgorpc.feign;

import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Component;
import tech.cooplink.springgorpc.dto.ApiResponse;
import tech.cooplink.springgorpc.dto.UserDTO;

import java.util.Collections;
import java.util.List;

/**
 * 用户服务Feign客户端降级处理类
 * 当服务调用失败时提供兜底响应
 */
@Component
@Slf4j
public class UserFeignClientFallback implements UserFeignClient {
    
    @Override
    public ApiResponse<List<UserDTO>> getAllUsers() {
        log.warn("调用用户服务获取所有用户失败，使用降级处理");
        return ApiResponse.error(503, "用户服务暂时不可用，请稍后重试");
    }
    
    @Override
    public ApiResponse<UserDTO> getUserById(Long id) {
        log.warn("调用用户服务获取用户失败，用户ID: {}", id);
        return ApiResponse.error(503, "用户服务暂时不可用，请稍后重试");
    }
    
    @Override
    public ApiResponse<UserDTO> createUser(UserDTO userDTO) {
        log.warn("调用用户服务创建用户失败，用户信息: {}", userDTO);
        return ApiResponse.error(503, "用户服务暂时不可用，无法创建用户");
    }
    
    @Override
    public ApiResponse<UserDTO> updateUser(Long id, UserDTO userDTO) {
        log.warn("调用用户服务更新用户失败，用户ID: {}, 用户信息: {}", id, userDTO);
        return ApiResponse.error(503, "用户服务暂时不可用，无法更新用户");
    }
    
    @Override
    public ApiResponse<Void> deleteUser(Long id) {
        log.warn("调用用户服务删除用户失败，用户ID: {}", id);
        return ApiResponse.error(503, "用户服务暂时不可用，无法删除用户");
    }
    
    @Override
    public ApiResponse<String> health() {
        log.warn("调用用户服务健康检查失败");
        return ApiResponse.error(503, "用户服务暂时不可用");
    }
} 