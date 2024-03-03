package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/ichatgpt/apps/file"
	"github.com/solodba/mcube/apps"
)

// handler结构体
type handler struct {
	svc file.Service
}

// 实现注册restful实例Name方法
func (h *handler) Name() string {
	return file.AppName
}

// 实现注册restful实例Conf方法
func (h *handler) Conf() error {
	// 从IOC获取token实例
	h.svc = apps.GetInternalApp(file.AppName).(file.Service)
	return nil
}

// 实现注册restful实例version方法
func (h *handler) Version() string {
	return "v1"
}

// 实现注册restful实例Registry方法
func (h *handler) RegistryHandler(ws *restful.WebService) {
	tags := []string{"File管理"}
	// webservice定义路由信息
	// 创建Pod
	ws.Route(ws.POST("/").To(h.UploadFile).
		Doc("上传文件").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(file.UploadFileRequest{}).
		Writes(file.FileResponseItem{}).
		Returns(200, "OK", file.FileResponseItem{}))
}

// 初始化函数注册restful实例
func init() {
	apps.RegistryRestfulApp(&handler{})
}
