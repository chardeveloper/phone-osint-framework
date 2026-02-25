// @author CHAR DEV QUANTUM
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	webMode := flag.Bool("web", false, "Start Local Web Server on port 5000")
	targetOpt := flag.String("target", "", "Target for CLI OSINT scan")
	flag.Parse()

	if *webMode {
		startWebServer()
		return
	}

	if *targetOpt != "" {
		runCLI(*targetOpt)
		return
	}

	fmt.Println("Usage: ./chardev-osint --web or ./chardev-osint --target <number>")
	os.Exit(1)
}

func startWebServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/scan/veriphone", handleVeriphone)
	mux.HandleFunc("/api/scan/ipqs", handleIPQS)

	port := "5000"
	fmt.Printf("[+] CharDev OSINT Framework (Web Mode) running on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, corsMiddleware(mux)))
}

func parseResponseBody(r *http.Response) ([]byte, error) {
	defer r.Body.Close()
	return io.ReadAll(r.Body)
}

func handleVeriphone(w http.ResponseWriter, r *http.Request) {
	key := r.Header.Get("X-Veriphone-Key")
	target := r.URL.Query().Get("target")

	if key == "" || target == "" {
		http.Error(w, `{"error": "Missing key or target"}`, http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf("https://api.veriphone.io/v2/verify?phone=%s&key=%s", target, key)
	resp, err := http.Get(url)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, fmt.Sprintf(`{"error": "%v"}`, err), http.StatusInternalServerError)
		return
	}

	data, _ := parseResponseBody(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func handleIPQS(w http.ResponseWriter, r *http.Request) {
	key := r.Header.Get("X-Ipqs-Key")
	target := r.URL.Query().Get("target")

	if key == "" || target == "" {
		http.Error(w, `{"error": "Missing key or target"}`, http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf("https://www.ipqualityscore.com/api/json/phone/%s/%s?strictness=0", key, target)
	resp, err := http.Get(url)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, fmt.Sprintf(`{"error": "%v"}`, err), http.StatusInternalServerError)
		return
	}

	data, _ := parseResponseBody(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func runCLI(target string) {
	fmt.Printf("[+] OSINT Engine Target: %s\n", target)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Veriphone-Key, X-Ipqs-Key, X-Groq-Key")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
