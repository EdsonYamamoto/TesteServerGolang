package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/static/model"
	srv "github.com/heroku/go-getting-started/static/service"
)

//GETAllCollection ...
func POSTSaveCollection(c *gin.Context) {
	var variaveis model.Pessoa
	err := c.Bind(&variaveis)
	if err != nil {
		log.Fatalln(err)
	}
	srvUtil := srv.UtilBaseService{}

	err = srvUtil.SaveUniqueInterfaceDataIntoCollection("Pessoa", variaveis)
	if err != nil {
		log.Println(err)
	}
}

//POSTSaveTesteCollection ...
func POSTSaveTesteCollection(c *gin.Context) {
	var variaveis model.TesteNome

	err := c.Bind(&variaveis)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERROR",
			"message": "Alguma coisa aconteceu e a culpa foi sua, concerta esse codigo aew.",
		})
	}

	log.Println(variaveis)
	srvUtil := srv.UtilBaseService{}

	err = srvUtil.SaveUniqueInterfaceDataIntoCollection("Teste", variaveis)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERROR",
			"message": "Deu algum erro no banco, a culpa nao foi sua",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "posted",
		"message": variaveis,
	})
}

//GETAllCollection ...
func GetTesteCollection(c *gin.Context) {
	srvUtil := srv.UtilBaseService{}

	mapString, err := srvUtil.GetAllDataFromCollection("Teste")
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, mapString)
}
