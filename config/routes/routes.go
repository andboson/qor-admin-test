package routes

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/qor/qor"
	"github.com/andboson/qor-admin-test/app/controllers"
	"github.com/andboson/qor-admin-test/config"
	"github.com/andboson/qor-admin-test/config/auth"
	"github.com/andboson/qor-admin-test/db"
	"github.com/qor/qor/utils"
)

func Router() *http.ServeMux {
	router := gin.Default()
	router.Use(func(ctx *gin.Context) {
		if locale := utils.GetLocale(&qor.Context{Request: ctx.Request, Writer: ctx.Writer}); locale != "" {
			ctx.Set("DB", db.DB.Set("l10n:locale", locale))
		}
	})
	gin.SetMode(gin.DebugMode)

	router.GET("/", controllers.HomeIndex)
	router.GET("/switch_locale", controllers.SwitchLocale)

	var mux = http.NewServeMux()
	mux.Handle("/", router)
	mux.Handle("/auth/", auth.Auth.NewRouter())
	publicDir := http.Dir(strings.Join([]string{config.Root, "public"}, "/"))
	mux.Handle("/dist/", http.FileServer(publicDir))
	mux.Handle("/vendors/", http.FileServer(publicDir))
	mux.Handle("/images/", http.FileServer(publicDir))
	mux.Handle("/fonts/", http.FileServer(publicDir))
	return mux
}
