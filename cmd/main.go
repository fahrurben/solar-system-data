package main

import (
	"fmt"
	"github.com/fahrurben/solar-system-data/internal/solarsystem/domain/body"
	"github.com/fahrurben/solar-system-data/internal/solarsystem/domain/orbitalparameters"
	"github.com/fahrurben/solar-system-data/internal/solarsystem/domain/physicalcharateristic"
	recover2 "github.com/fahrurben/solar-system-data/internal/solarsystem/domain/recover"
	"github.com/fahrurben/solar-system-data/internal/solarsystem/platform/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

// LoadConfig reads configuration f`rom file or environment variables.
func LoadConfig(path string) {
	viper.AddConfigPath(path)
	viper.SetConfigName("test")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("fatal: %+v", err))
	}
}

func main() {
	var bodyRepository *body.RepositoryImpl
	var orbitalRepository *orbitalparameters.RepositoryImpl
	var physicalRepository *physicalcharateristic.RepositoryImpl
	var mysql *sqlx.DB

	LoadConfig("configs/")
	var err error
	mysql, err = db.New(viper.GetString("DATABASE_URL"))
	if err != nil {
		log.Fatal(fmt.Errorf("fatal: %+v", err))
	}
	bodyRepository = body.NewRepository(mysql)
	orbitalRepository = orbitalparameters.NewRepository(mysql)
	physicalRepository = physicalcharateristic.NewRepository(mysql)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/api/v1/recover", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		if err := r.ParseMultipartForm(1024); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		bodyFile, _, err := r.FormFile("body")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer bodyFile.Close()

		orbitalFile, _, err := r.FormFile("orbitalData")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer orbitalFile.Close()

		physicalFile, _, err := r.FormFile("physicalData")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer physicalFile.Close()

		serviceImpl := recover2.NewService(bodyRepository, orbitalRepository, physicalRepository)
		err = serviceImpl.Recover(bodyFile, orbitalFile, physicalFile)

		if err != nil {
			panic(err)
		}

		w.Write([]byte("done"))
	})
	http.ListenAndServe(":3000", r)
}
