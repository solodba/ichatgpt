package impl_test

import (
	"testing"

	"github.com/solodba/ichatgpt/apps/finetune"
	"github.com/solodba/ichatgpt/test/tools"
)

func TestCreateFineTuneJob(t *testing.T) {
	fineTuneJobReq := finetune.NewCreateFineTuneJobRequest()
	fineTuneJobReq.Model = "gpt-3.5-turbo"
	fineTuneJobReq.TrainingFile = "file-XcsjryvD4YbWxwyblju5FZ0H"
	fineTuneJobResp, err := svc.CreateFineTuneJob(ctx, fineTuneJobReq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(fineTuneJobResp))
}

func TestListFineTuneJobEvents(t *testing.T) {
	fineTuneJobEventsReq := finetune.NewListFineTuneJobEventsRequest()
	fineTuneJobEventsReq.FineTuningJobId = "ftjob-EV9nv185osolVAis9qCi9Pfy"
	fineTuneJobEventsResp, err := svc.ListFineTuneJobEvents(ctx, fineTuneJobEventsReq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(fineTuneJobEventsResp))
}
