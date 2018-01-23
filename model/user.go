package model

import (
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id           bson.ObjectId
	Username     string
	Email        string
	PasswordHash string
}

func (u *User) GeneratePasswordHash(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hashedPassword)
	return nil
}

func (u User) ValidatePassword(password string) bool {
	hashedPassword := []byte(u.PasswordHash)
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	return err == nil
}
