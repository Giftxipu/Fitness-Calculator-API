package calculations

// CalculateMacros computes macronutrient distribution in grams
func CalculateMacros(tdee float64) map[string]float64 {
	proteinCalories := 0.3 * tdee // 30% protein
	carbsCalories := 0.4 * tdee   // 40% carbs
	fatCalories := 0.3 * tdee     // 30% fat

	return map[string]float64{
		"protein_grams": proteinCalories / 4,
		"carbs_grams":   carbsCalories / 4,
		"fat_grams":     fatCalories / 9,
	}
}

