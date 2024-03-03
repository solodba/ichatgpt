package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/ichatgpt/apps/model"
	"github.com/solodba/mcube/response"
)

// 创建chat
func (h *handler) ListModels(r *restful.Request, w *restful.Response) {
	req := model.NewListModelsRequest()
	err := r.ReadEntity(req)
	if err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	resp, err := h.svc.ListModels(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, resp))
}
