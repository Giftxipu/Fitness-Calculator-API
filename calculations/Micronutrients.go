package calculations

// CalculateMicronutrients provides basic micronutrient recommendations
func CalculateMicronutrients(weight float64) map[string]float64 {
	return map[string]float64{
		"calcium_mg":  1000,
		"iron_mg":     8,
		"fiber_g":     14 * weight / 1000,
		"vitamin_c_mg": 90,
	}
}

