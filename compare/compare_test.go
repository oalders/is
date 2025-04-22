package compare_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/oalders/is/compare"
	"github.com/oalders/is/ops"
	"github.com/oalders/is/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type compareTest struct {
	Op      string
	Got     string
	Want    string
	Segment uint
	Error   bool
	Success bool
	Debug   bool
}

func TestVersions(t *testing.T) {
	t.Parallel()
	tests := []compareTest{
		{
			Op:      ops.Gt,
			Got:     "3.3",
			Want:    "3.3",
			Success: false,
		},
		{Op: ops.Ne, Got: "3.3", Want: "3.3", Success: false},
		{Op: ops.Eq, Got: "3.3", Want: "3.3", Success: true},
		{Op: ops.Gte, Got: "3.3", Want: "3.3", Success: true},
		{Op: ops.In, Got: "3.3", Want: "3.3,4.4", Success: true},
		{Op: ops.In, Got: "3.3", Want: "4.4", Success: false},
		{Op: ops.Lte, Got: "3.3", Want: "3.3", Success: true},
		{Op: ops.Lt, Got: "3.3", Want: "3.3", Success: false},
		{Op: ops.Like, Got: "3.3", Want: "3.3", Success: true},
		{Op: ops.Unlike, Got: "3.3", Want: "4", Success: true},

		{Op: ops.Gt, Got: "3.3a", Want: "3.3a", Success: false},
		{Op: ops.Ne, Got: "3.3a", Want: "3.3a", Success: false},
		{Op: ops.Eq, Got: "3.3a", Want: "3.3a", Success: true},
		{Op: ops.Gte, Got: "3.3a", Want: "3.3a", Success: true},
		{Op: ops.Lte, Got: "3.3a", Want: "3.3a", Success: true},
		{Op: ops.Lt, Got: "3.3a", Want: "3.3a", Success: false},
		{Op: ops.Like, Got: "3.3a", Want: "3.3a", Success: true},
		{Op: ops.Unlike, Got: "3.3a", Want: "4", Success: true},

		{Op: ops.Gt, Got: "2", Want: "1", Success: true},
		{Op: ops.Ne, Got: "2", Want: "1", Success: true},
		{Op: ops.Eq, Got: "2", Want: "1", Success: false},
		{Op: ops.Gte, Got: "2", Want: "1", Success: true},
		{Op: ops.Lte, Got: "2", Want: "1", Success: false},
		{Op: ops.Lt, Got: "2", Want: "1", Success: false},
		{Op: ops.Like, Got: "2", Want: "1", Success: false},
		{Op: ops.Unlike, Got: "2", Want: "1", Success: true},

		{Op: ops.Gt, Got: "1", Want: "2", Success: false},
		{Op: ops.Ne, Got: "1", Want: "2", Success: true},
		{Op: ops.Eq, Got: "1", Want: "2", Success: false},
		{Op: ops.Gte, Got: "1", Want: "2", Success: false},
		{Op: ops.Lte, Got: "1", Want: "2", Success: true},
		{Op: ops.Lt, Got: "1", Want: "2", Success: true},
		{Op: ops.Like, Got: "1", Want: "2", Success: false},
		{Op: ops.Unlike, Got: "1", Want: "2", Success: true},
	}

	testTable(t, tests,
		func(ctx *types.Context, this compareTest) (bool, error) {
			return compare.Versions(ctx, this.Op, this.Got, this.Want)
		},
	)
	{
		ctx := &types.Context{}
		ok, err := compare.Versions(ctx, ops.In, "3.3", strings.Repeat("3,", 100))
		assert.Error(t, err)
		assert.False(t, ok)
	}
	{
		ctx := &types.Context{}
		ok, err := compare.Versions(ctx, ops.In, "3.3", "!!x")
		assert.Error(t, err)
		assert.False(t, ok)
	}
}

