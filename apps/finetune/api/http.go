package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/ichatgpt/apps/finetune"
	"github.com/solodba/mcube/apps"
)

// handler结构体
type handler struct {
	svc finetune.Service
}

// 实现注册restful实例Name方法
func (h *handler) Name() string {
	return finetune.AppName
}

// 实现注册restful实例Conf方法
func (h *handler) Conf() error {
	// 从IOC获取token实例
	h.svc = apps.GetInternalApp(finetune.AppName).(finetune.Service)
	return nil
}

// 实现注册restful实例version方法
func (h *handler) Version() string {
	return "v1"
}

// 实现注册restful实例Registry方法
func (h *handler) RegistryHandler(ws *restful.WebService) {
	tags := []string{"FineTuning管理"}
	// webservice定义路由信息
	// 创建FineTuneJob
	ws.Route(ws.POST("/").To(h.CreateFineTuneJob).
		Doc("创建FineTuneJob").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(finetune.CreateFineTuneJobRequest{}).
		Writes(finetune.FineTuneJobItem{}).
		Returns(200, "OK", finetune.FineTuneJobItem{}))
}

// 初始化函数注册restful实例
func init() {
	apps.RegistryRestfulApp(&handler{})
}
