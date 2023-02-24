package service

import (
	"domain"
	"repository"

	"cloud.google.com/go/firestore"
)

func FindPostsByID(ID string) ([]domain.Post, error) {
	return repository.FindPostByID(ID)
}

func JoinPost(Post domain.Post) (*firestore.DocumentRef, *firestore.WriteResult, error) {
	return repository.SavePost(Post)
}

func FindPosts() ([]domain.Post, error) {
	return repository.FindPosts()
}

func ModifyPost(ID string, Post domain.Post) (*firestore.WriteResult, error) {
	return repository.SetPost(ID, Post)
}

func DelPost(ID string) (*firestore.WriteResult, error) {
	return repository.DelPost(ID)
}
