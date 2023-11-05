package service

import (
	"testing"
)

const (
	EMAIL    = "test_email@gmail.com"
	PASSWORD = "test_password"
)

func clean_up(email string, t *testing.T) {
	result, err := DeleteAdmin(email)
	if err != nil {
		panic(err)
	}
	if result != "Admin deleted" {
		t.Errorf("Admin should be deleted")
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
	// Clean up
	clean_up(email, t)
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

	// Clean up
	clean_up(email, t)
}

func TestDeleteAdminAdminIsDeleted(t *testing.T) {
	// This test checks if an admin is deleted
	// Arrange
	email := EMAIL
	password := PASSWORD
	// Act
	admin_save, _ := SaveAdmin(email, password)

	if admin_save == nil {
		t.Errorf("Admin should be saved")
	}

	admin := GetAdmin(email)

	if admin == nil {
		t.Errorf("Admin should be found")
	}

	message, _ := DeleteAdmin(email)
	// Assert
	if message != "Admin deleted" {
		t.Errorf("Message should be Admin deleted")
	}
}

func TestDeleteAdminNotFound(t *testing.T) {
	// This test checks if an admin is deleted
	// Arrange
	email := EMAIL
	// Act
	message, _ := DeleteAdmin(email)
	// Assert
	if message != "Admin not found" {
		t.Errorf("Message should be Admin not found and was " + message)
	}
}
