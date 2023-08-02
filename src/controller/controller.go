package controller

import "context"

type Controller interface {
	Request() interface{}
	Response() interface{}
	Process(ctx context.Context, req interface{}) (res interface{}, err error)
}
type GetControllerFunc func() Controller

var HandlerMap map[string]GetControllerFunc

func init() {
	HandlerMap = make(map[string]GetControllerFunc, 0)
}

func Register(urlName string, controllerFunc GetControllerFunc) {
	HandlerMap[urlName] = controllerFunc
}

type Res struct {
	Code     string
	Error    interface{}
	Response interface{}
}
