package category

import (
	"strings"
	"testing"
)

func TestNewName(t *testing.T) {
	t.Run("Empty name returns error", func(t *testing.T) {
		name, err := NewName("    ")
		if err != ErrorEmptyName {
			t.Errorf("expected ErrorEmptyName, got %v", err)
		}
		if name != nil {
			t.Errorf("expected nil name for empty input, got %v", *name)
		}
	})

	t.Run("Valid name returns pointer to trimmed name", func(t *testing.T) {
		input := "  Electronics  "
		expected := "Electronics"
		name, err := NewName(input)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if *name != expected {
			t.Errorf("expected %s, got %s", expected, *name)
		}
	})

	t.Run("Name too long returns error", func(t *testing.T) {
		longName := strings.Repeat("a", 257)
		name, err := NewName(longName)
		if err != ErrorNameTooLong {
			t.Errorf("expected ErrorNameTooLong, got %v", err)
		}
		if name != nil {
			t.Error("expected nil name for name too long")
		}
	})
}
