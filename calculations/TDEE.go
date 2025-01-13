package calculations

// ActivityLevelMap contains multipliers for different activity levels
var ActivityLevelMap = map[string]float64{
	"sedentary":   1.2,
	"moderate":    1.55,
	"active":      1.725,
	"very_active": 1.9,
}

// CalculateTDEE computes Total Daily Energy Expenditure
func CalculateTDEE(bmr float64, activityLevel string) float64 {
	multiplier, exists := ActivityLevelMap[activityLevel]
	if !exists {
		return 0
	}
	return bmr * multiplier
}

