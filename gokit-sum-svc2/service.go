package main

import (
	"errors"
)

type SumService interface {
	Sum(int64, int64) (int64, error)
}

type sumService struct{}

func (sumService) Sum(num1 int64, num2 int64) (int64, error) {
	if num1 == 0 && num2 == 0 {
		return -1, errors.New("num1 and num2 are 0")
	}

	return num1 + num2, nil
}

var ErrEmpty = errors.New("empty input")
