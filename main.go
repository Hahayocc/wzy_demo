package main

import (
	"code.byted.org/douyincloud-open/configcenter-sdk-golang/base"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Ccc-me/for-golang-test/db/mongodb"
	"github.com/Ccc-me/for-golang-test/db/mysql"
	"github.com/Ccc-me/for-golang-test/db/redis"
)

var sdkClient base.Client

func Init() {
	mysql.InitMysql()
	redis.InitRedis()
	mongodb.InitMongoDB()
	sdkClient, _ = base.Start()
}

func main() {

	Init()
	http.HandleFunc("/get_key", GetKey)
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
	http.HandleFunc("/config", func(w http.ResponseWriter, request *http.Request) {
		domain := "http://180.184.81.65:8000"
		path := "/config/ping"

		client := http.Client{Timeout: 4 * time.Second}

		req, err := http.NewRequest(http.MethodGet, domain+path, nil)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}
		resp, err := client.Do(req)

		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}
		defer resp.Body.Close()
		followBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprint(w, "err 2 "+err.Error())
			return
		}
		fmt.Fprintf(w, string(followBody))
	})

	http.ListenAndServe(":8000", nil)
}
