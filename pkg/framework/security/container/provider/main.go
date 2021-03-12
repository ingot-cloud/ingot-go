package provider

import (
	"reflect"

	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/pkg/framework/log"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"

	coreUtils "github.com/ingot-cloud/ingot-go/pkg/framework/core/utils"
)

// All 所有实例
var All = wire.NewSet(
	CommonContainerFields,
	CommonContainer,
	OAuth2ContainerFields,
	OAuth2Container,
	AuthorizationServerContainerFields,
	AuthorizationServerContainer,
	ResourceServerContainerFields,
	ResourceServerContainer,
	AuthProvidersContainer,
	AuthProvidersContainerFields,
	SecurityContainer,
	PrintInjectInstance,
)

// SecurityContainer 安全容器
var SecurityContainer = wire.NewSet(
	wire.Struct(new(container.SecurityContainerImpl), "*"),
	wire.Bind(new(container.SecurityContainer), new(*container.SecurityContainerImpl)),
)

// PrintInjectInstance 打印注入
func PrintInjectInstance(sc container.SecurityContainer) container.PrintSecurityInjector {
	log.Debug("===========================================")
	log.Debug("== 开始打印 SecurityContainer 中所有注入实例 ==")
	log.Debug("===========================================")

	printContainer(sc)

	log.Debug("===========================================")
	log.Debug("== 结束打印 SecurityContainer 中所有注入实例 ==")
	log.Debug("===========================================")
	var result struct{}
	return &result
}

func printContainer(c interface{}) {
	value := reflect.Indirect(reflect.ValueOf(c))
	targetType := value.Type()
	len := targetType.NumField()
	for i := 0; i < len; i++ {
		sf := targetType.Field(i)
		printContainerChild(value.FieldByName(sf.Name).Interface())
	}
}

func printContainerChild(c interface{}) {
	value := reflect.Indirect(reflect.ValueOf(c))
	targetType := value.Type()
	len := targetType.NumField()
	log.Infof("-------> 容器 %s 注入字段打印开始", targetType.Name())
	for i := 0; i < len; i++ {
		sf := targetType.Field(i)
		field := value.FieldByName(sf.Name)
		log.Infof("字段名: %s, %s", sf.Name, coreUtils.GetType(field.Interface()))
	}
	log.Infof("<------- 容器 %s 注入字段打印结束", targetType.Name())
}
