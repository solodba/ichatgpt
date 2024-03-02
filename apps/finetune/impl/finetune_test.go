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

func TestRetrieveFineTuneJob(t *testing.T) {
	retrieveFineTuneJobReq := finetune.NewRetrieveFineTuneJobRequest()
	retrieveFineTuneJobReq.FineTuningJobId = "ftjob-EV9nv185osolVAis9qCi9Pfy"
	retrieveFineTuneJobResp, err := svc.RetrieveFineTuneJob(ctx, retrieveFineTuneJobReq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(retrieveFineTuneJobResp))
}

func TestCancelFineTuneJob(t *testing.T) {
	cancelFineTuneJobReq := finetune.NewCancelFineTuneJobRequest()
	cancelFineTuneJobReq.FineTuningJobId = "ftjob-vr09yqo7VysvOusukbbNQnUX"
	cancelFineTuneJobResp, err := svc.CancelFineTuneJob(ctx, cancelFineTuneJobReq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(cancelFineTuneJobResp))
}
