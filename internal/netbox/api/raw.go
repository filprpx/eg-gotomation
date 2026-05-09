package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/filprpx/eg-gotomation/internal/restapi"
)

type RawAPI struct {
	doer restapi.Doer
}

func NewRawAPI(d restapi.Doer) *RawAPI {
	return &RawAPI{doer: d}
}

func (r *RawAPI) List(ctx context.Context, path string, filters url.Values) ([]map[string]any, error) {
	res, err := r.doer.Do(ctx, restapi.Request{Method: "GET", Path: path, Query: filters})
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if restapi.IsError(res) {
		return nil, restapi.APIError(res)
	}

	var result struct {
		Results []map[string]any `json:"results"`
	}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Results, nil
}

func (r *RawAPI) Get(ctx context.Context, path string, id int) (map[string]any, error) {
	if id == 0 {
		return nil, fmt.Errorf("id can't be zero")
	}

	res, err := r.doer.Do(ctx, restapi.Request{Method: "GET", Path: fmt.Sprintf("%s%d/", path, id)})
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if restapi.IsError(res) {
		return nil, restapi.APIError(res)
	}

	var result map[string]any
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *RawAPI) Create(ctx context.Context, path string, data map[string]any) (map[string]any, error) {
	res, err := r.doer.Do(ctx, restapi.Request{Method: "POST", Path: path, Body: data})
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if restapi.IsError(res) {
		return nil, restapi.APIError(res)
	}

	var result map[string]any
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *RawAPI) Update(ctx context.Context, path string, id int, data map[string]any) (map[string]any, error) {
	if id == 0 {
		return nil, fmt.Errorf("id can't be zero")
	}

	res, err := r.doer.Do(ctx, restapi.Request{Method: "PATCH", Path: fmt.Sprintf("%s%d/", path, id), Body: data})
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if restapi.IsError(res) {
		return nil, restapi.APIError(res)
	}

	var result map[string]any
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *RawAPI) Delete(ctx context.Context, path string, id int) error {
	if id == 0 {
		return fmt.Errorf("id can't be zero")
	}

	res, err := r.doer.Do(ctx, restapi.Request{Method: "DELETE", Path: fmt.Sprintf("%s%d/", path, id)})
	if err != nil {
		return err
	}

	if restapi.IsError(res) {
		return restapi.APIError(res)
	}

	return nil
}
