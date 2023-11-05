package service

import (
	"testing"
)

const (
	EMAIL    = "test_email@gmail.com"
	PASSWORD = "test_password"
)

func TestDummy(t *testing.T) {
	// This is a dummy test that always passes
	if 1+1 != 2 {
		t.Errorf("1+1 should equal 2")
	}
}

func TestSaveAdminAdminIsSaved(t *testing.T) {
	// This test checks if an admin is saved
	// Arrange
	email := EMAIL
	password := PASSWORD
	// Act
	admin, _ := SaveAdmin(email, password)
	// Assert
	if admin.Email != email {
		t.Errorf("Email should be %s", email)
	}
}

func TestGetAdminAdminIsFound(t *testing.T) {
	// This test checks if an admin is found
	// Arrange
	email := EMAIL
	password := PASSWORD
	// Act
	admin_save, _ := SaveAdmin(email, password)

	if admin_save == nil {
		t.Errorf("Admin should be saved")
	}

	admin := GetAdmin(email)
	// Assert
	if admin.Email != email {
		t.Errorf("Email should be %s", email)
	}
}
