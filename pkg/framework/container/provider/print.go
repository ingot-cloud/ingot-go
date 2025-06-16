package provider

import (
	"reflect"

	"github.com/ingot-cloud/ingot-go/pkg/framework/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/utils"
	"github.com/ingot-cloud/ingot-go/pkg/framework/log"
)

// PrintInjectInstance 打印注入
func PrintInjectInstance(con container.ContainerPrint) container.Container {
	log.Debug("===========================================")
	log.Debug("====== 开始打印 Container 中所有注入实例 ======")
	log.Debug("===========================================")

	containerValue := reflect.Indirect(reflect.ValueOf(con))
	containerType := containerValue.Type()
	fieldLen := containerType.NumField()

	var field reflect.StructField
	var tag string
	for i := 0; i < fieldLen; i++ {
		field = containerType.Field(i)
		tag = field.Tag.Get("container")
		if tag == "true" {
			printContainer(containerValue.Field(i).Interface())
		}
	}

	log.Debug("===========================================")
	log.Debug("====== 结束打印 Container 中所有注入实例 ======")
	log.Debug("===========================================")
	return con
}

func printContainer(c any) {
	value := reflect.Indirect(reflect.ValueOf(c))
	targetType := value.Type()
	len := targetType.NumField()
	for i := 0; i < len; i++ {
		sf := targetType.Field(i)
		printContainerChild(value.FieldByName(sf.Name).Interface())
	}
}

func printContainerChild(c any) {
	value := reflect.Indirect(reflect.ValueOf(c))
	targetType := value.Type()
	len := targetType.NumField()
	log.Infof("-------> 容器 %s 注入字段打印开始", targetType.Name())
	for i := 0; i < len; i++ {
		sf := targetType.Field(i)
		field := value.FieldByName(sf.Name)
		log.Infof("字段名: %s, %s", sf.Name, utils.GetType(field.Interface()))
	}
	log.Infof("<------- 容器 %s 注入字段打印结束", targetType.Name())
}
