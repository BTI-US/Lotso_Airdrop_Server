package cron

import (
	"Lotso_Airdrop_Server/model/base"
	"Lotso_Airdrop_Server/service"
	"github.com/labstack/gommon/log"
	"github.com/robfig/cron/v3"
)

func StartWorker(spec string) {
	c := cron.New()
	_, _ = c.AddFunc(spec, setAirdropWorker)
	c.Start()
}

func setAirdropWorker() {
	response := service.DistributeAirdrops()
	if response.Code != base.SUCCESS.Code {
		log.Error(response.Code, response.Message, response.Error)
	} else {
		log.Info("AirdropWorker Success: ", response.Data)
	}
}
