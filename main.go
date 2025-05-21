package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/atotto/clipboard"
)

func refreshClipboardContentHandler(w http.ResponseWriter, _ *http.Request) {
	// I love CORS <3
	w.Header().Set("Access-Control-Allow-Origin", "*")

	clipboardContent, err := clipboard.ReadAll()
	if err != nil {
		log.Printf("Error reading from clipboard: %v", err)
		http.Error(w, "Error reading from clipboard", http.StatusInternalServerError)
		return
	}

	var js json.RawMessage
	if err := json.Unmarshal([]byte(clipboardContent), &js); err == nil {
		log.Println("Serving JSON from clipboard.")
		w.Header().Set("Content-Type", "application/json")
		_, writeErr := w.Write([]byte(clipboardContent))
		if writeErr != nil {
			log.Printf("Error writing JSON response: %v", writeErr)
		}
	} else {
		// Not valid JSON, serve as plain text
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, writeErr := w.Write([]byte(clipboardContent))
		if writeErr != nil {
			log.Printf("Error writing plain text response: %v", writeErr)
		}
	}
}

func main() {
	host := flag.String("host", "localhost", "Host address to serve on")
	port := flag.Int("port", 6901, "Port to serve on")
	apiPath := flag.String("api-path", "/api", "API endpoint path for serving clipboard content")

	programName := os.Args[0]

	flag.Usage = func() {
		fmt.Printf("ClipServa is a simple HTTP server that serves the content of your clipboard.\n")
		fmt.Printf("It attempts to serve content as JSON if it's valid JSON, otherwise as plain text.\n\n")
		fmt.Printf("Flags:\n")
		flag.PrintDefaults()
		fmt.Printf("\nExample:\n")
		fmt.Printf("  %s -host 0.0.0.0 -p 8080 -api-path /clipboard_data\n", programName)
	}

	flag.Parse()

	fmt.Print("   _____ _ _       _____                      \n  / ____| (_)     / ____|                     \n | |    | |_ _ __| (___   ___ _ ____   ____ _ \n | |    | | | '_ \\\\___ \\ / _ \\ '__\\ \\ / / _` |\n | |____| | | |_) |___) |  __/ |   \\ V / (_| |\n  \\_____|_|_| .__/_____/ \\___|_|    \\_/ \\__,_|\n            | | By SleepyStew                  \n            |_|                               \n\n")

	// Perform an initial check if clipboard is accessible.
	if _, err := clipboard.ReadAll(); err != nil {
		log.Printf("Warning: Initial clipboard check failed: %v. Clipboard might not be accessible or is empty. The server will still attempt to run.", err)
	}

	if *apiPath == "" || (*apiPath)[0] != '/' {
		log.Fatalf("API path must start with a '/'. Provided: %s", *apiPath)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		message := fmt.Sprintf("ClipServa is running! Send a GET request to %s to get current clipboard content (JSON or plain text).", *apiPath)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, err := fmt.Fprint(w, message)
		if err != nil {
			log.Printf("Error writing root response: %v", err)
		}
	})
	http.HandleFunc(*apiPath, refreshClipboardContentHandler)

	addr := fmt.Sprintf("%s:%d", *host, *port)
	fullAPIURL := fmt.Sprintf("http://%s%s", addr, *apiPath)

	log.Printf("Starting ClipServa on %s", fullAPIURL)
	log.Printf("Copy some text (including JSON) to your clipboard!\n")
	log.Printf("Run with --help to see available options.\n")

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
