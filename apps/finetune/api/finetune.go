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
