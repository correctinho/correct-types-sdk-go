package chain

import "github.com/gin-gonic/gin"

// Use - Função que adicina middlewares antes da api ser executada
func Use(handler gin.HandlerFunc, midds ...Middleware) gin.HandlerFunc {
	adapters := []Middleware{}
	adapters = append(adapters, midds...)
	adapters = append(adapters, ResponseWith())

	for i := len(adapters); i > 0; i-- {
		handler = adapters[i-1](handler)
	}

	return handler
}
