package netbox

import ()

type ApiListResponse[T any] struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []T    `json:"results"`
}

type ApiBaseFields struct {
	id          int
	url         string
	display_url string
	display     string
}
