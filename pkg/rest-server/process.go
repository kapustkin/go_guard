package rest

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kapustkin/go_guard/pkg/rest-server/config"
	storage "github.com/kapustkin/go_guard/pkg/rest-server/dal/storage"
	"github.com/kapustkin/go_guard/pkg/rest-server/dal/storage/inmemory"
	"github.com/kapustkin/go_guard/pkg/rest-server/handlers"
	"github.com/kapustkin/go_guard/pkg/utils/logger"
	log "github.com/sirupsen/logrus"
)

// Run основной обработчик
func Run() error {
	// logger init
	logger.Init("rest-server", "0.0.1")
	log.Info("starting app...")
	conf := config.InitConfig()
	log.Infof("use config: %v", conf)

	r := chi.NewRouter()
	// middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	// Logging
	switch conf.Logging {
	case 1:
		r.Use(middleware.Logger)
	case 2:
		r.Use(logger.NewChiLogger())
	default:
		log.Warn("starting without request logging...")
	}

	handler := handlers.Init(getStorage(conf.Storage))
	// Healthchecks
	r.Route("/", func(r chi.Router) {
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte("OK"))
			if err != nil {
				log.Fatal(err)
			}
		})
	})

	// Checker
	r.Route("/checker", func(r chi.Router) {
		r.Post("/", handler.RequestChecker)
	})

	// Adminka
	r.Route("/admin", func(r chi.Router) {
		r.Post("/reset", handler.ResetBucket)
		r.Post("/whitelist", handler.AddToWhiteList)
		r.Delete("/whitelist", handler.RemoveFromWhiteList)
		r.Post("/blacklist", handler.AddToBlackList)
		r.Delete("/blacklist", handler.RemoveFromBlackList)
	})

	log.Infof("listner started...")

	err := http.ListenAndServe(conf.Host, r)
	if err != nil {
		log.Error(err)
	}
	return err
}

func getStorage(storageType int) *storage.Storage {
	switch storageType {
	case 0:
		var db storage.Storage
		db = inmemory.Init()
		return &db
	default:
		log.Panicf("storage type %d not supported", storageType)
	}
	return nil
}
