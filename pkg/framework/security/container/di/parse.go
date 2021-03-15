package di

import (
	"reflect"
)

// 构建提供者构建方法集合
// 通过wire构建injector和container，解析拿到所有实例，先将injector中自定义的实例替换到container中相应类型
// 然后依次判断container中替换后的实例是否依赖（深判断）了刚刚自定义的实例，如果依赖了，那么使用新的实例逐一构建并且替换之前的值

// NewSet 实例化数据集
func NewSet(items ...interface{}) *ProviderSet {
	set := &ProviderSet{}

	for _, i := range items {
		switch v := i.(type) {
		case *Provider:
			set.AddProvider(v)
		case *IfaceBinding:
			set.AddBinding(v)
		}
	}

	return set
}

// Struct 构建结构体provider
func Struct(target interface{}) *Provider {
	t := indirect(reflect.TypeOf(target))

	if t.Kind() != reflect.Struct {
		panic("使用Struct方法构建Provider，入参Kind必须为reflect.Struct")
	}

	num := t.NumField()
	args := make([]reflect.Type, 0, num)

	for i := 0; i < num; i++ {
		args = append(args, t.Field(i).Type)
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
	args := make([]reflect.Type, 0, num)

	for i := 0; i < num; i++ {
		args = append(args, indirect(t.In(i)))
	}

	return &Provider{
		Type:     t,
		IsStruct: false,
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

	// 构建所需要的参数
	Args []reflect.Type
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
func (p *Provider) DependsOn(target reflect.Type) bool {
	for _, t := range p.Args {
		if t == target {
			return true
		}
	}
	return false
}

// IfaceBinding 接口绑定关系
type IfaceBinding struct {
	// 接口类型
	Iface reflect.Type

	// 接口对应的提供者类型
	Provider reflect.Type
}

// CustomInjector 自定义注入
type CustomInjector struct {
	// 类型
	Type reflect.Type

	// 值
	Value reflect.Value
}

// ProviderSet 提供者集合
type ProviderSet struct {
	// 所有提供者
	Providers map[reflect.Type]*Provider

	// 所有绑定关系
	Bindings []*IfaceBinding
}

// AddProvider 添加类型提供者
func (set *ProviderSet) AddProvider(p *Provider) {
	if set.Providers == nil {
		set.Providers = make(map[reflect.Type]*Provider)
	}
	set.Providers[p.Type] = p
}

// AddBinding 添加 interface 和 struct 绑定关系
func (set *ProviderSet) AddBinding(b *IfaceBinding) {
	set.Bindings = append(set.Bindings, b)
}

// Parse 解析 CustomInjector
func (set *ProviderSet) Parse(cj []*CustomInjector) {
	// todo
	// for _, in := range cj {

	// }
}

// 获取依赖指定类型的所有 Provider
func (set *ProviderSet) getDependsOn(t reflect.Type) []*Provider {
	var result []*Provider
	for _, p := range set.Providers {
		if p.DependsOn(t) {
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
