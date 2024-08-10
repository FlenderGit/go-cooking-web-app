package utils

import "net/http"

type Middleware func(http.Handler) http.HandlerFunc

func Chain(f http.HandlerFunc, fns ...Middleware) http.HandlerFunc {
	for _, fn := range fns {
		f = fn(f)
	}
	return f
}
