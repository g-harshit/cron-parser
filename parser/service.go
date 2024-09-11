package parser

type Service interface {
	Parse() error
	Print()
}

var defaultObject Service

func Init(obj Service) {
	defaultObject = obj
}

func GetService() Service {
	return defaultObject
}
