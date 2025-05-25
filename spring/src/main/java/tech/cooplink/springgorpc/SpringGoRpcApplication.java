package tech.cooplink.springgorpc;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.client.discovery.EnableDiscoveryClient;
import org.springframework.cloud.openfeign.EnableFeignClients;

@SpringBootApplication
@EnableDiscoveryClient  // 启用服务发现
@EnableFeignClients     // 启用OpenFeign客户端
public class SpringGoRpcApplication {

    public static void main(String[] args) {
        SpringApplication.run(SpringGoRpcApplication.class, args);
    }

}
