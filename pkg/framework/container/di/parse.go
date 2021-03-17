package di

import (
	"reflect"

	"github.com/ingot-cloud/ingot-go/pkg/framework/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/log"
)

// NewSet 实例化数据集
func NewSet(items ...interface{}) *ProviderSet {
	set := &ProviderSet{}

	for _, i := range items {
		switch v := i.(type) {
		case *Provider:
			set.AddProvider(v)
		case *IfaceBinding:
			set.AddBinding(v)
		case *ProviderSet:
			for _, p := range v.Providers {
				set.AddProvider(p)
			}
			for _, b := range v.Bindings {
				set.AddBinding(b)
			}
		}
	}

	return set
}

// Struct 构建结构体provider
func Struct(target interface{}, fieldNames ...string) *Provider {
	t := indirect(reflect.TypeOf(target))

	if t.Kind() != reflect.Struct {
		panic("使用Struct方法构建Provider，入参Kind必须为reflect.Struct")
	}

	num := t.NumField()
	args := make([]*ProviderParams, 0, num)

	// 判断是否注入所有字段
	isAllField := false
	if len(fieldNames) == 0 || fieldNames[0] == "*" {
		isAllField = true
	}
	var contains bool
	var field reflect.StructField
	for i := 0; i < num; i++ {
		field = t.Field(i)
		if !isAllField {
			for _, includeFiled := range fieldNames {
				if includeFiled == field.Name {
					contains = true
				}
			}
			if !contains {
				continue
			}
			contains = false
		}
		args = append(args, &ProviderParams{
			Type:      field.Type,
			FieldName: field.Name,
		})
	}

	return &Provider{
		Type:     t,
		IsStruct: true,
		Args:     args,
	}
}

// Func 构建方法provider
func Func(target interface{}) *Provider {
	t := indirect(reflect.TypeOf(target))

	if t.Kind() != reflect.Func {
		panic("使用Func方法构建Provider，入参Kind必须为reflect.Func")
	}

	num := t.NumIn()
	args := make([]*ProviderParams, 0, num)

	var field reflect.Type
	for i := 0; i < num; i++ {
		field = t.In(i)
		args = append(args, &ProviderParams{
			Type: indirect(field),
		})
	}

	return &Provider{
		Type:     t,
		IsStruct: false,
		Return:   t.Out(0),
		Args:     args,
	}
}

// Bind 绑定接口和结构体的关系
func Bind(iface, impl interface{}) *IfaceBinding {
	return &IfaceBinding{
		Iface:    indirect(reflect.TypeOf(iface)),
		Provider: indirect(reflect.TypeOf(impl)),
	}
}

// Provider 类型提供者
type Provider struct {
	// 提供者构建的类型
	Type reflect.Type

	// 是否为struct，如果不是则为func
	IsStruct bool

	// 如果是 func 构建的类型
	Return reflect.Type

	// 构建所需要的参数
	Args []*ProviderParams
}

// GetBuildType 获取构建类型
func (p *Provider) GetBuildType() reflect.Type {
	if p.IsStruct {
		return p.Type
	}
	return p.Return
}

// New 构建
func (p *Provider) New() reflect.Value {
	value := reflect.New(p.Type)

	if p.IsStruct {
		// todo 设置参数

	}

	return value
}

// DependsOn 是否依赖指定类型
// isImplIface代表是否判断指定类型是否实现了参数中的接口类型，一般自定义注入的类型实现了参数中的接口
func (p *Provider) DependsOn(target reflect.Type, isImplIface bool) bool {
	for _, t := range p.Args {
		if isImplIface && t.Type.Kind() == reflect.Interface && target.Implements(t.Type) {
			return true
		}
		if t.Type == target {
			return true
		}
	}
	return false
}

// ProviderParams 类型提供者构建所需要的参数
type ProviderParams struct {
	// 参数类型
	Type reflect.Type

	// 参数名称，struct使用，方法可以为空
	FieldName string
}

// IfaceBinding 接口绑定关系
type IfaceBinding struct {
	// 接口类型
	Iface reflect.Type

	// 接口对应的提供者类型
	Provider reflect.Type
}

// Injector 注入参数
type Injector struct {
	// 类型
	Type reflect.Type

	// 值
	Value reflect.Value
}

// InjectorNode 注入节点
type InjectorNode struct {
	*Injector

	//
	Parent []*Injector
}

// ProviderSet 提供者集合
type ProviderSet struct {
	// 所有提供者
	Providers []*Provider

	// 自定义绑定关系
	// todo 该关系暂时没有用到，因为在执行Parse方法的时候传入了SecurityContainer
	// 并且根据该容器填充了类型实例映射表，所以暂时没用到自定义类型关系
	Bindings []*IfaceBinding

	// 构建类型和对应实例映射
	TypeInstance map[reflect.Type]reflect.Value
}

