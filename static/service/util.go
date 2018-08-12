package service

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

//UtilBaseService ...
type UtilBaseService struct {
}

//DataBaseAccess ...
func (srvUtil *UtilBaseService) DataBaseAccess() (*firestore.Client, error) {
	opt := option.WithCredentialsFile("firebase.json")

	//opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return client, err
}

//GetOneDataByDoc ...
func (srvUtil *UtilBaseService) GetOneDataByDoc(collection string, document string) (map[string]interface{}, error) {
	//retorna as informações do banco
	client, err := srvUtil.DataBaseAccess()
	if err != nil {
		return nil, err
	}
	result, err := client.Collection(collection).Doc(document).Get(context.Background())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	m := result.Data()
	defer client.Close()

	return m, err
}

//GetAllDataFromCollection ...
func (srcUtil *UtilBaseService) GetAllDataFromCollection(collection string) ([]map[string]interface{}, error) {
	//retorna as informações do banco
	var results []map[string]interface{}

	client, err := srcUtil.DataBaseAccess()
	if err != nil {
		return nil, err
	}
	//comando where é o mesmo do SQL SERVER
	//result := client.Collection(collection).Where("ativo", "==", true).Documents(context.Background())
	result := client.Collection(collection).Documents(context.Background())
	for {
		doc, err := result.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		results = append(results, doc.Data())
	}
	return results, err
}

//SaveUniqueInterfaceDataIntoCollection ...
func (srcUtil *UtilBaseService) SaveUniqueInterfaceDataIntoCollection(collection string, model ...interface{}) error {
	var mapString map[string]interface{}
	mapstructure.WeakDecode(model, &mapString)
	client, err := srcUtil.DataBaseAccess()
	if err != nil {
		return err
	}
	_, _, err = client.Collection(collection).Add(context.Background(), mapString)
	if err != nil {
		return err
	}
	return err
}

//SaveUniqueMapStringDataIntoCollection ...
func (srcUtil *UtilBaseService) SaveUniqueMapStringDataIntoCollection(collection string, mapString map[string]interface{}) error {
	client, err := srcUtil.DataBaseAccess()
	if err != nil {
		return err
	}
	_, _, err = client.Collection(collection).Add(context.Background(), mapString)
	if err != nil {
		return err
	}
	return err
}
