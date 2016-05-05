package highbrow

import "time"

func Try(times int, fn func() error) error {
	var errors []error

	for i := 0; i < times; i++ {
		err := fn()
		if err != nil {
			errors = append(errors, err)
			backoff := time.Duration(400*(i+1)) * time.Millisecond
			time.Sleep(backoff)
			continue
		}
		return nil
	}

	return &RetryError{"Giving up on Network failures", errors}
}
