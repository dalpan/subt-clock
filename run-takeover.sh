#!/bin/bash

# Check dependencies
command -v subfinder >/dev/null 2>&1 || { echo >&2 "subfinder is required but not found."; exit 1; }
command -v httpx >/dev/null 2>&1 || { echo >&2 "httpx is required but not found."; exit 1; }
command -v ./subt-cloak >/dev/null 2>&1 || { echo >&2 "subt-cloak binary not found in current directory."; exit 1; }

# Input domain
TARGET=$1
if [[ -z "$TARGET" ]]; then
    echo "[!] Usage: $0 example.com"
    exit 1
fi

# Output files
TIMESTAMP=$(date +%Y%m%d_%H%M%S)
DOMAINS="domains_$TARGET.txt"
LIVE="live_$TARGET.txt"
RAW_OUTPUT="takeover_results_$TARGET_$TIMESTAMP.yaml"
FILTERED_OUTPUT="flagged_takeovers_$TARGET_$TIMESTAMP.yaml"

echo "[+] Enumerating subdomains for: $TARGET"
subfinder -d "$TARGET" -silent > "$DOMAINS"

echo "[+] Probing live subdomains..."
cat "$DOMAINS" | httpx -silent -cname -status-code -mc 200,403,404 > "$LIVE"

echo "[+] Extracting domains for evasive scan..."
cat "$LIVE" | awk '{print $1}' > targets.txt

echo "[+] Running evasive takeover detection with subt-cloak..."
./subt-cloak --detect-evade --input targets.txt --output "$RAW_OUTPUT"

echo "[+] Filtering flagged takeovers (fake-200-response, stripped-headers)..."
grep -B 5 -A 5 'evasion_signatures' "$RAW_OUTPUT" | grep -Ei 'subdomain:|cname:|evasion_signatures|fake-200|stripped-headers' > "$FILTERED_OUTPUT"

echo "[✓] Done."
echo "    Full results:     $RAW_OUTPUT"
echo "    Flagged findings: $FILTERED_OUTPUT"
