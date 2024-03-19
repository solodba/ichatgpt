package api

import (
	"io"
	"os"

	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/ichatgpt/apps/file"
	"github.com/solodba/mcube/response"
)

// 上传文件
func (h *handler) UploadFile(r *restful.Request, w *restful.Response) {
	req := file.NewUploadFileRequest()
	err := r.ReadEntity(req)
	if err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	resp, err := h.svc.UploadFile(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, resp))
}

// 查看文件
func (h *handler) ListFile(r *restful.Request, w *restful.Response) {
	req := file.NewListFileRequestFromRestful(r)
	resp, err := h.svc.ListFile(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, resp))
}

// 查看文件详情
func (h *handler) RetrieveFile(r *restful.Request, w *restful.Response) {
	req := file.NewRetrieveFileRequestFromRestful(r)
	resp, err := h.svc.RetrieveFile(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, resp))
}

// 查询文件内容
func (h *handler) RetrieveFileContent(r *restful.Request, w *restful.Response) {
	req := file.NewRetrieveFileContentRequestFromRestful(r)
	resp, err := h.svc.RetrieveFileContent(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, resp))
}

// 删除文件
func (h *handler) DeleteFile(r *restful.Request, w *restful.Response) {
	req := file.NewDeleteFileRequestFromRestful(r)
	resp, err := h.svc.DeleteFile(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, resp))
}

// web上传文件
func (h *handler) WebUploadFile(r *restful.Request, w *restful.Response) {
	f, err := os.OpenFile("file/finetuning.jsonl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	defer f.Close()
	_, err = io.Copy(f, r.Request.Body)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
}
