package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/ichatgpt/apps/image"
	"github.com/solodba/mcube/response"
)

// 创建image
func (h *handler) CreateImage(r *restful.Request, w *restful.Response) {
	req := image.NewCreateImageRequest()
	err := r.ReadEntity(req)
	if err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	resp, err := h.svc.CreateImage(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, resp))
}
