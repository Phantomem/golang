package helpers

import "example/http/types"

func ResolveSuccessHttpJsonResponse[T any](data T) types.SuccessHttpJsonResponse[string, T] {
	return types.SuccessHttpJsonResponse[string, T]{
		Status: "success",
		Data:   data,
	}
}

func ResolveErrorHttpJsonResponse(err error) types.ErrorHttpJsonResponse[string] {
	return types.ErrorHttpJsonResponse[string]{
		Status: "error",
		Error:  err,
	}
}
