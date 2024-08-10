package ingredient

type Ingredient struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	MonthStart int
	MonthEnd   int
}

type IngredientDTO struct {
	Name          string  `json:"name"`
	Quantity      float64 `json:"quantity"`
	QuantityValue string  `json:"quantity_value"`
}

type IngredientRecipePost struct {
	ID            uint    `json:"id"`
	Name          string  `json:"name"`
	Quantity      float64 `json:"quantity"`
	QuantityValue string  `json:"quantity_value"`
}
