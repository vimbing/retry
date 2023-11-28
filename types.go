package retry

import "time"

type Logger struct {
	LogFunc func(string, ...any)
	Format  string
}

type Retrier struct {
	// A maximum number of retry
	Max int
	// Delay which should be executed between trys
	Delay time.Duration
	// An error types which should be returned and not retryied
	FatalErrors []error
	Logger      Logger

	IncreaseRetryChan chan int
	DecreaseRetryChan chan int
	SetRetryChan      chan int
}

type RetryFunc func() error
