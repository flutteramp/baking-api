package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/flutteramp/baking-api/comment"
	"github.com/flutteramp/baking-api/entity"
	"github.com/gorilla/mux"
)

type CommentHandler struct {
	commentService comment.CommentService
}

func NewCommentHandler(cmtService comment.CommentService) *CommentHandler {
	fmt.Println("comment handler created")
	return &CommentHandler{commentService: cmtService}
}

func (ch *CommentHandler) GetSingleComment(w http.ResponseWriter,
	r *http.Request) {

	params := mux.Vars(r)
	idParam, exists := params["id"]
	if !exists {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	id, err := strconv.Atoi(idParam)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	comment, errs := ch.commentService.Comment(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(comment, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

//Gets all halls
func (ch *CommentHandler) GetComments(w http.ResponseWriter,
	r *http.Request) {

	comments, errs := ch.commentService.Comments()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(comments, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

func (ch *CommentHandler) PostComment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("comment handelr")

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	comment := &entity.Comment{}
	fmt.Println("in post comment 2")

	err := json.Unmarshal(body, comment)
	fmt.Println(comment)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	comment, errs := ch.commentService.StoreComment(comment)

	if len(errs) > 0 {
		fmt.Println("HEEEEEEEEEEEEEEEE")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// p := fmt.Sprintf("/api/recipe/%d", recipe.ID)
	// w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}

func (ch *CommentHandler) GetCommentsByRecipe(w http.ResponseWriter,
	r *http.Request) {
	params := mux.Vars(r)
	idParam, exists := params["rid"]
	if !exists {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	rid, err := strconv.Atoi(idParam)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	comments, errs := ch.commentService.RetrieveComments(uint(rid))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(comments, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

func (ch *CommentHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	idParam, exists := params["id"]
	if !exists {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	id, err := strconv.Atoi(idParam)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := ch.commentService.DeleteComment(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}

func (ch *CommentHandler) PutComment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam, exists := params["id"]
	if !exists {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	comment, errs := ch.commentService.Comment(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	json.Unmarshal(body, &comment)
	comment, errs = ch.commentService.UpdateComment(comment)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(comment, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
