package impl_test

import (
	"testing"

	"github.com/solodba/ichatgpt/apps/model"
	"github.com/solodba/ichatgpt/test/tools"
)

func TestListModels(t *testing.T) {
	listModelsReq := model.NewListModelsRequest()
	listModelsResp, err := svc.ListModels(ctx, listModelsReq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(listModelsResp))
}
