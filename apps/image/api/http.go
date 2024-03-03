package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/ichatgpt/apps/image"
	"github.com/solodba/mcube/apps"
)

// handler结构体
type handler struct {
	svc image.Service
}

// 实现注册restful实例Name方法
func (h *handler) Name() string {
	return image.AppName
}

// 实现注册restful实例Conf方法
func (h *handler) Conf() error {
	// 从IOC获取token实例
	h.svc = apps.GetInternalApp(image.AppName).(image.Service)
	return nil
}

// 实现注册restful实例version方法
func (h *handler) Version() string {
	return "v1"
}

// 实现注册restful实例Registry方法
func (h *handler) RegistryHandler(ws *restful.WebService) {
	tags := []string{"Image管理"}
	// webservice定义路由信息
	// 创建Image
	ws.Route(ws.POST("/").To(h.CreateImage).
		Doc("创建Image").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(image.CreateImageRequest{}).
		Writes(image.CreateImageResponse{}).
		Returns(200, "OK", image.CreateImageResponse{}))
}

// 初始化函数注册restful实例
func init() {
	apps.RegistryRestfulApp(&handler{})
}
