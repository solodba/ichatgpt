package impl_test

import (
	"testing"

	"github.com/solodba/ichatgpt/apps/model"
	"github.com/solodba/ichatgpt/test/tools"
)

func TestListModels(t *testing.T) {
	listModelsReq := model.NewListModelsRequest()
	modelResp, err := svc.ListModels(ctx, listModelsReq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(modelResp))
}

func TestRetrieveModel(t *testing.T) {
	retrieveModelReq := model.NewRetrieveModelRequest()
	retrieveModelReq.Model = "gpt-4-0125-preview"
	modelResp, err := svc.RetrieveModel(ctx, retrieveModelReq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(modelResp))
}
