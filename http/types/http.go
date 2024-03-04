package types

type SuccessHttpJsonResponse[S string, T any] struct {
	Status S `json:"status"`
	Data   T `json:"data,omitempty"`
}

type ErrorHttpJsonResponse[S string] struct {
	Status S     `json:"status"`
	Error  error `json:"error,omitempty"`
}
