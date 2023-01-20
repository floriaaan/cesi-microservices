package main

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   SumService
}

func (mw loggingMiddleware) Sum(num1 int64, num2 int64) (output int64, err error) {
	inputs := []int64{num1, num2}
	inputs_as_string := fmt.Sprintf("%v", inputs)

	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "Sum",
			// log inputs as string
			"inputs", inputs_as_string,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Sum(num1, num2)
	return
}
