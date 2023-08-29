package age

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func StringToDuration(val, rawUnit string) (*time.Duration, error) {
	units := map[string]string{
		"s":       "s",
		"second":  "s",
		"seconds": "s",
		"m":       "m",
		"minute":  "m",
		"minutes": "m",
		"h":       "h",
		"hour":    "h",
		"hours":   "h",
		"d":       "d",
		"day":     "d",
		"days":    "d",
	}

	unit := units[rawUnit]
	unitMultiplier := -1
	if unit == "d" {
		unitMultiplier = -24
		unit = "h"
	}

	value, err := strconv.Atoi(val)
	if err != nil {
		return nil, errors.Join(fmt.Errorf(
			"%s does not appear to be an integer",
			val,
		), err)
	}
	durationString := fmt.Sprintf("%d%s", value*unitMultiplier, unit)
	dur, err := time.ParseDuration(durationString)
	if err != nil {
		err = errors.Join(errors.New("cannot parse duration"), err)
	}
	return &dur, err
}
