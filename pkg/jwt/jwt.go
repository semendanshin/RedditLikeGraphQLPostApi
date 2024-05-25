package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// Generator is a JWT token generator and parser
type Generator struct {
	secret     []byte
	accessTTL  time.Duration
	refreshTTL time.Duration
}

// NewGenerator creates a new Generator
func NewGenerator(secret string, accessTTL time.Duration, refreshTTL time.Duration) *Generator {
	return &Generator{
		secret:     []byte(secret),
		accessTTL:  accessTTL,
		refreshTTL: refreshTTL,
	}
}

// GeneratePair generates a pair of access and refresh tokens
func (g *Generator) GeneratePair(sub string) (accessToken string, refreshToken string, err error) {
	iat := time.Now().Unix()

	accessClaims := jwt.MapClaims{
		"sub": sub,
		"exp": time.Now().Add(g.accessTTL).Unix(),
		"iat": iat,
	}

	refreshClaims := jwt.MapClaims{
		"sub":  sub,
		"exp":  time.Now().Add(g.refreshTTL).Unix(),
		"iat":  iat,
		"type": "refresh",
	}

	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(g.secret)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(g.secret)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// ParseToken parses a token string
func (g *Generator) ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return g.secret, nil
	})
}

// EnsurePair ensures that the access and refresh tokens are a pair
func (g *Generator) EnsurePair(accessToken *jwt.Token, refreshToken *jwt.Token) error {
	accessTokenClaims := accessToken.Claims.(jwt.MapClaims)
	refreshTokenClaims := refreshToken.Claims.(jwt.MapClaims)

	if accessTokenClaims["sub"] != refreshTokenClaims["sub"] || accessTokenClaims["iat"] != refreshTokenClaims["iat"] {
		return jwt.ErrInvalidKey
	}

	if refreshTokenClaims["type"] != "refresh" {
		return jwt.ErrInvalidKey
	}

	return nil
}
