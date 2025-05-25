package tech.cooplink.springgorpc.feign;

import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.web.bind.annotation.*;
import tech.cooplink.springgorpc.dto.ApiResponse;
import tech.cooplink.springgorpc.dto.UserDTO;

import java.util.List;

/**
 * 用户服务Feign客户端
 * 用于调用用户服务的接口
 */
@FeignClient(name = "spring-go-rpc-go.http", fallback = UserFeignClientFallback.class)
public interface UserFeignClient {
    
    /**
     * 获取所有用户
     * @return 用户列表
     */
    @GetMapping("/api/users")
    ApiResponse<List<UserDTO>> getAllUsers();
    
    /**
     * 根据ID获取用户
     * @param id 用户ID
     * @return 用户信息
     */
    @GetMapping("/api/users/{id}")
    ApiResponse<UserDTO> getUserById(@PathVariable("id") Long id);
    
    /**
     * 创建用户
     * @param userDTO 用户信息
     * @return 创建的用户
     */
    @PostMapping("/api/users")
    ApiResponse<UserDTO> createUser(@RequestBody UserDTO userDTO);
    
    /**
     * 更新用户
     * @param id 用户ID
     * @param userDTO 用户信息
     * @return 更新后的用户
     */
    @PutMapping("/api/users/{id}")
    ApiResponse<UserDTO> updateUser(@PathVariable("id") Long id, @RequestBody UserDTO userDTO);
    
    /**
     * 删除用户
     * @param id 用户ID
     * @return 删除结果
     */
    @DeleteMapping("/api/users/{id}")
    ApiResponse<Void> deleteUser(@PathVariable("id") Long id);
    
    /**
     * 健康检查
     * @return 服务状态
     */
    @GetMapping("/api/users/health")
    ApiResponse<String> health();
} 