package api

import (
	"github.com/qor/qor"
	"github.com/andboson/qor-admin-test/app/models"
	"github.com/andboson/qor-admin-test/db"
	"github.com/qor/admin"
)

var API *admin.Admin

func init() {
	API = admin.New(&qor.Config{DB: db.DB})
	API.AddResource(&models.Category{})
}
