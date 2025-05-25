# Spring Cloud 微服务项目

这是一个基于Spring Cloud Alibaba的微服务示例项目，演示了如何使用Nacos进行服务注册与发现，以及如何使用OpenFeign进行服务间调用。

## 功能特性

- ✅ Spring Boot 3.4.6
- ✅ Spring Cloud 2023.0.1
- ✅ Spring Cloud Alibaba 2023.0.1.3
- ✅ Nacos服务注册与发现
- ✅ OpenFeign远程服务调用
- ✅ 完整的RESTful API
- ✅ 统一响应格式
- ✅ 服务降级处理
- ✅ 详细的中文注释

## 项目结构

```
src/main/java/tech/cooplink/springgorpc/
├── SpringGoRpcApplication.java          # 主启动类
├── controller/
│   ├── UserController.java              # 用户REST控制器
│   └── RemoteController.java            # 远程调用演示控制器
├── service/
│   └── UserService.java                 # 用户业务服务
├── feign/
│   ├── UserFeignClient.java             # Feign客户端接口
│   └── UserFeignClientFallback.java     # Feign降级处理
├── dto/
│   ├── UserDTO.java                     # 用户数据传输对象
│   └── ApiResponse.java                 # 统一响应格式
└── config/
    └── FeignConfig.java                  # Feign配置类
```

## 快速开始

### 1. 启动Nacos服务器

首先需要启动Nacos服务器。如果没有安装Nacos，可以通过以下方式之一启动：

**方式一：使用Docker启动Nacos**
```bash
docker run --name nacos-standalone -e MODE=standalone -e JVM_XMS=512m -e JVM_XMX=512m -e JVM_XMN=256m -p 8848:8848 -d nacos/nacos-server:latest
```

**方式二：下载并启动Nacos**
1. 从 [Nacos官网](https://nacos.io/zh-cn/) 下载最新版本
2. 解压后执行启动命令：
   ```bash
   # Linux/Mac
   sh startup.sh -m standalone
   
   # Windows
   startup.cmd -m standalone
   ```

### 2. 启动应用

```bash
# 使用Maven启动
./mvnw spring-boot:run

# 或者先编译再运行
./mvnw clean package
java -jar target/spring-go-rpc-0.0.1-SNAPSHOT.jar
```

应用启动后会自动注册到Nacos服务注册中心。

### 3. 验证服务注册

访问Nacos控制台：http://localhost:8848/nacos

- 用户名：nacos
- 密码：nacos

在"服务管理 > 服务列表"中应该能看到名为`user-service`的服务。

## API接口

### 用户管理接口

| 方法 | 路径 | 描述 |
|------|------|------|
| GET | `/api/users` | 获取所有用户 |
| GET | `/api/users/{id}` | 根据ID获取用户 |
| POST | `/api/users` | 创建用户 |
| PUT | `/api/users/{id}` | 更新用户 |
| DELETE | `/api/users/{id}` | 删除用户 |
| GET | `/api/users/health` | 健康检查 |

### 远程调用演示接口

| 方法 | 路径 | 描述 |
|------|------|------|
| GET | `/api/remote/users` | 通过Feign获取所有用户 |
| GET | `/api/remote/users/{id}` | 通过Feign根据ID获取用户 |
| POST | `/api/remote/users` | 通过Feign创建用户 |
| PUT | `/api/remote/users/{id}` | 通过Feign更新用户 |
| DELETE | `/api/remote/users/{id}` | 通过Feign删除用户 |
| GET | `/api/remote/users/health` | 通过Feign检查服务健康状态 |
| GET | `/api/remote/test` | 综合测试远程调用 |

## 测试示例

### 1. 获取所有用户
```bash
curl -X GET http://localhost:8080/api/users
```

### 2. 创建用户
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "username": "测试用户",
    "email": "test@example.com",
    "age": 25
  }'
```

### 3. 通过Feign调用获取用户
```bash
curl -X GET http://localhost:8080/api/remote/users
```

### 4. 测试远程调用
```bash
curl -X GET http://localhost:8080/api/remote/test
```

## 响应格式

所有API接口都返回统一的响应格式：

```json
{
  "code": 200,
  "message": "操作成功",
  "data": {
    // 具体数据
  }
}
```

- `code`: 响应状态码（200表示成功，其他表示错误）
- `message`: 响应消息
- `data`: 响应数据

## 配置说明

### Nacos配置

在`application.properties`中配置Nacos连接信息：

```properties
# Nacos服务注册与发现配置
spring.cloud.nacos.discovery.server-addr=127.0.0.1:8848
spring.cloud.nacos.discovery.namespace=
spring.cloud.nacos.discovery.group=DEFAULT_GROUP
```

### Feign配置

```properties
# Feign配置
feign.hystrix.enabled=false
feign.client.config.default.connect-timeout=5000
feign.client.config.default.read-timeout=5000
```

## 故障排除

### 1. 无法连接到Nacos
- 确保Nacos服务器已启动
- 检查端口8848是否被占用
- 验证防火墙设置

### 2. Feign调用失败
- 检查服务是否已注册到Nacos
- 确认服务名称配置正确
- 查看应用日志获取详细错误信息

### 3. 服务注册失败
- 检查网络连接
- 验证Nacos服务器地址配置
- 确认应用名称配置正确

## 技术栈

- **Spring Boot**: 3.4.6
- **Spring Cloud**: 2023.0.1
- **Spring Cloud Alibaba**: 2023.0.1.3
- **Nacos**: 服务注册与发现
- **OpenFeign**: 声明式HTTP客户端
- **Lombok**: 简化Java代码
- **SLF4J**: 日志框架

## 许可证

MIT License 