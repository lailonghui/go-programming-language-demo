/*
@Time : 2021/3/8 9:24
@Author : lai
@Description :
@File : register
*/
package register

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"os"
	"strconv"
)

func Register(consulHost, consulPort, svcHost, svcPort string, logger log.Logger) (registrar sd.Registrar) {
	// 创建Consul 客户端连接
	var client consul.Client
	{
		consulCfg := api.DefaultConfig()
		consulCfg.Address = consulHost + ":" + consulPort
		consulClient, err := api.NewClient(consulCfg)
		if err != nil {
			logger.Log("create consul client error:", err)
			os.Exit(1)
		}
		client = consul.NewClient(consulClient)
	}

	// 设置Consul对服务健康检查的参数
	check := api.AgentServiceCheck{
		HTTP:     "http://" + svcHost + ":" + svcPort + "/health",
		Interval: "10s",
		Timeout:  "1s",
		Notes:    "Consul check service health status",
	}
	port, _ := strconv.Atoi(svcPort)

	//设置微服务项Consul的注册信息
	reg := api.AgentServiceRegistration{
		ID:      "arithmetic" + uuid.New().String(),
		Name:    "arithmetic",
		Address: svcHost,
		Port:    port,
		Tags:    []string{"arithmetic", "raysonxin"},
		Check:   &check,
	}

	//执行注册
	registrar = consul.NewRegistrar(client, &reg, logger)

	return
}
