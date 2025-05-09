package battery

import (
	"errors"
	"fmt"
	"log"
	"math"
	"reflect"
	"strings"

	"github.com/distatus/battery"
	"github.com/oalders/is/types"
)

//nolint:tagliatelle
type Battery struct {
	State            string  `json:"state"`
	BatteryNumber    int     `json:"battery-number"`
	Count            int     `json:"count"`
	ChargeRate       float64 `json:"charge-rate"`
	CurrentCapacity  float64 `json:"current-capacity"`
	CurrentCharge    float64 `json:"current-charge"`
	DesignCapacity   float64 `json:"design-capacity"`
	DesignVoltage    float64 `json:"design-voltage"`
	LastFullCapacity float64 `json:"last-full-capacity"`
	Voltage          float64 `json:"voltage"`
}

var attributeMap = map[string]string{ //nolint: gochecknoglobals
	"charge-rate":        "ChargeRate",
	"count":              "Count",
	"current-capacity":   "CurrentCapacity",
	"current-charge":     "CurrentCharge",
	"design-capacity":    "DesignCapacity",
	"design-voltage":     "DesignVoltage",
	"last-full-capacity": "LastFullCapacity",
	"state":              "State",
	"voltage":            "Voltage",
}

func Get(ctx *types.Context, nth int) (*Battery, error) {
	if nth == 0 {
		return nil, errors.New("use --nth 1 to get the first battery")
	}
	batteries, err := battery.GetAll()
	if err != nil {
		return nil, fmt.Errorf("get battery info: %w", err)
	}
	count := len(batteries)

	// All other attribute checks should generate an error message if no
	// batteries are found.
	// https://github.com/distatus/battery/issues/34
	//nolint:lll
	if count == 0 || (count == 1 && batteries[0].Current == 0 && batteries[0].Full == 0 && batteries[0].Design == 0) {
		return &Battery{Count: 0}, nil
	}
	if nth > count {
		return nil, fmt.Errorf(
			"battery %d requested, but only %d batteries found",
			nth,
			len(batteries),
		)
	}
	battery := batteries[nth-1]

	batt := Battery{
		BatteryNumber:    nth,
		Count:            len(batteries),
		ChargeRate:       battery.ChargeRate,
		CurrentCapacity:  battery.Current,
		CurrentCharge:    math.Round(battery.Current / battery.Full * 100),
		DesignCapacity:   battery.Design,
		DesignVoltage:    battery.DesignVoltage,
		LastFullCapacity: battery.Full,
		State:            strings.ToLower(battery.State.String()),
		Voltage:          battery.Voltage,
	}

	if ctx.Debug {
		log.Printf(
			"Run %q to see all available battery data\n",
			"is known summary battery",
		)
	}
	return &batt, nil
}

func GetAttrAsString(ctx *types.Context, attr string, round bool, nth int) (string, error) {
	fieldValue, err := GetAttr(ctx, attr, nth)
	if err != nil {
		return "", err
	}

	switch value := fieldValue.(type) {
	case float64:
		if round {
			return fmt.Sprintf("%d", int(math.Round(value))), nil
		}
		return fmt.Sprintf("%f", value), nil
	default:
		return fmt.Sprintf("%v", value), nil
	}
}

func GetAttr(ctx *types.Context, attr string, nth int) (any, error) {
	batt, err := Get(ctx, nth)
	if err != nil {
		return "", err
	}

	fieldName, ok := attributeMap[attr]
	if !ok {
		return "", fmt.Errorf("attr %s not recognized", attr)
	}

	return reflect.ValueOf(*batt).FieldByName(fieldName).Interface(), nil
}
