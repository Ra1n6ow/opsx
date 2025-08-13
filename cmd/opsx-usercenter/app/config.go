// Copyright 2025 JingFeng Du <jeffduuu@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Ra1n6ow/opsx.

package app

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	// defaultConfigDir 定义放置 opsx 项目配置的默认目录.
	defaultConfigDir = ".opsx"

	// defaultConfigName 指定 opsx-usercenter 服务的默认配置文件名.
	defaultConfigName = "opsx-usercenter.yaml"
)

// onInitialize 设置需要读取的配置文件名、环境变量，并将其内容读取到 viper 中.
func onInitialize() {
	if configFile != "" {
		// 从命令行参数加载配置文件
		viper.SetConfigFile(configFile)
	} else {
		// 使用默认配置文件路径和名称
		for _, dir := range searchDirs() {
			// 将 dir 目录加入到配置文件的搜索路径
			viper.AddConfigPath(dir)
		}

		// 设置配置文件格式为 YAML
		viper.SetConfigType("yaml")

		// 配置文件名称（没有文件扩展名）
		viper.SetConfigName(defaultConfigName)
	}

	// 读取环境变量并设置前缀
	setupEnvironmentVariables()

	// 读取配置文件.如果指定了配置文件名，则使用指定的配置文件，否则在注册的搜索路径中搜索
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Failed to read viper configuration file, err: %v", err)
	}

	// 打印当前使用的配置文件，方便调试
	log.Printf("Using config file: %s", viper.ConfigFileUsed())
}

// setupEnvironmentVariables 配置环境变量规则.
func setupEnvironmentVariables() {
	// 允许 viper 自动匹配环境变量
	viper.AutomaticEnv()
	// 设置环境变量前缀
	viper.SetEnvPrefix("OPSX")

	// 替换环境变量 key 中的分隔符 '.' 和 '-' 为 '_'
	// 配置项 mysql.port → 环境变量 OPSX_MYSQL_PORT
	// 配置项 db-name → 环境变量 OPSX_DB_NAME
	// 配置项 auth.jwt-key → 环境变量 OPSX_AUTH_JWT_KEY
	// 配置项 auth.expiration → 环境变量 OPSX_AUTH_EXPIRATION
	replacer := strings.NewReplacer(".", "_", "-", "_")
	viper.SetEnvKeyReplacer(replacer)
}

// searchDirs 返回默认的配置文件搜索目录.
func searchDirs() []string {
	// 获取用户主目录
	homeDir, err := os.UserHomeDir()
	// 如果获取用户主目录失败，则打印错误信息并退出程序
	cobra.CheckErr(err)
	return []string{filepath.Join(homeDir, defaultConfigDir), "."}
}

// filePath 获取默认配置文件的完整路径.
func filePath() string {
	home, err := os.UserHomeDir()
	// 如果不能获取用户主目录，则记录错误并返回空路径
	cobra.CheckErr(err)
	return filepath.Join(home, defaultConfigDir, defaultConfigName)
}
