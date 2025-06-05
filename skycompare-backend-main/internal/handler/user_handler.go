package handler

import (
	"encoding/json"
	"net/http"
	"skycompare-backend-main/internal/models"
	"skycompare-backend-main/internal/service"
)

type UserHandler struct {
	Service *service.UserService
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	nickname := r.Form.Get("nickname")
	password := r.Form.Get("password")

	ok, err := h.Service.Login(nickname, password)
	if err != nil || !ok {
		http.Error(w, "User or password incorrect", http.StatusUnauthorized)
		return
	}
	w.Write([]byte("ok"))
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	nickname := r.Form.Get("nickname")

	exists, _ := h.Service.Login(nickname, "")
	if exists {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	user := models.User{
		Nickname: nickname,
		FullName: r.Form.Get("full_name"),
		Email:    r.Form.Get("email"),
		Password: r.Form.Get("password"),
	}

	ok, err := h.Service.Register(user)
	if err != nil || !ok {
		http.Error(w, "Error registering user", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("User registration confirmed"))
}

func (h *UserHandler) UpdateFavourite(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	nickname := r.Form.Get("user")
	fav := r.Form.Get("fav")

	if fav == "No favourite airport" {
		fav = ""
	}

	if err := h.Service.UpdateFavourite(nickname, fav); err != nil {
		http.Error(w, "Error updating favourite", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Favourite airport saved correctly"))
}

func (h *UserHandler) GetFavourite(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	nickname := r.Form.Get("nickName")

	fav, err := h.Service.GetFavourite(nickname)
	if err != nil {
		http.Error(w, "Error getting favourite", http.StatusInternalServerError)
		return
	}

	resp := struct {
		Airport string `json:"airport"`
	}{Airport: fav}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
