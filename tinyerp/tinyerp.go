package tinyerp

import (
	"fmt"
	"log"
	"net/url"
	"os"
)

// Processing Status
const (
	NotProcesssed = iota
	ProcessedWithErrors
	ProcessedOk
)

var statusText = map[int]string{
	NotProcesssed:       "Solicitação não processada",
	ProcessedWithErrors: "Solicitação processada, mas possui erros de validação",
	ProcessedOk:         "Solicitação processada corretamente",
}

// StatusText returns a text for the tiny processing status. It returns the empty
// string if the code is unknown.
func StatusText(code int) string {
	return statusText[code]
}

// Error codes
const (
	MissingToken = iota
	InvalidToken
	XMLError
	XMLProcessError
	NotAccessOrBlocked
	TemporaryBlocked
	NotEnoughSpace
	BlockedBusiness
	DuplicatedSequenceNumbers
	MissingParameter
	EmptyQueryReturn
	ExcessReturn
	XMLRecordsExceededBatch
	PageNotFound
	DuplicatedResourceError
	ValidationError
	ResourceNotFound
	DuplicatedResource
	MaintenanceSystem = 99
)

var errorCode = map[int]string{
	MissingToken:              "token não informado",
	InvalidToken:              "token inválido ou não encontrado",
	XMLError:                  "XML mal formado ou com erros",
	XMLProcessError:           "Erro de procesamento de XML",
	NotAccessOrBlocked:        "API bloqueada ou sem acesso",
	TemporaryBlocked:          "API bloqueada momentaneamente - muitos acessos no último minuto",
	NotEnoughSpace:            "Espaço da empresa Esgotado",
	BlockedBusiness:           "Empresa Bloqueada",
	DuplicatedSequenceNumbers: "Números de sequencia em duplicidade",
	MissingParameter:          "Parametro não informado",
	EmptyQueryReturn:          "A Consulta não retornou registros",
	ExcessReturn:              "A Consulta retornou muitos registros",
	XMLRecordsExceededBatch:   "O xml tem mais registros que o permitido por lote de envio",
	PageNotFound:              "A página que você está tentanto obter não existe",
	DuplicatedResourceError:   "Erro de Duplicidade de Registro",
	ValidationError:           "Erros de Validação",
	ResourceNotFound:          "Registro não localizado",
	DuplicatedResource:        "Registro localizado em duplicidade",
	MaintenanceSystem:         "Sistema em manutenção",
}

// ErrorCode returns a text for the tiny errors code. It returns the empty
// string if the code is unknown.
func ErrorCode(code int) string {
	return errorCode[code]
}

type TinyERP struct {
	baseURL string
	token   string
	format  string
}

func NewTinyERP() *TinyERP {
	token := mustEnv("TINY_TOKEN_API")
	return &TinyERP{
		baseURL: "https://api.tiny.com.br/api2",
		token:   token,
		format:  "json",
	}
}

func mustEnv(key string) (value string) {
	if value = os.Getenv(key); value == "" {
		log.Fatalf("ENV %q is not set.", key)
	}
	return value
}

func (t *TinyERP) apiURI(route string) (string, url.Values) {
	v := url.Values{}
	v.Set("token", t.token)
	v.Set("formato", t.format)
	apiURI := fmt.Sprintf("%s/%s", t.baseURL, route)
	return apiURI, v
}
