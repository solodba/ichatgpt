package tools

import (
	_ "github.com/solodba/ichatgpt/apps/all"
	"github.com/solodba/ichatgpt/conf"
	"github.com/solodba/mcube/apps"
	"github.com/solodba/mcube/logger"
)

func DevelopmentSet() {
	err := conf.LoadConfigFromEnv()
	if err != nil {
		logger.L().Panic().Msgf("load global config error, err: %s", err.Error())
	}
	err = apps.InitInternalApps()
	if err != nil {
		logger.L().Panic().Msgf("initial object config error, err: %s", err.Error())
	}
}
