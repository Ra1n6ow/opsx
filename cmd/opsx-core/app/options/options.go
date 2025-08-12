// Copyright 2025 JingFeng Du <jeffduuu@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Ra1n6ow/opsx.

package options

import (
	"errors"
	"fmt"
	"time"

	genericoptions "github.com/ra1n6ow/opsxstack/pkg/options"
	stringsutil "github.com/ra1n6ow/opsxstack/pkg/util/strings"
	"github.com/spf13/pflag"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/ra1n6ow/opsx/internal/core"
)

// 定义支持的服务器模式集合.
var availableServerModes = sets.New(
	core.GRPCServerMode,
	core.GRPCGatewayServerMode,
	core.GinServerMode,
)

// ServerOptions 包含服务器配置选项.
// mapstructure 标签用于将配置文件中的配置项与 Go 结构体字段进行映射，
// 在调用 viper.Unmarshal 函数时，viper 会将配置文件中配置项的值赋值给对应的结构体字段
type ServerOptions struct {
	// ServerMode 定义服务器模式：gRPC、Gin HTTP、HTTP Reverse Proxy.
	ServerMode string `json:"server-mode" mapstructure:"server-mode"`
	// JWTKey 定义 JWT 密钥.
	JWTKey string `json:"jwt-key" mapstructure:"jwt-key"`
	// Expiration 定义 JWT Token 的过期时间.
	Expiration time.Duration `json:"expiration" mapstructure:"expiration"`
	// GRPC 配置
	GRPCOptions *genericoptions.GRPCOptions `json:"grpc" mapstructure:"grpc"`
	// HTTP 配置
	HTTPOptions *genericoptions.HTTPOptions `json:"http" mapstructure:"http"`
}

// NewServerOptions 创建带有默认值的 ServerOptions 实例.
func NewServerOptions() *ServerOptions {
	opts := &ServerOptions{
		ServerMode:  "grpc-gateway",
		JWTKey:      "Rtg8BPKNEf2mB4mgvKONGPZZQSaJWNLijxR42qRgq0iB31",
		Expiration:  2 * time.Hour,
		GRPCOptions: genericoptions.NewGRPCOptions(),
		HTTPOptions: genericoptions.NewHTTPOptions(),
	}
	opts.GRPCOptions.Addr = ":7701"
	opts.HTTPOptions.Addr = ":7700"
	return opts
}

// AddFlags 将 ServerOptions 的选项绑定到命令行标志.
// 通过使用 pflag 包，可以实现从命令行中解析这些选项的功能.
func (o *ServerOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.ServerMode, "server-mode", o.ServerMode, fmt.Sprintf("Server mode, available options: %v", availableServerModes.UnsortedList()))
	fs.StringVar(&o.JWTKey, "jwt-key", o.JWTKey, "JWT signing key. Must be at least 6 characters long.")
	// 绑定 JWT Token 的过期时间选项到命令行标志。
	// 参数名称为 `--expiration`，默认值为 o.Expiration
	fs.DurationVar(&o.Expiration, "expiration", o.Expiration, "The expiration duration of JWT tokens.")
	o.GRPCOptions.AddFlags(fs)
	o.HTTPOptions.AddFlags(fs)
}

// Validate 校验 ServerOptions 中的选项是否合法.
func (o *ServerOptions) Validate() error {
	errs := []error{}

	// 校验 ServerMode 是否有效
	if !availableServerModes.Has(o.ServerMode) {
		errs = append(errs, fmt.Errorf("invalid server mode: must be one of %v", availableServerModes.UnsortedList()))
	}

	// 校验 JWTKey 长度
	if len(o.JWTKey) < 6 {
		errs = append(errs, errors.New("JWTKey must be at least 6 characters long"))
	}

	// 如果是 gRPC 或 gRPC-Gateway 模式，校验 gRPC 配置
	if stringsutil.StringIn(o.ServerMode, []string{core.GRPCServerMode, core.GRPCGatewayServerMode}) {
		errs = append(errs, o.GRPCOptions.Validate()...)
	}

	// 如果是 gRPC-Gateway 模式或 Gin 模式，校验 HTTP 配置
	if stringsutil.StringIn(o.ServerMode, []string{core.GRPCGatewayServerMode, core.GinServerMode}) {
		errs = append(errs, o.HTTPOptions.Validate()...)
	}

	// 合并所有错误并返回
	return utilerrors.NewAggregate(errs)
}

// Config 将初始化配置 ServerOptions 转换为运行时配置 core.Config.
func (o *ServerOptions) Config() (*core.Config, error) {
	return &core.Config{
		ServerMode:  o.ServerMode,
		JWTKey:      o.JWTKey,
		Expiration:  o.Expiration,
		GRPCOptions: o.GRPCOptions,
		HTTPOptions: o.HTTPOptions,
	}, nil
}
