package age

import (
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
		return nil, fmt.Errorf("%s does not appear to be an integer: %w", val, err)
	}
	durationString := fmt.Sprintf("%d%s", value*unitMultiplier, unit)
	dur, err := time.ParseDuration(durationString)
	if err != nil {
		err = fmt.Errorf("cannot parse duration: %w", err)
	}
	return &dur, err
}
