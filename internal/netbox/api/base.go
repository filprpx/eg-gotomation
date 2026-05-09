// Package api is used to configure the comunication with the netbox api
// base.go is the common implementation to work with the api actions
// the direct resources (e.g. devices) only specificy the desired endpoint, but at any point it is possible to
// override the generic functions into specific ones for each resource.
package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/filprpx/eg-gotomation/internal/netbox/models"
	"github.com/filprpx/eg-gotomation/internal/restapi"
)

type BaseAPI[T models.APIResource] struct {
	doer restapi.Doer
	path string
}

func NewBaseAPI[T models.APIResource](d restapi.Doer, path string) *BaseAPI[T] {
	return &BaseAPI[T]{
		doer: d,
		path: path,
	}
}

func (a *BaseAPI[T]) List(ctx context.Context, filters url.Values) ([]T, error) {
	res, err := a.doer.Do(ctx, restapi.Request{
		Method: "GET",
		Path:   a.path, Query: filters,
	})
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if restapi.IsError(res) {
		return nil, restapi.APIError(res)
	}

	var result models.APIListResponse[T]
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Results, nil
}

func (a *BaseAPI[T]) Get(ctx context.Context, id int) (*T, error) {
	if id == 0 {
		return nil, fmt.Errorf("id can't be zero")
	}

	res, err := a.doer.Do(ctx, restapi.Request{
		Method: "GET",
		Path:   fmt.Sprintf("%s%d/", a.path, id),
	})
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if restapi.IsError(res) {
		return nil, restapi.APIError(res)
	}

	var result T
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (a *BaseAPI[T]) Create(ctx context.Context, obj *T) (*T, error) {
	if obj == nil {
		return nil, fmt.Errorf("entity can't be nil")
	}

	entity := (*obj)

	res, err := a.doer.Do(ctx, restapi.Request{
		Method: "POST",
		Path:   a.path,
		Body:   entity.MapToWrite(),
	})
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if restapi.IsError(res) {
		return nil, restapi.APIError(res)
	}

	var result T
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (a *BaseAPI[T]) Update(ctx context.Context, obj *T) (*T, error) {
	if obj == nil {
		return nil, fmt.Errorf("obj can't be nil")
	}

	entity := (*obj)

	if entity.GetId() == 0 {
		return nil, fmt.Errorf("obj.Id can't be zero")
	}

	res, err := a.doer.Do(ctx, restapi.Request{
		Method: "PATCH",
		Path:   fmt.Sprintf("%s%d/", a.path, entity.GetId()),
		Body:   entity.MapToWrite(),
	})
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if restapi.IsError(res) {
		return nil, restapi.APIError(res)
	}

	var result T
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (a *BaseAPI[T]) Delete(ctx context.Context, obj *T) error {
	if obj == nil {
		return fmt.Errorf("obj can't be nil")
	}

	entity := (*obj)

	if entity.GetId() == 0 {
		return fmt.Errorf("obj.Id can't be zero")
	}

	res, err := a.doer.Do(ctx, restapi.Request{
		Method: "DELETE",
		Path:   fmt.Sprintf("%s%d/", a.path, entity.GetId()),
	})
	if err != nil {
		return err
	}

	if restapi.IsError(res) {
		return restapi.APIError(res)
	}

	return nil
}
