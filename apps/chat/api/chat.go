package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/ichatgpt/apps/chat"
	"github.com/solodba/mcube/response"
)

// 创建chat
func (h *handler) CreateChat(r *restful.Request, w *restful.Response) {
	req := chat.NewCreateChatRequest()
	err := r.ReadEntity(req)
	if err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	pod, err := h.svc.CreateChat(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, pod))
}
