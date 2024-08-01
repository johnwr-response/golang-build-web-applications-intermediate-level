package main

import (
	"database/sql"
	"fmt"
	"github.com/johnwr-response/golang-build-web-applications-intermediate-level/go-stripe/internal/config"
	"github.com/johnwr-response/golang-build-web-applications-intermediate-level/go-stripe/internal/driver"
	"github.com/johnwr-response/golang-build-web-applications-intermediate-level/go-stripe/internal/models"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

//type config struct {
//	port          int
//	hostInterface string
//	env           string
//	db            struct {
//		dsn string
//	}
//	stripe struct {
//		secret string
//		key    string
//	}
//	smtp struct {
//		host     string
//		port     int
//		username string
//		password string
//	}
//	secretKey string
//	frontend  string
//	invoice   string
//}

type application struct {
	config   *config.Config
	infoLog  *log.Logger
	errorLog *log.Logger
	version  string
	DB       models.DBModel
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf("%s:%d", app.config.Api.HostInterface, app.config.Api.Port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}
	app.infoLog.Println(fmt.Sprintf("Starting Back end server in %s mode on port %d", app.config.Env, app.config.Api.Port))
	return srv.ListenAndServe()
}

func main() {
	// Read configuration
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err.Error())
	}
	//var cfg config
	//flag.StringVar(&cfg.hostInterface, "interface", "localhost", "Server interface to listen to")
	//flag.IntVar(&cfg.port, "port", 4001, "Server port to listen on")
	//flag.StringVar(&cfg.env, "env", "development", "Application environment {development|production|maintenance}")
	//flag.StringVar(&cfg.db.dsn, "dsn", "widgets:secret@tcp(localhost:3306)/widgets?parseTime=true&tls=false", "Database connection string")
	//flag.StringVar(&cfg.smtp.host, "smtp-host", "sandbox.smtp.mailtrap.io", "smtp host")
	//flag.IntVar(&cfg.smtp.port, "smtp-port", 587, "smtp port")
	//flag.StringVar(&cfg.smtp.username, "smtp-username", "25853d08526311", "smtp username")
	//flag.StringVar(&cfg.smtp.password, "smtp-password", "399982fbb4cbe9", "smtp password")
	//flag.StringVar(&cfg.secretKey, "secret-key", "JustABe2BeBlockOfBe2BeVeryRandom", "secret key")
	//flag.StringVar(&cfg.frontend, "frontend", "http://localhost:4000", "frontend url")
	//flag.StringVar(&cfg.invoice, "invoice-url", "http://localhost:5000/invoice/create-and-send", "invoice microservice url")
	////flag.StringVar(&cfg.db.dsn, "dsn", "", "datasource")
	//flag.Parse()

	//cfg.stripe.key = os.Getenv("STRIPE_KEY")
	//cfg.stripe.secret = os.Getenv("STRIPE_SECRET")
	//cfg.stripe.key = "pk_test_51PbNQJAmpQVYH1go2dhZHbpjNtORcVyaGAAEiuKI0Gy8Uk3vuXRLCOy5YGqYTLohNEmkph9fMiQwZVvHsRLiz09m00TrybUDVX"
	//cfg.stripe.secret = "sk_test_51PbNQJAmpQVYH1goK1AAe0OzKXyXwOMneDmKEG9gC2TvIJh5kRolff1ph5QkCyUsGX3foNk1BEPUzffSSxQUUTME00PibhFZo0"

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	dsn := cfg.Api.Dsn
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

	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		version:  version,
		DB:       models.DBModel{DB: conn},
	}
	err = app.serve()
	if err != nil {
		log.Fatal(err)
	}

}
