package main

import (
	"fmt"
	"net/http"

	"github.com/andboson/qor-admin-test/config"
	"github.com/andboson/qor-admin-test/config/admin"
	"github.com/andboson/qor-admin-test/config/api"
	_ "github.com/andboson/qor-admin-test/config/i18n"
	"github.com/andboson/qor-admin-test/config/routes"
	_ "github.com/andboson/qor-admin-test/db/migrations"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", routes.Router())
	admin.Admin.MountTo("/admin", mux)
	api.API.MountTo("/api", mux)
	config.Filebox.MountTo("/downloads", mux)

	for _, path := range []string{"system", "javascripts", "stylesheets", "images"} {
		mux.Handle(fmt.Sprintf("/%s/", path), http.FileServer(http.Dir("public")))
	}

	fmt.Printf("Listening on: %v\n", config.Config.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), mux); err != nil {
		panic(err)
	}
}
