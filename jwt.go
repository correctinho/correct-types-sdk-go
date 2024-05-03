package types

import "github.com/dgrijalva/jwt-go"

// JwtToken representa um token JWT.
type JwtToken struct {
	*jwt.StandardClaims             // Contém os campos padrão do token JWT, como expiração, emissor, etc.
	Data                interface{} `json:"data,omitempty"`   // Dados adicionais que podem ser incluídos no token.
	Extras              interface{} `json:"extras,omitempty"` // Informações extras que podem ser incluídas no token.
}
