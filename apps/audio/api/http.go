package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/ichatgpt/apps/audio"
	"github.com/solodba/mcube/apps"
)

// handler结构体
type handler struct {
	svc audio.Service
}

// 实现注册restful实例Name方法
func (h *handler) Name() string {
	return audio.AppName
}

// 实现注册restful实例Conf方法
func (h *handler) Conf() error {
	// 从IOC获取token实例
	h.svc = apps.GetInternalApp(audio.AppName).(audio.Service)
	return nil
}

// 实现注册restful实例version方法
func (h *handler) Version() string {
	return "v1"
}

// 实现注册restful实例Registry方法
func (h *handler) RegistryHandler(ws *restful.WebService) {
	tags := []string{"Audio管理"}
	// webservice定义路由信息
	// 创建Speech
	ws.Route(ws.POST("/speech").To(h.CreateSpeech).
		Doc("文字转语音").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(audio.CreateSpeechRequest{}).
		Writes(audio.CreateSpeechResponse{}).
		Returns(200, "OK", audio.CreateSpeechResponse{}))

	// 创建Transcription
	ws.Route(ws.POST("/transcription").To(h.CreateTranscription).
		Doc("语音转文字").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(audio.CreateTranscriptionRequest{}).
		Writes(audio.CreateAudioResponse{}).
		Returns(200, "OK", audio.CreateAudioResponse{}))

	// 创建Translation
	ws.Route(ws.POST("/translation").To(h.CreateTranslation).
		Doc("语音翻译成英文").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(audio.CreateTranslationRequest{}).
		Writes(audio.CreateAudioResponse{}).
		Returns(200, "OK", audio.CreateAudioResponse{}))
}

// 初始化函数注册restful实例
func init() {
	apps.RegistryRestfulApp(&handler{})
}
