package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/inari111/go-sandbox/xo/models"
	_ "github.com/lib/pq"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/", createUser)
		r.Get("/search", searchUser)
	})

	http.ListenAndServe(":3333", r)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	id, err := uuid.NewUUID()
	if err != nil {
		log.Printf("gen uuid err: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	u := &models.User{
		ID:        id,
		Name:      "testuser",
		CreatedAt: now,
		UpdatedAt: now,
	}

	db := models.NewDB()
	if err := u.Insert(context.Background(), db); err != nil {
		log.Printf("insert err: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func searchUser(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db := models.NewDB()
	u := models.User{}
	rows, err := u.GetByName(context.Background(), db, name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(rows)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
