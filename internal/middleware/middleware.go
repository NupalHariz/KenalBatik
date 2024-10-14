package middleware

import "kenalbatik-be/internal/infra/jwt"

type Middleware struct {
	jwt jwt.JWTInterface
}

func NewMiddleware(jwt jwt.JWTInterface) *Middleware {
	return &Middleware{
		jwt: jwt,
	}
}
