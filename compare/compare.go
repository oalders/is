// package compare compares versions
package compare

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/oalders/is/ops"
	"github.com/oalders/is/types"
	"github.com/oalders/is/version"
)

type Number interface {
	int | float32 | float64
}

func IntegersOrFloats[T Number](ctx *types.Context, operator string, got, want T) {
	switch operator {
	case ops.Eq:
		ctx.Success = got == want
	case ops.Ne:
		ctx.Success = got != want
	case ops.Gt:
		ctx.Success = got > want
	case ops.Gte:
		ctx.Success = got >= want
	case ops.Lt:
		ctx.Success = got < want
	case ops.Lte:
		ctx.Success = got <= want
	}
}

func Floats(ctx *types.Context, operator, g, w string) error { //nolint:varnamelen
	if operator == ops.In {
		wantList, err := want2List(w)
		if err != nil {
			return err
		}
		for _, v := range wantList {
			err := Floats(ctx, ops.Eq, g, v)
			if err != nil {
				return err
			}
			if ctx.Success {
				return nil
			}
		}
		return nil
	}
	got, err := strconv.ParseFloat(g, 32)
	if err != nil {
		return fmt.Errorf("wanted result must be a float: %w", err)
	}
	want, err := strconv.ParseFloat(w, 32)
	if err != nil {
		return fmt.Errorf("command output is not a float: %w", err)
	}

	if ctx.Debug {
		log.Printf("compare floats %f %s %f", got, operator, want)
	}
	IntegersOrFloats(ctx, operator, got, want)
	return nil
}

func Integers(ctx *types.Context, operator, g, w string) error { //nolint:varnamelen
	if operator == ops.In {
		wantList, err := want2List(w)
		if err != nil {
			return err
		}

		for _, v := range wantList {
			err := Integers(ctx, ops.Eq, g, v)
			if err != nil {
				return err
			}
			if ctx.Success {
				return nil
			}
		}
		return nil
	}
	got, err := strconv.Atoi(g)
	if err != nil {
		return fmt.Errorf("wanted result must be an integer: %w", err)
	}
	want, err := strconv.Atoi(w)
	if err != nil {
		return fmt.Errorf("command output is not an integer: %w", err)
	}

	if ctx.Debug {
		log.Printf("compare integers %d %s %d", got, operator, want)
	}
	IntegersOrFloats(ctx, operator, got, want)
	return nil
}

func VersionSegment(ctx *types.Context, operator, gotStr, wantStr string, segment uint) error {
	if operator == ops.In {
		wantList, err := want2List(wantStr)
		if err != nil {
			return err
		}
		for _, v := range wantList {
			err := VersionSegment(ctx, ops.Eq, gotStr, v, segment)
			if err != nil {
				return err
			}
			if ctx.Success {
				return nil
			}
		}
		return nil
	}
	got, err := version.NewVersion(gotStr)
	if err != nil {
		return fmt.Errorf("parse version from output: %w", err)
	}

	segments := got.Segments()
	gotSegment := segments[segment]

	switch operator {
	case ops.Like, ops.Unlike:
		return Strings(ctx, operator, fmt.Sprint(gotSegment), wantStr)
	}
	return Integers(ctx, operator, fmt.Sprint(gotSegment), wantStr)
}

func Versions( //nolint:cyclop
	ctx *types.Context,
	operator, gotStr, wantStr string,
) error {
	var success bool
	maybeDebug(ctx, "versions", operator, gotStr, wantStr)

	switch operator {
	case ops.In:
		wantList, err := want2List(wantStr)
		if err != nil {
			return err
		}
		for _, v := range wantList {
			err := Versions(ctx, ops.Eq, gotStr, v)
			if err != nil {
				return err
			}
			if ctx.Success {
				return nil
			}
		}
		return nil
	case ops.Like, ops.Unlike:
		return Strings(ctx, operator, gotStr, wantStr)
	}

	got, err := version.NewVersion(gotStr)
	if err != nil {
		return err
	}
	want, err := version.NewVersion(wantStr)
	if err != nil {
		return err
	}

	switch operator {
	case ops.Eq:
		success = got.Equal(want)
	case ops.Ne:
		success = got.Compare(want) != 0
	case ops.Lt:
		success = got.LessThan(want)
	case ops.Lte:
		success = got.Compare(want) <= 0
	case ops.Gt:
		success = got.GreaterThan(want)
	case ops.Gte:
		success = got.Compare(want) >= 0
	}

	ctx.Success = success
	return nil
}

func Strings(ctx *types.Context, operator, got, want string) error {
	var err error
	var success bool

	maybeDebug(ctx, "strings", operator, got, want)

	switch operator {
	case ops.Eq:
		success = got == want
	case ops.In:
		wantList, err := want2List(want)
		if err != nil {
			return err
		}
		success = slices.Contains(wantList, got)
	case ops.Ne:
		success = got != want
	case ops.Like, ops.Unlike:
		success, err = regexp.MatchString(want, got)
	}

	if err != nil {
		ctx.Success = false
		return fmt.Errorf(`compare strings "%s" %s "%s": %w`, got, operator, want, err)
	}
	if operator == ops.Unlike {
		success = !success
	}
	ctx.Success = success
	return nil
}

//nolint:cyclop
func Optimistic(ctx *types.Context, operator, got, want string) error {
	stringy := []string{ops.Eq, ops.In, ops.Ne, ops.Like, ops.Unlike}
	reg := []string{ops.Like, ops.Unlike}
	if slices.Contains(stringy, operator) {
		err := Strings(ctx, operator, got, want)
		if err != nil || ctx.Success || slices.Contains(reg, operator) {
			return err
		}
	}

	// We are being optimistic here and we can't know if the intention was a
	// string or a numeric comparison, so we'll suppress the error message
	// unless debugging is enabled.

	if err := Integers(ctx, operator, got, want); err == nil {
		return nil
	} else if ctx.Debug {
		log.Printf("cannot compare integers: %s", err)
	}

	if err := Floats(ctx, operator, got, want); err == nil {
		return nil
	} else if ctx.Debug {
		log.Printf("cannot compare floats: %s", err)
	}

	err := Versions(ctx, operator, got, want)
	if err != nil && ctx.Debug {
		log.Printf("cannot compare versions: %s", err)
	}
	return nil
}

func want2List(want string) ([]string, error) {
	wantList := strings.Split(want, ",")
	for i := range wantList {
		wantList[i] = strings.TrimSpace(wantList[i])
	}
	if len(wantList) > 100 {
		return []string{}, errors.New("\"in\" takes a maximum of 100 arguments")
	}
	return wantList, nil
}

func maybeDebug(ctx *types.Context, comparisonType, operator, got, want string) {
	if !ctx.Debug {
		return
	}

	log.Printf(`compare %s: "%s" %s "%s"\n`, comparisonType, got, operator, want)
}
