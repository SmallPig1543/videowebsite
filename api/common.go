package api

import (
	"encoding/json"
	"videoweb/pkg/ctl"
	"videoweb/pkg/e"
)

func ErrorResponse(err error) *ctl.Response {
	_, ok := err.(*json.UnmarshalTypeError)
	if ok {
		return ctl.RespError(err, e.InvalidParams)
	}
	return ctl.RespError(err, e.ERROR)
}
