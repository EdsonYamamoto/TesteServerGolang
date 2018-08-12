package model

import (
	"time"
)

//API ...
type API struct {
	APIName        string `json: "apiTeste"`
	APIDescription string `json: "apiDescription"`
	APIServer      string `json: "apiServer"`
	APIUser        string `json: "apiUser"`
	APIToken       string `json: "apiToken"`
}

//DocJSON ...
type DocJSON map[string]interface{}

//DocStringJSON ...
type DocStringJSON map[string]string

//Pessoa ...
type Pessoa struct {
	Nome       string    `json:"nome" mapstructure:"Nome" form:"nome"`
	Faculdade  string    `json:"faculdade" mapstructure:"Faculdade" form:"faculdade"`
	Nascimento time.Time `json:"nascimento" mapstructure:"Nascimento" form:"nascimento"`
}

//TesteNome ...
type TesteNome struct {
	Nome string `json:"nome" mapstructure:"nome" form:"nome"`
}

//PostCollectionConsulta ...
type PostCollectionConsulta struct {
	Collection string `json:"collection" form:"collection"`
	ObjectID   string `json:"objectID" form:"objectID"`
}
