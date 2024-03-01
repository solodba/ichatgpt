package impl_test

import (
	"testing"

	"github.com/solodba/ichatgpt/apps/file"
	"github.com/solodba/ichatgpt/test/tools"
)

func TestUploadFile(t *testing.T) {
	uploadFileReq := file.NewUploadFileRequest()
	uploadFileReq.FilePath = "file"
	uploadFileReq.FileName = "mydata.jsonl"
	uploadFileReq.Purpose = "fine-tune"
	uploadFileResp, err := svc.UploadFile(ctx, uploadFileReq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(uploadFileResp))
}

func TestListFile(t *testing.T) {
	listFileReq := file.NewListFileRequest()
	listFileResp, err := svc.ListFile(ctx, listFileReq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(listFileResp))
}

func TestRetrieveFile(t *testing.T) {
	retrieveFileReq := file.NewRetrieveFileRequest()
	retrieveFileReq.FileId = "file-XyrHCz4GYmcwOzAyVIWfF7I2"
	retrieveFileResp, err := svc.RetrieveFile(ctx, retrieveFileReq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(retrieveFileResp))
}

func TestDeleteFile(t *testing.T) {
	deleteFileReq := file.NewDeleteFileRequest()
	deleteFileReq.FileId = "file-XyrHCz4GYmcwOzAyVIWfF7I2"
	deleteFileResp, err := svc.DeleteFile(ctx, deleteFileReq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(deleteFileResp))
}
