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
	// 上传文件
	ws.Route(ws.POST("/").To(h.UploadFile).
		Doc("上传文件").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(file.UploadFileRequest{}).
		Writes(file.FileResponseItem{}).
		Returns(200, "OK", file.FileResponseItem{}))

	// 查询文件
	ws.Route(ws.GET("/").To(h.ListFile).
		Doc("查询文件").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(file.ListFileRequest{}).
		Writes(file.FileResponse{}).
		Returns(200, "OK", file.FileResponse{}))

	// 查询文件详情
	ws.Route(ws.GET("/detail/{file_id}").To(h.RetrieveFile).
		Doc("查询文件详情").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(file.RetrieveFileRequest{}).
		Writes(file.FileResponseItem{}).
		Returns(200, "OK", file.FileResponseItem{}))

	// 查询文件内容
	ws.Route(ws.GET("/content/{file_id}").To(h.RetrieveFileContent).
		Doc("查询文件内容").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(file.RetrieveFileContentRequest{}).
		Writes(file.RetrieveFileContentResponse{}).
		Returns(200, "OK", file.RetrieveFileContentResponse{}))

	// 删除文件
	ws.Route(ws.DELETE("/{file_id}").To(h.DeleteFile).
		Doc("删除文件").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(file.DeleteFileRequest{}).
		Writes(file.DeleteFileRequest{}).
		Returns(200, "OK", file.DeleteFileResponse{}))
}

// 初始化函数注册restful实例
func init() {
	apps.RegistryRestfulApp(&handler{})
}
