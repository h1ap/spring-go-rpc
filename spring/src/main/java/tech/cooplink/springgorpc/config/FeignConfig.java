package tech.cooplink.springgorpc.config;

import feign.Logger;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

/**
 * Feign配置类
 * 用于自定义Feign客户端的行为
 */
@Configuration
public class FeignConfig {
    
    /**
     * 配置Feign日志级别
     * @return 日志级别
     */
    @Bean
    Logger.Level feignLoggerLevel() {
        // NONE：不记录任何日志（默认）
        // BASIC：只记录请求方法、URL以及响应状态码和执行时间
        // HEADERS：记录BASIC级别的基础上，还记录请求和响应的头信息
        // FULL：记录所有请求和响应的明细，包括头信息、请求体、元数据
        return Logger.Level.BASIC;
    }
} 