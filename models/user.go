package models

import (
	"time"

	"github.com/pkg/errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Name     string
	Email    string
	Address  string
	Password string
	Phone    string
}

func (s *Store) CreateUser(user User) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = sq.Insert("users").
		Columns("name", "email", "address", "password", "phone_number").
		Values(user.Name, user.Email, user.Address, hashedPassword, user.Phone).
		RunWith(s.DB).
		PlaceholderFormat(sq.Dollar).Query()
	if err != nil {
		return err
	}
	return nil
}

type Creds struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func (s *Store) AuthUser(creds Creds) (tokenString string, err error) {
	var (
		userID         uint
		userName       string
		hashedPassword string
	)

	query := sq.Select("id", "name", "password").
		From("users").
		Where(sq.Eq{"emal": creds.Login}).
		RunWith(s.DB).
		PlaceholderFormat(sq.Dollar)

	err = query.QueryRow().Scan(&userID, &userName, &hashedPassword)
	if err != nil {
		return "", errors.Wrap(err, "finding recoder in db")
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(creds.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return "", errors.Wrap(err, "invalid login credentials")
		}
		return "", errors.Wrap(err, "comparing the given password with the hashed one")
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		UserID:   userID,
		Username: userName,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err = token.SignedString([]byte("my_secret_key"))
	if err != nil {
		return "", errors.Wrap(err, "creating tokenString")
	}

	return tokenString, nil
}
