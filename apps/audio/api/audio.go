package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/ichatgpt/apps/audio"
	"github.com/solodba/mcube/response"
)

// 创建Speech
func (h *handler) CreateSpeech(r *restful.Request, w *restful.Response) {
	req := audio.NewCreateSpeechRequest()
	err := r.ReadEntity(req)
	if err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	resp, err := h.svc.CreateSpeech(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, resp))
}

// 创建Transcription
func (h *handler) CreateTranscription(r *restful.Request, w *restful.Response) {
	req := audio.NewCreateTranscriptionRequest()
	err := r.ReadEntity(req)
	if err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	resp, err := h.svc.CreateTranscription(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, resp))
}

// 创建Translation
func (h *handler) CreateTranslation(r *restful.Request, w *restful.Response) {
	req := audio.NewCreateTranslationRequest()
	err := r.ReadEntity(req)
	if err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	resp, err := h.svc.CreateTranslation(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, resp))
}
