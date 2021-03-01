package comment

import "github.com/flutter-amp/baking-api/entity"

type CommentService interface {
	Comments() ([]entity.Comment, []error)
	Comment(id uint) (*entity.Comment, []error)
	UpdateComment(comment *entity.Comment) (*entity.Comment, []error)
	DeleteComment(id uint) (*entity.Comment, []error)
	StoreComment(comment *entity.Comment) (*entity.Comment, []error)
	RetrieveComments(recipeId uint) ([]entity.Comment, []error)
}
