package handlers

import (
	"net/http"
	"fitnessApi/calculations"
	"fitnessApi/models"
	"github.com/gin-gonic/gin"
)

// FitnessHandler processes user data and returns fitness metrics
func FitnessHandler(c *gin.Context) {
	var user models.User

	// Bind JSON to User struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Perform calculations
	bmi := calculations.CalculateBMI(user.Weight, user.Height)
	bmr := calculations.CalculateBMR(user.Weight, user.Height, user.Age, user.Gender)
	tdee := calculations.CalculateTDEE(bmr, user.ActivityLevel)
	macros := calculations.CalculateMacros(tdee)
	micronutrients := calculations.CalculateMicronutrients(user.Weight)

	// Respond with the calculated metrics
	c.JSON(http.StatusOK, gin.H{
		"bmi":            bmi,
		"bmr":            bmr,
		"tdee":           tdee,
		"macros":         macros,
		"micronutrients": micronutrients,
	})
}

