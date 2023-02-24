package service

import (
	"domain"
	"errors"
	"repository"

	"cloud.google.com/go/firestore"
)

func FindUser(ID string) ([]domain.Account, error) {
	return repository.FindUserByID(ID)
}

func JoinUser(User domain.Account) (*firestore.DocumentRef, *firestore.WriteResult, error) {
	return repository.SaveUser(User)
}

func FindUserByIDAndPassword(user domain.Account) (bool, error) {
	users, err := repository.FindUserByID(user.ID)

	if err != nil {
		return false, err
	}

	if len(users) == 0 {
		return false, errors.New("no users")
	}

	if users[0].Password != user.Password {
		return false, errors.New("not match password")
	}

	return true, nil
}
