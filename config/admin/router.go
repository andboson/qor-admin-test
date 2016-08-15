package admin

import (
	"github.com/qor/admin"
)


func ReportsDataHandler(context *admin.Context) {

	var b []byte
	context.Writer.Write(b)
	return
}

func initRouter() {
	Admin.GetRouter().Get("/reports", ReportsDataHandler)
}
