package database

import (
	"cooking-web-app/services/ingredient"
	"cooking-web-app/services/recipe"
	"cooking-web-app/services/relation"
	"cooking-web-app/services/user"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetDB() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic("failed to connect database" + err.Error())
	}

	// Set Join Tables

	// Migrate the schema
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&ingredient.Ingredient{})
	db.AutoMigrate(&relation.RelationIngredientUser{})
	db.AutoMigrate(&recipe.Recipe{})
	db.AutoMigrate(&relation.RelationIngredientRecipe{})

	//db.SetupJoinTable(&model.User{}, "Ingredients", &model.RelationIngredientUser{})

	//populateDB(db)
	log.Println("Database connected")
	return db
}

func populateDB(db *gorm.DB) {

	// Clean
	db.Exec("DELETE FROM relation_ingredient_users")
	db.Exec("DELETE FROM ingredients")
	db.Exec("DELETE FROM users")

	db.Create(&user.User{Login: "admin", HashPassword: "admin", Email: "admin@mail.com"})
	db.Create(&user.User{Login: "user", HashPassword: "user", Email: "user@mail.com"})
	db.Create(&ingredient.Ingredient{Name: "Tomato", Price: 1.0, MonthStart: 1, MonthEnd: 12})
	db.Create(&ingredient.Ingredient{Name: "Potato", Price: 0.5, MonthStart: 1, MonthEnd: 12})
	db.Create(&ingredient.Ingredient{Name: "Onion", Price: 0.5, MonthStart: 1, MonthEnd: 12})
	db.Create(&ingredient.Ingredient{Name: "Garlic", Price: 0.5, MonthStart: 1, MonthEnd: 12})

	db.Create(&relation.RelationIngredientUser{UserID: 1, IngredientID: 1, Quantity: 2, QuantityValue: "kg"})
	db.Create(&relation.RelationIngredientUser{UserID: 1, IngredientID: 2, Quantity: 2, QuantityValue: "kg"})

	// Create new recipe
	db.Create(&recipe.Recipe{
		Name:        "Tomato soup",
		Description: "Very good soup",
		Steps:       "1. Cut tomato\n2. Cook tomato\n3. Eat tomato",
		NbPeople:    4,
	})
	db.Create(&relation.RelationIngredientRecipe{RecipeID: 1, IngredientID: 1, Quantity: 2, QuantityValue: "kg"})
	db.Create(&relation.RelationIngredientRecipe{RecipeID: 1, IngredientID: 2, Quantity: 2, QuantityValue: "kg"})
	db.Create(&relation.RelationIngredientRecipe{RecipeID: 1, IngredientID: 3, Quantity: 1, QuantityValue: "kg"})

	// Create new recipe
	db.Create(&recipe.Recipe{
		Name:        "Potato soup",
		Description: "Very good soup",
		Steps:       "1. Cut potato\n2. Cook potato\n3. Eat potato",
		NbPeople:    4,
	})
	db.Create(&relation.RelationIngredientRecipe{RecipeID: 2, IngredientID: 2, Quantity: 2, QuantityValue: "kg"})
	db.Create(&relation.RelationIngredientRecipe{RecipeID: 2, IngredientID: 3, Quantity: 1, QuantityValue: "kg"})
	db.Create(&relation.RelationIngredientRecipe{RecipeID: 2, IngredientID: 4, Quantity: 1, QuantityValue: "kg"})

}
