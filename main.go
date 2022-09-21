package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"po_go/conf"
	"po_go/db"
	"po_go/redis"
	"po_go/router"
	"po_go/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	defer db.Db.Close()
	defer func() {
		redis.RedisClient.FlushDB()
		redis.RedisClient.Close()
	}()

	//loading log
	log := utils.Log()

	gin.SetMode(conf.Conf.Server.Model)

	router := router.InitRouter()

	srv := &http.Server{
		Addr:    conf.Conf.Server.Address,
		Handler: router,
	}

	go func() {
		//start server
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("listen: %s \n", err)
		}
		log.Fatal("listen: %s \n", conf.Conf.Server.Address)
	}()

	quit := make(chan os.Signal)
	//listen signal to shut sown
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	// Save shut down
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}
