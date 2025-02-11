package config

import (
	"github.com/go-ll/mrpc"
	"mall-go/common/conf"
)

type Config struct {
	conf.ApiConf `mapstructure:",squash"`
	conf.JwtAuth
	UserRpcConf    mrpc.RpcClientConf
	BalanceRpcConf mrpc.RpcClientConf
}
