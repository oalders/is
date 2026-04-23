package age

import (
	"testing"
)

func TestStringToDuration(t *testing.T) {
	t.Run("valid value succeeds", func(t *testing.T) {
		dur, err := StringToDuration("10", "days")
		if err != nil {
			t.Fatalf("unexpected error for value 10: %v", err)
		}
		if dur == nil {
			t.Fatal("expected non-nil duration")
		}
	})

	t.Run("value above 36500 returns error", func(t *testing.T) {
		_, err := StringToDuration("36501", "days")
		if err == nil {
			t.Fatal("expected error for value 36501, got nil")
		}
	})

	t.Run("value below 1 returns error", func(t *testing.T) {
		_, err := StringToDuration("0", "days")
		if err == nil {
			t.Fatal("expected error for value 0, got nil")
		}
	})

	t.Run("negative value returns error", func(t *testing.T) {
		_, err := StringToDuration("-1", "days")
		if err == nil {
			t.Fatal("expected error for value -1, got nil")
		}
	})

	t.Run("boundary value 36500 succeeds", func(t *testing.T) {
		dur, err := StringToDuration("36500", "days")
		if err != nil {
			t.Fatalf("unexpected error for boundary value 36500: %v", err)
		}
		if dur == nil {
			t.Fatal("expected non-nil duration")
		}
	})

	t.Run("boundary value 1 succeeds", func(t *testing.T) {
		dur, err := StringToDuration("1", "hours")
		if err != nil {
			t.Fatalf("unexpected error for boundary value 1: %v", err)
		}
		if dur == nil {
			t.Fatal("expected non-nil duration")
		}
	})
}
