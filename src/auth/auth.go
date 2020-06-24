package auth

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

// TokenDetails Represents a token
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}

var accessSecretKey = []byte("our_secret-key_is-very_strongue")
var refreshSecretKey = []byte("our_secret-key_is-very_strongue_ohlalala")

// GenerateToken Given an username as parameter, returns a JWT token
func GenerateToken(salt string) (*TokenDetails, error) {
	td := &TokenDetails{}

	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUUID = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUUID = uuid.NewV4().String()

	var err error

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUUID
	atClaims["salt"] = salt
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString(accessSecretKey)
	if err != nil {
		return nil, err
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUUID
	rtClaims["salt"] = salt
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString(refreshSecretKey)
	if err != nil {
		return nil, err
	}

	return td, nil
}

// VerifyToken Given a token, checks if it's valid
func VerifyToken(tokenToVerify string) bool {
	token, err := jwt.Parse(tokenToVerify, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return accessSecretKey, nil
	})

	if err != nil {
		return false
	}

	return token.Valid
}
