package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/ichatgpt/apps/model"
	"github.com/solodba/mcube/apps"
)

// handler结构体
type handler struct {
	svc model.Service
}

// 实现注册restful实例Name方法
func (h *handler) Name() string {
	return model.AppName
}

// 实现注册restful实例Conf方法
func (h *handler) Conf() error {
	// 从IOC获取token实例
	h.svc = apps.GetInternalApp(model.AppName).(model.Service)
	return nil
}

// 实现注册restful实例version方法
func (h *handler) Version() string {
	return "v1"
}

// 实现注册restful实例Registry方法
func (h *handler) RegistryHandler(ws *restful.WebService) {
	tags := []string{"Model管理"}
	// webservice定义路由信息
	// 查看所有Model
	ws.Route(ws.GET("/").To(h.ListModels).
		Doc("查看所有Model").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(model.ListModelsRequest{}).
		Writes(model.ModelsResponse{}).
		Returns(200, "OK", model.ModelsResponse{}))

	// 查询指定Model
	ws.Route(ws.GET("/{model}").To(h.RetrieveModel).
		Doc("查询指定Model").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(model.RetrieveModelRequest{}).
		Writes(model.ModelsItem{}).
		Returns(200, "OK", model.ModelsItem{}))

	// 删除FineTuneModel
	ws.Route(ws.DELETE("/{model}").To(h.DeleteFineTunedModel).
		Doc("删除FineTune模型").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(model.DeleteFineTunedModelRequest{}).
		Writes(model.DeleteFineTunedModelResponse{}).
		Returns(200, "OK", model.DeleteFineTunedModelResponse{}))
}

// 初始化函数注册restful实例
func init() {
	apps.RegistryRestfulApp(&handler{})
}
