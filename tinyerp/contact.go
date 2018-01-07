package tinyerp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ContactResponse struct {
	Response struct {
		ProcessState string              `json:"status_processamento"`
		Status       string              `json:"status"`
		ErrorCode    string              `json:"codigo_erro,omitempty"`
		Errors       []map[string]string `json:"erros,omitempty"`
		Contact      struct {
			ID      string `json:"id"`
			Name    string `json:"nome"`
			Type    string `json:"tipo_pessoa"`
			CPFCNPJ string `json:"cpf_cnpj"`
		} `json:"contato"`
	} `json:"retorno"`
}

func (t *TinyERP) GetContact(id string) (*ContactResponse, error) {
	uri, data := t.apiURI("contato.obter.php")
	data.Set("id", id)
	resp, err := http.PostForm(uri, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var contact ContactResponse
	if err := json.NewDecoder(resp.Body).Decode(&contact); err != nil {
		return nil, err
	}
	return &contact, nil
}
