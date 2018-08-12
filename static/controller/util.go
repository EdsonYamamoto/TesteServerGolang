package controller

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/static/model"
	srv "github.com/heroku/go-getting-started/static/service"
	"github.com/mitchellh/mapstructure"
)

func GetAPI(apiName string) (model.API, error) {
	var allAPI []model.API
	var API model.API
	raw, err := ioutil.ReadFile("./appDoc.json")
	if err != nil {
		log.Fatalln(err)
	}
	json.Unmarshal(raw, &allAPI)
	for i := 0; i < len(allAPI); i++ {
		if allAPI[i].APIName == apiName {
			API = allAPI[i]
			return API, nil
		}
	}
	return API, errors.New("API nao encontrada")
}

//GETAllCollection ...
func GETAllCollection(c *gin.Context) {
	name := c.Param("name")
	log.Println(name)
	serviceUtil := srv.UtilBaseService{}
	mapString, err := serviceUtil.GetAllDataFromCollection(name)
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, mapString)
}

//POSTIndex ...
func POSTIndex(c *gin.Context) {
	var variaveis model.PostCollectionConsulta
	err := c.Bind(&variaveis)
	if err != nil {
		log.Fatalln(err)
	}

	serviceUtil := srv.UtilBaseService{}

	m, err := serviceUtil.GetOneDataByDoc(variaveis.Collection, variaveis.ObjectID)
	if err != nil {
		log.Println(err)
	}
	var p model.Pessoa
	err = mapstructure.Decode(m, &p)
	log.Println(p)
	c.JSON(http.StatusOK, p)
}

//JSONIndex ...
func JSONIndex(c *gin.Context) {

	var t model.Pessoa
	t.Nome = "Pessoa1"
	t.Faculdade = "Alguma Faculdade"
	t.Nascimento = time.Now()

	log.Println("JSON teste")
	c.JSON(http.StatusOK, t)
}
