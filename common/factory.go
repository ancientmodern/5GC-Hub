package common

type Method interface {
	Run()
	GetInterval() int
}

var factorySet = make(map[string]Method)

func Register(name string, factory Method) {
	factorySet[name] = factory
}

func GetMethodSet() map[string]Method {
	return factorySet
}

func init() {
	// getCommonConfig()
}
