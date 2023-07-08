package schema

type SuccessResponse[T any] struct {
	Status     string `json:"status"`
	StatusCode int    `json:"statusCode"`
	Data       T      `json:"data"`
}

func NewSuccessResponse[T any](Data T, Status string, StatusCode int) SuccessResponse[T] {
	return SuccessResponse[T]{
		Data:       Data,
		Status:     Status,
		StatusCode: StatusCode,
	}
}
