package sDAGraph_mongo

import(
	"fmt"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
)

func GetDB(ip string, dbName string) (*mgo.Database,*mgo.Session) {
	session, err := mgo.Dial(ip)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	db := session.DB(dbName)
	return db, session
}

func DefaultGetDB() *mgo.Database {
	session, err := mgo.Dial("mongodb://192.168.51.202:27017")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("testdatabase")
	return db
}

func Insert(db *mgo.Database, collection string, content interface{}) (error){
	c := db.C(collection)
	err := c.Insert(content)
	return err
}

func Update(db *mgo.Database, collection string, target interface{}, content interface{}) {
        c := db.C(collection)
	err := c.Update(&target, &content)
        if err != nil {
                fmt.Println(err)
        }
}

func FindOne(db *mgo.Database, collection string, content interface{}) (interface{}){
        c := db.C(collection)
	var users bson.M
        c.Find(content).One(&users)
	return users
}

func FindAll(db *mgo.Database, collection string, content interface{}) ([]bson.M){//interface{}){
        c := db.C(collection)
        var users []bson.M
        c.Find(content).All(&users)
        return users
}


