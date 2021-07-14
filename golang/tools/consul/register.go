package consul_register

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

//consul服务注册相关代码

type register struct {
	consulHost string
	consulPort int
}

type RegisterClient interface {
	Register(address string, port int, name string, tags []string, id string) error
	DeRegister(serverId string) error
}

func (r *register) NewRegisterClient(host string, port int) *register {
	return &register{consulHost: host, consulPort: port}
}

func (r *register) Register(address string, port int, name string, tags []string, id string) error {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", r.consulHost, r.consulPort)
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	//生成对应的检查对象
	check := &api.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%d/health", address, port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}
	//生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	registration.Port = port
	registration.Tags = tags
	registration.Address = address
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
	return nil
}

func (r *register) DeRegister(serverId string) error {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", r.consulHost, r.consulPort)
	client, err := api.NewClient(cfg)
	if err != nil {
		return err
	}
	err = client.Agent().ServiceDeregister(serverId)
	return err
}
