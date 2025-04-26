package utils

func FormMemberPermissions(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}
