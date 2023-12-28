package e

const (
	SUCCESS = 200
	ERROR   = 500

	InvalidParams int = iota

	ErrorUserExist
	ErrorUserNotExist
	TokenGeneratedFail

	ErrorDataBase
)
