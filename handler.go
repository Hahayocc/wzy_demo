package main

import (
	"fmt"
	"github.com/Ccc-me/for-golang-test/db/mongodb"
	"github.com/Ccc-me/for-golang-test/db/mysql"
	"github.com/Ccc-me/for-golang-test/db/redis"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func GetKey(w http.ResponseWriter, req *http.Request) {
	key, _ := ioutil.ReadAll(req.Body)
	value, _ := sdkClient.Get(string(key))
	fmt.Fprintf(w, "value = %v\n", value)
}

func Hello(w http.ResponseWriter, req *http.Request) {
	os.Stdout.WriteString("Msg to hello\n")
	fmt.Fprintf(w, "hello\n")
}

func Headers(w http.ResponseWriter, req *http.Request) {
	os.Stdout.WriteString("Msg to headers\n")
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func Gray(w http.ResponseWriter, req *http.Request) {
	os.Stderr.WriteString("Msg to gray err\n")
	http.Error(w, "this is an gray err", 500)
}

func Err(w http.ResponseWriter, req *http.Request) {
	c := req.FormValue("code")
	code, _ := strconv.ParseInt(c, 10, 64)
	os.Stderr.WriteString("Msg to err\n")
	http.Error(w, "this is an err interface", int(code))
}

func Ping(w http.ResponseWriter, req *http.Request) {
	os.Stdout.WriteString("Msg to ping\n")
	fmt.Fprintf(w, "pong!\n")
}

func Body(w http.ResponseWriter, req *http.Request) {
	b, e := ioutil.ReadAll(req.Body)
	fmt.Fprintf(w, "body: %v, err: %v", string(b), e)
}

func Panic(w http.ResponseWriter, req *http.Request) {
	os.Stderr.WriteString("Msg to Panic\n")
	panic(req)
}

func Log(w http.ResponseWriter, req *http.Request) {
	os.Stdout.WriteString("FATAL 1658217911838250000 example.go:66 10.79.163.90 fatal level test!\n")
	os.Stdout.WriteString("WARN 1658217911838250001 example.go:66 10.79.163.90 warn level test!\n")
	os.Stdout.WriteString("ERROR 1658217911838250002 example.go:66 10.79.163.90 error level test!\n")
	os.Stdout.WriteString("NOTICE 1658217911838250002 example.go:66 10.79.163.90 notice level test!\n")
	os.Stdout.WriteString("INFO 1658217911838250002 example.go:66 10.79.163.90 notice level test!\n")
	os.Stdout.WriteString("DEBUG 1658217911838250002 example.go:66 10.79.163.90 debug level test!\n")
}

func OutLog(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	age, _ := strconv.ParseInt(req.FormValue("age"), 10, 64)
	salary, _ := strconv.ParseFloat(req.FormValue("salary"), 10)

	log.SetLevel(log.InfoLevel) // 设置输出警告级别
	if req.FormValue("format") == "text" {
		log.SetFormatter(&log.TextFormatter{}) // 设置 fromat text
	} else {
		log.SetFormatter(&log.JSONFormatter{}) // 设置 format json
	}
	if req.FormValue("type") == "stdout" {
		log.SetOutput(os.Stdout)
	} else {
		log.SetOutput(os.Stderr)
	}
	contextLogger := log.WithFields(log.Fields{
		"name":   name,
		"age":    age,
		"salary": salary,
	})

	contextLogger.Warn("这是一个warn级别日志")
	contextLogger.Info("这是一个info级别日志")
}

func MysqlSelect(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	model, err := mysql.Select(id)
	fmt.Fprintf(w, "Response: %+v, err: %v\n", model, err)
}

func MysqlSelectList(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	models, err := mysql.SelectList(name)
	fmt.Fprintf(w, "Response: %+v, err: %v\n", models, err)
}

func MysqlCreate(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	model, err := mysql.Create(name)
	fmt.Fprintf(w, "Response: %+v, err: %v\n", model, err)
}

func MysqlCreateLockTable(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	model, err := mysql.CreateLockTable(name)
	fmt.Fprintf(w, "Response: %+v, err: %v\n", model, err)
}

func MysqlUpdate(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	model, err := mysql.Update(id)
	fmt.Fprintf(w, "Response: %+v, err: %v\n", model, err)
}

func MysqlUpdateCounts(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	count := req.FormValue("count")
	err := mysql.UpdateCounts(name, count)
	fmt.Fprintf(w, "err: %v\n", err)
}

func MysqlDelete(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	err := mysql.Delete(id)
	fmt.Fprintf(w, "err: %v\n", err)
}

func MysqlDeleteRollback(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	err := mysql.DeleteRollback(id)
	fmt.Fprintf(w, "err: %v\n", err)
}

func RedisSet(w http.ResponseWriter, req *http.Request) {
	key := req.FormValue("key")
	value := req.FormValue("value")
	expireTime := req.FormValue("expireTime")
	res, err := redis.Set(key, value, expireTime)
	fmt.Fprintf(w, "Response: %+v, err: %v\n", res, err)
}

func RedisGet(w http.ResponseWriter, req *http.Request) {
	key := req.FormValue("key")
	res, err := redis.Get(key)
	fmt.Fprintf(w, "Response: %+v, err: %v\n", res, err)
}

func RedisDel(w http.ResponseWriter, req *http.Request) {
	key := req.FormValue("key")
	res, err := redis.Del(key)
	fmt.Fprintf(w, "Response: %+v, err: %v\n", res, err)
}

func MongoInsert(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	count := req.FormValue("count")
	res, err := mongodb.InsertOne(name, count)
	fmt.Fprintf(w, "Response: %+v, err: %v\n", res, err)
}

func MongoFind(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	count := req.FormValue("count")
	res, err := mongodb.FindOne(name, count)
	fmt.Fprintf(w, "Response: %+v, err: %v\n", res, err)
}

func MongoDelete(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	count := req.FormValue("count")
	res, err := mongodb.DeleteOne(name, count)
	fmt.Fprintf(w, "Response: %+v, err: %v\n", res, err)
}
