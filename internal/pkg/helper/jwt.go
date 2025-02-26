package helper

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/pebruwantoro/movie-festival-backend/config"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
	"time"
)

type Claims struct {
	UUID      string    `json:"uuid"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	jwt.RegisteredClaims
}

func GenerateJWT(user entities.User) (string, error) {
	expirationTime := time.Now().Add(time.Duration(config.Get().JwtExpiredTime) * time.Hour)

	claims := &Claims{
		UUID:      user.UUID,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.Get().AppName,
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(config.Get().JwtSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