func TestCompareVersionSegments(t *testing.T) {
	t.Parallel()
	tests := []compareTest{
		{
			Op:      ops.Eq,
			Got:     "3.3",
			Want:    "3",
			Segment: 0,
			Error:   false,
			Success: true,
		},
		{
			Op:      ops.Eq,
			Got:     "3.3",
			Want:    "3",
			Segment: 1,
			Error:   false,
			Success: true,
		},
		{
			Op:      ops.Eq,
			Got:     "3.3",
			Want:    "0",
			Segment: 2,
			Error:   false,
			Success: true,
		},
		{
			Op:      ops.In,
			Got:     "3.3",
			Want:    "1,2,3,4",
			Segment: 0,
			Error:   false,
			Success: true,
		},
		{
			Op:      ops.In,
			Got:     "3.3",
			Want:    "4,5,6",
			Segment: 0,
			Error:   false,
			Success: false,
		},
		{
			Op:      ops.In,
			Got:     "3.3",
			Want:    "4.0,5,6",
			Segment: 0,
			Error:   true,
			Success: false,
		},
		{
			Op:      ops.In,
			Got:     "3.3",
			Want:    strings.Repeat("X,", 100),
			Segment: 0,
			Error:   true,
			Success: false,
		},
		{
			Op:      ops.Like,
			Got:     "3.3",
			Want:    "3",
			Segment: 0,
			Error:   false,
			Success: true,
		},
		{
			Op:      ops.Like,
			Got:     "3.3",
			Want:    "3",
			Segment: 1,
			Error:   false,
			Success: true,
		},
		{
			Op:      ops.Like,
			Got:     "3.3",
			Want:    "0",
			Segment: 2,
			Error:   false,
			Success: true,
		},
		{
			Op:      ops.Like,
			Got:     "!!x]",
			Want:    "0",
			Segment: 2,
			Error:   true,
			Success: false,
		},
	}

	for _, v := range tests { //nolint:varnamelen
		label := fmt.Sprintf("%s %s %s %d", v.Got, v.Op, v.Want, v.Segment)
		ctx := &types.Context{Debug: false}
		success, err := compare.VersionSegment(ctx, v.Op, v.Got, v.Want, v.Segment)
		if v.Error {
			assert.Error(t, err)
		} else {
			require.NoError(t, err)
		}
		if v.Success {
			assert.True(t, success, label+"success ")
		} else {
			assert.False(t, success, label+"failure")
		}
	}
}

func TestStrings(t *testing.T) {
	t.Parallel()

	tests := []compareTest{
		{
			Op:      ops.Like,
			Got:     "delboy trotter",
			Want:    "delboy",
			Error:   false,
			Success: true,
			Debug:   false,
		},
		{
			Op:      ops.Unlike,
			Got:     "delboy trotter",
			Want:    "delboy",
			Error:   false,
			Success: false,
			Debug:   false,
		},
		{
			Op:      ops.Like,
			Got:     "delboy trotter",
			Want:    "Zdelboy",
			Error:   false,
			Success: false,
			Debug:   false,
		},
		{
			Op:      ops.Unlike,
			Got:     "delboy trotter",
			Want:    "Zdelboy",
			Error:   false,
			Success: true,
			Debug:   false,
		},
		{
			Op:      ops.Like,
			Got:     "delboy trotter",
			Want:    "/[/",
			Error:   true,
			Success: false,
			Debug:   false,
		},
		{
			Op:      ops.Unlike,
			Got:     "delboy trotter",
			Want:    "/[/",
			Error:   true,
			Success: false,
			Debug:   false,
		},
		{
			Op:      ops.Like,
			Got:     "delboy trotter",
			Want:    "delboy",
			Error:   false,
			Success: true,
			Debug:   true,
		},
		{
			Op:      ops.Like,
			Got:     "delboy trotter",
			Want:    "delboyD",
			Error:   false,
			Success: false,
			Debug:   true,
		},
		{
			Op:      ops.In,
			Got:     "delboy trotter",
			Want:    "delboy trotter, rodney trotter",
			Error:   false,
			Success: true,
			Debug:   false,
		},
		{
			Op:      ops.In,
			Got:     "X",
			Want:    strings.Repeat("X,", 99),
			Error:   false,
			Success: true,
			Debug:   false,
		},
		{
			Op:      ops.In,
			Got:     "X",
			Want:    strings.Repeat("X,", 100),
			Error:   true,
			Success: false,
			Debug:   false,
		},
	}

	testTable(t, tests,
		func(ctx *types.Context, this compareTest) (bool, error) {
			return compare.Strings(ctx, this.Op, this.Got, this.Want)
		},
	)
}

