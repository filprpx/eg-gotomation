package netbox

import ()

type ApiListResponse[T any] struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []T    `json:"results"`
}

type ApiBaseFields struct {
	Id          int    `json:"id,omitempty"`
	Url         string `json:"url,omitempty"`
	Display_url string `json:"display_url,omitempty"`
	Display     string `json:"display,omitempty"`
	Name        string `json:"name,omitempty"`
	Slug        string `json:"slug,omitempty"`
	Description string `json:"description,omitempty"`
}
