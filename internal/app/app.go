package app

import (
	"ulab3/config"
	"ulab3/internal/controller"
	"ulab3/internal/controller/http"

	"github.com/gin-gonic/gin"
	"log"
	"ulab3/pkg/logger"
	"ulab3/pkg/mongo"
)

func Run(cfg config.Config) {

	logger1 := logger.NewLogger()

	db, err := mongo.Connection(cfg)
	if err != nil {
		log.Fatal(err)
	}

	controller1 := controller.NewController(db, logger1, cfg.DB_NAME)

	engine := gin.Default()
	http.NewRouter(engine, controller1)

	log.Fatal(engine.Run(cfg.RUN_PORT))
}
