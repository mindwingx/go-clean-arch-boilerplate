package driver

import (
	"fmt"
	"github.com/fatih/color"
	rpcEntity "github.com/mindwingx/go-clean-arch-boilerplate/app/rpc"
	"github.com/mindwingx/go-clean-arch-boilerplate/helper"
	"net"
	"net/rpc"
)

type RpcAbstraction interface {
	InitRpcService()
	StartRpc()
	Caller(destinationPort string, rpcMethod string, args interface{}, reply interface{}) error
}

type (
	rpcEngine struct {
		config   rpcConfig
		locale   LocaleAbstraction
		entities []interface{}
	}

	rpcConfig struct {
		Network string
		Port    string
	}
)

func NewRpc(registry RegistryAbstraction, locale LocaleAbstraction) RpcAbstraction {
	serviceRpc := new(rpcEngine)
	registry.Parse(&serviceRpc.config)
	serviceRpc.locale = locale
	return serviceRpc
}

func (r *rpcEngine) InitRpcService() {
	r.entities = append(r.entities,
		new(rpcEntity.UserRpc),
		//todo: append more rpc entities
	)
}

func (r *rpcEngine) StartRpc() {
	color.Cyan(r.locale.Get("rpc_init"))

	for _, entity := range r.entities {
		err := rpc.Register(entity)

		if err != nil {
			helper.CustomPanic(r.locale.Get("rpc_init_err"), err)
		}
	}

	listener, err := net.Listen(r.config.Network, fmt.Sprintf(":%s", r.config.Port))
	if err != nil {
		helper.CustomPanic(r.locale.Get("rpc_listen_err"), err)
	}

	defer listener.Close()

	for {
		rpcConn, acceptErr := listener.Accept()

		if acceptErr != nil {
			//todo: handle logger
			continue
		}

		go rpc.ServeConn(rpcConn)
	}
}

func (r *rpcEngine) Caller(destinationPort string, rpcMethod string, args interface{}, reply interface{}) (err error) {
	port := fmt.Sprintf(":%s", destinationPort)
	dial, err := rpc.Dial(r.config.Network, port)
	if err != nil {
		//todo: call logger
		return
	}

	defer dial.Close()

	err = dial.Call(rpcMethod, args, reply)
	if err != nil {
		//todo: handle logger
		return
	}

	return
}
