package conf

import (
	"fmt"

	"github.com/solodba/mcube/logger"
)

// 全局配置Config参数
var (
	c *Config
)

func C() *Config {
	if c == nil {
		logger.L().Panic().Msgf("please initial global config")
	}
	return c
}

// Http结构体
type Http struct {
	Host string `toml:"host" env:"HTTP_HOST"`
	Port int    `toml:"port" env:"HTTP_PORT"`
}

// App结构体
type App struct {
	Name string `toml:"name" env:"APP_NAME"`
	Http *Http  `toml:"http"`
}

// ChatGPT结构体
type ChatGPT struct {
	OpenaiApiKey string `toml:"openai_api_key" env:"CHATGPT_OPENAIAPIKEY"`
}

// 全局配置Config结构体
type Config struct {
	App     *App     `toml:"app"`
	ChatGPT *ChatGPT `toml:"chatgpt"`
}

// Http初始化函数
func NewDefaultHttp() *Http {
	return &Http{
		Host: "127.0.0.1",
		Port: 8080,
	}
}

// App初始化函数
func NewDefaultApp() *App {
	return &App{
		Name: "ichatgpt",
		Http: NewDefaultHttp(),
	}
}

// ChatGPT初始化函数
func NewDefaultChatGPT() *ChatGPT {
	return &ChatGPT{}
}

// Config初始化函数
func NewDefaultConfig() *Config {
	return &Config{
		App:     NewDefaultApp(),
		ChatGPT: NewDefaultChatGPT(),
	}
}

func (h *Http) Addr() string {
	return fmt.Sprintf("%s:%d", h.Host, h.Port)
}
