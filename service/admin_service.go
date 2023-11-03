package service

import "admins/repository"

func SaveNumber(number int32) {
	repository.SaveNumber(number)
}

func GetNumbers() []int32 {
	return repository.GetNumbers()
}
