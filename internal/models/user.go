package models

import (
	"errors"

	db "github.com/cp-Coder/khelo/pkg/platform/database"
	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

// User struct to define the user model
type User struct {
	ExternalID         uuid.UUID `gorm:"column:external_id;type:uuid;default:gen_random_uuid();primarykey" json:"external_id"`
	Username           string    `gorm:"column:username;unique" json:"username" binding:"required"`
	Name               string    `gorm:"column:name" json:"name" binding:"required"`
	Email              string    `gorm:"column:email;unique" json:"email" binding:"required,email"`
	Password           string    `gorm:"column:password" json:"password" binding:"required"`
	Phone              string    `gorm:"column:phone;unique" json:"phone" binding:"required"`
	Gender             string    `gorm:"column:gender" json:"gender" binding:"required"`
	Age                int       `gorm:"column:age" json:"age" binding:"required"`
	DistrictExternalID uuid.UUID `gorm:"column:district_external_id;type:uuid" json:"district_external_id"`
	District           District  `gorm:"foreignkey:DistrictExternalID" json:"district"`
	Deleted            bool      `gorm:"column:deleted;default:false;index" json:"-"`
	UpdatedAt          int64     `gorm:"column:updated_at" json:"-"`
	CreatedAt          int64     `gorm:"column:created_at" json:"-"`
}

// LoginForm struct defines the login form
type LoginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UserModel struct to define methods on user model
type UserModel struct{}

var authModel = &AuthModel{}

// Init migrates the user model to the database
func (m *UserModel) Init() {
	db.GetDBClient().Migrate(&User{})
}

// Login login the user and return the user and token
func (m *UserModel) Login(form *LoginForm) (User, Token, error) {
	var user User
	var token Token
	//	Check if the user exists in database
	err := db.GetDBClient().Find(&user, "username = ?", form.Username).Error
	if err != nil {
		return user, token, errors.New("invalid credentials")
	}

	bytePassword := []byte(form.Password)
	byteHashedPassword := []byte(user.Password)

	err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)

	if err != nil {
		return user, token, err
	}

	//Generate the JWT auth token
	tokenDetails, err := authModel.CreateToken(user.ExternalID)
	if err != nil {
		return user, token, err
	}

	saveErr := authModel.CreateAuth(user.ExternalID, tokenDetails)
	if saveErr == nil {
		token.AccessToken = tokenDetails.AccessToken
		token.RefreshToken = tokenDetails.RefreshToken
	}

	return user, token, nil
}

// Register ...
func (m *UserModel) Register(user *User) error {
	bytePassword := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return errors.New("something went wrong, please try again later")
	}

	user.Password = string(hashedPassword)
	if err := db.GetDBClient().Create(&user); err != nil {
		return errors.New("something went wrong, please try again later")
	}

	return err
}
