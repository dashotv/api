package models

import (
	"crypto/md5"
	"fmt"
	"io"

	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	bongo.DocumentBase `bson:",inline"`
	Name               string
	Email              string
	Password           string `bson:",omitempty"`
	PasswordHash       string
	Token              string
}

func UserFind(email string) (*User, error) {
	user := &User{}

	err := DB.Users.FindOne(bson.M{"email": email}, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) Save() error {
	u.PasswordHash = cryptPassword(u.Password)
	u.Password = "" // clear password, so we don't store in plain text
	return DB.Users.Save(u)
}

func (u *User) CheckPassword(password string) bool {
	crypt := cryptPassword(password)
	return u.PasswordHash == crypt
}

func cryptPassword(password string) string {
	h := md5.New()
	io.WriteString(h, password)

	pwmd5 := fmt.Sprintf("%x", h.Sum(nil))

	// Specify two salt: salt1 = @#$% salt2 = ^&*()
	salt1 := "@#$%"
	salt2 := "^&*()"

	// salt1 + username + salt2 + MD5 splicing
	io.WriteString(h, salt1)
	io.WriteString(h, "abc")
	io.WriteString(h, salt2)
	io.WriteString(h, pwmd5)

	return fmt.Sprintf("%x", h.Sum(nil))
}
