package main

import (
	"database/sql"
	"encoding/gob"
	"flag"
	"fmt"
	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
	"github.com/johnwr-response/golang-build-web-applications-intermediate-level/go-stripe/internal/config"
	"github.com/johnwr-response/golang-build-web-applications-intermediate-level/go-stripe/internal/driver"
	"github.com/johnwr-response/golang-build-web-applications-intermediate-level/go-stripe/internal/models"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"
const cssVersion = "1"

var session *scs.SessionManager

//type config struct {
//	port          int
//	hostInterface string
//	env           string
//	api           string
//	db            struct {
//		dsn string
//	}
//	stripe struct {
//		secret string
//		key    string
//	}
//	secretKey string
//	frontend  string
//	invoice   string
//}

type application struct {
	config        *config.Config
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache map[string]*template.Template
	version       string
	cssVersion    string
	DB            models.DBModel
	Session       *scs.SessionManager
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf("%s:%d", app.config.Web.HostInterface, app.config.Web.Port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}
	app.infoLog.Println(fmt.Sprintf("Starting HTTP server in %s mode on port %d", app.config.Env, app.config.Web.Port))
	return srv.ListenAndServe()
}

func main() {
	gob.Register(TransactionData{})
	// Read configuration
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err.Error())
	}
	//var cfg config

	//flag.StringVar(&cfg.hostInterface, "interface", "localhost", "Server interface to listen to")
	//flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	//flag.StringVar(&cfg.env, "env", "development", "Application environment {development|production}")
	//flag.StringVar(&cfg.db.dsn, "dsn", "widgets:secret@tcp(127.0.0.1:3306)/widgets?parseTime=true&tls=false", "Database connection string")
	//flag.StringVar(&cfg.api, "api", "http://localhost:4001", "URL to API")
	//flag.StringVar(&cfg.secretKey, "secret-key", "JustABe2BeBlockOfBe2BeVeryRandom", "secret key")
	//flag.StringVar(&cfg.frontend, "frontend", "http://localhost:4000", "frontend url")
	//flag.StringVar(&cfg.invoice, "invoice-url", "http://localhost:5000/invoice/create-and-send", "invoice microservice url")

	flag.Parse()

	//cfg.stripe.key = os.Getenv("STRIPE_KEY")
	//cfg.stripe.secret = os.Getenv("STRIPE_SECRET")
	//cfg.stripe.key = "pk_test_51PbNQJAmpQVYH1go2dhZHbpjNtORcVyaGAAEiuKI0Gy8Uk3vuXRLCOy5YGqYTLohNEmkph9fMiQwZVvHsRLiz09m00TrybUDVX"
	//cfg.stripe.secret = "sk_test_51PbNQJAmpQVYH1goK1AAe0OzKXyXwOMneDmKEG9gC2TvIJh5kRolff1ph5QkCyUsGX3foNk1BEPUzffSSxQUUTME00PibhFZo0"

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	if cfg.Web.Dsn == "" {

	} else {

	}
	dsn := cfg.Web.Dsn
	if dsn == "" {
		dsn = cfg.Dsn
	}
	conn, err := driver.OpenDB(dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer func(conn *sql.DB) {
		_ = conn.Close()
	}(conn)

	// set up session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Store = mysqlstore.New(conn)

	tc := make(map[string]*template.Template)
	app := &application{
		config:        cfg,
		infoLog:       infoLog,
		errorLog:      errorLog,
		templateCache: tc,
		version:       version,
		cssVersion:    cssVersion,
		DB:            models.DBModel{DB: conn},
		Session:       session,
	}

	go app.ListenToWSChannel()

	err = app.serve()
	if err != nil {
		app.errorLog.Println(err)
		log.Fatal(err)
	}
}
