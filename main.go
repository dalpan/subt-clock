// subt-cloak/main.go
package main

import (
	"flag"
	"fmt"
	"os"

	"subt-cloak/scanner"
)

func main() {
	cloak := flag.Bool("cloak", false, "Enable attacker-side cloaking simulation")
	detect := flag.Bool("detect-evade", false, "Enable evasive takeover detection")
	input := flag.String("input", "", "Input file with subdomains")
	template := flag.String("template", "", "Service template to use in --cloak mode")
	rotateStatus := flag.Bool("rotate-status", false, "Rotate HTTP status codes in --cloak mode")
	fakeHeaders := flag.Bool("fake-headers", false, "Apply fake headers in --cloak mode")
	dnsRotate := flag.Bool("dns-rotate", false, "Enable rotating DNS responses (NXDOMAIN/Valid)")
	output := flag.String("output", "result.yaml", "Output file (YAML format)")

	flag.Parse()

	if *cloak {
		scanner.StartCloakMode(*template, *rotateStatus, *fakeHeaders, *dnsRotate)
	} else if *detect {
		if *input == "" {
			fmt.Println("--input is required for --detect-evade mode")
			os.Exit(1)
		}
		scanner.StartDetectEvadeMode(*input, *output)
	} else {
		fmt.Println("[!] Please specify either --cloak or --detect-evade mode")
		flag.Usage()
		os.Exit(1)
	}
}
