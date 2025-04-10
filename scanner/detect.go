// subt-cloak/scanner/detect.go
package scanner

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type Result struct {
	Subdomain  string   `yaml:"subdomain"`
	CNAME      string   `yaml:"cname"`
	StatusCode int      `yaml:"status_code"`
	Headers    []string `yaml:"headers"`
	ContentMD5 string   `yaml:"content_hash"`
	Flags      []string `yaml:"evasion_signatures"`
}

func StartDetectEvadeMode(inputPath, outputPath string) {
	file, err := os.Open(inputPath)
	if err != nil {
		fmt.Printf("[!] Failed to read input: %v\n", err)
		return
	}
	defer file.Close()

	var results []Result
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		domain := scanner.Text()
		fmt.Printf("[>] Scanning %s...\n", domain)

		res, err := http.Get("http://" + domain)
		if err != nil {
			fmt.Printf("[-] Error fetching: %v\n", err)
			continue
		}
		defer res.Body.Close()

		// Hash body
		body, _ := io.ReadAll(res.Body)
		hash := md5.Sum(body)
		contentHash := hex.EncodeToString(hash[:])

		// Header profiling
		headers := []string{}
		for k, v := range res.Header {
			headers = append(headers, fmt.Sprintf("%s: %s", k, strings.Join(v, ", ")))
		}

		// Detect evasive flags
		flags := []string{}
		if res.StatusCode == 200 && strings.Contains(strings.ToLower(string(body)), "error") {
			flags = append(flags, "fake-200-response")
		}
		if !hasCommonHeaders(res.Header) {
			flags = append(flags, "stripped-headers")
		}

		results = append(results, Result{
			Subdomain:  domain,
			CNAME:      ResolveCNAME(domain),
			StatusCode: res.StatusCode,
			Headers:    headers,
			ContentMD5: contentHash,
			Flags:      flags,
		})

		time.Sleep(1 * time.Second)
	}

	writeYAML(outputPath, results)
	fmt.Printf("[✓] Scan complete. Output saved to %s\n", outputPath)
}

func hasCommonHeaders(h http.Header) bool {
	for key := range h {
		if strings.Contains(strings.ToLower(key), "server") ||
			strings.Contains(strings.ToLower(key), "x-powered-by") {
			return true
		}
	}
	return false
}
