package mlr

import (
	"context"
	"net/http"
)

type Service struct {
	client    *http.Client
	UserAgent string
	Matches   *MatchesService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return UserAgent
	}
	return UserAgent + " " + s.UserAgent
}

func NewService() (*Service, error) {
	s := &Service{client: &http.Client{}}
	s.Matches = NewMatchesService(s)
	return s, nil
}

func SendRequest(client *http.Client, req *http.Request) (*http.Response, error) {
	ctx := context.Background()
	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		select {
		case <-ctx.Done():
			err = ctx.Err()
		default:
		}
	}
	return resp, err
}
