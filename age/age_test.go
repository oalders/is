package age_test

import (
	"testing"

	"github.com/oalders/is/age"
)

func TestStringToDuration(t *testing.T) {
	t.Parallel()

	t.Run("valid value succeeds", func(t *testing.T) {
		t.Parallel()
		dur, err := age.StringToDuration("10", "days")
		if err != nil {
			t.Fatalf("unexpected error for value 10: %v", err)
		}
		if dur == nil {
			t.Fatal("expected non-nil duration")
		}
	})

	t.Run("value above 36500 returns error", func(t *testing.T) {
		t.Parallel()
		_, err := age.StringToDuration("36501", "days")
		if err == nil {
			t.Fatal("expected error for value 36501, got nil")
		}
	})

	t.Run("value below 1 returns error", func(t *testing.T) {
		t.Parallel()
		_, err := age.StringToDuration("0", "days")
		if err == nil {
			t.Fatal("expected error for value 0, got nil")
		}
	})

	t.Run("negative value returns error", func(t *testing.T) {
		t.Parallel()
		_, err := age.StringToDuration("-1", "days")
		if err == nil {
			t.Fatal("expected error for value -1, got nil")
		}
	})

	t.Run("boundary value 36500 succeeds", func(t *testing.T) {
		t.Parallel()
		dur, err := age.StringToDuration("36500", "days")
		if err != nil {
			t.Fatalf("unexpected error for boundary value 36500: %v", err)
		}
		if dur == nil {
			t.Fatal("expected non-nil duration")
		}
	})

	t.Run("boundary value 1 succeeds", func(t *testing.T) {
		t.Parallel()
		dur, err := age.StringToDuration("1", "hours")
		if err != nil {
			t.Fatalf("unexpected error for boundary value 1: %v", err)
		}
		if dur == nil {
			t.Fatal("expected non-nil duration")
		}
	})
}
