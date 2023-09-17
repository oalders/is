// package compare compares versions
package compare

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"slices"
	"strconv"

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

func Floats(ctx *types.Context, operator, g, w string) error {
	got, err := strconv.ParseFloat(g, 32)
	if err != nil {
		return errors.Join(errors.New("wanted result must be a float"), err)
	}
	want, err := strconv.ParseFloat(w, 32)
	if err != nil {
		return errors.Join(errors.New("command output is not a float"), err)
	}

	if ctx.Debug {
		log.Printf("compare floats %f %s %f", got, operator, want)
	}
	IntegersOrFloats(ctx, operator, got, want)
	return nil
}

func Integers(ctx *types.Context, operator, g, w string) error {
	got, err := strconv.Atoi(g)
	if err != nil {
		return errors.Join(errors.New("wanted result must be an integer"), err)
	}
	want, err := strconv.Atoi(w)
	if err != nil {
		return errors.Join(errors.New("command output is not an integer"), err)
	}

	if ctx.Debug {
		log.Printf("compare integers %d %s %d", got, operator, want)
	}
	IntegersOrFloats(ctx, operator, got, want)
	return nil
}

func Versions(ctx *types.Context, operator, gotString, wantString string) error {
	var success bool
	switch operator {
	case ops.Like, ops.Unlike:
		return Strings(ctx, operator, gotString, wantString)
	}

	maybeDebug(ctx, "versions", operator, gotString, wantString)

	got, err := version.NewVersion(gotString)
	if err != nil {
		return err
	}
	want, err := version.NewVersion(wantString)
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
	case ops.Ne:
		success = got != want
	case ops.Like, ops.Unlike:
		success, err = regexp.MatchString(want, got)
	}

	if err != nil {
		ctx.Success = false
		comparison := fmt.Sprintf(`compare strings "%s" %s "%s"`, got, operator, want)
		return errors.Join(errors.New(comparison), err)
	}
	if operator == ops.Unlike {
		success = !success
	}
	ctx.Success = success
	return nil
}

func Optimistic(ctx *types.Context, operator, got, want string) error {
	stringy := []string{ops.Eq, ops.Ne, ops.Like, ops.Unlike}
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
	err := Versions(ctx, operator, got, want)
	if err != nil && ctx.Debug {
		log.Printf("cannot compare versions: %s", err)
	}
	return nil
}

func maybeDebug(ctx *types.Context, comparisonType, operator, got, want string) {
	if !ctx.Debug {
		return
	}

	log.Printf(`compare %s: "%s" %s "%s"\n`, comparisonType, got, operator, want)
}