// AddProvider 添加类型提供者
func (set *ProviderSet) AddProvider(p *Provider) {
	set.Providers = append(set.Providers, p)
}

// AddBinding 添加 interface 和 struct 绑定关系
func (set *ProviderSet) AddBinding(b *IfaceBinding) {
	set.Bindings = append(set.Bindings, b)
}

// Parse 执行操作如下：
// 1. 解析容器，填充容器中所有类型和实例映射表
// 2. 解析 Injector 获取所有需要重新构建的类型
// 判断标准如下：
// 	  在自定义接口实现数组中，如果存在自定义实现A和B，且A深度依赖B，那么此依赖链上的所有实例均需要重新构建，
//    并且如果B也依赖了其他自定义实现，那么B也需要重新构建，以此类推，直到依赖的对象为默认构建的对象。
// 3. 将需要重构的类型和自定义的注入类型一起判断，生成依赖关系树，并从叶子节点逐一重新构建（根据依赖关系）
// 4. 用重新构建好的实例替换Container容器中之前的实例，并且返回新的Container
func (set *ProviderSet) Parse(cj []*Injector, c container.Container) container.Container {

	// 填充当前类型实例映射
	set.paddingTypeInstanceWithContainer(c)

	// 收集需要重新构建的类型
	rebuildMap := make(map[*Provider]int)
	for _, in := range cj {
		// 由于自定义注入类型，都是实现了相关接口，所有在校验依赖关系的时候，需要判断是否为接口实现
		set.mergeRebuild(rebuildMap, in.Type, true)
	}

	for p := range rebuildMap {
		log.Errorf("需要重新构建的实例，类型：%s", p.GetBuildType())
		for _, in := range cj {
			if p.GetBuildType().Kind() == reflect.Interface && in.Type.Implements(p.GetBuildType()) {
				log.Errorf("自定义类型 %s 实现了接口 %s", in.Type, p.GetBuildType())
			}
		}
	}

	// todo

	return c
}

// 填充Container中的子容器
func (set *ProviderSet) paddingTypeInstanceWithContainer(con container.Container) {
	containerValue := reflect.Indirect(reflect.ValueOf(con))
	containerType := containerValue.Type()
	fieldLen := containerType.NumField()

	var field reflect.StructField
	var tag string
	for i := 0; i < fieldLen; i++ {
		field = containerType.Field(i)
		tag = field.Tag.Get("container")
		if tag == "true" {
			set.paddingTypeInstance(containerValue.Field(i).Interface())
		}
	}
}

// 填充类型实例映射表
func (set *ProviderSet) paddingTypeInstance(con interface{}) {
	set.TypeInstance = make(map[reflect.Type]reflect.Value)

	paddingChild := func(con interface{}) {
		originValue := reflect.ValueOf(con)
		value := reflect.Indirect(originValue)
		t := value.Type()
		set.TypeInstance[t] = originValue

		log.Errorf("padding type: %s, value: %s", reflect.TypeOf(con), value.Interface())
		len := t.NumField()
		for i := 0; i < len; i++ {
			log.Errorf("padding type: %s, value: %s", t.Field(i).Type, reflect.ValueOf(value.Field(i).Interface()).Type())
			// value.Field(i) 获取的 reflect.Value 为接口，需要获取具体值然后在拿到对应的 reflect.Value
			set.TypeInstance[t.Field(i).Type] = reflect.ValueOf(value.Field(i).Interface())
		}
	}

	containerValue := reflect.Indirect(reflect.ValueOf(con))
	containerType := containerValue.Type()
	fieldLen := containerType.NumField()
	for i := 0; i < fieldLen; i++ {
		paddingChild(containerValue.Field(i).Interface())
	}
}

// 构建需要重新创建的实例
func (set *ProviderSet) mergeRebuild(rebuildMap map[*Provider]int, t reflect.Type, isImplIface bool) {
	providers := set.getDependsOn(t, isImplIface)
	if len(providers) != 0 {
		for _, p := range providers {
			rebuildMap[p] = 1
			// 校验构建当前类型所需要参数的依赖关系时，不需要校验是否为接口实现
			set.mergeRebuild(rebuildMap, p.GetBuildType(), false)
		}
	}
}

// 获取依赖指定类型的所有 Provider
func (set *ProviderSet) getDependsOn(t reflect.Type, isImplIface bool) []*Provider {
	var result []*Provider
	for _, p := range set.Providers {
		if p.DependsOn(t, isImplIface) {
			result = append(result, p)
		}
	}
	return result
}

func indirect(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		return t.Elem()
	}
	return t
}