func TestOptimistic(t *testing.T) {
	t.Parallel()

	tests := []compareTest{
		{
			Op:      ops.Like,
			Got:     "delboy trotter",
			Want:    "delboy",
			Error:   false,
			Success: true,
			Debug:   false,
		},
		{
			Op:      ops.Unlike,
			Got:     "delboy trotter",
			Want:    "delboy",
			Error:   false,
			Success: false,
			Debug:   false,
		},
		{
			Op:      ops.Like,
			Got:     "delboy trotter",
			Want:    "Zdelboy",
			Error:   false,
			Success: false,
			Debug:   false,
		},
		{
			Op:      ops.Unlike,
			Got:     "delboy trotter",
			Want:    "Zdelboy",
			Error:   false,
			Success: true,
			Debug:   false,
		},
		{
			Op:      ops.Like,
			Got:     "delboy trotter",
			Want:    "/[/",
			Error:   false,
			Success: false,
			Debug:   false,
		},
		{
			Op:      ops.Unlike,
			Got:     "delboy trotter",
			Want:    "/[/",
			Error:   false,
			Success: false,
			Debug:   false,
		},
		{
			Op:      ops.Like,
			Got:     "delboy trotter",
			Want:    "delboy",
			Error:   false,
			Success: true,
			Debug:   true,
		},
		{
			Op:      ops.Like,
			Got:     "delboy trotter",
			Want:    "delboyD",
			Error:   false,
			Success: false,
			Debug:   true,
		},
		{
			Op:      ops.Gte,
			Got:     "1",
			Want:    "1",
			Error:   false,
			Success: true,
			Debug:   true,
		},
		{
			Op:      ops.Eq,
			Got:     "1.0",
			Want:    "1",
			Error:   false,
			Success: true,
			Debug:   true,
		},
		{
			Op:      ops.Ne,
			Got:     "1",
			Want:    "2",
			Error:   false,
			Success: true,
			Debug:   true,
		},
		{
			Op:      ops.Ne,
			Got:     "a",
			Want:    "2",
			Error:   false,
			Success: true,
			Debug:   true,
		},
		{
			Op:      ops.Gte,
			Got:     "/[/",
			Want:    "1",
			Error:   false,
			Success: false,
			Debug:   true,
		},
		{
			Op:      ops.Gte,
			Got:     "1",
			Want:    "/[/",
			Error:   false,
			Success: false,
			Debug:   true,
		},
		{
			Op:      ops.In,
			Got:     "X",
			Want:    strings.Repeat("X,", 100),
			Error:   false,
			Success: false,
			Debug:   false,
		},
	}

	testTable(t, tests,
		func(ctx *types.Context, this compareTest) (bool, error) {
			return compare.Optimistic(ctx, this.Op, this.Got, this.Want), nil
		},
	)
}

func TestIntegers(t *testing.T) {
	t.Parallel()

	tests := []compareTest{
		{
			Op:      ops.Eq,
			Got:     "1",
			Want:    "1",
			Error:   false,
			Success: true,
			Debug:   true,
		},
		{
			Op:      ops.Gte,
			Got:     "1",
			Want:    "1",
			Error:   false,
			Success: true,
			Debug:   true,
		},
		{
			Op:      ops.Gt,
			Got:     "1",
			Want:    "1",
			Error:   false,
			Success: false,
			Debug:   true,
		},
		{
			Op:      ops.Gte,
			Got:     "2",
			Want:    "1",
			Error:   false,
			Success: true,
			Debug:   true,
		},
		{
			Op:      ops.In,
			Got:     "1",
			Want:    "0,1",
			Error:   false,
			Success: true,
			Debug:   true,
		},
		{
			Op:      ops.In,
			Got:     "1",
			Want:    "2,3",
			Error:   false,
			Success: false,
			Debug:   true,
		},
		{
			Op:      ops.In,
			Got:     "1",
			Want:    "2.0,3.0",
			Error:   true,
			Success: false,
			Debug:   true,
		},
		{
			Op:      ops.Lt,
			Got:     "1",
			Want:    "1",
			Error:   false,
			Success: false,
			Debug:   true,
		},
		{
			Op:      ops.Lte,
			Got:     "1",
			Want:    "1",
			Error:   false,
			Success: true,
			Debug:   true,
		},
		{
			Op:      ops.Ne,
			Got:     "1",
			Want:    "2",
			Error:   false,
			Success: true,
			Debug:   true,
		},
		{
			Op:      ops.Ne,
			Got:     "a",
			Want:    "2",
			Error:   true,
			Success: false,
			Debug:   true,
		},
		{
			Op:      ops.Gte,
			Got:     "/[/",
			Want:    "1",
			Error:   true,
			Success: false,
			Debug:   true,
		},
		{
			Op:      ops.Gte,
			Got:     "1",
			Want:    "/[/",
			Error:   true,
			Success: false,
			Debug:   true,
		},
		{
			Op:      ops.In,
			Got:     "X",
			Want:    strings.Repeat("X,", 100),
			Error:   true,
			Success: false,
			Debug:   false,
		},
	}

	testTable(t, tests,
		func(ctx *types.Context, this compareTest) (bool, error) {
			return compare.Integers(ctx, this.Op, this.Got, this.Want)
		},
	)
}

