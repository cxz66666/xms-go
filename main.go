package main

import (
	"fmt"
	"net/http"
	"xms/models"
	"xms/pkg/logging"
	"xms/pkg/setting"
	"xms/routers"
)

// @title xms
// @license.name MIT
// @version 1.0
// @host xms.zjueva.net

func main() {
	// packages init, you can also use `init` function to init package one by one, but
	// init function will be called in order of dependency, so much time it's not very obviously
	// so we rename `init` to `Setup` and call them in our needed orders
	setting.Setup()

	models.Setup()
	logging.Setup()

	router:=routers.InitRouter()

	fmt.Println("running on ",setting.ServerSetting.HttpPort)

	s:=&http.Server{
		Addr:  fmt.Sprintf(":%d",setting.ServerSetting.HttpPort),
		Handler: router,
		ReadTimeout: setting.ServerSetting.ReadTimeout,
		WriteTimeout: setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1<<20,
	}
	s.ListenAndServe()
}
