package main

import (
	"net/http"

	"github.com/flutter-amp/baking-api/baking/http/handler"
	"github.com/flutter-amp/baking-api/entity"

	resrep "github.com/flutter-amp/baking-api/recipe/repository"
	resser "github.com/flutter-amp/baking-api/recipe/service"

	//userhand "github.com/flutter-amp/baking-api/recipe/service"
	comrep "github.com/flutter-amp/baking-api/comment/repository"
	comser "github.com/flutter-amp/baking-api/comment/service"

	userrep "github.com/flutter-amp/baking-api/user/repository"
	userser "github.com/flutter-amp/baking-api/user/service"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
)

// func createTables(dbconn *gorm.DB) []error {
// 	//e := dbconn.DropTable(&entity.User{}, &entity.Recipe{}).GetErrors()
// 	// if e != nil {
// 	// 	return e
// 	// }
// 	errs := dbconn.CreateTable(&entity.User{}, &entity.Recipe{}, &entity.Comment{}).GetErrors()
// 	if errs != nil {
// 		return errs
// 	}
// 	return nil
// }

func main() {

	// csrfSignKey := []byte(rtoken.GenerateRandomID(32))
	// tmpl := template.Must(template.ParseGlob("ui/templates/*"))

	dbconn, err := gorm.Open("postgres", "postgres://postgres:admin@localhost/bakingdb?sslmode=disable")
	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	dbconn.AutoMigrate(&entity.Recipe{})
	dbconn.AutoMigrate(&entity.User{})
	dbconn.AutoMigrate(&entity.Comment{})
	dbconn.AutoMigrate(&entity.Ingredient{})
	dbconn.AutoMigrate(&entity.Step{})
	recipeRepo := resrep.NewRecipeGormRepo(dbconn)
	recipeService := resser.NewRecipeService(recipeRepo)
	recipeHandler := handler.NewRecipeHandler(recipeService)

	commentRepo := comrep.NewCommentGormRepo(dbconn)
	commentService := comser.NewCommentService(commentRepo)
	commentHandler := handler.NewCommentHandler(commentService)

	userRepo := userrep.NewUserGormRepo(dbconn)
	userService := userser.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	router := httprouter.New()
	router.GET("/recipes", recipeHandler.GetRecipes)
	router.GET("/ingredients/:id", recipeHandler.GetIngredients)
	router.GET("/steps/:id", recipeHandler.GetSteps)
	router.POST("/recipes/new", recipeHandler.PostRecipe)
	router.POST("/recipes/newImage/:id", recipeHandler.PostImage)
	router.GET("/recipes/get/:id", recipeHandler.GetSingleRecipe)
	router.GET("/recipe/image/:id", recipeHandler.GetSingleRecipe)
	router.DELETE("/recipes/delete/:id", recipeHandler.DeleteRecipe)
	router.PUT("/recipes/update/:id", recipeHandler.PutRecipe)
	router.GET("/user/:uid/recipes", recipeHandler.GetUserRecipes)

	router.GET("/comments", commentHandler.GetComments)
	router.POST("/comments/new", commentHandler.PostComment)
	router.GET("/comments/get/:id", commentHandler.GetSingleComment)
	router.GET("/recipe/comments/:rid", commentHandler.GetCommentsByRecipe)
	router.PUT("/comments/update/:id", commentHandler.PutComment)
	router.DELETE("/comments/delete/:id", commentHandler.DeleteComment)

	router.POST("/signup", userHandler.SignUp)
	router.POST("/login", userHandler.Login)
	router.GET("/users/:id", userHandler.GetSingleUser)
	//http.HandleFunc("/login", uh.Login)
	//	http.HandleFunc("/signup", uh.SignUp)

	http.ListenAndServe(":8181", router)
}

// func configSess() *entity.Session {
// 	tokenExpires := time.Now().Add(time.Minute * 30).Unix()
// 	sessionID := rtoken.GenerateRandomID(32)
// 	signingString, err := rtoken.GenerateRandomString(32)
// 	if err != nil {
// 		panic(err)
// 	}
// 	signingKey := []byte(signingString)

// 	return &entity.Session{
// 		Expires:    tokenExpires,
// 		SigningKey: signingKey,
// 		UUID:       sessionID,
// 	}
// }
