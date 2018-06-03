package elasticex

import (
	"errors"
	"strconv"
)

// ConvertIntervalToSeconds convert interval to seconds
func ConvertIntervalToSeconds(interval string) (int, error) {
	var unitSeconds int
	unit := interval[len(interval)-1:]
	if unit == "h" {
		unitSeconds = 3600
	} else if unit == "d" {
		unitSeconds = 86400
	} else if unit == "M" {
		unitSeconds = 2592000
	} else {
		return 0, errors.New("Failed to convert interval")
	}
	number, err := strconv.Atoi(interval[:len(interval)-1])
	if err != nil {
		return 0, err
	}

	return number * unitSeconds, nil
}
