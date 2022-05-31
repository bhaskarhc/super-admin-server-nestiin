package main

import (
	"log"
	"net/http"

	loginhandler "github.com/bhaskarhc/admin-nestiin/Handler/LoginHandler"
	"github.com/bhaskarhc/admin-nestiin/Handler/fabricHandler"
	"github.com/bhaskarhc/admin-nestiin/config"
	"github.com/bhaskarhc/admin-nestiin/db"
	jwt_service "github.com/bhaskarhc/admin-nestiin/services/jwt"
	"github.com/gorilla/mux"
)

func main() {
	Router := mux.NewRouter()
	viperConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	err = db.InitDB(viperConfig)
	if err != nil {
		panic(err)
	}
	log.Print("\n Database Connected .. :)")
	publicEP := Router.PathPrefix("/v1").Subrouter()
	publicEP.HandleFunc("/login", loginhandler.LoginHandler).Methods(http.MethodPost)

	adminEP := Router.PathPrefix("/v2").Subrouter()
	adminEP.Use(jwt_service.JwtAuthentication)

	adminEP.HandleFunc("/fabric/sellar", fabricHandler.SellarDetailsHandler)
	adminEP.HandleFunc("/fabric/quality", fabricHandler.QualityRequirements)
	adminEP.HandleFunc("/fabric/technical", fabricHandler.TechnicalRequirement)

	log.Fatal(http.ListenAndServe(":"+viperConfig.GetString("serverConfig.port"), Router))
}
