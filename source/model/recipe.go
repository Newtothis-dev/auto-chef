package model

const (
	// CalPerFat - Calories per Fat gram
	CalPerFat = 9

	// CalPerProt - Calories per protein gram
	CalPerProt = 4

	// CalPerCarb - Calories per carb gram
	CalPerCarb = 4

	// CalPerAlch - Calories per Alcohol gram
	CalPerAlch = 7
)

// HealthData of a recipe and all it's contents.
type HealthData struct {
	// Carbs - grams of carbs in the food
	Carbs int `json:"carbs"`

	// Sugar - grams of sugar in the food
	Sugar int `json:"sugars"`

	// Fiber - grams of fiber in the food
	Fiber int `json:"fiber"`

	// Fat - grams of fats in the food
	Fat int `json:"fat"`

	// SatFat - grams of saturated fats in food
	SatFat int `json:"sat fats"`

	// Protein - grams of protein in the food
	Protein int `json:"proteins"`

	// Calories just in case there are non-macro calories to be
	// accounted for (e.g. alcohol calories)
	Calories int `json:"calories"`
}

// Summary of a recipe with generic prep time
type Summary struct {
	Prep     int `json:"perptime"`
	Wait     int `json:"waittime"`
	Cook     int `json:"cooktime"`
	Servings int `json:"servings"`
}

// CookStep when prepairing a recipe
type CookStep struct {
	Action string `json:"action"`
	Dur    int    `json:"duration"`
	Units  string `json:"dur units"`
}

// Food specific food to be used in a a recipe. This food
// instance represents 1 serving of the food
type Food struct {
	// Name of the food that is being used
	Food string `json:"food"`
	// Measureable amount of food used in the recipe. this is -1 if the
	// food is not used during prep and is added in after the fact.
	// In that case, the qty per serving is used.
	Qty int `json:"quantity"`

	// Units of the measurement of 1 serving
	Units string `json:"units"`

	// How to prepare the food that is being used in the meal
	Prep string `json:"prep type"`

	// Data for 1 serving of this food
	Data HealthData `json:"health data"`

	// QtyPerServ of servings of food per serving in recipe
	QtyPerServ int `json:"qty per serving"`
}

// GenericIngredients in a recipe
type GenericIngredients struct {
	Options []Food `json:"food"`
}

// GenericRecipe a container class for a variable recipes. variable
// recipes have N permutations of a given recipe, based off the
// ingredients/cooking steps/tools available.
type GenericRecipe struct {
	UniqueID      string               `json:"id"`
	RecipeName    string               `json:"name"`
	Src           string               `json:"source"`
	Comments      []string             `json:"comments"`
	RecipeSummary Summary              `json:"summary"`
	Instructions  []CookStep           `json:"instructions"`
	Ingredients   []GenericIngredients `json:"Ingredients"`
	Tools         []string             `json:"kitchenware"`
	Tags          []string             `json:"tags"`
}

var (
	// @TODO:
	// Use a slice of recipes for now, will need to convert
	// over to a map using the GenericRecipes ID's. But we don't
	// have any good GUID generation atm.
	recipes []*GenericRecipe
)
