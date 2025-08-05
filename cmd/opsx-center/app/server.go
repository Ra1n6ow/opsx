package app

import (
	"github.com/ra1n6ow/opsx/cmd/opsx-center/app/options"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// 配置文件绝对路径
var configFile string

// NewOpsxCenterCommand 创建一个 *cobra.Command 对象，用于启动应用程序.
func NewOpsxCenterCommand() *cobra.Command {
	// 创建默认的应用命令行选项
	opts := options.NewServerOptions()

	cmd := &cobra.Command{
		// 指定命令的名字，该名字会出现在帮助信息中
		Use: "opsx-center",
		// 命令的简短描述
		Short: "opsx center",
		// 命令的详细描述
		Long: "opsx-center is a bff service for opsx",
		// 命令出错时，不打印帮助信息。设置为 true 可以确保命令出错时一眼就能看到错误信息
		SilenceUsage: true,
		// 指定调用 cmd.Execute() 时，执行的 Run 函数
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(opts)
		},
		// 设置命令运行时的参数检查，不需要指定命令行参数。例如：./opsx-center param1 param2
		Args: cobra.NoArgs,
	}

	// 注册一个回调函数，该回调函数在每次运行任意命令时都会被调用。
	// 借助该方法，来注册一个配置文件加载函数，在每次运行命令时，加载配置文件。
	cobra.OnInitialize(onInitialize)

	// cobra 支持持久性标志(PersistentFlag)，该标志可用于它所分配的命令以及该命令下的每个子命令
	// 推荐使用配置文件来配置应用，便于管理配置项
	cmd.PersistentFlags().StringVarP(&configFile, "config", "c", filePath(), "Path to the miniblog configuration file.")

	// 将 ServerOptions 中的选项绑定到命令标志
	opts.AddFlags(cmd.PersistentFlags())

	return cmd
}

// run 是主运行逻辑，负责初始化日志、解析配置、校验选项并启动服务器。
func run(opts *options.ServerOptions) error {
	// 将 viper 中的配置解析到 opts.
	if err := viper.Unmarshal(opts); err != nil {
		return err
	}

	// 校验命令行选项
	if err := opts.Validate(); err != nil {
		return err
	}

	// 获取应用配置.
	cfg, err := opts.Config()
	if err != nil {
		return err
	}

	// 创建服务器实例.
	// 注意这里是联合服务器，因为可能同时启动多个不同类型的服务器.
	server, err := cfg.NewUnionServer()
	if err != nil {
		return err
	}

	// 启动服务器
	return server.Run()
}
