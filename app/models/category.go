package models

import (
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/qor/l10n"
	"github.com/qor/validations"
	"github.com/qor/sorting"
	"github.com/qor/slug"
	//"github.com/qor/publish"
)

type Category struct {
	gorm.Model
	l10n.Locale
	//publish.Status
	sorting.Sorting  `json:"sort_id"`
	Name         string    `json:"name"`
	NameWithSlug slug.Slug    `l10n:"sync" json:"slug"`
	Parent       uint   `l10n:"sync" json:"parent_id"`
}

func (category Category) Validate(db *gorm.DB) {
	if strings.TrimSpace(category.Name) == "" {
		db.AddError(validations.NewError(category, "Name", "Name can not be empty"))
	}
}
