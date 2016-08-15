package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/andboson/qor-admin-test/config"
	"github.com/andboson/qor-admin-test/config/admin"
	"github.com/andboson/qor-admin-test/config/auth"
	"github.com/qor/seo"
	"gopkg.in/authboss.v0"
)

func isEditMode(ctx *gin.Context) bool {
	return admin.ActionBar.EditMode(ctx.Writer, ctx.Request)
}

func HomeIndex(ctx *gin.Context) {


	config.View.Funcs(I18nFuncMap(ctx)).Execute(
		"home_index",
		gin.H{
			"ActionBarTag":           admin.ActionBar.Render(ctx.Writer, ctx.Request),
			authboss.FlashSuccessKey: auth.Auth.FlashSuccess(ctx.Writer, ctx.Request),
			authboss.FlashErrorKey:   auth.Auth.FlashError(ctx.Writer, ctx.Request),
			"MicroSearch": seo.MicroSearch{
				URL:    "http://demo.getqor.com",
				Target: "http://demo.getqor.com/search?q={keyword}",
			}.Render(),
			"MicroContact": seo.MicroContact{
				URL:         "http://demo.getqor.com",
				Telephone:   "080-0012-3232",
				ContactType: "Customer Service",
			}.Render(),
			"CurrentLocale": CurrentLocale(ctx),
		},
		ctx.Request,
		ctx.Writer,
	)
}
