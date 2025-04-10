// subt-cloak/scanner/cloak.go
package scanner

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var fakeHeadersList = map[string]string{
	"Server": "nginx",
	"X-Powered-By": "PHP/7.4.3",
	"X-Custom-Fake": "cloak-engine",
}

var statusPool = []int{200, 403, 404, 503}

func StartCloakMode(template string, rotateStatus, fakeHeaders, dnsRotate bool) {
	fmt.Println("[+] Cloak mode activated")
	fmt.Printf("[+] Using template: %s\n", template)
	if rotateStatus {
		fmt.Println("[~] Status rotation enabled")
	}
	if fakeHeaders {
		fmt.Println("[~] Fake headers applied")
	}
	if dnsRotate {
		fmt.Println("[~] DNS simulation mode active (NXDOMAIN/Valid toggle)")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if rotateStatus {
			rand.Seed(time.Now().UnixNano())
			status := statusPool[rand.Intn(len(statusPool))]
			w.WriteHeader(status)
		} else {
			w.WriteHeader(200)
		}

		if fakeHeaders {
			for key, value := range fakeHeadersList {
				w.Header().Set(key, value)
			}
		}

		// Load template HTML (mocked for now)
		body := LoadTemplate(template)
		fmt.Fprint(w, body)
	})

	fmt.Println("[!] Hosting cloaked takeover on http://0.0.0.0:8080")
	http.ListenAndServe(":8080", nil)
}

func LoadTemplate(service string) string {
	// TODO: Replace with file loading from templates/github.html, etc.
	switch service {
	case "github":
		return `<html><head><title>GitHub Pages</title></head><body><h1>This site is not yet published.</h1></body></html>`
	case "heroku":
		return `<html><head><title>Heroku</title></head><body><h1>No such app</h1></body></html>`
	default:
		return `<html><body><h1>Cloaked Subdomain</h1></body></html>`
	}
}
