package compare_test

import (
	"fmt"
	"testing"

	"github.com/oalders/is/compare"
	"github.com/oalders/is/ops"
	"github.com/oalders/is/types"
	"github.com/stretchr/testify/assert"
)

type compareTest struct {
	Op      string
	Got     string
	Want    string
	Error   bool
	Success bool
	Debug   bool
}

func TestCompareVersions(t *testing.T) {
	t.Parallel()
	type test struct {
		Got     string
		Op      string
		Want    string
		Success bool
	}
	tests := []test{
		{"3.3", ops.Gt, "3.3", false},
		{"3.3", ops.Ne, "3.3", false},
		{"3.3", ops.Eq, "3.3", true},
		{"3.3", ops.Gte, "3.3", true},
		{"3.3", ops.Lte, "3.3", true},
		{"3.3", ops.Lt, "3.3", false},
		{"3.3", ops.Like, "3.3", true},
		{"3.3", ops.Unlike, "4", true},

		{"3.3a", ops.Gt, "3.3a", false},
		{"3.3a", ops.Ne, "3.3a", false},
		{"3.3a", ops.Eq, "3.3a", true},
		{"3.3a", ops.Gte, "3.3a", true},
		{"3.3a", ops.Lte, "3.3a", true},
		{"3.3a", ops.Lt, "3.3a", false},
		{"3.3a", ops.Like, "3.3a", true},
		{"3.3a", ops.Unlike, "4", true},

		{"2", ops.Gt, "1", true},
		{"2", ops.Ne, "1", true},
		{"2", ops.Eq, "1", false},
		{"2", ops.Gte, "1", true},
		{"2", ops.Lte, "1", false},
		{"2", ops.Lt, "1", false},
		{"2", ops.Like, "1", false},
		{"2", ops.Unlike, "1", true},

		{"1", ops.Gt, "2", false},
		{"1", ops.Ne, "2", true},
		{"1", ops.Eq, "2", false},
		{"1", ops.Gte, "2", false},
		{"1", ops.Lte, "2", true},
		{"1", ops.Lt, "2", true},
		{"1", ops.Like, "2", false},
		{"1", ops.Unlike, "2", true},
	}

	for _, v := range tests {
		ctx := &types.Context{Debug: false}
		err := compare.Versions(ctx, v.Op, v.Got, v.Want)
		assert.NoError(t, err)
		if v.Success {
			assert.True(t, ctx.Success, "success")
		} else {
			assert.False(t, ctx.Success, "failure")
		}
	}
}

