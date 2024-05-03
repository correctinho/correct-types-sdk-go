package chain

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/correctinho/correct-mlt-go/qlog"
	"github.com/correctinho/correct-types-sdk-go/err"
	"github.com/gin-gonic/gin"
)

// Response é uma função que define a resposta da requisição.
// Recebe o contexto `ctx`, o status code `statusCode` e o payload `payload` como argumentos.
// O payload é armazenado no contexto com a chave "payload" e o status code é armazenado com a chave "status_code".
func Response(ctx *gin.Context, statusCode int, payload interface{}) {
	ctx.Set("payload", payload)
	ctx.Set("status_code", statusCode)
}

// ResponseWith - Middleware responsável por escrever o response da aplicação
func ResponseWith() Middleware {
	return func(next gin.HandlerFunc) gin.HandlerFunc {
		return func(ctx *gin.Context) {
			next(ctx)
			WriteResponse(ctx)
		}
	}
}

// WriteResponse é uma função que escreve a resposta no contexto da requisição.
// Define o cabeçalho "Content-Type" como "application/json".
// Obtém o status code do contexto com a chave "status_code".
// Se o status code estiver definido, define o status code da resposta como o status code obtido.
// Caso contrário, define o status code da resposta como 500 (Internal Server Error).
// Obtém o erro do contexto com a chave "error".
// Se o erro estiver definido, codifica o erro como JSON e escreve na resposta.
// Caso contrário, obtém o payload do contexto com a chave "payload".
// Se o payload estiver definido e não for um ponteiro nulo, codifica o payload como JSON e escreve na resposta.
func WriteResponse(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	statusCode, ok := ctx.MustGet("status_code").(int)

	if ok {
		ctx.Writer.WriteHeader(statusCode)
	} else {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
	}

	err, ok := ctx.Get("error")
	if ok && err != nil {
		json.NewEncoder(ctx.Writer).Encode(&err)
		return
	}

	val, ok := ctx.Get("payload")
	if !ok {
		return
	}

	if reflect.ValueOf(val).Kind() == reflect.Ptr && reflect.ValueOf(val).IsNil() {
		return
	}

	if val != nil {
		json.NewEncoder(ctx.Writer).Encode(&val)
	}
}

// ResponseError é uma função que define a resposta de erro da requisição.
// Recebe o contexto `ctx` e o erro `err` como argumentos.
// Registra um log de erro usando um logger de produção.
// O status code do erro é obtido usando o método `HTTPCode` do erro.
// A resposta de erro é armazenada no contexto com a chave "error" e o status code é armazenado com a chave "status_code".
func ResponseError(ctx *gin.Context, err *err.CustomError) {
	logger := qlog.NewProduction(ctx)
	defer logger.Sync()

	logger.Error(err.Error())

	ctx.Set("status_code", err.HTTPCode())
	ctx.Set("error", err.Response())
}
