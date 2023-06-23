package postcontroller

import (
	"errors"
	"go-api/config"
	"go-api/helper"
	"go-api/models"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var post []models.Post

	if err := config.DB.Find(&post).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "List of Posts", post)

}

func Show(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]

	var post models.Post

	if err := config.DB.First(&post, "slug = ?", slug).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Post not found", nil)
			return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "Post Detail", post)
}
