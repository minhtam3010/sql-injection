package handler

import (
	"log"
	"net/http"

	"github.com/minhtam3010/sql-injection/db/entity"
	"github.com/minhtam3010/sql-injection/middleware"
)

type Handler struct {
	middleware *middleware.UserUsecase
}

func NewHandler() *Handler {
	return &Handler{
		middleware: middleware.NewUserUsecase(),
	}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	username := r.Form.Get("email")
	password := r.Form.Get("password")

	user := entity.User{
		UserName: username,
		Password: password,
	}

	_, err = h.middleware.CreateUser(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	username := r.FormValue("email")
	password := r.FormValue("password")

	check, err := h.middleware.Login(username, password)
	if err != nil {
		log.Fatal(err)
	}

	if check {
		http.Redirect(w, r, "/login.html", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/fail.html", http.StatusSeeOther)
	}
}
