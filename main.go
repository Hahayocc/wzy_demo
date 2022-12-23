package main

import (
	"code.byted.org/douyincloud-open/configcenter-sdk-golang/base"
	"github.com/Ccc-me/for-golang-test/db/mongodb"
	"github.com/Ccc-me/for-golang-test/db/mysql"
	"github.com/Ccc-me/for-golang-test/db/redis"
	"net/http"
)

var sdkClient base.Client

func Init() {
	mysql.InitMysql()
	redis.InitRedis()
	mongodb.InitMongoDB()
	sdkClient = base.Start()
}

func main() {

	Init()
	http.HandleFunc("/get_key", GetKey)
	http.HandleFunc("/get_keys", GetKeys)
	http.HandleFunc("/update_cache", UpdateCache)
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/headers", Headers)
	http.HandleFunc("/v1/ping", Ping)
	http.HandleFunc("/err", Err)
	http.HandleFunc("/vi/body", Body)
	http.HandleFunc("/err/panic", Panic)
	http.HandleFunc("/log", Log)
	http.HandleFunc("/outlog", OutLog)
	http.HandleFunc("/gray", Gray)

	http.HandleFunc("/mysql/select", MysqlSelect)
	http.HandleFunc("/mysql/select_list", MysqlSelectList)
	http.HandleFunc("/mysql/create", MysqlCreate)
	http.HandleFunc("/mysql/create_lock_table", MysqlCreateLockTable)
	http.HandleFunc("/mysql/update", MysqlUpdate)
	http.HandleFunc("/mysql/update_counts", MysqlUpdateCounts)
	http.HandleFunc("/mysql/delete", MysqlDelete)
	http.HandleFunc("/mysql/delete_rollback", MysqlDeleteRollback)

	http.HandleFunc("/redis/set", RedisSet)
	http.HandleFunc("/redis/get", RedisGet)
	http.HandleFunc("/redis/del", RedisDel)

	http.HandleFunc("/mongodb/insert", MongoInsert)
	http.HandleFunc("/mongodb/find", MongoFind)
	http.HandleFunc("/mongodb/delete", MongoDelete)

	http.ListenAndServe(":8000", nil)
}
