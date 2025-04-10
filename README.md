# SubT-Cloak

**SubT-Cloak** is an advanced subdomain takeover analysis and cloaking simulation tool designed for red teamers, bug bounty hunters, and defenders alike. It expands traditional takeover scanning by introducing stealth, persistence, and deception into the equation.

---

## ✨ Features

🔹 `--cloak` Mode (Attacker Simulation)
- Simulates evasive takeover behavior
- Rotate status codes (200/403/404/503)
- Fake HTTP headers (remove fingerprints)
- Serve realistic spoofed HTML templates
- DNS cloaking simulation toggle (coming soon)

🔹 `--detect-evade` Mode (Defender Scanner)
- Detect evasive takeovers missed by traditional tools
- Analyze header anomalies, content hashes, status code rotation
- CNAME resolution and fingerprinting
- YAML output format

---

## 📦 Install

```bash
git clone https://github.com/tegal1337/subt-cloak.git
cd subt-cloak
go build -o subt-cloak main.go
```

---

## 🚀 Usage

### Attacker Cloak Mode
```bash
./subt-cloak --cloak --template github --rotate-status --fake-headers
```

### Defender Detect Mode
```bash
./subt-cloak --detect-evade --input domains.txt --output result.yaml
```

---

## 🔁 Automation Example

```bash
chmod +x run-takeover.sh
./run-takeover.sh targetdomain.com
```

Generates:
- All subdomains from subfinder
- Probes live with httpx
- Filters evasive takeovers

---

## 🧪 Fingerprint Format (YAML)
```yaml
- cname: github.io
  status_code: 200
  status: vulnerable can be takeover!
  headers:
    - Server: GitHub.com
  content:
    - "There isn't a GitHub Pages site here."
  evasion_signatures:
    - fake-200-response
    - stripped-headers
```

---

## 📁 Directory Structure
```
subt-cloak/
├── main.go
├── scanner/
│   ├── cloak.go
│   ├── detect.go
│   └── utils.go
├── templates/
│   ├── github.html
│   └── heroku.html
├── fingerprints/
│   └── services.yaml
├── run-takeover.sh
└── README.md
```

---

## 📚 Paper Reference
This tool is based on the research paper:
> **"SubT-Cloak: Persistent Subdomain Takeover via Stealth and Deception"** (For Black Hat Submission 2026)

---

## 💬 License
MIT License. Feel free to fork, adapt, or contribute.

---

## 🧠 Author
[@tegal1337](https://github.com/tegal1337) | Security Researcher | TegalSec
