package repository

import (
	"domain"
	"log"
	"module"

	"cloud.google.com/go/firestore"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/iterator"
)

func SaveUser(user domain.Account) (*firestore.DocumentRef, *firestore.WriteResult, error) {
	ref, wr, err := module.GetFirestore().Collection("users").Add(module.Ctx, user)
	if err != nil {
		log.Printf("error save user: %v\n", err)
	}

	return ref, wr, err
}

func FindUserByUID(UID string) (domain.Account, error) {
	user := domain.Account{}

	dsnap, err := module.GetFirestore().Collection("users").Doc(UID).Get(module.Ctx)
	if err != nil {
		log.Printf("error find user by id: %v\n", err)
		return domain.Account{}, err
	}

	err = mapstructure.Decode(dsnap.Data(), &user)
	if err != nil {
		log.Printf("error find user by id: %v\n", err)
		return domain.Account{}, err
	}

	return user, err
}

func FindUserByID(ID string) ([]domain.Account, error) {
	accounts := []domain.Account{}

	iter := module.GetFirestore().Collection("users").Where("ID", "==", ID).Documents(module.Ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("error find users by UID: %v\n", err)
			return accounts, err
		}

		account := domain.Account{}
		err = mapstructure.Decode(doc.Data(), &account)
		if err != nil {
			log.Printf("error find users by UID: %v\n", err)
			return accounts, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func SetUser(ID string, user domain.Account) (*firestore.WriteResult, error) {
	wr, err := module.GetFirestore().Collection("users").Doc(ID).Set(module.Ctx, user)
	if err != nil {
		log.Printf("error set user: %v\n", err)
	}

	return wr, err
}

func DelUser(ID string) (*firestore.WriteResult, error) {
	wr, err := module.GetFirestore().Collection("users").Doc(ID).Delete(module.Ctx)
	if err != nil {
		log.Printf("error delete user: %v\n", err)
	}
	return wr, err
}
