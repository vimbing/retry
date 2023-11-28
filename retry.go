package retry

import (
	"errors"
	"time"
)

func (r Retrier) Retry(retryFunc RetryFunc) error {
	var latestError error

	if r.Max <= 0 {
		r.Max = 1
	}

	for i := 0; i < r.Max; i++ {
		if i != 0 {
			time.Sleep(r.Delay)
		}

		latestError = retryFunc()

		if latestError == nil {
			return nil
		}

		if r.Logger.LogFunc != nil {
			r.Logger.LogFunc(r.Logger.Format, latestError.Error())
		}

		for _, err := range r.FatalErrors {
			if errors.Is(latestError, err) {
				return latestError
			}
		}
	}

	return latestError
}
