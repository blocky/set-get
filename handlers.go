package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/redis/go-redis/v9"
	"gopkg.in/go-playground/validator.v9"
)

func ValidateStruct(data interface{}) error {
	v := validator.New()
	err := v.Struct(data)
	if err != nil {
		var validationErrs validator.ValidationErrors
		if errors.As(err, &validationErrs) {
			var missingFields []string
			for _, e := range validationErrs {
				missingFields = append(missingFields, e.Field())
			}
			return fmt.Errorf("missing required fields: %s", strings.Join(missingFields, ", "))
		}
		return fmt.Errorf("validating struct fields: %w", err)
	}
	return nil
}

type SetValue struct {
	Key   string `json:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
}

// handlerSetValue sets a key-value pair in Redis.
func handlerSetValue(rdb *redis.Client) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := io.ReadAll(req.Body)
		defer func() { _ = req.Body.Close() }()
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}

		var data SetValue
		err = json.Unmarshal(body, &data)
		if err != nil {
			http.Error(w, "Error unmarshalling request body", http.StatusBadRequest)
			return
		}

		err = ValidateStruct(data)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error validating request: %s", err), http.StatusBadRequest)
			return
		}

		ctx := req.Context()
		err = rdb.Set(ctx, data.Key, data.Value, 0).Err()
		if err != nil {
			http.Error(w, "Error setting value", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

type GetValue struct {
	Key string `json:"key" validate:"required"`
}

// handlerGetValue looks up a given key in Redis and writes the return value to the response.
func handlerGetValue(rdb *redis.Client) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		var data GetValue
		body, err := io.ReadAll(req.Body)
		defer func() { _ = req.Body.Close() }()
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(body, &data)
		if err != nil {
			http.Error(w, "Error unmarshalling request body", http.StatusBadRequest)
			return
		}

		err = ValidateStruct(data)
		if err != nil {
			http.Error(w, "Missing required field", http.StatusBadRequest)
			return
		}

		ctx := req.Context()
		val, err := rdb.Get(ctx, data.Key).Result()
		if err != nil {
			if errors.Is(err, redis.Nil) {
				http.Error(w, "Key not found", http.StatusNotFound)
			} else {
				http.Error(w, "Error getting value", http.StatusInternalServerError)
			}
			return
		}

		w.Write([]byte(val))
	}
}
