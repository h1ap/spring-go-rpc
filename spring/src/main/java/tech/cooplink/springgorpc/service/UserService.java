package tech.cooplink.springgorpc.service;

import com.alibaba.nacos.common.utils.JacksonUtils;
import jakarta.annotation.PostConstruct;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;
import tech.cooplink.springgorpc.dto.UserDTO;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.atomic.AtomicLong;

/**
 * 用户服务类
 */
@Service
@Slf4j
public class UserService {

    // 模拟用户数据存储
    private final List<UserDTO> users = new ArrayList<>();
    private final AtomicLong idGenerator = new AtomicLong(1);

    @Value("${spring.go.config.spring.user}")
    private String userConfig;

    @PostConstruct
    public void init() {
        // 初始化一些测试数据
        log.info("用户服务启动，读取配置: {}", userConfig);
        initTestData();
    }

    /**
     * 初始化测试数据
     */
    private void initTestData() {
        users.add(new UserDTO(1L, "张三", "zhangsan@example.com", 25));
        users.add(new UserDTO(2L, "李四", "lisi@example.com", 30));
        users.add(new UserDTO(3L, "王五", "wangwu@example.com", 28));
        idGenerator.set(4);
        log.info("初始化用户测试数据完成，共{}条记录", users.size());
    }

    /**
     * 获取所有用户
     *
     * @return 用户列表
     */
    public List<UserDTO> getAllUsers() {
        log.info("获取所有用户，当前用户数量: {}", users.size());
        return new ArrayList<>(users);
    }

    /**
     * 根据ID获取用户
     *
     * @param id 用户ID
     * @return 用户信息
     */
    public UserDTO getUserById(Long id) {
        log.info("根据ID获取用户: {}", id);
        return users.stream()
                .filter(user -> user.getId().equals(id))
                .findFirst()
                .orElse(null);
    }

    /**
     * 创建用户
     *
     * @param userDTO 用户信息
     * @return 创建的用户
     */
    public UserDTO createUser(UserDTO userDTO) {
        userDTO.setId(idGenerator.getAndIncrement());
        users.add(userDTO);
        log.info("创建用户成功: {}", userDTO);
        return userDTO;
    }

    /**
     * 更新用户
     *
     * @param id      用户ID
     * @param userDTO 用户信息
     * @return 更新后的用户
     */
    public UserDTO updateUser(Long id, UserDTO userDTO) {
        for (int i = 0; i < users.size(); i++) {
            UserDTO existingUser = users.get(i);
            if (existingUser.getId().equals(id)) {
                userDTO.setId(id);
                users.set(i, userDTO);
                log.info("更新用户成功: {}", userDTO);
                return userDTO;
            }
        }
        log.warn("更新用户失败，用户不存在: {}", id);
        return null;
    }

    /**
     * 删除用户
     *
     * @param id 用户ID
     * @return 是否删除成功
     */
    public boolean deleteUser(Long id) {
        boolean removed = users.removeIf(user -> user.getId().equals(id));
        if (removed) {
            log.info("删除用户成功: {}", id);
        } else {
            log.warn("删除用户失败，用户不存在: {}", id);
        }
        return removed;
    }
}