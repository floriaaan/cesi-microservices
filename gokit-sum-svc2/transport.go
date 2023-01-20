package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-kit/kit/endpoint"
)

func makeSumEndpoint(svc SumService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(sumRequest)
		num1, err := strconv.ParseInt(req.Num1, 10, 64)
		if err != nil {
			return sumResponse{-1, err.Error()}, nil
		}
		num2, err := strconv.ParseInt(req.Num2, 10, 64)
		if err != nil {
			return sumResponse{-1, err.Error()}, nil
		}

		sum, err := svc.Sum(num1, num2)
		if err != nil {
			return sumResponse{-1, err.Error()}, nil
		}
		return sumResponse{sum, ""}, nil
	}
}

func decodeSumRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request sumRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	// if err in response is not nil, return error status code
	if f, ok := response.(sumResponse); ok {
		if f.Err != "" {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	return json.NewEncoder(w).Encode(response)
}

// For each method, we define request and response structs
type sumRequest struct {
	Num1 string `json:"num1"`
	Num2 string `json:"num2"`
}

type sumResponse struct {
	Sum int64  `json:"sum"`
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}
