package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
)

func enforceJSONHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")

		if contentType != "" {
			mt, _, err := mime.ParseMediaType(contentType)
			if err != nil {
				http.Error(w, "Malformed Content-Type header", http.StatusBadRequest)
				return
			}

			if mt != "application/json" {
				http.Error(w, "Content-Type header must be application/json", http.StatusUnsupportedMediaType)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

func processAlertmanagerWebhook(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error reading body!")
	}
	defer r.Body.Close()

	log.Printf("Raw body: %s\n", string(b))

	alertData := &AlertWebhook{}
	if err := json.Unmarshal(b, &alertData); err != nil {
		log.Fatal("Error unmarshalling Alertmanager webhook data!")
	}

	switch alertData.Status {
	case "resolved":
		// TODO: process resolved alerts
	case "firing":
		// TODO: process active alerts
	default:
		log.Fatalf("Unknown alert status: %s", alertData.Status)
	}
}

func main() {
	initKubernetesClient()

	mux := http.NewServeMux()

	alertmanagerHandler := http.HandlerFunc(processAlertmanagerWebhook)
	mux.Handle("/", enforceJSONHandler(alertmanagerHandler))

	log.Println("Listening on :3000 port...")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
