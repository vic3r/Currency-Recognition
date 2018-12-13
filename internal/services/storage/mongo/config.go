package mongo

import (
	"fmt"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/vic3r/Currency-Recognition/internal/services/storage"

	mgo "gopkg.in/mgo.v2"
)

// mongoConfig is struct config for connecting to Mongo
type mongoConfig struct {
	Hosts      string
	Database   string
	Username   string
	Password   string
	Collection string
}

var (
	db                   *mgo.Session
	database, collection string
)

func connect(conf storage.Config) (err error) {
	if db != nil {
		return nil
	}

	var mConfig mongoConfig
	mapstructure.Decode(conf, &mConfig)
	db, err = mConfig.initMongo()
	if err != nil {
		return fmt.Errorf("Mongo did not respond: %s", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("Mongo did not respond: %s", err)
	}

	return nil
}

func (config *mongoConfig) initMongo() (*mgo.Session, error) {
	info := &mgo.DialInfo{
		Addrs:    []string{config.Hosts},
		Timeout:  60 * time.Second,
		Database: config.Database,
		Username: config.Username,
		Password: config.Password,
	}

	database = config.Database
	collection = config.Collection

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		return nil, err
	}

	return session, nil
}
