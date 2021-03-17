package provider

import (
	"reflect"
	"time"

	"github.com/ingot-cloud/ingot-go/pkg/framework/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/container/di"
	securityContainer "github.com/ingot-cloud/ingot-go/pkg/framework/container/security"

	"github.com/ingot-cloud/ingot-go/pkg/framework/log"
)

// BuildContainerProcess 构建容器过程
func BuildContainerProcess(pre container.ContainerPre, ij *di.ProviderSet) container.ContainerPrint {
	log.Debug("											  ")
	log.Debug("===========================================")
	log.Debug("====== BuildContainerProcess 开始执行 ======")
	log.Debug("===========================================")
	startNanosecond := time.Now().Nanosecond()

	injectFields := paddingInjectFields(pre.GetContainerInjector().GetSecurityInjector())
	con := ij.Parse(injectFields, pre)

	log.Debug("======================================================")
	log.Debugf("====== BuildContainerProcess 执行结束，用时%d毫秒 ======", (time.Now().Nanosecond()-startNanosecond)/1e6)
	log.Debug("======================================================")
	log.Debug("											  ")
	return con
}

func paddingInjectFields(injector securityContainer.SecurityInjector) []*di.Injector {
	var injectFields []*di.Injector

	inValue := reflect.Indirect(reflect.ValueOf(injector))
	inType := inValue.Type()
	len := inType.NumField()

	var field reflect.StructField
	var injectTag string
	for i := 0; i < len; i++ {
		field = inType.Field(i)
		injectTag = field.Tag.Get("inject")
		if injectTag == "true" {
			injectFields = append(injectFields, &di.Injector{
				Value: inValue.FieldByName(field.Name),
				Type:  field.Type,
			})
		}
	}

	return injectFields
}
