package util

import (
	"testing"
)

func TestHashAndPassword(t *testing.T) {
	password := "my_secret_password"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Expected no error when hashing password, got %v", err)
	}

	if hash == "" {
		t.Fatal("Expected hash to not be empty")
	}

	if hash == password {
		t.Fatal("Expected hash to be different from password")
	}

	// Test correct password
	err = CheckPassword(hash, password)
	if err != nil {
		t.Errorf("Expected password to match hash, got error: %v", err)
	}

	// Test incorrect password
	err = CheckPassword(hash, "wrong_password")
	if err == nil {
		t.Error("Expected error when checking wrong password, got nil")
	}
}
