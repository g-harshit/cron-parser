package stratergy

import (
	"errors"
	"strconv"
	"strings"

	"github.com/cron-parser/lib"
)

type simpleStatergy struct {
}

func NewSimpleServiceStatergy() *simpleStatergy {
	return &simpleStatergy{}
}

func (s *simpleStatergy) Parse(field string, startLimit, endLimit int) ([]int, error) {
	var targetList []int
	if field == "*" {
		for i := startLimit; i <= endLimit; i++ {
			targetList = append(targetList, i)
		}
	} else {
		data := strings.Split(field, ",")
		for _, val := range data {
			var err error
			data, err := parseRanges(val, startLimit, endLimit)
			if err != nil {
				return nil, err
			}
			targetList = append(targetList, data...)
		}
	}
	return targetList, nil
}

// parseRanges parses a string range (e.g., "1-5", "10", "*/2") into a list of integers.
func parseRanges(val string, startLimit, endLimit int) ([]int, error) {
	var tempData []int

	if len(val) <= 2 {
		value, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		if value >= startLimit && value <= endLimit {
			tempData = append(tempData, value)
		}
	} else if strings.Contains(val, "-") {
		data := strings.Split(val, "-")
		if len(data) != 2 {
			return nil, errors.New("cannot send more than start range and end range")
		}

		startVal, err := strconv.Atoi(data[0])
		if err != nil {
			return nil, err
		}
		endVal, err := strconv.Atoi(data[1])
		if err != nil {
			return nil, err
		}

		startVal = lib.Max(startVal, startLimit)
		endVal = lib.Min(endVal, endLimit)

		if startVal > endVal {
			return nil, errors.New("invalid cron expression")
		}

		for i := startVal; i <= endVal; i++ {
			tempData = append(tempData, i)
		}
	} else if strings.Contains(val, "/") {
		data := strings.Split(val, "/")
		if len(data) != 2 {
			return nil, errors.New("invalid cron expression")
		}

		startVal := startLimit
		if data[0] != "*" {
			var err error
			startVal, err = strconv.Atoi(data[0])
			if err != nil {
				return nil, err
			}
			if startVal < startLimit {
				return nil, errors.New("start value is less that start limit")
			}
		}

		stepUp, err := strconv.Atoi(data[1])
		if err != nil {
			return nil, err
		}

		for i := startVal; i <= endLimit; i += stepUp {
			tempData = append(tempData, i)
		}
	} else {
		return nil, errors.New("invalid cron expression")
	}
	return tempData, nil

}
