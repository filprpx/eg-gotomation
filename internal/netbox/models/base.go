// Package models is as the common definitions for the resources in netbox api
package models

import (
	"fmt"
	"strconv"
	"time"
)

type APIResource interface {
	GetId() int
	MapToWrite() APIResourceWrite
}

type APIResourceWrite interface{}

type APIListResponse[T any] struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []T    `json:"results"`
}

type APIBaseFields struct {
	Id          int       `json:"id,omitempty"`
	Url         string    `json:"url,omitempty"`
	DisplayUrl  string    `json:"display_url,omitempty"`
	Display     string    `json:"display,omitempty"`
	Name        string    `json:"name,omitempty"`
	Slug        string    `json:"slug,omitempty"`
	Description string    `json:"description,omitempty"`
	Comments    string    `json:"comments,omitempty"`
	Created     time.Time `json:"created"`
	LastUpdated time.Time `json:"last_updated"`
}

func (a APIBaseFields) GetId() int {
	return a.Id
}

func (a APIBaseFields) MapToWrite() *APIWriteBaseFields {
	return &APIWriteBaseFields{
		Id:          a.Id,
		Name:        a.Name,
		Slug:        a.Slug,
		Description: a.Description,
	}
}

type APIWriteBaseFields struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Slug        string `json:"slug,omitempty"`
	Description string `json:"description,omitempty"`
}

type Choice struct {
	Value any
	Label string
}

func SafeGetId[T interface{ GetId() int }](obj *T) int {
	if obj == nil {
		return 0
	}
	return (*obj).GetId()
}

func safeChoiceValue(choice *Choice) string {
	if choice == nil {
		return ""
	}

	switch v := choice.Value.(type) {
	case string:
		return v
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case int:
		return strconv.Itoa(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	default:
		return fmt.Sprint(v)
	}
}

func safeChoiceIntValue(choice *Choice) int {
	if choice == nil {
		return 0
	}

	switch v := choice.Value.(type) {
	case float64:
		return int(v)
	case int:
		return v
	case int64:
		return int(v)
	case uint64:
		return int(v)
	case string:
		n, err := strconv.Atoi(v)
		if err != nil {
			return 0
		}
		return n
	default:
		return 0
	}
}
