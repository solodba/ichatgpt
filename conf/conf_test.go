package conf_test

import (
	"testing"

	"github.com/solodba/ichatgpt/conf"
)

func TestLoadConfigFromToml(t *testing.T) {
	err := conf.LoadConfigFromToml("test/test.toml")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.C().App.Name)
	t.Log(conf.C().App.Http)
	t.Log(conf.C().ChatGPT.OpenaiApiKey)
	t.Log(conf.C().ChatGPT.HttpProxy)
	t.Log(conf.C().ChatGPT.HttpsProxy)
}

func TestLoadConfigFromEnv(t *testing.T) {
	err := conf.LoadConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.C().App.Name)
	t.Log(conf.C().App.Http)
	t.Log(conf.C().ChatGPT.OpenaiApiKey)
	t.Log(conf.C().ChatGPT.HttpProxy)
	t.Log(conf.C().ChatGPT.HttpsProxy)
}
