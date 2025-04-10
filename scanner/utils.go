// subt-cloak/scanner/utils.go
package scanner

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"net"
	"os"
)

// writeYAML writes scan results to a YAML file.
func writeYAML(output string, data interface{}) {
	file, err := os.Create(output)
	if err != nil {
		fmt.Printf("[!] Failed to write YAML output: %v\n", err)
		return
	}
	defer file.Close()

	yamlBytes, err := yaml.Marshal(data)
	if err != nil {
		fmt.Printf("[!] YAML marshal error: %v\n", err)
		return
	}

	file.Write(yamlBytes)
}

// ResolveCNAME attempts to resolve CNAME for a domain
func ResolveCNAME(domain string) string {
	cname, err := net.LookupCNAME(domain)
	if err != nil {
		return ""
	}
	return cname
}
