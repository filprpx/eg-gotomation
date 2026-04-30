package restapi

import (
	"context"
	"io"
	"net/http"
)

type Doer interface {
	Get(ctx context.Context, path string) (*http.Response, error)
	Post(ctx context.Context, path string, body io.Reader) (*http.Response, error)
	Patch(ctx context.Context, path string, body io.Reader) (*http.Response, error)
	Delete(ctx context.Context, path string) (*http.Response, error)
}
