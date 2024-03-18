package api

import (
	"fmt"
	v1 "go-recommendation-system/api/v1"
	"log"

	"github.com/beego/beego/v2/server/web"
)

func RunApiServer() {

	ns_v1 := web.NewNamespace("/v1",
		web.NSRouter("/reg", &v1.RegCtrl{}),
		web.NSRouter("/confirm", &v1.ConfirmCtrl{}),
		web.NSRouter("/auth", &v1.AuthCtrl{}),
		web.NSRouter("/recommendation", &v1.GetRecommendationCtrl{}),
	)

	web.BConfig.WebConfig.DirectoryIndex = true
	web.AddNamespace(ns_v1)

	address := fmt.Sprintf("0.0.0.0:8888")
	log.Printf("api server is listening to %v", address)
	web.Run(address)
}
