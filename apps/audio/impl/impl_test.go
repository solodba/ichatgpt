package impl_test

import (
	"context"

	"github.com/solodba/ichatgpt/apps/audio"
	"github.com/solodba/ichatgpt/test/tools"
	"github.com/solodba/mcube/apps"
)

// 全局变量
var (
	svc audio.Service
	ctx = context.Background()
)

// 初始化函数
func init() {
	tools.DevelopmentSet()
	svc = apps.GetInternalApp(audio.AppName).(audio.Service)
}
