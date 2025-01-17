package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"main/database"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/acme/autocert"
	"gorm.io/gorm"
)

func main() {
	if loadErr := godotenv.Load(); loadErr != nil {
		panic(loadErr)
	}

	database.SetupDatabase()

	if migrationErr := database.Db().AutoMigrate(&database.CHandle{}); migrationErr != nil {
		panic(migrationErr)
	}

	manager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(os.Getenv("DOMAIN")),
		Cache:      autocert.DirCache("certs/"),
	}

	go func() {
		httpServer := &http.Server{
			Addr:              ":80",
			Handler:           manager.HTTPHandler(nil),
			ReadTimeout:       30 * time.Second,
			ReadHeaderTimeout: 10 * time.Second,
			WriteTimeout:      30 * time.Second,
			IdleTimeout:       time.Minute,
		}

		if httpListenErr := httpServer.ListenAndServe(); httpListenErr != nil {
			panic(httpListenErr)
		}
	}()

	sMux := http.NewServeMux()
	sMux.HandleFunc("GET /", indexPage)
	sMux.HandleFunc("GET /{username}/.well-known/atproto-did", getProtogen)
	sMux.HandleFunc("GET /{username}/.well-known/discord", getDiscord)

	httpsServer := &http.Server{
		Addr:              ":443",
		Handler:           sMux,
		TLSConfig:         manager.TLSConfig(),
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       time.Minute,
	}

	if httpsListenErr := httpsServer.ListenAndServeTLS("", ""); httpsListenErr != nil {
		panic(httpsListenErr)
	}
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	// For now, just redirect to Bluesky
	// TODO eventually: actually create a frontend
	http.Redirect(w, r, "https://bsky.app", http.StatusFound)
}

func getProtogen(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	username := r.PathValue("username")

	var bHandle database.CHandle
	dbErr := database.Db().Model(&database.CHandle{}).Where("handle = ?", username).First(&bHandle).Error

	if errors.Is(dbErr, gorm.ErrRecordNotFound) {
		http.Error(w, "Handle not found", http.StatusNotFound)
		return
	}

	if dbErr != nil {
		http.Error(w, "Failed to get handle", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, bHandle.DID)
}

func getDiscord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	username := r.PathValue("username")

	var dHandle database.CHandle
	dbErr := database.Db().Model(&database.CHandle{}).Where("handle = ?", username).First(&dHandle).Error

	if errors.Is(dbErr, gorm.ErrRecordNotFound) {
		http.Error(w, "Handle not found", http.StatusNotFound)
		return
	}

	if dbErr != nil {
		http.Error(w, "Failed to get handle", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, dHandle.DHCode)
}
