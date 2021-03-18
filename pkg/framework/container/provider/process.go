package provider

import (
	"time"

	"github.com/ingot-cloud/ingot-go/pkg/framework/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/container/di"

	"github.com/ingot-cloud/ingot-go/pkg/framework/log"
)

// BuildContainerProcess 构建容器过程
func BuildContainerProcess(pre container.ContainerPre, ij *di.ProviderSet) container.ContainerPrint {
	log.Debug("											  ")
	log.Debug("===========================================")
	log.Debug("====== BuildContainerProcess 开始执行 ======")
	log.Debug("===========================================")
	startNanosecond := time.Now().Nanosecond()

	con := ij.Parse(pre.GetContainerInjector(), pre)

	log.Debug("======================================================")
	log.Debugf("====== BuildContainerProcess 执行结束，用时%d毫秒 ======", (time.Now().Nanosecond()-startNanosecond)/1e6)
	log.Debug("======================================================")
	log.Debug("											  ")
	return con
}
