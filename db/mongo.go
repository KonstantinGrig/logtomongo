package db

import (
	"github.com/globalsign/mgo"
	"logtomongo/config"
	"time"
)

type MongoItemLog struct {
	Log_time   time.Time
	Log_msg    string
	File_name  string
	Log_format string
}

func LogCollection() *mgo.Collection {
	conf := config.Config{}
	info := &mgo.DialInfo{
		Addrs:    []string{conf.GetString("db.hosts")},
		Timeout:  6 * time.Second,
		Username: conf.GetString("db.username"),
		Password: conf.GetString("db.password"),
	}

	session, err1 := mgo.DialWithInfo(info)
	if err1 != nil {
		panic(err1)
	}

	database := conf.GetString("db.database")
	session.DB(database).DropDatabase()

	return session.DB(database).C(conf.GetString("db.collection"))
}
