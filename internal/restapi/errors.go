package restapi

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func IsSuccess(res *http.Response) bool {
	return res.StatusCode > 199 && res.StatusCode < 300
}

func IsError(res *http.Response) bool {
	return res.StatusCode > 399
}

// TODO: give this a better name, make sure it has the right level of abstraction on how to deal with the api errors
func ApiError(r *http.Response) error {
	url := r.Request.URL.String()

	limitedReader := io.LimitReader(r.Body, 1024)
	body, err := io.ReadAll(limitedReader)
	if err != nil {
		return fmt.Errorf("api error %s: %d (failed to read body: %v)", url, r.StatusCode, err)
	}

	bodyStr := string(body)
	if len(bodyStr) == 1024 {
		bodyStr += "...(truncated)"
	}

	if strings.Contains(bodyStr, "<html") || strings.Contains(bodyStr, "<!DOCTYPE") {
		return fmt.Errorf("api error %s: %d: (received HTML instead of JSON, check endpoint URL)", url, r.StatusCode)
	}

	return fmt.Errorf("api error %s: %d, %s", url, r.StatusCode, bodyStr)
}
