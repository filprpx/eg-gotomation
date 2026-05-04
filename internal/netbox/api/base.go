// Package api is used to configure the comunication with the netbox api
// base.go is the common implementation to work with the api actions
// the direct resources (e.g. devices) only specificy the desired endpoint, but at any point it is possible to
// override the generic functions into specific ones for each resource.
package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

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

func (a *BaseAPI[T]) List(ctx context.Context) ([]T, error) {
	res, err := a.doer.Get(ctx, a.path)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if restapi.IsError(res) {
		return nil, restapi.APIError(res)
	}

	var result models.APIListResponse[T]
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Results, nil
}

func (a *BaseAPI[T]) Get(ctx context.Context, id int) (*T, error) {
	if id == 0 {
		return nil, fmt.Errorf("id can't be zero")
	}

	res, err := a.doer.Get(ctx, fmt.Sprintf("%s%d/", a.path, id))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if restapi.IsError(res) {
		return nil, restapi.APIError(res)
	}

	var result T
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (a *BaseAPI[T]) Create(ctx context.Context, obj *T) (*T, error) {
	if obj == nil {
		return nil, fmt.Errorf("entity.ct can't be nil")
	}

	entity := (*obj)

	payload, err := json.Marshal(entity.MapToWrite())
	if err != nil {
		return nil, fmt.Errorf("failed to marshal: %w", err)
	}

	res, err := a.doer.Post(
		ctx,
		a.path,
		bytes.NewReader(payload),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if restapi.IsError(res) {
		return nil, restapi.APIError(res)
	}

	var result T
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (a *BaseAPI[T]) Update(ctx context.Context, obj *T) (bool, error) {
	if obj == nil {
		return false, fmt.Errorf("obj can't be nil")
	}

	entity := (*obj)

	if entity.GetId() == 0 {
		return false, fmt.Errorf("obj.Id can't be zero")
	}

	payload, err := json.Marshal(entity.MapToWrite())
	if err != nil {
		return false, fmt.Errorf("failed to marshal: %w", err)
	}

	res, err := a.doer.Patch(
		ctx,
		fmt.Sprintf("%s%d/", a.path, entity.GetId()),
		bytes.NewReader(payload),
	)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	if restapi.IsError(res) {
		return false, restapi.APIError(res)
	}

	return true, nil
}

func (a *BaseAPI[T]) Delete(ctx context.Context, obj *T) (bool, error) {
	if obj == nil {
		return false, fmt.Errorf("obj can't be nil")
	}

	entity := (*obj)

	if entity.GetId() == 0 {
		return false, fmt.Errorf("obj.Id can't be zero")
	}

	res, err := a.doer.Delete(
		ctx,
		fmt.Sprintf("%s%d/", a.path, entity.GetId()),
	)
	if err != nil {
		return false, err
	}

	if restapi.IsError(res) {
		return false, restapi.APIError(res)
	}

	return true, nil
}
