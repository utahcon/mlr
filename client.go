package mlr

import "net/http"

const (
	Version   = "0.0.1"
	UserAgent = "mlr-api-go-client/" + Version
)

func NewClient() (*http.Client, error) {
	return &http.Client{}, nil
}
