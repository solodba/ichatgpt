package cmd

import (
	"fmt"

	_ "github.com/solodba/ichatgpt/apps/all"
	"github.com/solodba/ichatgpt/cmd/initial"
	"github.com/solodba/ichatgpt/cmd/start"
	"github.com/solodba/ichatgpt/conf"
	"github.com/solodba/mcube/apps"
	"github.com/solodba/mcube/logger"
	"github.com/solodba/mcube/version"
	"github.com/spf13/cobra"
)

// 全局参数
var (
	showVersion bool
	configType  string
	filePath    string
)

// 根命令
var RootCmd = &cobra.Command{
	Use:     "ichatgpt [init|start]",
	Short:   "ichatgpt service",
	Long:    "ichatgpt service",
	Example: "ichatgpt-api -v",
	RunE: func(cmd *cobra.Command, args []string) error {
		if showVersion {
			logger.L().Info().Msgf(version.ShortVersion())
			return nil
		}
		return cmd.Help()
	},
}

// 加载全局配置
func LoadGlobalConfig(configType string) error {
	switch configType {
	case "file":
		if err := conf.LoadConfigFromToml(filePath); err != nil {
			return err
		}
	case "env":
		if err := conf.LoadConfigFromEnv(); err != nil {
			return err
		}
	case "etcd":
		return fmt.Errorf("load global from etcd is not implement")
	default:
		return fmt.Errorf("this config type is not support")
	}
	return nil
}

// 初始化函数
func Initial() {
	err := LoadGlobalConfig(configType)
	cobra.CheckErr(err)
	err = apps.InitInternalApps()
	cobra.CheckErr(err)
}

// 执行函数
func Execute() {
	cobra.OnInitialize(Initial)
	RootCmd.AddCommand(initial.Cmd, start.Cmd)
	err := RootCmd.Execute()
	cobra.CheckErr(err)
}

// 初始化函数
func init() {
	RootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "show project ichatgpt version")
	RootCmd.PersistentFlags().StringVarP(&configType, "config-type", "t", "file", "project ichatgpt config type")
	RootCmd.PersistentFlags().StringVarP(&filePath, "file-path", "f", "etc/config.toml", "project ichatgpt config file path")
}
