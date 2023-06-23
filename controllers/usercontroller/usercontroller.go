package usercontroller

import (
	"encoding/json"
	"errors"
	"go-api/config"
	"go-api/helper"
	"go-api/models"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var user []models.User

	if err := config.DB.Find(&user).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "List Users", user)

}

func Create(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if err := config.DB.Create(&user).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "success created user", nil)
}


func Show(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]


	var user models.User

	if err := config.DB.First(&user, "username=?", username).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			helper.Response(w, 400, "User not found", nil)
			return
		}


		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "User Detail", user)
}