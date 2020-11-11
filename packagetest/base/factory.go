package base

//这个包是base 负责处理注册和使用工厂
//类接口
type Class interface {
	Do()
}

var (
	//保存注册好的工厂信息
	factory_By_Name = make(map[string]func() Class)
)

//注册一个类生产工厂，提供给工厂方注册使用，所谓工厂，就是一个定义为func() Class 的普通函数，调用此函数，创建一个类实例，实现的工程内部
//结构体会实现class接口
func Register(name string, factory func() Class) {
	factory_By_Name[name] = factory
}

/*
通过名字创建类实例的函数，该函数会在注册好后调用
*/
func Create(name string) Class {
	//已经注册好信息中查找名字对应的工厂函数，找到以后，在下面行调用并返回接口
	if f, ok := factory_By_Name[name]; ok {
		return f()
	} else {
		panic("name not found")
	}
}
