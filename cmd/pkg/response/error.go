package response

func Error(statusCode int) string {
	switch statusCode {
	case 400:
		return "Bad Request"
	case 401:
		return "Unauthorised Request"
	case 404:
		return "Not Found"
	case 500:
		return "Internal Server error"
	case 503:
		return "Service Unavailable"
	default:
		return "Something went wrong"
	}
}
