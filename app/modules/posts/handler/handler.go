package handler

import (
	post "github.com/vkhoa145/facebook-mini-api/app/modules/posts/usecase"
)

type PostHandler struct {
	postUsecase post.PostUseCaseInterface
}

func NewPostHandler(postUsecase post.PostUseCaseInterface) *PostHandler {
	return &PostHandler{postUsecase: postUsecase}
}