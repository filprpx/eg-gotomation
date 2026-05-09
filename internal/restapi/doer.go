package restapi

import (
	"context"
	"net/http"
	"net/url"
)

type Request struct {
	Method string
	Path   string
	Query  url.Values
	Body   any
}

type Doer interface {
	Do(ctx context.Context, req Request) (*http.Response, error)
}