func TestCompareVersionSegments(t *testing.T) {
	t.Parallel()
	type test struct {
		Got     string
		Op      string
		Want    string
		Segment uint
		Error   bool
		Success bool
	}
	tests := []test{
		{"3.3", ops.Eq, "3", 0, false, true},
		{"3.3", ops.Eq, "3", 1, false, true},
		{"3.3", ops.Eq, "0", 2, false, true},
		{"3.3", ops.Like, "3", 0, false, true},
		{"3.3", ops.Like, "3", 1, false, true},
		{"3.3", ops.Like, "0", 2, false, true},
		{"!!x]", ops.Like, "0", 2, true, false},
	}

	for _, v := range tests { //nolint:varnamelen
		ctx := &types.Context{Debug: false}
		err := compare.VersionSegment(ctx, v.Op, v.Got, v.Want, v.Segment)
		if v.Error {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
		if v.Success {
			assert.True(t, ctx.Success, "success")
		} else {
			assert.False(t, ctx.Success, "failure")
		}
	}
}

func TestStrings(t *testing.T) {
	t.Parallel()

	tests := []compareTest{
		{ops.Like, "delboy trotter", "delboy", false, true, false},
		{ops.Unlike, "delboy trotter", "delboy", false, false, false},
		{ops.Like, "delboy trotter", "Zdelboy", false, false, false},
		{ops.Unlike, "delboy trotter", "Zdelboy", false, true, false},
		{ops.Like, "delboy trotter", "/[/", true, false, false},
		{ops.Unlike, "delboy trotter", "/[/", true, false, false},
		{ops.Like, "delboy trotter", "delboy", false, true, true},
		{ops.Like, "delboy trotter", "delboyD", false, false, true},
	}

	testTable(t, tests,
		func(ctx *types.Context, this compareTest) error {
			return compare.Strings(ctx, this.Op, this.Got, this.Want)
		},
	)
}

func TestOptimistic(t *testing.T) {
	t.Parallel()

	tests := []compareTest{
		{ops.Like, "delboy trotter", "delboy", false, true, false},
		{ops.Unlike, "delboy trotter", "delboy", false, false, false},
		{ops.Like, "delboy trotter", "Zdelboy", false, false, false},
		{ops.Unlike, "delboy trotter", "Zdelboy", false, true, false},
		{ops.Like, "delboy trotter", "/[/", true, false, false},
		{ops.Unlike, "delboy trotter", "/[/", true, false, false},
		{ops.Like, "delboy trotter", "delboy", false, true, true},
		{ops.Like, "delboy trotter", "delboyD", false, false, true},
		{ops.Gte, "1", "1", false, true, true},
		{ops.Eq, "1.0", "1", false, true, true},
		{ops.Ne, "1", "2", false, true, true},
		{ops.Ne, "a", "2", false, true, true},
		{ops.Gte, "/[/", "1", false, false, true},
		{ops.Gte, "1", "/[/", false, false, true},
	}

	testTable(t, tests,
		func(ctx *types.Context, this compareTest) error {
			return compare.Optimistic(ctx, this.Op, this.Got, this.Want)
		},
	)
}

func TestIntegers(t *testing.T) {
	t.Parallel()

	tests := []compareTest{
		{ops.Eq, "1", "1", false, true, true},
		{ops.Gte, "1", "1", false, true, true},
		{ops.Gt, "1", "1", false, false, true},
		{ops.Gte, "2", "1", false, true, true},
		{ops.Lt, "1", "1", false, false, true},
		{ops.Lte, "1", "1", false, true, true},
		{ops.Ne, "1", "2", false, true, true},
		{ops.Ne, "a", "2", true, false, true},
		{ops.Gte, "/[/", "1", true, false, true},
		{ops.Gte, "1", "/[/", true, false, true},
	}

	testTable(t, tests,
		func(ctx *types.Context, this compareTest) error {
			return compare.Integers(ctx, this.Op, this.Got, this.Want)
		},
	)
}

func TestFloats(t *testing.T) {
	t.Parallel()

	tests := []compareTest{
		{ops.Eq, "1", "1", false, true, true},
		{ops.Eq, "1.0", "1", false, true, true},
		{ops.Eq, "1", "1.0", false, true, true},
		{ops.Gte, "1", "1", false, true, true},
		{ops.Gte, "2", "1", false, true, true},
		{ops.Ne, "1", "2", false, true, true},
		{ops.Ne, "a", "2", true, false, true},
		{ops.Gte, "/[/", "1", true, false, true},
		{ops.Gte, "1", "/[/", true, false, true},
	}

	testTable(t, tests,
		func(ctx *types.Context, this compareTest) error {
			return compare.Floats(ctx, this.Op, this.Got, this.Want)
		},
	)
}

func testTable( //nolint:thelper
	t *testing.T,
	tests []compareTest,
	comparison func(ctx *types.Context, this compareTest) error,
) {
	for _, this := range tests {
		ctx := &types.Context{Debug: this.Debug}
		err := comparison(ctx, this)
		label := fmt.Sprintf("%s %s %s", this.Got, this.Op, this.Want)
		if this.Success {
			assert.True(t, ctx.Success, label)
		} else {
			assert.False(t, ctx.Success, label)
		}
		if this.Error {
			assert.Error(t, err, label)
		} else {
			assert.NoError(t, err, label)
		}
	}
}
