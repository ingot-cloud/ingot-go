package process

import (
	"reflect"
	"time"

	"github.com/ingot-cloud/ingot-go/pkg/framework/container/di"
	securityContainer "github.com/ingot-cloud/ingot-go/pkg/framework/container/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/log"
)

type InjectField struct {
	Value reflect.Value
	Type  reflect.Type
}

var injectFields []*InjectField

func DoPre(injector securityContainer.SecurityInjector, sc securityContainer.SecurityContainerPre) securityContainer.SecurityContainerCombine {
	log.Debug(">>>>>> DoPre 开始执行")
	injectFields = nil
	startNanosecond := time.Now().Nanosecond()
	// 填充需要替换的自定义实现
	paddingInjectFields(injector)

	doChangeContainer(sc)

	var rebuildArray []*di.Injector
	for _, item := range injectFields {
		rebuildArray = append(rebuildArray, &di.Injector{
			Type:  item.Type,
			Value: item.Value,
		})
	}

	log.Errorf("rebuildArray len = %d", len(rebuildArray))
	// ProviderSet.Parse(rebuildArray, sc)

	log.Debugf(">>>>>> DoPre 执行结束，用时%d毫秒", (time.Now().Nanosecond()-startNanosecond)/1e6)
	return sc
}

func DoPost(injector securityContainer.SecurityInjector, sc securityContainer.SecurityContainerPost) securityContainer.SecurityContainer {
	log.Debug(">>>>>> DoPost 开始执行")
	injectFields = nil
	startNanosecond := time.Now().Nanosecond()
	// 填充需要替换的自定义实现
	paddingInjectFields(injector)

	doChangeContainer(sc)

	log.Debugf(">>>>>> DoPost 执行结束，用时%d毫秒", (time.Now().Nanosecond()-startNanosecond)/1e6)
	return sc
}

func doChangeContainer(sc interface{}) {
	value := reflect.Indirect(reflect.ValueOf(sc))
	targetType := value.Type()
	len := targetType.NumField()
	for i := 0; i < len; i++ {
		sf := targetType.Field(i)
		doChangeChild(value.FieldByName(sf.Name).Interface())
	}
}

func doChangeChild(c interface{}) {
	value := reflect.Indirect(reflect.ValueOf(c))
	t := value.Type()
	len := t.NumField()
	for i := 0; i < len; i++ {
		changeField(t.Field(i), value)
	}
}

func paddingInjectFields(injector securityContainer.SecurityInjector) {
	inValue := reflect.Indirect(reflect.ValueOf(injector))
	inType := inValue.Type()
	len := inType.NumField()

	var field reflect.StructField
	var injectTag string
	for i := 0; i < len; i++ {
		field = inType.Field(i)
		injectTag = field.Tag.Get("inject")
		if injectTag == "true" {
			injectFields = append(injectFields, &InjectField{
				Value: inValue.FieldByName(field.Name),
				Type:  field.Type,
			})
		}
	}

}

func changeField(field reflect.StructField, target reflect.Value) {
	fieldName := field.Name
	fieldValue := target.FieldByName(fieldName)

	for _, injectField := range injectFields {
		// 只处理接口
		if field.Type.Kind() != reflect.Interface {
			continue
		}
		if injectField.Type.Implements(field.Type) {
			fieldValue.Set(injectField.Value)
		}
	}
}
