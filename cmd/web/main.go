package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP service address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	mux := app.routes()

	srv := &http.Server{
		Addr:     *addr,
		Handler:  mux,
		ErrorLog: app.errorLog,
	}

	infoLog.Printf("Listening on port %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)

}
