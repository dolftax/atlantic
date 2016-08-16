package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	_ = viper.ReadInConfig()
}

func main() {
	router := httprouter.New()
	router.GET("/", queuePass)

	log.Println("Atlantic server listening at port" + ":" + viper.Get("appPort").(string))

	log.Fatal(http.ListenAndServe(":"+viper.Get("appPort").(string), middleware(router)))
}
