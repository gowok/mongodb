package mongo

import (
	"github.com/gowok/gowok/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	*mongo.Client
}

var _ MongoClient = Mongo{}

func New(conf config.Database) (*Mongo, error) {
	db, err := mongo.NewClient(options.Client().ApplyURI(conf.DSN))
	if err != nil {
		return nil, err
	}

	return &Mongo{db}, err
}
