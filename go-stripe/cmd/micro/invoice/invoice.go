package main

import (
	"fmt"
	"github.com/johnwr-response/golang-build-web-applications-intermediate-level/go-stripe/cmd/micro/config"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

//type oldConfig struct {
//	port          int
//	hostInterface string
//	smtp          struct {
//		host     string
//		port     int
//		username string
//		password string
//	}
//	frontend string
//}

type application struct {
	//config   oldConfig
	cfg      *config.Config
	infoLog  *log.Logger
	errorLog *log.Logger
	version  string
}

func main() {
	// Read configuration
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err.Error())
	}

	//viper.SetConfigName("default")
	//viper.SetConfigType("yaml")
	//viper.AddConfigPath(".")
	//viper.AddConfigPath("./config/")
	//viper.AutomaticEnv()
	//err := viper.ReadInConfig()
	//if err != nil {
	//	fmt.Println("fatal error config file: default \n", err)
	//	os.Exit(1)
	//}

	//var oldCfg oldConfig
	//flag.StringVar(&oldCfg.hostInterface, "interface", "localhost", "Server interface to listen to")
	//flag.IntVar(&oldCfg.port, "port", 5000, "Server port to listen on")
	//flag.StringVar(&oldCfg.smtp.host, "smtp-host", "sandbox.smtp.mailtrap.io", "smtp host")
	//flag.IntVar(&oldCfg.smtp.port, "smtp-port", 587, "smtp port")
	//flag.StringVar(&oldCfg.smtp.username, "smtp-username", "25853d08526311", "smtp username")
	//flag.StringVar(&oldCfg.smtp.password, "smtp-password", "399982fbb4cbe9", "smtp password")
	//flag.StringVar(&oldCfg.frontend, "frontend", "http://localhost:4000", "frontend url")
	////flag.StringVar(&oldCfg.db.dsn, "dsn", "", "datasource")
	//flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		//config:   oldCfg,
		cfg:      cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		version:  version,
	}

	_ = app.CreateDirIfNotExist("./invoices")

	err = app.serve()
	if err != nil {
		log.Fatal(err)
	}
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf("%s:%d", app.cfg.HostInterface, app.cfg.Port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}
	app.infoLog.Println(fmt.Sprintf("Starting invoice microservice on port %d", app.cfg.Port))
	return srv.ListenAndServe()
}
