package model

// HealthData of a recipe and all it's contents.
type HealthData struct {
	Calories int
	Fats     int
	Carbs    int
	Protein  int
	Fiber    int
	SatFat   int
	Sugar    int
}

// Summary of a recipe with generic prep time
type Summary struct {
	Preptime int
	Waittime int
	Cooktime int
	Servings int
}

// CookStep when prepairing a recipe
type CookStep struct {
	Action        string
	Duration      int
	DurationUnits string
}

// Ingredient in a recipe
type Ingredient struct {
	Food          string
	Quantity      int
	QuantityUnits string
	PrepAction    string
}

// Recipe to be catalogued and cooked
type Recipe struct {
	UniqueID         string
	RecipeName       string
	Source           string
	Comments         []string
	RecipeSummary    Summary
	RecipeHealthData HealthData
	Instructions     []CookStep
	Ingredients      []Ingredient
	Tags             []string
}

var (
	recipes []*Recipe
)
