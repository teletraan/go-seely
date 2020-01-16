package seely

import "fmt"

func errFactory(message string) func(error) error {
	return func(err error) error {
		if err != nil {
			err = fmt.Errorf("%v: %w", message, err)
		}
		return err
	}
}

func stringInSlice(a string, s []string) bool {
	for _, b := range s {
		if a == b {
			return true
		}
	}
	return false
}
