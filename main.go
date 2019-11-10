package main

import (
	"chi-domain-go/database"
	"chi-domain-go/models"
	"chi-domain-go/router"
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

func initialiseMongoDB() (*mgo.Session, error) {
	urlFromJSON := "mongodb://localhost:27017"
	return mgo.Dial(urlFromJSON)
}

func onCloseHandler(sigs chan os.Signal, done chan bool, dbSession database.Session) {

}

func main() {
	// All initialise here
	mongoDBSession, err := initialiseMongoDB()
	if err != nil {
		panic(err)
	}

	dbSession := &database.Session{MongoDBSession: mongoDBSession}
	routerHandler := router.Initialise(dbSession)

	// os.Signal Handler
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	port := models.ServerPort("3000").String()
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
		color.FgLightGreen.Render("Server Started"),
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
		// dbSession.CloseAllDatabases()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}

	log.Print(color.FgLightBlue.Render("Server Exited Properly"))

	// fmt.Printf("%s\n", color.FgGreen.Render("Listening to port ", port))
	// http.ListenAndServe(port.String(), r)

}
