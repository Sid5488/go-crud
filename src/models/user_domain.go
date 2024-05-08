package models

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type UserDomainInterface interface {
	GetEmail() string

	GetPassword() string

	GetName() string

	GetAge() int8

	GetJSONValue() (string, error)

	SetId(string)

	EncryptPassword()
}

func NewUserDomain(
	email, password, name string,
	age int8,
) *userDomain {
	return &userDomain{
		Name:     name,
		Email:    email,
		Password: password,
		Age:      age,
	}
}

type userDomain struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int8   `json:"age"`
}

func (ud *userDomain) GetJSONValue() (string, error) {
	b, err := json.Marshal(ud)

	if err != nil {
		fmt.Println(err)

		return "", nil
	}

	return string(b), nil
}

func (ud *userDomain) SetId(id string) {
	ud.Id = id
}

func (ud *userDomain) GetEmail() string {
	return ud.Email
}

func (ud *userDomain) GetPassword() string {
	return ud.Password
}

func (ud *userDomain) GetName() string {
	return ud.Name
}

func (ud *userDomain) GetAge() int8 {
	return ud.Age
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}
