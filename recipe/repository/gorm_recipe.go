package repository

import (
	"fmt"

	"github.com/flutter-amp/baking-api/entity"
	"github.com/flutter-amp/baking-api/recipe"
	"github.com/jinzhu/gorm"
)

type RecipeGormRepo struct {
	conn *gorm.DB
}

func NewRecipeGormRepo(db *gorm.DB) recipe.RecipeRepository {
	return &RecipeGormRepo{conn: db}
}

func (recipeRepo *RecipeGormRepo) Recipes() ([]entity.Recipe, []error) {
	recipes := []entity.Recipe{}

	errs := recipeRepo.conn.Find(&recipes).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return recipes, errs
}
func (recipeRepo *RecipeGormRepo) Ingredients(id uint) ([]entity.Ingredient, []error) {
	ingredients := []entity.Ingredient{}
	err := recipeRepo.conn.Where("recipe_id = ?", id).Find(&ingredients).GetErrors()

	if len(err) > 0 {
		return nil, err
	}

	return ingredients, err
}
func (recipeRepo *RecipeGormRepo) Recipe(id uint) (*entity.Recipe, []error) {
	recipe := entity.Recipe{}
	errs := recipeRepo.conn.First(&recipe, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &recipe, errs
}
func (recipeRepo *RecipeGormRepo) Steps(id uint) ([]entity.Step, []error) {
	steps := []entity.Step{}
	err := recipeRepo.conn.Where("recipe_id = ?", id).Find(&steps).GetErrors()

	if len(err) > 0 {
		return nil, err
	}

	return steps, err
}

// func (recipeRepo *RecipeGormRepo) Recipe(id uint) (*entity.Recipe, []error) {
// 	recipe := entity.Recipe{}
// 	errs := recipeRepo.conn.First(&recipe, id).GetErrors()
// 	if len(errs) > 0 {
// 		return nil, errs
// 	}
// 	return &recipe, errs
// }
func (recipeRepo *RecipeGormRepo) DeleteRecipe(id uint) (*entity.Recipe, []error) {
	rcpe, errs := recipeRepo.Recipe(id)

	//err := recipeRepo.conn.Delete("recipe_id=?", id).GetErrors()
	err := recipeRepo.conn.Where("recipe_id=?", id).Delete(entity.Ingredient{}).GetErrors()
	if len(err) > 0 {
		print(err)
		return nil, errs

	}
	if len(errs) > 0 {
		return nil, errs
	}
	errs = recipeRepo.conn.Delete(rcpe, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return rcpe, errs
}
func (recipeRepo *RecipeGormRepo) UpdateRecipe(recipe *entity.Recipe) (*entity.Recipe, []error) {
	rsp := recipe
	err := recipeRepo.conn.Where("recipe_id=?", rsp.ID).Delete(entity.Ingredient{}).GetErrors()
	errt := recipeRepo.conn.Where("recipe_id=?", rsp.ID).Delete(entity.Step{}).GetErrors()
	if len(err) > 0 {
		print(err)
		return nil, err

	}
	if len(errt) > 0 {
		print(err)
		return nil, err

	}
	errs := recipeRepo.conn.Save(rsp).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return rsp, errs
}

func (recipeRepo *RecipeGormRepo) StoreRecipe(recipe *entity.Recipe) (*entity.Recipe, []error) {
	rcpe := recipe
	errs := recipeRepo.conn.Create(rcpe).GetErrors()
	if len(errs) > 0 {
		fmt.Println("found errors........................")
		return nil, errs
	}
	return rcpe, errs
}

// func (recipeRepo *RecipeGormRepo) updateImage(recipe *entity.Recipe, imagePath string) (*entity.Recipe, []error) {

// 	rcpe := recipe
// 	errs := recipeRepo.conn.Model(&rcpe).UpdateColumn("imageUrl", imagePath).GetErrors()
// 	if len(errs) > 0 {
// 		return nil, errs
// 	}
// 	return rcpe, errs
// }

func (recipeRepo *RecipeGormRepo) UserRecipes(uid uint) ([]entity.Recipe, []error) {
	usrRecipes := []entity.Recipe{}
	errs := recipeRepo.conn.Where("recipe_user = ?", uid).Find(&usrRecipes).GetErrors()
	//errs := recipeRepo.conn.Model(user).Related(&usrRecipes, "Orders").GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usrRecipes, errs
}
