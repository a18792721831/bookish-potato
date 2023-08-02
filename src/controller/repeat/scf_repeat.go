package repeat

import (
	"context"
	. "scf/src/controller"
)

func newRepeatController() Controller {
	return &repeatController{}
}

type repeatReq struct {
	Msg string `json:"msg"`
}

type repeatRes struct {
	Msg string `json:"msg"`
}

type repeatController struct {
}

func init() {
	Register("/repeat", newRepeatController)
}

func (this *repeatController) Request() interface{} {
	return &repeatReq{}
}
func (this *repeatController) Response() interface{} {
	return &repeatRes{}
}
func (this *repeatController) Process(ctx context.Context, req interface{}) (interface{}, error) {
	request := req.(repeatReq)
	response := &repeatRes{}
	response.Msg = request.Msg
	return response, nil
}
