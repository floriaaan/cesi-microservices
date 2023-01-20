package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// SumService provides operations on strings.
type SumService interface {
	Sum(int64, int64) (int64, error)
}

// stringService is a concrete implementation of StringService
type sumService struct{}

func (sumService) Sum(num1 int64, num2 int64) (int64, error) {
	if num1 == 0 && num2 == 0 {
		return -1, errors.New("num1 and num2 are 0")
	}

	return num1 + num2, nil
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

// Endpoints are a primary abstraction in go-kit. An endpoint represents a single RPC (method in our service interface)
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

// Transports expose the service to the network. In this first example we utilize JSON over HTTP.
func main() {
	svc := sumService{}

	sumHandler := httptransport.NewServer(
		makeSumEndpoint(svc),
		decodeSumRequest,
		encodeResponse,
	)

	http.Handle("/sum", sumHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func decodeSumRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request sumRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
