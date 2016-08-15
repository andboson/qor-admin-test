// +build ignore

package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jinzhu/now"
	"github.com/qor/i18n/backends/database"
	"github.com/qor/media_library"
	"github.com/qor/publish"
	"go-cat/app/models"
	"go-cat/config/admin"
	"go-cat/db"
	"go-cat/db/seeds"
	"github.com/qor/seo"
	"github.com/qor/slug"
	"github.com/qor/sorting"
)

/* How to upload file
 * $ brew install s3cmd
 * $ s3cmd --configure (Refer https://github.com/theplant/qor-example)
 * $ s3cmd put local_file_path s3://qor3/
 */

var (
	fake           = seeds.Fake
	truncateTables = seeds.TruncateTables

	Seeds  = seeds.Seeds
	Tables = []interface{}{
		&models.User{}, &models.Address{},
		&models.Category{}, &models.Color{}, &models.Size{}, &models.Collection{},
		&models.Product{}, &models.ColorVariation{}, &models.ColorVariationImage{}, &models.SizeVariation{},
		&models.Store{},
		&models.Order{}, &models.OrderItem{},
		&models.Setting{},
		&models.SEOSetting{},

		&media_library.AssetManager{},
		&publish.PublishEvent{},
		&database.Translation{},
		&admin.QorWidgetSetting{},
	}
)

func main() {
	truncateTables(Tables...)
	createRecords()
}

func findCategoryByName(name string) *models.Category {
	category := &models.Category{}
	if err := db.DB.Where(&models.Category{Name: name}).First(category).Error; err != nil {
		log.Fatalf("can't find category with name = %q, got err %v", name, err)
	}
	return category
}

func findCollectionByName(name string) *models.Collection {
	collection := &models.Collection{}
	if err := db.DB.Where(&models.Collection{Name: name}).First(collection).Error; err != nil {
		log.Fatalf("can't find collection with name = %q, got err %v", name, err)
	}
	return collection
}

func findColorByName(name string) *models.Color {
	color := &models.Color{}
	if err := db.DB.Where(&models.Color{Name: name}).First(color).Error; err != nil {
		log.Fatalf("can't find color with name = %q, got err %v", name, err)
	}
	return color
}

func findSizeByName(name string) *models.Size {
	size := &models.Size{}
	if err := db.DB.Where(&models.Size{Name: name}).First(size).Error; err != nil {
		log.Fatalf("can't find size with name = %q, got err %v", name, err)
	}
	return size
}

func findProductByColorVariationID(colorVariationID uint) *models.Product {
	colorVariation := models.ColorVariation{}
	product := models.Product{}

	if err := db.DB.Find(&colorVariation, colorVariationID).Error; err != nil {
		log.Fatalf("query colorVariation (%v) failure, got err %v", colorVariation, err)
		return &product
	}
	if err := db.DB.Find(&product, colorVariation.ProductID).Error; err != nil {
		log.Fatalf("query product (%v) failure, got err %v", product, err)
		return &product
	}
	return &product
}

func randTime() time.Time {
	num := rand.Intn(10)
	return time.Now().Add(-time.Duration(num*24) * time.Hour)
}

func openFileByURL(rawURL string) (*os.File, error) {
	if fileURL, err := url.Parse(rawURL); err != nil {
		return nil, err
	} else {
		path := fileURL.Path
		segments := strings.Split(path, "/")
		fileName := segments[len(segments)-1]

		filePath := filepath.Join("/tmp", fileName)

		if _, err := os.Stat(filePath); err == nil {
			return os.Open(filePath)
		}

		file, err := os.Create(filePath)
		if err != nil {
			return file, err
		}

		check := http.Client{
			CheckRedirect: func(r *http.Request, via []*http.Request) error {
				r.URL.Opaque = r.URL.Path
				return nil
			},
		}
		resp, err := check.Get(rawURL) // add a filter to check redirect
		if err != nil {
			return file, err
		}
		defer resp.Body.Close()
		fmt.Printf("----> Downloaded %v\n", rawURL)

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			return file, err
		}
		return file, nil
	}
}
