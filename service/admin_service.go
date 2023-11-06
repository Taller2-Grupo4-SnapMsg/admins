package service

import (
	"admins/repository"
	"admins/structs"
	"time"
)

/**
 * This function saves an admin on the database.
 * @param email The email of the admin.
 * @param password The password of the admin.
 * @return the admin if it was saved, nil otherwise.
 */
func SaveAdmin(email string, password string) (*structs.AdminStruct, error) {
	// First we check the admin doesn't exist:
	admin := GetAdmin(email)
	if admin != nil {
		return nil, nil
	}
	admin = &structs.AdminStruct{
		Email:     email,
		Password:  password,
		TimeStamp: time.Now().Unix(),
	} // refference for efficiency
	return repository.SaveAdmin(admin)
}

/**
 * This function gets an admin from the database.
 * @param email The email of the admin.
 * @return the admin if it was found, nil otherwise.
 */
func GetAdmin(email string) *structs.AdminStruct {
	return repository.GetAdmin(email)
}

/**
 * This function deletes an admin from the database.
 * @param email The email of the admin.
 * @return a string with the result of the operation, and an error if something went wrong.
 */
func DeleteAdmin(email string) (string, error) {
	delete_count, error := repository.DeleteAdmin(email)
	if error != nil {
		panic(error)
	}
	if delete_count == 0 {
		return "Admin not found", nil
	}
	return "Admin deleted", nil
}
