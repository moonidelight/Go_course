package module

type TwoNumbers struct {
	A, B int64
}

func Sum(numbers TwoNumbers) int64 {
	return numbers.A + numbers.B
}
func Product(numbers TwoNumbers) int64 {
	return numbers.A * numbers.B
}
