package chain

import "github.com/gin-gonic/gin"

// Middleware - Middleware da aplicação
type Middleware func(gin.HandlerFunc) gin.HandlerFunc
