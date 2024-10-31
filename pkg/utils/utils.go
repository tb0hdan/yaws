package utils

import (
	"fmt"

	"yaws/internal/server/api"
)

func Index[T any](arr []T, f func(T) bool) int {
	for i, item := range arr {
		if f(item) {
			return i
		}
	}

	return -1
}

func Ptr[T any](t T) *T {
	return &t
}

func ToAPIError[T any](t T) api.Error {
	return api.Error{
		Error: Ptr(fmt.Sprint(t)),
	}
}
