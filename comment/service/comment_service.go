package service

import (
	"github.com/flutter-amp/baking-api/comment"
	"github.com/flutter-amp/baking-api/entity"
)

type CommentService struct {
	commentRepo comment.CommentRepository
}

func NewCommentService(commRepo comment.CommentRepository) comment.CommentService {
	return &CommentService{commentRepo: commRepo}
}

func (cs *CommentService) Comment(id uint) (*entity.Comment, []error) {
	cmnt, errs := cs.commentRepo.Comment(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}

func (cs *CommentService) Comments() ([]entity.Comment, []error) {
	cmnts, errs := cs.commentRepo.Comments()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnts, errs
}

func (cs *CommentService) StoreComment(comment *entity.Comment) (*entity.Comment, []error) {
	cmnt, errs := cs.commentRepo.StoreComment(comment)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}
func (cs *CommentService) RetrieveComments(recipeId uint) ([]entity.Comment, []error) {
	cmnts, errs := cs.commentRepo.RetrieveComments(recipeId)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnts, errs
}

func (cs *CommentService) UpdateComment(comment *entity.Comment) (*entity.Comment, []error) {
	cmnt, errs := cs.commentRepo.UpdateComment(comment)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}

func (cs *CommentService) DeleteComment(id uint) (*entity.Comment, []error) {
	cmnt, errs := cs.commentRepo.DeleteComment(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}
