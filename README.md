# Fitness Calculator API

This API calculates various fitness metrics such as BMI, BMR, TDEE, and macronutrient/micronutrient requirements based on user input. It uses the Gin framework for routing and a modular file structure for organization.

---

## Features

- Calculate **BMI** (Body Mass Index)
    
- Calculate **BMR** (Basal Metabolic Rate)
    
- Calculate **TDEE** (Total Daily Energy Expenditure)
    
- Generate macronutrient and micronutrient breakdowns
    

---

## Folder Structure

```
fitness-calculator-api/
|
├── main.go              // Entry point
├── handlers/            // API handlers
│   ├── fitness.go       // Handlers for all fitness calculations
│
├── calculations/        // Core calculation logic
│   ├── bmi.go           // Logic for BMI calculation
│   ├── bmr.go           // Logic for BMR calculation
│   ├── tdee.go          // Logic for TDEE calculation
│   ├── macros.go        // Logic for macronutrient calculation
│   ├── micronutrients.go// Logic for micronutrient calculation
│
└── models/              // Data models
    ├── user.go          // User struct definition
```

---

## Installation

1. **Clone the repository**:
    

```
git clone <repository-url>
cd fitness-calculator-api
```

2. **Initialize the project**:
    

```
go mod init fitnessApi
```

3. **Install Gin**:
    

```
go get -u github.com/gin-gonic/gin
```

4. **Run the application**:
    

```
go run .
```

---

## API Endpoints

### POST `/fitness`

**Description**: Calculate fitness metrics based on user input.

**Request Body**:

```
{
  "weight": 70,
  "height": 175,
  "age": 25,
  "gender": "male",
  "activity_level": "moderate"
}
```

**Response**:

```
{
  "bmi": 22.86,
  "bmr": 1667.5,
  "tdee": 2584.63,
  "macros": {
    "protein_grams": 194.1,
    "carbs_grams": 258.46,
    "fat_grams": 86.15
  },
  "micronutrients": {
    "calcium_mg": 1000,
    "iron_mg": 8,
    "fiber_g": 0.98,
    "vitamin_c_mg": 90
  }
}
```

---

## Testing with CURL

Use the following command to test the API:

```
curl -X POST http://localhost:8080/fitness \
-H "Content-Type: application/json" \
-d '{
  "weight": 70,
  "height": 175,
  "age": 25,
  "gender": "male",
  "activity_level": "moderate"
}'
```

**Expected Output**:

```
{
  "bmi": 22.86,
  "bmr": 1667.5,
  "tdee": 2584.63,
  "macros": {
    "protein_grams": 194.1,
    "carbs_grams": 258.46,
    "fat_grams": 86.15
  },
  "micronutrients": {
    "calcium_mg": 1000,
    "iron_mg": 8,
    "fiber_g": 0.98,
    "vitamin_c_mg": 90
  }
}
```

---

## Dependencies

- [Gin Web Framework](https://github.com/gin-gonic/gin)
    
- Go 1.18 or later
    

---

## Notes

- Ensure `go` is installed and set up properly on your system.
    
- Use proper activity levels (e.g., "sedentary," "light," "moderate," "active," "very active").
    
- Extend the `calculations` module for additional features if needed.
    

---
