// Package config provides the configuration
package config

import (
	"strings"

	"github.com/spf13/viper"
)

var (
	// LogDirPath string 日志目录路径
	//	@update 2025-06-09 19:24:01
	LogDirPath string

	// LogLevel string 日志级别
	//	@update 2025-06-09 19:24:02
	LogLevel string

	// APIKey string 认证key
	//	@update 2025-06-13 16:47:01
	APIKey string
)

func init() {
	initEnvironment()
}

func initEnvironment() {
	config := viper.New()
	config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	config.SetDefault("log.dir.path", "./logs")
	config.SetDefault("log.level", "debug")

	config.AutomaticEnv()

	LogDirPath = config.GetString("log.dir.path")
	LogLevel = config.GetString("log.level")

	APIKey = config.GetString("mcp.api.key")
}
