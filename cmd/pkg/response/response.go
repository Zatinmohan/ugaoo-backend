package response

type Response[T any] struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Result     T      `json:"result"`
}

func GetFinalResponse[T any](statusCode int, message string, result T) Response[T] {
	return Response[T]{
		StatusCode: statusCode,
		Message:    message,
		Result:     result,
	}
}
