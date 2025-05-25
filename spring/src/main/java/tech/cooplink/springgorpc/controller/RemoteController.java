package tech.cooplink.springgorpc.controller;

import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.*;
import tech.cooplink.springgorpc.dto.ApiResponse;
import tech.cooplink.springgorpc.dto.UserDTO;
import tech.cooplink.springgorpc.feign.UserFeignClient;

import java.util.List;

/**
 * 远程调用控制器
 * 演示如何使用OpenFeign调用其他服务
 */
@RestController
@RequestMapping("/api/remote")
@RequiredArgsConstructor
@Slf4j
public class RemoteController {
    
    private final UserFeignClient userFeignClient;
    
    /**
     * 通过Feign客户端获取所有用户
     * @return 用户列表
     */
    @GetMapping("/users")
    public ApiResponse<List<UserDTO>> getAllUsersViaFeign() {
        log.info("通过Feign客户端获取所有用户");
        try {
            return userFeignClient.getAllUsers();
        } catch (Exception e) {
            log.error("通过Feign调用用户服务失败", e);
            return ApiResponse.error("远程调用失败: " + e.getMessage());
        }
    }
    
    /**
     * 通过Feign客户端根据ID获取用户
     * @param id 用户ID
     * @return 用户信息
     */
    @GetMapping("/users/{id}")
    public ApiResponse<UserDTO> getUserByIdViaFeign(@PathVariable Long id) {
        log.info("通过Feign客户端获取用户，用户ID: {}", id);
        try {
            return userFeignClient.getUserById(id);
        } catch (Exception e) {
            log.error("通过Feign调用用户服务失败，用户ID: {}", id, e);
            return ApiResponse.error("远程调用失败: " + e.getMessage());
        }
    }
    
    /**
     * 通过Feign客户端创建用户
     * @param userDTO 用户信息
     * @return 创建的用户
     */
    @PostMapping("/users")
    public ApiResponse<UserDTO> createUserViaFeign(@RequestBody UserDTO userDTO) {
        log.info("通过Feign客户端创建用户: {}", userDTO);
        try {
            return userFeignClient.createUser(userDTO);
        } catch (Exception e) {
            log.error("通过Feign创建用户失败", e);
            return ApiResponse.error("远程调用失败: " + e.getMessage());
        }
    }
    
    /**
     * 通过Feign客户端更新用户
     * @param id 用户ID
     * @param userDTO 用户信息
     * @return 更新后的用户
     */
    @PutMapping("/users/{id}")
    public ApiResponse<UserDTO> updateUserViaFeign(@PathVariable Long id, @RequestBody UserDTO userDTO) {
        log.info("通过Feign客户端更新用户，用户ID: {}, 用户信息: {}", id, userDTO);
        try {
            return userFeignClient.updateUser(id, userDTO);
        } catch (Exception e) {
            log.error("通过Feign更新用户失败，用户ID: {}", id, e);
            return ApiResponse.error("远程调用失败: " + e.getMessage());
        }
    }
    
    /**
     * 通过Feign客户端删除用户
     * @param id 用户ID
     * @return 删除结果
     */
    @DeleteMapping("/users/{id}")
    public ApiResponse<Void> deleteUserViaFeign(@PathVariable Long id) {
        log.info("通过Feign客户端删除用户，用户ID: {}", id);
        try {
            return userFeignClient.deleteUser(id);
        } catch (Exception e) {
            log.error("通过Feign删除用户失败，用户ID: {}", id, e);
            return ApiResponse.error("远程调用失败: " + e.getMessage());
        }
    }
    
    /**
     * 通过Feign客户端检查用户服务健康状态
     * @return 服务状态
     */
    @GetMapping("/users/health")
    public ApiResponse<String> checkUserServiceHealth() {
        log.info("通过Feign客户端检查用户服务健康状态");
        try {
            return userFeignClient.health();
        } catch (Exception e) {
            log.error("通过Feign检查用户服务健康状态失败", e);
            return ApiResponse.error("远程调用失败: " + e.getMessage());
        }
    }
    
    /**
     * 测试远程调用的综合接口
     * @return 测试结果
     */
    @GetMapping("/test")
    public ApiResponse<String> testRemoteCall() {
        log.info("开始测试远程调用");
        
        try {
            // 检查服务健康状态
            ApiResponse<String> healthResponse = userFeignClient.health();
            if (healthResponse.getCode() == 200) {
                log.info("用户服务健康检查通过");
                
                // 获取用户列表
                ApiResponse<List<UserDTO>> usersResponse = userFeignClient.getAllUsers();
                if (usersResponse.getCode() == 200) {
                    int userCount = usersResponse.getData().size();
                    log.info("成功获取用户列表，用户数量: {}", userCount);
                    return ApiResponse.success("远程调用测试成功", 
                        String.format("用户服务正常运行，当前用户数量: %d", userCount));
                } else {
                    return ApiResponse.error("获取用户列表失败: " + usersResponse.getMessage());
                }
            } else {
                return ApiResponse.error("用户服务健康检查失败: " + healthResponse.getMessage());
            }
        } catch (Exception e) {
            log.error("远程调用测试失败", e);
            return ApiResponse.error("远程调用测试失败: " + e.getMessage());
        }
    }
} 