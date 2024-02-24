package impl

import (
	openapi "github.com/sashabaranov/go-openai"
	"github.com/solodba/ichatgpt/apps/audio"
	"github.com/solodba/ichatgpt/conf"
	"github.com/solodba/mcube/apps"
)

var (
	svc = &impl{}
)

// 业务实现类
type impl struct {
	client *openapi.Client
}

// 实现Ioc中心Name方法
func (i *impl) Name() string {
	return audio.AppName
}

// 实现Ioc中心Conf方法
func (i *impl) Conf() error {
	i.client = conf.C().ChatGPT.GetClient()
	return nil
}

// 注册到Ioc中心
func init() {
	apps.RegistryInternalApp(svc)
}
