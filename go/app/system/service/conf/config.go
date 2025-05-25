package conf

import (
	"flag"
	"fmt"
	config "github.com/go-kratos/kratos/contrib/config/nacos/v2"
	oc "github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"spring-go-rpc/app/gen/go/system/service/conf"
)

// 启动的配置文件
var bootstrapConf *conf.Bootstrap

// 全局动态配置
var globalConfig oc.Config

func ResolveLocalConf() {
	var flagconf string
	flag.StringVar(&flagconf, "conf", "/Users/heap/Developer/code/spring-go-rpc/go/app/system/service/configs", "config path, eg: -conf config.yaml")

	fc := oc.New(
		oc.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer fc.Close()

	if err := fc.Load(); err != nil {
		panic(err)
	}

	if err := fc.Scan(&bootstrapConf); err != nil {
		panic(err)
	}
}

func GetConfig() (*conf.Bootstrap, oc.Config) {
	if globalConfig == nil {
		ResolveRemoteConf()
	}
	return bootstrapConf, globalConfig
}

func ResolveRemoteConf() oc.Config {
	nacosConf := NewNacosConf()
	client, err := clients.NewConfigClient(nacosConf)
	if err != nil {
		panic(fmt.Errorf("nacos初始化错误: %s \n", err))
	}

	sc := bootstrapConf.Server.ServiceConfig
	if sc == nil {
		panic("未找到配置中心")
	}

	globalConfig = reloadRemoteConfig(client, sc)
	if globalConfig == nil {
		panic("解析远程配置失败")
	}

	//配置监听
	err = client.ListenConfig(vo.ConfigParam{
		DataId: sc.DataId,
		Group:  sc.Group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("config changed group:" + group + ", dataId:" + dataId + ", content:" + data)
			globalConfig = reloadRemoteConfig(client, sc)
			if globalConfig == nil {
				panic("解析远程配置失败")
			}
		},
	})
	return globalConfig
}

func reloadRemoteConfig(client config_client.IConfigClient, sc *conf.Server_ServiceConfig) oc.Config {
	source := config.NewConfigSource(client, config.WithGroup(sc.Group), config.WithDataID(sc.DataId))
	fc := oc.New(oc.WithSource(source))

	if err := fc.Load(); err != nil {
		panic(fmt.Errorf("nacos读取配置错误: %s \n", err))
	}
	return fc
}

// NewNacosConf 初始化nacos客户端服务
func NewNacosConf() vo.NacosClientParam {
	if bootstrapConf == nil {
		ResolveLocalConf()
	}
	discovery := bootstrapConf.Server.ServiceDiscovery

	sc := []constant.ServerConfig{
		*constant.NewServerConfig(discovery.ServerAddr, discovery.Port),
	}
	cc := &constant.ClientConfig{
		NamespaceId:         discovery.Namespace, //namespace id
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	return vo.NacosClientParam{
		ClientConfig:  cc,
		ServerConfigs: sc,
	}
}
