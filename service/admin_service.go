package service

import "admins/repository"

func SaveNumber(number int) {
	repository.SaveNumber(number)
}

func GetNumbers() []int {
	return repository.GetNumbers()
}
