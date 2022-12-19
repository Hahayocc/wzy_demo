package mongodb

const DataBase = "mongo_count"

const Collection = "count"

type MongoCount struct {
	Name  string `bson:"name"`
	Count int64  `bson:"count"`
}
