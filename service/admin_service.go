package service

import (
	"admins/repository"
	"admins/structs"
)

/**
 * This function saves an admin on the database.
 * @param email The email of the admin.
 * @param password The password of the admin.
 * @return the admin if it was saved, nil otherwise.
 */
func SaveAdmin(email string, password string) (*structs.AdminStruct, error) {
	admin := &structs.AdminStruct{
		Email:    email,
		Password: password,
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

/**
 * This function saves a number on the database.
 * @param number The number to be saved.
 * @return void
 */
func SaveNumber(number int32) {
	repository.SaveNumber(number)
}

/**
 * This function gets a list of numbers from the database.
 * @return []int32 The list of numbers.
 */
func GetNumbers() []int32 {
	return repository.GetNumbers()
}
