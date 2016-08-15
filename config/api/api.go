package api

import (
	"github.com/qor/qor"
	"go-cat/app/models"
	"go-cat/db"
	"github.com/qor/admin"
)

var API *admin.Admin

func init() {
	API = admin.New(&qor.Config{DB: db.DB})

	 API.AddResource(&models.Category{})


}
