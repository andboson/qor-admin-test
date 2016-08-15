package migrations

import (
	"github.com/qor/activity"
	"github.com/qor/media_library"
	"github.com/qor/publish"
	"github.com/andboson/qor-admin-test/app/models"
	"github.com/andboson/qor-admin-test/db"
	"github.com/qor/transition"
)

func init() {
	AutoMigrate(&media_library.AssetManager{})
	AutoMigrate(&models.Category{})
	AutoMigrate(&transition.StateChangeLog{})
	AutoMigrate(&activity.QorActivity{})
	AutoMigrate(&models.User{})
}

func AutoMigrate(values ...interface{}) {
	for _, value := range values {
		db.DB.AutoMigrate(value)

		if publish.IsPublishableModel(value) {
			db.Publish.AutoMigrate(value)
		}
	}
}
