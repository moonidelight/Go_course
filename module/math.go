package module

type TwoNumbers struct {
	a, b int64
}

func Sum(numbers TwoNumbers) int64 {
	return numbers.a + numbers.b
}
func Product(numbers TwoNumbers) int64 {
	return numbers.a * numbers.b
}
