package tinyerp

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func TestGetContact(t *testing.T) {
	assert := assert.New(t)
	tiny := TinyERP{
		baseURL: "https://tiny-api/test",
		token:   "45bc28e1-24e8-41de-92fa-6bcccaf0cd80",
		format:  "json",
	}
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	fixture, err := ioutil.ReadFile("./fixtures/contact-ok.json")
	if err != nil {
		log.Fatal("Fixture file not found")
	}
	httpmock.RegisterResponder("POST", "https://tiny-api/test/contato.obter.php",
		httpmock.NewStringResponder(200, string(fixture)))
	cr, err := tiny.GetContact("122334456")
	assert.Nil(err)
	assert.Equal("OK", cr.Response.Status)
	assert.Equal("Contato Teste 3", cr.Response.Contact.Name)
}

func TestGetContactError(t *testing.T) {
	assert := assert.New(t)
	tiny := TinyERP{
		baseURL: "https://tiny-api/test",
		token:   "45bc28e1-24e8-41de-92fa-6bcccaf0cd80",
		format:  "json",
	}
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	fixture, err := ioutil.ReadFile("./fixtures/contact-error.json")
	if err != nil {
		log.Fatal("Fixture file not found")
	}
	httpmock.RegisterResponder("POST", "https://tiny-api/test/contato.obter.php",
		httpmock.NewStringResponder(200, string(fixture)))
	cr, err := tiny.GetContact("122334ERR")
	assert.Nil(err)
	assert.Equal("Erro", cr.Response.Status)
	assert.Equal("2", cr.Response.ErrorCode)
}
