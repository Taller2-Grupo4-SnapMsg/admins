package main

import (
	"admins/service"
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
	admin, _ := service.SaveAdmin(email, password)
	// Assert
	if admin.Email != email {
		t.Errorf("Email should be %s", email)
	}
}
