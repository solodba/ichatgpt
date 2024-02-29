package impl

import (
	"os"

	openapi "github.com/sashabaranov/go-openai"
	"github.com/solodba/ichatgpt/apps/image"
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
	return image.AppName
}

// 实现Ioc中心Conf方法
func (i *impl) Conf() error {
	err := os.Setenv("http_proxy", conf.C().ChatGPT.HttpProxy)
	if err != nil {
		return err
	}
	err = os.Setenv("https_proxy", conf.C().ChatGPT.HttpsProxy)
	if err != nil {
		return err
	}
	i.client = conf.C().ChatGPT.GetClient()
	return nil
}

// 注册到Ioc中心
func init() {
	apps.RegistryInternalApp(svc)
}
