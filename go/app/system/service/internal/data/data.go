package data

import (
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"spring-go-rpc/app/system/service/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewUserRepo,
	conf.NewNacosConf, NewDiscovery, NewRegistrar)

// Data .
type Data struct {
}

// NewData .
func NewData(logger log.Logger) (*Data, func(), error) {

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}

// NewDiscovery nacos服务发现注入
func NewDiscovery(param vo.NacosClientParam) registry.Discovery {
	client, err := clients.NewNamingClient(param)
	if err != nil {
		panic(err)
	}
	return nacos.New(client, nacos.WithDefaultKind("http"))
}

// NewRegistrar 服务注册业务注入
func NewRegistrar(param vo.NacosClientParam) registry.Registrar {
	client, err := clients.NewNamingClient(param)
	if err != nil {
		panic(err)
	}
	return nacos.New(client, nacos.WithDefaultKind("http"))
}
