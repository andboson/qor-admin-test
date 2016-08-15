package admin

import (
	"fmt"
	"github.com/qor/action_bar"
	"github.com/qor/admin"
	"github.com/qor/i18n/exchange_actions"
	"github.com/qor/l10n/publish"
	"github.com/qor/qor"
	"go-cat/app/models"
	"go-cat/config"
	"go-cat/config/admin/bindatafs"
	"go-cat/config/auth"
	"go-cat/config/i18n"
	"go-cat/db"
	"github.com/qor/roles"
)

var Admin *admin.Admin
var ActionBar *action_bar.ActionBar
var Countries = []string{"China", "Japan", "USA"}
var Categories = make(map[int]string)

func init() {
	Admin = admin.New(&qor.Config{DB: db.DB.Set("publish:draft_mode", true)})
	Admin.SetSiteName("Qor DEMO")
	Admin.SetAuth(auth.AdminAuth{})
	Admin.SetAssetFS(bindatafs.AssetFS)
	config.Filebox.SetAuth(auth.AdminAuth{})
	dir := config.Filebox.AccessDir("/")
	dir.SetPermission(roles.Allow(roles.Read, "admin"))

	cat := Admin.AddResource(&models.Category{}, &admin.Config{Menu: []string{"Product Management"}})
	cat.Meta(&admin.Meta{Name: "Parent", Type:"select_one",
		FormattedValuer: func(cat interface{}, ctx *qor.Context) interface{} {
			var category = new(models.Category)
			current := cat.(*models.Category)
			if current.ID == 0 {
				return "--"
			}

			err := db.DB.Find(category, "id = ?", int(current.Parent)).Error
			if err != nil {
				return "none"
			}

			return category.Name
		}, Collection: func(value interface{}, context *qor.Context) [][]string {
			var collectionValues = [][]string{{"0", "none"}}
			var cats []*models.Category
			current := value.(*models.Category)
			db.DB.Find(&cats)
			for _, cat := range cats {
				if cat.ID == current.ID {
					continue
				}
				collectionValues = append(collectionValues, []string{fmt.Sprint(cat.ID), cat.Name})
			}
			return collectionValues
		}})



	// Add User
	user := Admin.AddResource(&models.User{})
	user.Meta(&admin.Meta{Name: "Gender", Config: &admin.SelectOneConfig{Collection: []string{"Male", "Female", "Unknown"}}})
	user.Meta(&admin.Meta{Name: "Role", Config: &admin.SelectOneConfig{Collection: []string{"Admin", "Maintainer", "Member"}}})
	user.Meta(&admin.Meta{Name: "Confirmed", Valuer: func(user interface{}, ctx *qor.Context) interface{} {
		if user.(*models.User).ID == 0 {
			return true
		}
		return user.(*models.User).Confirmed
	}})

	user.IndexAttrs("ID", "Email", "Name", "Gender", "Role")
	user.ShowAttrs(
		&admin.Section{
			Title: "Basic Information",
			Rows: [][]string{
				{"Name"},
				{"Email", "Password"},
				{"Gender", "Role"},
				{"Confirmed"},
			}},
		"Addresses",
	)
	user.EditAttrs(user.ShowAttrs())




	// Add Worker
	Worker := getWorker()
	Admin.AddResource(Worker)

	db.Publish.SetWorker(Worker)
	exchange_actions.RegisterExchangeJobs(i18n.I18n, Worker)

	// Add Publish
	Admin.AddResource(db.Publish, &admin.Config{Singleton: true})
	publish.RegisterL10nForPublish(db.Publish, Admin)

	// Add Search Center Resources

	// Add ActionBar
	ActionBar = action_bar.New(Admin, auth.AdminAuth{})
	ActionBar.RegisterAction(&action_bar.Action{Name: "Admin Dashboard", Link: "/admin"})

	initFuncMap()
	initRouter()
}

