package models

// User represents the structure for user input data
type User struct {
	Weight        float64 `json:"weight" binding:"required"`
	Height        float64 `json:"height" binding:"required"`
	Age           int     `json:"age" binding:"required"`
	Gender        string  `json:"gender" binding:"required"`
	ActivityLevel string  `json:"activity_level" binding:"required"` // Add this field
}

