package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/ichatgpt/apps/finetune"
	"github.com/solodba/mcube/response"
)

// 创建FineTuneJob
func (h *handler) CreateFineTuneJob(r *restful.Request, w *restful.Response) {
	req := finetune.NewCreateFineTuneJobRequest()
	err := r.ReadEntity(req)
	if err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	resp, err := h.svc.CreateFineTuneJob(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, resp))
}

// 查询FineTuneJob事件
func (h *handler) ListFineTuneJobEvents(r *restful.Request, w *restful.Response) {
	req := finetune.NewListFineTuneJobEventsRequestFromRestful(r)
	resp, err := h.svc.ListFineTuneJobEvents(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, resp))
}

// 查询FineTuneJob
func (h *handler) RetrieveFineTuneJob(r *restful.Request, w *restful.Response) {
	req := finetune.NewRetrieveFineTuneJobRequestFromRestful(r)
	resp, err := h.svc.RetrieveFineTuneJob(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, resp))
}

// 停止FineTuneJob
func (h *handler) CancelFineTuneJob(r *restful.Request, w *restful.Response) {
	req := finetune.NewCancelFineTuneJobRequestFromRestful(r)
	resp, err := h.svc.CancelFineTuneJob(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, resp))
}
