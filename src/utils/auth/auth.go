package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
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

// AccessDetails Represents an access with its details
type AccessDetails struct {
	AccessUUID string
	UserID     string
}

// GenerateToken Given an user id as parameter, returns a JWT token
func GenerateToken(userID string) (*TokenDetails, error) {
	env, _ := godotenv.Read()
	accessSecretKey := []byte(env["ACCESS_SECRET"])
	refreshSecretKey := []byte(env["REFRESH_SECRET"])

	td := &TokenDetails{}

	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUUID = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUUID = uuid.NewV4().String()

	var err error

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUUID
	atClaims["user_id"] = userID
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString(accessSecretKey)
	if err != nil {
		return nil, err
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUUID
	rtClaims["user_id"] = userID
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString(refreshSecretKey)
	if err != nil {
		return nil, err
	}

	return td, nil
}

// VerifyToken Given a token, checks if it's valid
func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenToVerify := ExtractToken(r)

	env, _ := godotenv.Read()
	accessSecretKey := []byte(env["ACCESS_SECRET"])

	token, err := jwt.Parse(tokenToVerify, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return accessSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

// TokenIsValid Checks if token in request header is valid
func TokenIsValid(r *http.Request) error {
	token, err := VerifyToken(r)

	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

// ExtractToken Given a request which contains "Authorization" field, extracts the token
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")

	// Normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}

	return ""
}

// ExtractTokenMetadata Given a token string, extracts the metadata contained in it
func ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := VerifyToken(r)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return &AccessDetails{
			AccessUUID: claims["access_uuid"].(string),
			UserID:     claims["user_id"].(string),
		}, nil
	}
	return nil, err
}

// RefreshTokenIsValid Just refresh tokens given a valid refresh token
func RefreshTokenIsValid(refreshToken string) (string, error) {
	// Verify the refresh token
	env, _ := godotenv.Read()
	refreshSecretKey := []byte(env["REFRESH_SECRET"])

	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(refreshSecretKey), nil
	})

	// If there is an error, the token must have expired
	if err != nil {
		return "", err
	}

	// Is token valid?
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return "", err
	}

	// Token is valid, get the UUID
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID := claims["user_id"].(string)

		return userID, nil
	}

	return "", errors.New("Refresh token expired")
}
