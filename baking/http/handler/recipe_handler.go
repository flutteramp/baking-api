package handler

import (
	"bufio"
	"encoding/json"
	"fmt"
	"image"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/flutter-amp/baking-api/entity"
	"github.com/flutter-amp/baking-api/recipe"

	"github.com/julienschmidt/httprouter"
)

type RecipeHandler struct {
	recipeService recipe.RecipeService
}

func NewRecipeHandler(rspService recipe.RecipeService) *RecipeHandler {
	fmt.Println("recipe handler created")
	return &RecipeHandler{recipeService: rspService}
}

func (rh *RecipeHandler) GetRecipes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	recipes, errs := rh.recipeService.Recipes()
	fmt.Println("jummmmpppp")
	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(recipes, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}
func (rh *RecipeHandler) GetIngredients(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, errr := strconv.Atoi(ps.ByName("id"))

	if errr != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	ingredients, errs := rh.recipeService.Ingredients(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(ingredients, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

func (rh *RecipeHandler) GetSteps(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, errr := strconv.Atoi(ps.ByName("id"))

	if errr != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	steps, errs := rh.recipeService.Steps(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(steps, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

func (rh *RecipeHandler) GetUserRecipes(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {
	uid, err := strconv.Atoi(ps.ByName("uid"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	recipes, errs := rh.recipeService.UserRecipes(uint(uid))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(recipes, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

func (rh *RecipeHandler) PostRecipe(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("recipe handelr")

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	recipe := &entity.Recipe{}
	fmt.Println("in post recipe 2")
	err := json.Unmarshal(body, recipe)
	// for i := 0; i < recipe..length; i++ {
	// 	fmt.Println("gooooooooooooooooooooooooo")
	// }
	fmt.Println(recipe)

	if err != nil {
		fmt.Println("HEEEEEEEEEEE222EEEEEE")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	recipe, errs := rh.recipeService.StoreRecipe(recipe)
	fmt.Println("my recipeee")
	fmt.Println(recipe)
	if len(errs) > 0 {
		//fmt.Println("HEEEEEEEEEEEEEEEE")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(recipe, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	p := fmt.Sprintf("/recipes/add/%d", recipe.ID)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	w.Write(output)
	return
}

//post image of recipe
func (rh *RecipeHandler) PostImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("image here")
	r.ParseForm()
	fmt.Println("steppp1")
	file, handler, err := r.FormFile("file")
	fmt.Println("steppp2")
	rid := ps.ByName("id")

	if err != nil {
		fmt.Println("HEEEEEEEEEEE222EEEEEE")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	id, _ := strconv.Atoi(rid)
	if err != nil {
		fmt.Println("conversion")
	}
	recipet, errs := rh.recipeService.Recipe(uint(id))
	if errs != nil {
		fmt.Println(errs)
	}

	fmt.Println(recipet)

	dst, err := os.Create(filepath.Join("./images/", filepath.Base(rid+""+handler.Filename)))
	recipet.ImageUrl = filepath.Base(rid + "" + handler.Filename)
	fmt.Println(recipet)
	recip, errs := rh.recipeService.UpdateRecipe(recipet)
	fmt.Println(recip)
	if len(errs) > 0 {
		fmt.Println("HEEEEEEEEEEEEEEEE")
	}
	defer dst.Close()
	if _, err = io.Copy(dst, file); err != nil {
		fmt.Println("erorrrrrrrrrrrrrrrrr")
		return
	}
	// recipe, errs := rh.recipeService.updateImage(recipet, dst)

	// if len(errs) > 0 {
	// 	//fmt.Println("HEEEEEEEEEEEEEEEE")
	// 	w.Header().Set("Content-Type", "application/json")
	// 	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	// 	return
	// }

	// output, err := json.MarshalIndent(recipe, "", "\t\t")

	// if err != nil {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	// 	return
	// }

	// p := fmt.Sprintf("/recipes/add/%d", recipe.ID)
	// w.Header().Set("Location", p)
	// w.WriteHeader(http.StatusCreated)
	// w.Write(output)
	// return
	return
} // GetSinglerecipe handles
func (rh *RecipeHandler) GetSingleRecipe(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	recipe, errs := rh.recipeService.Recipe(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	file, _ := os.Open("./images/1image_picker4453127669479362428.jpg")
	defer file.Close()
	rdr := bufio.NewReader(file)
	bts, _ := rdr.Peek(512)
	contentType := http.DetectContentType(bts)
	fmt.Println(contentType)
	img, _, errt := image.Decode(rdr)
	if errt != nil {
		fmt.Println("error ")

	}
	fmt.Println("done")
	// existingImageFile, err := os.Open("./images/1image_picker4453127669479362428.jpg")
	// if err != nil {
	// 	fmt.Println("something happened")
	// }
	// existingImageFile.Seek(0, 0)
	// defer existingImageFile.Close()
	// image, _, err := image.Decode(existingImageFile)
	// if errs != nil {
	// 	fmt.Println(errs)
	// }

	// // Calling the generic image.Decode() will tell give us the data
	// // and type of image it is as a string. We expect "png"
	// //imageData, _, err := image.Decode(existingImageFile)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	fmt.Println(img)

	// // We only need this because we already read from the file
	// // We have to reset the file pointer back to beginning
	// existingImageFile.Seek(0, 0)

	// // Alternatively, since we know it is a png already
	// // we can call png.Decode() directly

	// fmt.Println(image)
	output, err := json.MarshalIndent(recipe, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
func (rh *RecipeHandler) GetImageRecipe(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	recipe, errs := rh.recipeService.Recipe(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(recipe, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func (rh *RecipeHandler) DeleteRecipe(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := rh.recipeService.DeleteRecipe(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}

func (rh *RecipeHandler) PutRecipe(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	recipe, errs := rh.recipeService.Recipe(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	json.Unmarshal(body, &recipe)
	recipe, errs = rh.recipeService.UpdateRecipe(recipe)
	fmt.Println(recipe)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(recipe, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
