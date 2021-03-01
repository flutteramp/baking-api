package recipe

import "github.com/flutter-amp/baking-api/entity"

type RecipeService interface {
	Recipes() ([]entity.Recipe, []error)
	Recipe(id uint) (*entity.Recipe, []error)
	Ingredients(id uint) ([]entity.Ingredient, []error)
	Steps(id uint) ([]entity.Step, []error)
	UserRecipes(uid uint) ([]entity.Recipe, []error)
	UpdateRecipe(recipe *entity.Recipe) (*entity.Recipe, []error)
	DeleteRecipe(id uint) (*entity.Recipe, []error)
	StoreRecipe(recipe *entity.Recipe) (*entity.Recipe, []error)
	//	updateImage(recipe *entity.Recipe, imagePath string) (*entity.Recipe, []error)
}
