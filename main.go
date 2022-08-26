package main

import (
	"demo/db"
	"demo/service"
	"fmt"
	"log"
	"net/http"
	"os"
)

func panic(v interface{}) {
}

func main() {

	if err := db.InitRedis(); err != nil {
		//panic(fmt.Sprintf("redis init failed with %+v", err))
	}

	if err := db.InitMongoDB(); err != nil {
		//panic(fmt.Sprintf("mongodb init failed with %+v", err))
	}

	http.HandleFunc("/", service.IndexHandler)
	http.HandleFunc("/api/count", service.CounterHandler)

	http.HandleFunc("/api/test", service.TestHandler)
	http.HandleFunc("/api/get_follow_list", service.FollowListHandler)
	http.HandleFunc("/api/get_follow_list_test", service.TestFollowListHandler)
	http.HandleFunc("/v1/ping", service.PingHandler)
	http.HandleFunc("/api/get_os_env", service.GetOsEnvHandler)

	//fmt.Println("have fun")
	fmt.Println("flagflag:", os.Getenv("flag"))
	listenPort := ":8000"
	if os.Getenv("flag") != "success" {
		panic("crash")
		listenPort = ""
		fmt.Println("fail")
	}

	fmt.Println(os.Getenv("flag"))

	if listenPort == "" {
		log.Fatal("failed to load _FAAS_RUNTIME_PORT")
	}
	fmt.Println("http ListenAndServe ", listenPort)
	log.Fatal(http.ListenAndServe(listenPort, nil))
}
