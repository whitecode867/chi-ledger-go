package main

import (
	"chi-ledger-go/conf"
	"chi-ledger-go/database"
	"chi-ledger-go/models"
	"chi-ledger-go/router"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/globalsign/mgo"
	"gopkg.in/gookit/color.v1"
)

func finaliseMongoDBURL() string {
	configs := conf.Configs
	url := configs.GetString("database.mongodb.url")
	port := configs.GetString("database.mongodb.port")
	username := configs.GetString("database.mongodb.username")
	password := configs.GetString("database.mongodb.password")

	if len(username) != 0 && len(password) != 0 {
		return fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, url, port)
	}
	return fmt.Sprintf("mongodb://%s:%s", url, port)
}

func initialiseMongoDB() (*mgo.Session, error) {
	return mgo.Dial(finaliseMongoDBURL())
}

func main() {
	// All initialisaion here
	conf.InitialiseConfigs()
	mongoDBSession, err := initialiseMongoDB()
	if err != nil {
		panic(err)
	}

	dbSession := &database.Session{MongoDBSession: mongoDBSession}
	routerHandler := router.Initialise(dbSession)

	// os.Signal Handler
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	port := models.ServerPort(conf.Configs.GetString("service.port")).String()
	srv := &http.Server{
		Addr:    port,
		Handler: routerHandler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Println(
		color.FgLightGreen.Render(fmt.Sprintf("Server %s Started", conf.Configs.GetString("service.name"))),
		"Listening to port",
		color.FgLightYellow.Render(port),
	)

	<-done

	fmt.Println()
	log.Println(color.FgLightRed.Render("Server Stopped"))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	// extra handling here
	defer func() {
		cancel()
		dbSession.CloseAllDatabases()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}

	log.Print(color.FgLightBlue.Render("Server Exited Properly"))
}
