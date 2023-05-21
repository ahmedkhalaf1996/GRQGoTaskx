// //go:generate go run github.com/99designs/gqlgen generate

// package graph

// // This file will not be regenerated automatically.
// //
// // It serves as dependency injection for your app, add any dependencies you require here.

// type Resolver struct{}

//go:generate go run github.com/99designs/gqlgen generate
package graph

import (
	"fmt"
	"os"
	"time"

	"github.com/ahmedkhalaf1996/GRQGoTaskx/graph/model"
	"github.com/ahmedkhalaf1996/GRQGoTaskx/models"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{ DB *gorm.DB }

func NewResolver(DB *gorm.DB) *Resolver {
	return &Resolver{DB}
}

// func GetUserByID(id string) (*User, error) {
func (r *Resolver) GetUserByID(id string) (*models.User, error) {
	// var db *gorm.DB

	var user models.User
	err := r.DB.First(&user, "id = ?", id).Error
	return &user, err

}

func GenToken(u *models.User) (*model.AuthToken, error) {
	expiredAt := time.Now().Add(time.Hour * 24 * 7) // a week

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: expiredAt.Unix(),
		Id:        fmt.Sprint(u.ID),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "meetmeup",
	})

	accessToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	return &model.AuthToken{
		AccessToken: accessToken,
		ExpiredAt:   expiredAt,
	}, nil
}
