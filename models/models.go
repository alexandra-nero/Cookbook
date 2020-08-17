package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Recipe is used to marshal and unmarshal from a document database
type Recipe struct {
	RecipeID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	RecipeName      string             `json:"recipeName,omitempty"`
	CreatedDate     string             `json:"createdDate,omitempty"`
	LastUpdatedDate string             `json:"lastUpdatedDate,omitempty"`
	Ingredients     []ingredient       `json:"ingredients,omitempty"`
	Author          string             `json:"author,omitempty"`
	PrepTime        int                `json:"prepTime,omitempty"`
	CookTime        int                `json:"cookTime,omitempty"`
	Steps           []step             `json:"steps,omitempty"`
	Rating          int                `json:"rating,omitempty"`
	Servings        int                `json:"servings,omitempty"`
	Calories        int                `json:"calories,omitempty"`
	UserID          string             `json:"userId,omitempty"`
	Private         bool               `json:"private,omitempty"`
}

// Ingredient is a component of a recipe consisting of the name, amount, and the measurement for that amount (cups, tbsp, lbs, etc)
type ingredient struct {
	Name        string  `json:"name,omitempty"`
	Amount      float32 `json:"amount,omitempty"`
	Measurement string  `json:"measurement,omitempty"`
}

// Step is what to do in order for a recipe
type step struct {
	Number int    `json:"number,omitempty"`
	Text   string `json:"text,omitEmpty"`
}

// CalorieLog is the object representing a calorie log
type CalorieLog struct {
	CalorieLogID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	EnteredDate  string             `json:"enteredDate,omitempty"`
	Calories     int                `json:"calories,omitempty"`
	Description  string             `json:"description,omitEmpty"`
	UserID       string             `json:"userId,omitEmpty"`
}

// UserData is data pertaining to the user's physical characteristics
type UserData struct {
	UserID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Height              int                `json:"height,omitEmpty"`
	Weight              int                `json:"weight,omitEmpty"`
	MaintenanceCalories int                `json:"maintenanceCalories,omitEmpty"`
}

// UserAuth is the authentication information for a user
type UserAuth struct {
	UserID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	PasswordHash string             `json:"passwordHash,omitEmpty"`
	AccessToken  string             `json:"accessToken,omitEmpty"`
	ExpiryDate   string             `json:"expiryDate,omitempty"`
}