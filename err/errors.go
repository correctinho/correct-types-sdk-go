package err

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// CustomError - tipo de error
type CustomError struct {
	Message  error
	Code     string
	httpCode int
	Args     []interface{}
}

// CustomErrorResponse - Representa uma resposta de erro personalizada.
type CustomErrorResponse struct {
	Message string `json:"message"` // O campo Message contém a mensagem de erro.
	Code    string `json:"code"`    // O campo Code contém o código de erro.
}

// Error implements error
func (e CustomError) Error() string {
	var r strings.Builder
	if len(e.Args) > 0 {
		e.Message = fmt.Errorf(e.Message.Error(), e.Args...)
	}
	r.WriteString(e.Code + ": " + e.Message.Error() + ".")
	return r.String()
}

// HTTPCode - Retorna o código HTTP associado ao erro.
func (e CustomError) HTTPCode() int {
	return e.httpCode
}

// Response - Retorna uma resposta personalizada de erro.
func (e CustomError) Response() CustomErrorResponse {
	var r CustomErrorResponse
	r.Message = e.Message.Error()
	r.Code = e.Code
	return r
}

// Mensagens de erro
var (
	ErrInternalService = CustomError{
		Message:  errors.New("Serviço indisponível no momento. Por favor, tente novamente em alguns instantes"),
		Code:     "00000",
		httpCode: http.StatusServiceUnavailable,
	}

	ErrJwtEnvNotfound = CustomError{
		Message:  errors.New("Configuração de URL do JWT não encontrada"),
		Code:     "00001",
		httpCode: http.StatusNotFound,
	}

	ErrRequiredData = CustomError{
		Message:  errors.New("Dados obrigatórios"),
		Code:     "00002",
		httpCode: http.StatusBadRequest,
	}

	ErrInvalidAuthenticationToken = CustomError{
		Message:  errors.New("Token de autenticação inválido"),
		Code:     "00003",
		httpCode: http.StatusUnauthorized,
	}
)
