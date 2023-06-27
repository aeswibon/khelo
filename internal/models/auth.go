package models

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/cp-Coder/khelo/pkg/platform/cache"
	db "github.com/cp-Coder/khelo/pkg/platform/database"
	jwt "github.com/golang-jwt/jwt/v4"
	uuid "github.com/google/uuid"
)

// TokenDetails struct definition for token details
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   uuid.UUID
	RefreshUUID  uuid.UUID
	AtExpires    int64
	RtExpires    int64
}

// AccessDetails ...
type AccessDetails struct {
	AccessUUID uuid.UUID
	UserID     uuid.UUID
}

// Token struct definition for token response
type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// AuthModel ...
type AuthModel struct{}

// Init method to migrate the auth model schema to the database
func (m *AuthModel) Init() {
	db.GetDBClient().Migrate(&TokenDetails{}, &AccessDetails{}, &Token{})
}

// CreateToken ...
func (m *AuthModel) CreateToken(userID uuid.UUID) (*TokenDetails, error) {

	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUUID = uuid.New()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUUID = uuid.New()

	var err error
	// Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUUID
	atClaims["user_id"] = userID
	atClaims["exp"] = td.AtExpires

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}
	// Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUUID
	rtClaims["user_id"] = userID
	rtClaims["exp"] = td.RtExpires

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return td, nil
}

// CreateAuth ...
func (m *AuthModel) CreateAuth(userid uuid.UUID, td *TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) // converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	// set the access token in redis
	client := cache.GetRedisClient()
	errAccess := client.Setx(td.AccessUUID.String(), userid, at.Sub(now))
	if errAccess != nil {
		return errAccess
	}
	// set the refresh token in redis
	errRefresh := client.Setx(td.RefreshUUID.String(), userid, rt.Sub(now))
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

// ExtractToken ...
func (m *AuthModel) ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	// normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// VerifyToken ...
func (m *AuthModel) VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := m.ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// TokenValid ...
func (m *AuthModel) TokenValid(r *http.Request) error {
	token, err := m.VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

// ExtractTokenMetadata ...
func (m *AuthModel) ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := m.VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(uuid.UUID)
		if !ok {
			return nil, err
		}
		userID := claims["user_id"].(uuid.UUID)
		return &AccessDetails{
			AccessUUID: accessUUID,
			UserID:     userID,
		}, nil
	}
	return nil, err
}

// FetchAuth ...
func (m *AuthModel) FetchAuth(authD *AccessDetails) (uuid.UUID, error) {
	userid, err := cache.GetRedisClient().Get(authD.AccessUUID.String())
	if err != nil {
		return uuid.Nil, err
	}
	userID, _ := uuid.Parse(userid)
	return userID, nil
}

// DeleteAuth ...
func (m *AuthModel) DeleteAuth(givenUUID uuid.UUID) error {
	err := cache.GetRedisClient().Del(givenUUID.String())
	if err != nil {
		return err
	}
	return nil
}
