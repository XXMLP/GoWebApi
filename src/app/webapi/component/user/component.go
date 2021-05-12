package user

import (
	"app/webapi/component"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

// New returns a new instance of the endpoint.
func New(bc component.Core) *Endpoint {
	return &Endpoint{
		Core: bc,
	}
}

// Endpoint contains the dependencies.
type Endpoint struct {
	component.Core
}

// Routes will set up the endpoints.
func (p *Endpoint) Routes(router component.IRouter) {
	router.Post("/v1/user", p.Create)
	router.Get("/v1/user/:user_id", p.Show)
	router.Get("/v1/userSelf", p.ShowSelf)
	router.Get("/v1/user", p.Index)
	router.Put("/v1/user/:user_id", p.Update)
	router.Delete("/v1/user/:user_id", p.Destroy)
	router.Delete("/v1/user", p.DestroyAll)
}

//ParseToken to userId
func ParseToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("jOJtWzbV6qiIuuC60iAnYTbnx7tz8NvxEdEWuaDw31Q="), nil
	})
	if claims, ok := token.Claims.(*jwt.StandardClaims); ok {
		return claims.Audience, nil
	} else {
		return "", err
	}
}
