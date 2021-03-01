package main

import (
	"net/http"

	"github.com/flutteramp/baking-api/baking/http/handler"
	"github.com/flutteramp/baking-api/entity"
	"github.com/gorilla/mux"

	resrep "github.com/flutteramp/baking-api/recipe/repository"
	resser "github.com/flutteramp/baking-api/recipe/service"

	//userhand "github.com/flutter-amp/baking-api/recipe/service"
	comrep "github.com/flutteramp/baking-api/comment/repository"
	comser "github.com/flutteramp/baking-api/comment/service"

	userrep "github.com/flutteramp/baking-api/user/repository"
	userser "github.com/flutteramp/baking-api/user/service"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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

	router1 := mux.NewRouter()
	//router := httprouter.New()

	router1.HandleFunc("/recipes", userHandler.Authenticated(recipeHandler.GetRecipes)).Methods("GET")
	router1.HandleFunc("/recipes/new", userHandler.Authenticated(recipeHandler.PostRecipe)).Methods("POST")
	router1.HandleFunc("/recipes/{id}/ingredients", userHandler.Authenticated(recipeHandler.GetIngredients)).Methods("GET")
	router1.HandleFunc("/recipes/{id}/steps", userHandler.Authenticated(recipeHandler.GetSteps)).Methods("GET")
	router1.HandleFunc("/recipe/{id}", userHandler.Authenticated(recipeHandler.GetSingleRecipe)).Methods("GET")
	router1.HandleFunc("/recipes/delete/{id}", userHandler.Authenticated(recipeHandler.DeleteRecipe)).Methods("DELETE")
	router1.HandleFunc("/recipes/update/{id}", userHandler.Authenticated(recipeHandler.PutRecipe)).Methods("PUT")
	router1.HandleFunc("/user/{uid}/recipes", userHandler.Authenticated(recipeHandler.GetUserRecipes)).Methods("GET")

	router1.HandleFunc("/comments", userHandler.Authenticated(commentHandler.GetComments)).Methods("GET")
	router1.HandleFunc("/comments/new", userHandler.Authenticated(commentHandler.PostComment)).Methods("POST")
	router1.HandleFunc("/comments/get/{id}", userHandler.Authenticated(commentHandler.GetSingleComment)).Methods("GET")
	router1.HandleFunc("/recipe/comments/{rid}", userHandler.Authenticated(commentHandler.GetCommentsByRecipe)).Methods("GET")
	router1.HandleFunc("/comments/update/{id}", userHandler.Authenticated(commentHandler.PutComment)).Methods("PUT")
	router1.HandleFunc("/comments/delete/{id}", userHandler.Authenticated(commentHandler.DeleteComment)).Methods("DELETE")
	router1.HandleFunc("/users/{id}", userHandler.Authenticated(userHandler.GetSingleUser)).Methods("GET")

	router1.HandleFunc("/login", userHandler.Login).Methods("POST")
	router1.HandleFunc("/signup", userHandler.SignUp).Methods("POST")
	http.ListenAndServe(":8181", router1)

	// router.GET("/recipes", recipeHandler.GetRecipes)
	//router.GET("/recipe/image/:id", recipeHandler.GetSingleRecipe)
	// router.DELETE("/recipes/delete/:id", recipeHandler.DeleteRecipe)
	//router.PUT("/recipes/update/:id", recipeHandler.PutRecipe)
	//router.GET("/user/:uid/recipes", recipeHandler.GetUserRecipes)
	//	router.GET("recipes/{id}/ingredients/", recipeHandler.GetIngredients)
	// router.POST("/recipes/newImage/:id", recipeHandler.PostImage)

	//	router.GET("/comments", commentHandler.GetComments)
	//router.POST("/comments/new", commentHandler.PostComment)
	//router.GET("/comments/get/:id", commentHandler.GetSingleComment)
	//router.GET("/recipe/comments/:rid", commentHandler.GetCommentsByRecipe)
	//router.PUT("/comments/update/:id", commentHandler.PutComment)
	//router.DELETE("/comments/delete/:id", commentHandler.DeleteComment)
	// router.POST("/signup", userHandler.SignUp)
	// router.POST("/login", userHandler.Login)
	//router.GET("/users/:id", userHandler.GetSingleUser)
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
