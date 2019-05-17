package db

import (
	"github.com/globalsign/mgo"
)

var GlobalDB *Database

const (
	CUser = "user"
)

type GetCollection func() (collection *mgo.Collection, closeConn func())

func InitializeGlobalDB(url string) error {
	d, err := NewDatabase(url)
	if err != nil {
		return err
	}
	err = d.EnsureIndex()
	if err != nil {
		return err
	}
	GlobalDB = d
	return nil
}

type Database struct {
	session *mgo.Session
}

func NewDatabase(url string) (*Database, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	d := &Database{
		session: session,
	}
	return d, nil
}

func (d *Database) DB() (*mgo.Database, func()) {
	conn := d.session.Copy()
	return conn.DB("esg"), func() {
		conn.Close()
	}
}

func (d *Database) EnsureIndex() (err error) {
	database, closeConn := d.DB()
	defer closeConn()
	err = database.C(CUser).EnsureIndex(mgo.Index{
		Key:    []string{"username"},
		Unique: true,
	})
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) User() (collection *mgo.Collection, closeConn func()) {
	database, closeConn := d.DB()
	collection = database.C(CUser)
	return collection, closeConn
}
