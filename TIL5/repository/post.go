package repository

import (
	"domain"
	"log"
	"module"

	"cloud.google.com/go/firestore"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/iterator"
)

func SavePost(post domain.Post) (*firestore.DocumentRef, *firestore.WriteResult, error) {
	ref, wr, err := module.GetFirestore().Collection("posts").Add(module.Ctx, post)
	if err != nil {
		log.Printf("error save post: %v\n", err)
	}

	return ref, wr, err
}

func FindPostByID(ID string) ([]domain.Post, error) {
	Posts := []domain.Post{}

	iter := module.GetFirestore().Collection("posts").Where("UID", "==", ID).Documents(module.Ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("error find posts by UID: %v\n", err)
			return Posts, err
		}

		Post := domain.Post{}
		err = mapstructure.Decode(doc.Data(), &Post)
		if err != nil {
			log.Printf("error find posts by UID: %v\n", err)
			return Posts, err
		}

		Posts = append(Posts, Post)
	}

	return Posts, nil
}

func FindPosts() ([]domain.Post, error) {
	Posts := []domain.Post{}

	iter := module.GetFirestore().Collection("posts").Documents(module.Ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("error find posts by UID: %v\n", err)
			return Posts, err
		}

		Post := domain.Post{}
		err = mapstructure.Decode(doc.Data(), &Post)
		if err != nil {
			log.Printf("error find posts by UID: %v\n", err)
			return Posts, err
		}

		Posts = append(Posts, Post)
	}

	return Posts, nil
}

func SetPost(ID string, post domain.Post) (*firestore.WriteResult, error) {
	wr, err := module.GetFirestore().Collection("posts").Doc(ID).Set(module.Ctx, post)
	if err != nil {
		log.Printf("error set post: %v\n", err)
	}

	return wr, err
}

func DelPost(ID string) (*firestore.WriteResult, error) {
	wr, err := module.GetFirestore().Collection("posts").Doc(ID).Delete(module.Ctx)
	if err != nil {
		log.Printf("error delete post: %v\n", err)
	}
	return wr, err
}
