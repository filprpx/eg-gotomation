package models

type ApiResource interface {
	GetId() int
	MapToWrite() ApiResourceWrite
}

type ApiResourceWrite interface{}

type ApiListResponse[T any] struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []T    `json:"results"`
}

type ApiBaseFields struct {
	Id          int    `json:"id,omitempty"`
	Url         string `json:"url,omitempty"`
	DisplayUrl  string `json:"display_url,omitempty"`
	Display     string `json:"display,omitempty"`
	Name        string `json:"name,omitempty"`
	Slug        string `json:"slug,omitempty"`
	Description string `json:"description,omitempty"`
}

func (a ApiBaseFields) GetId() int {
	return a.Id
}

func (a ApiBaseFields) MapToWrite() *ApiWriteBaseFields {
	return &ApiWriteBaseFields{
		Id:          a.Id,
		Name:        a.Name,
		Slug:        a.Slug,
		Description: a.Description,
	}
}

type ApiWriteBaseFields struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Slug        string `json:"slug,omitempty"`
	Description string `json:"description,omitempty"`
}

type Choice struct {
	Value string
	Label string
}

// Idea I had, but I am not sure it is possible to achieve cleanly
// type IApiEntity[T any] interface {
// 	List(ctx context.Context) (*[]T, error)
// 	Get(ctx context.Context, id int) (*T, error)
// 	Create(ctx context.Context, e *T) (*T, error)
// 	Update(ctx context.Context, e *T) (bool, error)
// 	Delete(ctx context.Context, id int) (bool, error)
// }

// type ApiModelDefaultActions[T any] struct {
// 	ApiEntity IApiEntity[T]
// }

// func (a *ApiModelDefaultActions) Update(ctx context.Context) (bool, error) {
// 	return a.Update(ctx, a)
// }
//
// func (a *ApiModelDefaultActions) Delete(ctx context.Context) (bool, error) {
// 	return a.Delete(ctx, a.Id)
// }
