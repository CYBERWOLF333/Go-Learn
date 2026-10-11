<div align="center">

# 🦫 GopherRecon
### Fast, Modular, Defensive OSINT & Attack Surface Discovery Engine

[![Go Report Card](https://img.shields.io/badge/go%20report-A%2B-brightgreen.svg)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/Go-1.22%2B-00ADD8?logo=go)](https://golang.org)
[![Security Audited](https://img.shields.io/badge/Security-OSINT%20Audited-red.svg)](#)

*A production-grade, zero-dependency cybersecurity tool written in pure Go for security engineers, penetration testers, and bug bounty hunters.*

```text
   ____             _               ____                          
  / ___| ___  _ __ | |__   ___ _ __|  _ \ ___  ___ ___  _ __  
 | |  _ / _ \| '_ \| '_ \ / _ \ '__| |_) / _ \/ __/ _ \| '_ \ 
 | |_| | (_) | |_) | | | |  __/ |  |  _ <  __/ (_| (_) | | | |
  \____|\___/| .__/|_| |_|\___|_|  |_| \_\___|\___\___/|_| |_|
             |_|                                              
       >>> Passive OSINT & Attack Surface Discovery Engine <<<
```

</div>

---

## ⚡ What is GopherRecon?

**GopherRecon** is an all-in-one attack surface recon engine designed to gather actionable defensive intelligence about a target domain **without touching intrusive endpoints or triggering defensive alarms**. 

By leveraging Go's lightweight concurrency (`goroutines` and `channels`), GopherRecon launches multi-vector discovery probes in parallel and delivers an executive security audit report in milliseconds.

---

## 🚀 Key Modules & Capabilities

| Module | What It Inspects | Cyber Value |
| :--- | :--- | :--- |
| **🌐 DNS & Email Security** | A, AAAA, MX, NS records, SPF and DMARC policies. | Detects missing SPF/DMARC configurations that allow attackers to spoof company emails. |
| **🔒 TLS/SSL Cryptography** | Active TLS version, cipher suite, cert issuer, expiration countdown, and SANs. | Flags obsolete TLS 1.0/1.1 protocols and mines **Subject Alternative Names** for hidden assets. |
| **🛡️ Defensive HTTP Headers** | OWASP Top Defensive Headers (HSTS, CSP, X-Frame-Options, etc.). | Scores target defense posture (0-100) and calculates letter grades (**A+** to **F**). |
| **🛰️ Passive CT Log Mining** | Certificate Transparency logs (crt.sh). | Discovers subdomains purely passively without sending any packets to target infrastructure! |
| **📊 Multi-Format Reporting** | ANSI colored Terminal UI, structured JSON, and executive Markdown. | Ready to attach directly to audit reports, bug bounty submissions, or tickets. |

---

## 🏗️ Concurrent Architecture

```text
                    +---------------------------+
                    |  Target: "example.com"    |
                    +-------------+-------------+
                                  |
              +-------------------+-------------------+
              |                   |                   |
              v                   v                   v
      [ Goroutine 1 ]     [ Goroutine 2 ]     [ Goroutine 3 ]
       DNS & SPF/DMARC      TLS & SANs        HTTP Headers
              |                   |                   |
              +-------------------+-------------------+
                                  |
                                  v
                      [ Aggregator & Reporter ]
                                  |
             +--------------------+--------------------+
             |                    |                    |
             v                    v                    v
      [ Terminal UI ]       [ JSON Export ]     [ Markdown Doc ]
```

---

## 📦 Installation

### Prerequisites
- Go 1.22 or newer installed on Linux, macOS, or Windows.

### Option 1: Direct Build
```bash
git clone https://github.com/CYBERWOLF333/gopherrecon.git
cd gopherrecon
make build
./bin/gopherrecon -demo
```

### Option 2: Go Install
```bash
go install github.com/CYBERWOLF333/gopherrecon/cmd/gopherrecon@latest
```

---

## 💻 Usage & CLI Options

```bash
gopherrecon [flags]
```

### CLI Flags Reference

| Flag | Default | Description |
| :--- | :--- | :--- |
| `-domain <string>` | `""` | Target domain to analyze (e.g., `github.com`) |
| `-demo` | `false` | Run interactive simulation demonstrating all recon modules |
| `-json <path>` | `""` | Export findings to a structured JSON file |
| `-markdown <path>` | `""` | Export findings to an executive Markdown report |
| `-timeout <duration>` | `5s` | Network timeout for probes (e.g., `3s`, `10s`) |
| `-skip-dns` | `false` | Disable DNS & Email security inspection |
| `-skip-tls` | `false` | Disable SSL/TLS certificate inspection |
| `-skip-headers` | `false` | Disable HTTP defensive headers audit |
| `-skip-ct` | `false` | Disable Certificate Transparency log enumeration |

---

## 🌟 Examples

### 1. Interactive Demo
```bash
./bin/gopherrecon -demo
```

### 2. Full Audit with JSON & Markdown Export
```bash
./bin/gopherrecon -domain targetcorp.com -json target_audit.json -markdown target_report.md
```

### 3. Fast Header-Only Audit
```bash
./bin/gopherrecon -domain targetcorp.com -skip-dns -skip-tls -skip-ct
```

---

## 🧪 Testing

Run all unit tests with:

```bash
make test
# or
go test -v ./...
```

---

## ⚖️ Ethics & Disclaimer

GopherRecon is intended strictly for authorized security assessments, defensive hardening, educational study, and security research. Only scan systems that you own or have explicit written authorization to evaluate. The authors assume no liability for misuse.
