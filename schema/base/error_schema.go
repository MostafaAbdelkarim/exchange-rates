package schema

type ErrorResponse struct {
	Status     string `json:"status"`
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
}
