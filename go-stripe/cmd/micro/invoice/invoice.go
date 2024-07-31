package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port          int
	hostInterface string
	smtp          struct {
		host     string
		port     int
		username string
		password string
	}
	frontend string
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
	version  string
}

func main() {
	var cfg config
	flag.StringVar(&cfg.hostInterface, "interface", "localhost", "Server interface to listen to")
	flag.IntVar(&cfg.port, "port", 5000, "Server port to listen on")
	flag.StringVar(&cfg.smtp.host, "smtp-host", "sandbox.smtp.mailtrap.io", "smtp host")
	flag.IntVar(&cfg.smtp.port, "smtp-port", 587, "smtp port")
	flag.StringVar(&cfg.smtp.username, "smtp-username", "25853d08526311", "smtp username")
	flag.StringVar(&cfg.smtp.password, "smtp-password", "399982fbb4cbe9", "smtp password")
	flag.StringVar(&cfg.frontend, "frontend", "http://localhost:4000", "frontend url")
	//flag.StringVar(&cfg.db.dsn, "dsn", "", "datasource")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		version:  version,
	}

	_ = app.CreateDirIfNotExist("./invoices")

	err := app.serve()
	if err != nil {
		log.Fatal(err)
	}
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf("%s:%d", app.config.hostInterface, app.config.port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}
	app.infoLog.Println(fmt.Sprintf("Starting invoice microservice on port %d", app.config.port))
	return srv.ListenAndServe()
}
