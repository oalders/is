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

func IntegersOrFloats[T Number](ctx *types.Context, operator string, got, want T) (bool, error) {
	if ctx.Debug {
		log.Printf("Evaluating %v %s %v\n", got, operator, want)
	}

	switch operator {
	case ops.Eq:
		return got == want, nil
	case ops.Ne:
		return got != want, nil
	case ops.Gt:
		return got > want, nil
	case ops.Gte:
		return got >= want, nil
	case ops.Lt:
		return got < want, nil
	case ops.Lte:
		return got <= want, nil
	default:
		return false, fmt.Errorf("unsupported operator: %s", operator)
	}
}

func Floats(ctx *types.Context, operator, g, w string) (bool, error) { //nolint:varnamelen
	if operator == ops.In {
		wantList, err := want2List(w)
		if err != nil {
			return false, err
		}
		for _, v := range wantList {
			success, err := Floats(ctx, ops.Eq, g, v)
			if err != nil {
				return false, err
			}
			if success {
				return true, nil
			}
		}
		return false, nil
	}
	got, err := strconv.ParseFloat(g, 32)
	if err != nil {
		return false, fmt.Errorf("wanted result must be a float: %w", err)
	}
	want, err := strconv.ParseFloat(w, 32)
	if err != nil {
		return false, fmt.Errorf("command output is not a float: %w", err)
	}

	if ctx.Debug {
		log.Printf("compare floats %f %s %f", got, operator, want)
	}
	success, err := IntegersOrFloats(ctx, operator, got, want)
	if err != nil {
		return false, err
	}
	return success, nil
}

func Integers(ctx *types.Context, operator, g, w string) (bool, error) { //nolint:varnamelen
	if operator == ops.In {
		wantList, err := want2List(w)
		if err != nil {
			return false, err
		}

		for _, v := range wantList {
			success, err := Integers(ctx, ops.Eq, g, v)
			if err != nil {
				return false, err
			}
			if success {
				return true, nil
			}
		}
		return false, nil
	}
	got, err := strconv.Atoi(g)
	if err != nil {
		return false, fmt.Errorf("wanted result must be an integer: %w", err)
	}
	want, err := strconv.Atoi(w)
	if err != nil {
		return false, fmt.Errorf("command output is not an integer: %w", err)
	}

	if ctx.Debug {
		log.Printf("compare integers %d %s %d", got, operator, want)
	}
	success, err := IntegersOrFloats(ctx, operator, got, want)
	if err != nil {
		return false, err
	}
	return success, nil
}

func VersionSegment(
	ctx *types.Context,
	operator, gotStr, wantStr string,
	segment uint,
) (bool, error) {
	if operator == ops.In {
		wantList, err := want2List(wantStr)
		if err != nil {
			return false, err
		}
		for _, v := range wantList {
			success, err := VersionSegment(ctx, ops.Eq, gotStr, v, segment)
			if err != nil {
				return false, err
			}
			if success {
				return true, nil
			}
		}
		return false, nil
	}
	got, err := version.NewVersion(gotStr)
	if err != nil {
		return false, fmt.Errorf("parse version from output: %w", err)
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
) (bool, error) {
	maybeDebug(ctx, "versions", operator, gotStr, wantStr)

	switch operator {
	case ops.In:
		wantList, err := want2List(wantStr)
		if err != nil {
			return false, err
		}
		for _, v := range wantList {
			success, err := Versions(ctx, ops.Eq, gotStr, v)
			if err != nil {
				return false, err
			}
			if success {
				return true, nil
			}
		}
		return false, nil
	case ops.Like, ops.Unlike:
		return Strings(ctx, operator, gotStr, wantStr)
	}

	got, err := version.NewVersion(gotStr)
	if err != nil {
		return false, err
	}
	want, err := version.NewVersion(wantStr)
	if err != nil {
		return false, err
	}

	switch operator {
	case ops.Eq:
		return got.Equal(want), nil
	case ops.Ne:
		return got.Compare(want) != 0, nil
	case ops.Lt:
		return got.LessThan(want), nil
	case ops.Lte:
		return got.Compare(want) <= 0, nil
	case ops.Gt:
		return got.GreaterThan(want), nil
	case ops.Gte:
		return got.Compare(want) >= 0, nil
	default:
		return false, fmt.Errorf("unsupported operator: %s", operator)
	}
}

//nolint:cyclop
func Strings(ctx *types.Context, operator, got, want string) (bool, error) {
	maybeDebug(ctx, "strings", operator, got, want)

	switch operator {
	case ops.Eq:
		return got == want, nil
	case ops.In:
		wantList, err := want2List(want)
		if err != nil {
			return false, err
		}
		return slices.Contains(wantList, got), nil
	case ops.Ne:
		matched, err := regexp.MatchString(want, got)
		if err != nil {
			return false, fmt.Errorf(`compare strings "%s" %s "%s"`, got, operator, want)
		}
		ctx.Success = matched
		if operator == ops.Unlike {
			ctx.Success = !matched
		}
		return got != want, nil
	case ops.Like:
		success, err := regexp.MatchString(want, got)
		if err != nil {
			return false, fmt.Errorf(`compare strings "%s" %s "%s"`, got, operator, want)
		}
		return success, nil
	case ops.Unlike:
		success, err := regexp.MatchString(want, got)
		if err != nil {
			return false, fmt.Errorf(`compare strings "%s" %s "%s"`, got, operator, want)
		}
		return !success, nil
	default:
		return false, fmt.Errorf("unsupported operator: %s", operator)
	}
}

//nolint:cyclop
func Optimistic(ctx *types.Context, operator, got, want string) bool {
	stringy := []string{ops.Eq, ops.In, ops.Ne, ops.Like, ops.Unlike}
	reg := []string{ops.Like, ops.Unlike}
	if slices.Contains(stringy, operator) {
		success, err := Strings(ctx, operator, got, want)
		if err != nil && ctx.Debug {
			log.Printf("cannot compare strings: %s", err)
		}
		if success || slices.Contains(reg, operator) {
			return success
		}
	}

	// We are being optimistic here and we can't know if the intention was a
	// string or a numeric comparison, so we'll suppress the error message
	// unless debugging is enabled.

	{
		success, err := Integers(ctx, operator, got, want)
		if err != nil && ctx.Debug {
			log.Printf("cannot compare integers: %s", err)
		}
		if success {
			return true
		}
	}

	{
		success, err := Floats(ctx, operator, got, want)
		if err != nil && ctx.Debug {
			log.Printf("cannot compare floats: %s", err)
		}
		if success {
			return true
		}
	}

	success, err := Versions(ctx, operator, got, want)
	if err != nil && ctx.Debug {
		log.Printf("cannot compare versions: %s", err)
	}
	return success
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