func TestFloats(t *testing.T) {
	t.Parallel()

	tests := []compareTest{
		{
			Op:      ops.Eq,
			Got:     "1",
			Want:    "1",
			Error:   false,
			Success: true,
			Debug:   true,
		},
		{
			Op:      ops.Eq,
			Got:     "1.0",
			Want:    "1",
			Error:   false,
			Success: true,
			Debug:   true,
		},
		{
			Op:      ops.Eq,
			Got:     "1",
			Want:    "1.0",
			Error:   false,
			Success: true,
			Debug:   true,
		},
		{
			Op:      ops.Gte,
			Got:     "1",
			Want:    "1",
			Error:   false,
			Success: true,
			Debug:   true,
		},
		{
			Op:      ops.Gte,
			Got:     "2",
			Want:    "1",
			Error:   false,
			Success: true,
			Debug:   true,
		},
		{
			Op:      ops.In,
			Got:     "1.0",
			Want:    "1.0,2.0",
			Error:   false,
			Success: true,
			Debug:   true,
		},
		{
			Op:      ops.In,
			Got:     "1.0",
			Want:    "2.0,3.0",
			Error:   false,
			Success: false,
			Debug:   false,
		},
		{
			Op:      ops.In,
			Got:     "1.0",
			Want:    "2.0,3.0,X",
			Error:   true,
			Success: false,
			Debug:   false,
		},
		{
			Op:      ops.Ne,
			Got:     "1",
			Want:    "2",
			Error:   false,
			Success: true,
			Debug:   true,
		},
		{
			Op:      ops.Ne,
			Got:     "a",
			Want:    "2",
			Error:   true,
			Success: false,
			Debug:   true,
		},
		{
			Op:      ops.Gte,
			Got:     "/[/",
			Want:    "1",
			Error:   true,
			Success: false,
			Debug:   true,
		},
		{
			Op:      ops.Gte,
			Got:     "1",
			Want:    "/[/",
			Error:   true,
			Success: false,
			Debug:   true,
		},
		{
			Op:      ops.In,
			Got:     "1.0",
			Want:    strings.Repeat("1.0,", 100),
			Error:   true,
			Success: false,
			Debug:   false,
		},
	}

	testTable(t, tests,
		func(ctx *types.Context, this compareTest) (bool, error) {
			return compare.Floats(ctx, this.Op, this.Got, this.Want)
		},
	)
}

func testTable( //nolint:thelper
	t *testing.T,
	tests []compareTest,
	comparison func(ctx *types.Context, this compareTest) (bool, error),
) {
	for _, this := range tests {
		ctx := &types.Context{Debug: this.Debug}
		ok, err := comparison(ctx, this)
		label := fmt.Sprintf("[%s %s %s]", this.Got, this.Op, this.Want)
		assert.Equal(t, this.Success, ok, label)
		if this.Error {
			require.Error(t, err, label)
		} else {
			require.NoError(t, err, label)
		}
	}
}
