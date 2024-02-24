package start

import (
	"github.com/solodba/ichatgpt/protocol"
	"github.com/spf13/cobra"
)

// 项目启动子命令
var Cmd = &cobra.Command{
	Use:     "start",
	Short:   "mkube start service",
	Long:    "mkube start service",
	Example: "mkube-api start -f etc/config.toml",
	RunE: func(cmd *cobra.Command, args []string) error {
		srv := NewServer()
		if err := srv.Start(); err != nil {
			return err
		}
		return nil
	},
}

// 服务结构体
type Server struct {
	HttpService *protocol.HttpService
}

// 服务结构体初始化函数
func NewServer() *Server {
	return &Server{
		HttpService: protocol.NewHttpService(),
	}
}

// Server服务启动方法
func (s *Server) Start() error {
	return s.HttpService.Start()
}

// Server服务停止方法
func (s *Server) Stop() error {
	err := s.HttpService.Stop()
	if err != nil {
		return err
	}
	return nil
}
