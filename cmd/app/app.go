package main

import (
	"asidikfauzi/go-gin-intikom/common/inject"
	"asidikfauzi/go-gin-intikom/common/log"
	"asidikfauzi/go-gin-intikom/route"
	"os"
	"time"
)

var stdoutLog string

func init() {
	stdoutLog = "log/intikom.log"
	envTz := os.Getenv("TZ")
	if envTz == "" {
		envTz = "Asia/Jakarta"
	}
	var err error
	time.Local, err = time.LoadLocation(envTz)
	if err != nil {
		log.Printf("error loading location '%s': %v\n", envTz, err)
	}
}

func main() {
	log.Init(stdoutLog)

	routes := route.InitPackage()
	inject.DependencyInjection(inject.InjectData{
		Routes: routes,
	})

	routes.InitRouter()
}
