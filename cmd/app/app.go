package main

import (
	"asidikfauzi/go-gin-intikom/common/inject"
	"asidikfauzi/go-gin-intikom/route"
)

func main() {
	routes := route.InitPackage()
	inject.DependencyInjection(inject.InjectData{
		Routes: routes,
	})

	routes.InitRouter()
}
