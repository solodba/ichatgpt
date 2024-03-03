package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/ichatgpt/apps/model"
	"github.com/solodba/mcube/response"
)

// 创建model
func (h *handler) ListModels(r *restful.Request, w *restful.Response) {
	req := model.NewListModelsRequest()
	// err := r.ReadEntity(req)
	// if err != nil {
	// 	w.WriteEntity(response.NewFail(400, err.Error()))
	// 	return
	// }
	resp, err := h.svc.ListModels(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, resp))
}

// 查询model
func (h *handler) RetrieveModel(r *restful.Request, w *restful.Response) {
	req := model.NewRetrieveModelRequestFromRestful(r)
	resp, err := h.svc.RetrieveModel(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, resp))
}

// 删除DeleteFineTuned模型
func (h *handler) DeleteFineTunedModel(r *restful.Request, w *restful.Response) {
	req := model.NewDeleteFineTunedModelRequestFromRestful(r)
	resp, err := h.svc.DeleteFineTunedModel(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, resp))
}
