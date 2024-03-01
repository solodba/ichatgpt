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
	retrieveFileReq.FileId = "file-Jee40QArmA7ltu5X4OwLTLK7"
	retrieveFileResp, err := svc.RetrieveFile(ctx, retrieveFileReq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(retrieveFileResp))
}
