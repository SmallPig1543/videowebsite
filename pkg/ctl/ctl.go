package ctl

import (
	"net/http"
	"videoweb/pkg/e"
)

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

type DataList struct {
	Item  interface{} `json:"item"`
	Total int64       `json:"total"`
}

func RespList(items interface{}, total int64) Response {
	return Response{
		Status: 200,
		Data: DataList{
			Item:  items,
			Total: total,
		},
		Msg: "ok",
	}
}

func RespSuccess() *Response {
	status := e.SUCCESS
	r := &Response{
		Status: status,
		Data:   "操作成功",
		Msg:    e.GetMsg(status),
	}

	return r
}

func RespSuccessWithData(data interface{}) *Response {
	status := e.SUCCESS
	r := &Response{
		Status: status,
		Data:   data,
		Msg:    e.GetMsg(status),
	}
	return r
}

// RespError 错误返回
func RespError(err error, code int) *Response {
	return &Response{
		Status: http.StatusOK,
		Data:   nil,
		Msg:    e.GetMsg(code),
		Error:  err.Error(),
	}
}

func RespErrorWithData(data interface{}, err error, code int) *Response {
	return &Response{
		Status: http.StatusOK,
		Data:   data,
		Msg:    e.GetMsg(code),
		Error:  err.Error(),
	}
}
