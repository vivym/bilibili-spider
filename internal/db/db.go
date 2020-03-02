package db

import (
	"github.com/Kamva/mgm/v2"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupDB(url string, name string) error {
	err := mgm.SetDefaultConfig(nil, name, options.Client().ApplyURI(url))
	if err != nil {
		return err
	}
	return nil
}
