package netbox

import ()

type ApiListResponse[T any] struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []T    `json:"results"`
}

type ApiBaseFields struct {
	Id          int    `json:"id"`
	Url         string `json:"url"`
	Display_url string `json:"display_url"`
	Display     string `json:"display"`
}
