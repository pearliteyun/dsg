package service

import (
	"github.com/xgpc/dsg/frame"
	"github.com/xgpc/dsg/models"
	"github.com/xgpc/dsg/service/cryptService"
	"github.com/xgpc/dsg/service/grpcService/proto"
	"github.com/xgpc/dsg/service/schedule"
	"github.com/xgpc/dsg/service/sysService"
	"github.com/xgpc/dsg/service/validatorService"
)

func LoadService() {
	if frame.Config.SysConfig.StartSchedule {
		schedule.Start()
		schedule.StartSchedules()
	}

	// 参数验证器配置
	validatorService.GetTranslations()

	//	系统时间版本
	sysService.InitSysVersion()

	//	默认user表
	models.InitUser()

	// rsa秘钥初始化
	cryptService.GenerateRSAKey(1024)

	// grpc连接初始化
	proto.GRPCConnect()

	// 启动缓存
	CacheStart()
}
