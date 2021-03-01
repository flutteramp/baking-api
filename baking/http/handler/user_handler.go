package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/flutter-amp/baking-api/baking/hash"
	Rtoken "github.com/flutter-amp/baking-api/baking/rtoken"
	"github.com/flutter-amp/baking-api/entity"
	"github.com/flutter-amp/baking-api/user"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	// "github.com/flutter-amp/Baking-API/form"
	// "github.com/flutter-amp/Baking-API/model"
	// "github.com/flutter-amp/Baking-API/rtoken"
	// "github.com/flutter-amp/Baking-API/session"
)

type UserHandler struct {
	UserService user.UserService
}

func NewUserHandler(us user.UserService) *UserHandler {
	return &UserHandler{UserService: us}
}

func (uh *UserHandler) GetSingleUser(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	user, errs := uh.UserService.User(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(user, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func (uh *UserHandler) Authenticated(next http.HandlerFunc) http.HandlerFunc {
	// validate the token
	fn := func(w http.ResponseWriter, r *http.Request) {
		// _token := r.Header.Get("Authorization")
		// _token = strings.Replace(_token, "Bearer ", "", 1)
		// // valid, err := uh.tokenService.ValidateToken(_token)
		// // if err != nil && !valid {
		// //   responses.ERROR(w, http.StatusUnauthorized, errors.New("unauthenticated: unauthorized to access the resource, log in again"))
		// //   return
		// // }
		next.ServeHTTP(w, r)
	}
	fmt.Println("authentication")
	return http.HandlerFunc(fn)

}

func (uh *UserHandler) SignUp(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("user handelr")

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	user := &entity.User{}
	fmt.Println("in post user 2")

	err := json.Unmarshal(body, user)
	fmt.Println(user)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	pass, err2 := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err2 != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	user.Password = string(pass)
	user, errs := uh.UserService.StoreUser(user)

	if len(errs) > 0 {
		fmt.Println("HEEEEEEEEEEEEEEEE")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, _ := json.MarshalIndent(user, "", "\t\t")
	// p := fmt.Sprintf("/api/recipe/%d", recipe.ID)
	// w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	w.Write(output)
	return

}
func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("user handelr")

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	user := &entity.User{}
	fmt.Println("in post user 2")

	err := json.Unmarshal(body, user)
	fmt.Println(user)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	user1, errs := uh.UserService.UserByEmail(user.Email)
	fmt.Println("sencond fffffffffffffff")
	fmt.Println(user1)
	if len(errs) > 0 || !hash.ArePasswordsSame(user1.Password, user.Password) {

		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

		return
	}
	fmt.Println("THE ID ISSSSSSSSSSSS")
	fmt.Println(user.ID)
	tokenString, err := Rtoken.GenerateJwtToken([]byte(Rtoken.GenerateRandomID(32)), Rtoken.CustomClaims{

		SessionId: strconv.Itoa(int(user1.ID)),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 1, 1).Unix(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
		},
	})

	output, _ := json.MarshalIndent(struct {
		Token string `json:"token"`
	}{
		Token: tokenString,
	}, "", "\t\t")
	// p := fmt.Sprintf("/api/recipe/%d", recipe.ID)
	// w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	w.Write(output)
	return
}

func (uh *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := uh.UserService.DeleteUser(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}

func (uh *UserHandler) PutUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	user, errs := uh.UserService.User(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	json.Unmarshal(body, &user)
	user, errs = uh.UserService.UpdateUser(user)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(user, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
