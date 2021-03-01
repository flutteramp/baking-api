package service

import (
	"fmt"

	"github.com/flutter-amp/baking-api/entity"
	"github.com/flutter-amp/baking-api/recipe"
)

type RecipeService struct {
	recipeRepo recipe.RecipeRepository
}

func NewRecipeService(recipeRepository recipe.RecipeRepository) recipe.RecipeService {
	return &RecipeService{recipeRepo: recipeRepository}
}

func (rs *RecipeService) Recipes() ([]entity.Recipe, []error) {
	rsps, errs := rs.recipeRepo.Recipes()
	if len(errs) > 0 {
		return nil, errs
	}
	return rsps, errs
}
func (rs *RecipeService) Ingredients(id uint) ([]entity.Ingredient, []error) {

	ing, err := rs.recipeRepo.Ingredients(id)

	if len(err) > 0 {
		return nil, err
	}

	return ing, err
}
func (rs *RecipeService) Steps(id uint) ([]entity.Step, []error) {

	steps, err := rs.recipeRepo.Steps(id)

	if len(err) > 0 {
		return nil, err
	}

	return steps, err
}
func (rs *RecipeService) StoreRecipe(recipe *entity.Recipe) (*entity.Recipe, []error) {
	res, errs := rs.recipeRepo.StoreRecipe(recipe)
	if len(errs) > 0 {
		fmt.Println("found errors herer 222........................")
		return nil, errs
	}
	return res, errs
}

// func (rs *RecipeService) updateImage(recipe *entity.Recipe, imagePath string) (*entity.Recipe, []error) {
// 	res, errs := rs.recipeRepo.updateImage(recipe*entity.Recipe, imagePath)
// 	if len(errs) > 0 {
// 		fmt.Println("found errors herer 222........................")
// 		return nil, errs
// 	}
// 	return res, errs
// }
func (rs *RecipeService) Recipe(id uint) (*entity.Recipe, []error) {
	rsp, errs := rs.recipeRepo.Recipe(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return rsp, errs
}

func (rs *RecipeService) UserRecipes(uid uint) ([]entity.Recipe, []error) {
	rsps, errs := rs.recipeRepo.UserRecipes(uid)
	if len(errs) > 0 {
		return nil, errs
	}
	return rsps, errs
}

func (rs *RecipeService) UpdateRecipe(recipe *entity.Recipe) (*entity.Recipe, []error) {
	rsp, errs := rs.recipeRepo.UpdateRecipe(recipe)
	if len(errs) > 0 {
		return nil, errs
	}
	return rsp, errs
}

func (rs *RecipeService) DeleteRecipe(id uint) (*entity.Recipe, []error) {
	rcp, errs := rs.recipeRepo.DeleteRecipe(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return rcp, errs
}
