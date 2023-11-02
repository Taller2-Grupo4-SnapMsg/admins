package repository

var numbers []int

func SaveNumber(number int) {
	numbers = append(numbers, number)
}

func GetNumbers() []int {
	return numbers
}
