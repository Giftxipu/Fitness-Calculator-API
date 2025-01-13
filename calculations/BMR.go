package calculations

func CalculateBMR(weight, height float64, age int, gender string) float64 {
	if gender == "male" {
		return 10*weight + 6.25*height - 5*float64(age) + 5
	}
	return 10*weight + 6.25*height - 5*float64(age) - 161
}
