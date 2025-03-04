package entity

import (
	"crud_fiber.com/m/config"
	"crud_fiber.com/m/pkg/errs"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strings"
	"time"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedBy string    `gorm:"default:null" json:"updated_by"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
}

type CustomClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func (*User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}

func (u *User) HashPassword() {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hashedPassword)
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) GenerateJwtToken() (string, error) {
	claims := CustomClaims{
		Email: u.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "crud_fiber.com",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(config.GetConfig().AuthSecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (u *User) ValidateToken(bearerToken string) error {
	isBearer := strings.HasPrefix(bearerToken, "Bearer")

	errInvalidToken := &errs.UnauthorizedError{
		Err: "Token Invalid",
	}

	if !isBearer {
		return errInvalidToken
	}

	splitToken := strings.Split(bearerToken, " ")

	if len(splitToken) != 2 {
		return errInvalidToken
	}

	tokenString := splitToken[1]

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errInvalidToken
		}
		return []byte(config.GetConfig().AuthSecretKey), nil
	})

	if err != nil {
		return err
	}

	var mapClaims jwt.MapClaims

	if claims, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return errInvalidToken
	} else {
		mapClaims = claims
	}

	if _, ok := mapClaims["email"]; !ok {
		return errInvalidToken
	}

	u.Email = mapClaims["email"].(string)

	return nil
}
