package main

import (
	"database/sql"
	"flag"
	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"kodesbox.snnafi.dev/internal/models"
	"log"
	"net/http"
	"os"
	"time"
)

type application struct {
	errorLog       *log.Logger
	infoLog        *log.Logger
	box            *models.KodesBox
	templateCache  map[string]*template.Template
	formDecoder    *form.Decoder
	sessionManager *scs.SessionManager
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP service address")

	dsn := flag.String("dsn", "kodesbox:kodesbox@/kodesbox?parseTime=true", "MySQL connection string")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	formDecoder := form.NewDecoder()

	sessionManager := scs.New()
	sessionManager.Store = mysqlstore.New(db)
	sessionManager.Lifetime = 12 * time.Hour

	app := application{
		errorLog:       errorLog,
		infoLog:        infoLog,
		box:            &models.KodesBox{DB: db},
		templateCache:  templateCache,
		formDecoder:    formDecoder,
		sessionManager: sessionManager,
	}

	handler := app.routes()

	srv := &http.Server{
		Addr:     *addr,
		Handler:  handler,
		ErrorLog: app.errorLog,
	}

	infoLog.Printf("Listening on port %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)

}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			db.Close()
		}
	}()

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
