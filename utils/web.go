package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func Encode[T any](w http.ResponseWriter, r *http.Request, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

func Decode[T any](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}
	return v, nil
}

func GetIdFromPath(r *http.Request) (int, error) {
	id := r.PathValue("id")
	if id == "" {
		return 0, fmt.Errorf("empty id")
	}
	return strconv.Atoi(id)
}
