package tech.cooplink.springgorpc.controller;

import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.*;
import tech.cooplink.springgorpc.dto.ApiResponse;
import tech.cooplink.springgorpc.dto.UserDTO;
import tech.cooplink.springgorpc.service.UserService;

import java.util.List;

/**
 * 用户控制器
 * 提供用户相关的RESTful API接口
 */
@RestController
@RequestMapping("/api/users")
@RequiredArgsConstructor
@Slf4j
public class UserController {
    
    private final UserService userService;
    
    /**
     * 获取所有用户
     * @return 用户列表
     */
    @GetMapping
    public ApiResponse<List<UserDTO>> getAllUsers() {
        log.info("接收到获取所有用户的请求");
        List<UserDTO> users = userService.getAllUsers();
        return ApiResponse.success("获取用户列表成功", users);
    }
    
    /**
     * 根据ID获取用户
     * @param id 用户ID
     * @return 用户信息
     */
    @GetMapping("/{id}")
    public ApiResponse<UserDTO> getUserById(@PathVariable Long id) {
        log.info("接收到获取用户的请求，用户ID: {}", id);
        UserDTO user = userService.getUserById(id);
        if (user != null) {
            return ApiResponse.success("获取用户信息成功", user);
        } else {
            return ApiResponse.error(404, "用户不存在");
        }
    }
    
    /**
     * 创建用户
     * @param userDTO 用户信息
     * @return 创建的用户
     */
    @PostMapping
    public ApiResponse<UserDTO> createUser(@RequestBody UserDTO userDTO) {
        log.info("接收到创建用户的请求: {}", userDTO);
        try {
            UserDTO createdUser = userService.createUser(userDTO);
            return ApiResponse.success("创建用户成功", createdUser);
        } catch (Exception e) {
            log.error("创建用户失败", e);
            return ApiResponse.error("创建用户失败: " + e.getMessage());
        }
    }
    
    /**
     * 更新用户
     * @param id 用户ID
     * @param userDTO 用户信息
     * @return 更新后的用户
     */
    @PutMapping("/{id}")
    public ApiResponse<UserDTO> updateUser(@PathVariable Long id, @RequestBody UserDTO userDTO) {
        log.info("接收到更新用户的请求，用户ID: {}, 用户信息: {}", id, userDTO);
        try {
            UserDTO updatedUser = userService.updateUser(id, userDTO);
            if (updatedUser != null) {
                return ApiResponse.success("更新用户成功", updatedUser);
            } else {
                return ApiResponse.error(404, "用户不存在");
            }
        } catch (Exception e) {
            log.error("更新用户失败", e);
            return ApiResponse.error("更新用户失败: " + e.getMessage());
        }
    }
    
    /**
     * 删除用户
     * @param id 用户ID
     * @return 删除结果
     */
    @DeleteMapping("/{id}")
    public ApiResponse<Void> deleteUser(@PathVariable Long id) {
        log.info("接收到删除用户的请求，用户ID: {}", id);
        try {
            boolean deleted = userService.deleteUser(id);
            if (deleted) {
                return ApiResponse.success("删除用户成功", null);
            } else {
                return ApiResponse.error(404, "用户不存在");
            }
        } catch (Exception e) {
            log.error("删除用户失败", e);
            return ApiResponse.error("删除用户失败: " + e.getMessage());
        }
    }
    
    /**
     * 健康检查接口
     * @return 服务状态
     */
    @GetMapping("/health")
    public ApiResponse<String> health() {
        return ApiResponse.success("用户服务运行正常", "healthy");
    }
} 