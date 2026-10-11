# 🐹 Go-Learn: From ELI5 Beginner to Cyber Security Tool Builder

[![Go Version](https://img.shields.io/badge/Go-1.22%2B-00ADD8?style=for-the-badge&logo=go)](https://golang.org)
[![Security Focus](https://img.shields.io/badge/Domain-Cybersecurity%20%26%20OSINT-red?style=for-the-badge&logo=shield)](https://github.com/CYBERWOLF333/Go-Learn)
[![Course Level](https://img.shields.io/badge/Level-5--Year--Old%20Friendly%20to%20Pro-success?style=for-the-badge)](#)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue?style=for-the-badge)](LICENSE)

Welcome to **Go-Learn**! This repository is an all-in-one hands-on learning roadmap for mastering the Go programming language from scratch. 

It is designed with two simple goals:
1. **Explain Go concepts like you're 5 years old (ELI5)**: No dry academic jargon. Every concept is explained using playful everyday analogies (toy boxes, treasure maps, lego blueprints, and armies of helper ants) backed by clean, runnable code.
2. **Step-by-step Cyber Engineering**: Progress through hands-on cybersecurity tools (cryptographic hashing, HTTP security auditors, concurrent port scanners, raw packet inspectors), culminating in a **production-ready OSINT reconnaissance engine ready to publish as its own standalone repository**!

---

## 🗺️ Curriculum Roadmap

```text
  [ Part 1: Go for 5-Year-Olds ]
     ├── 01. Hello Gopher (The Recipe Card & Megaphone)
     ├── 02. Variables & Boxes (Data Types & Strict Toy Boxes)
     ├── 03. Traffic Lights & Loops (If/Else & The One True For Loop)
     ├── 04. Vending Machines (Functions & Multiple Returns)
     ├── 05. Treasure Maps (Pointers & Memory Addresses)
     ├── 06. Backpacks & Blueprints (Slices, Maps & Structs)
     ├── 07. Superhero Badges (Methods & Universal Interfaces)
     ├── 08. Army of Ants (Goroutines, Channels & Concurrency)
     ├── 09. Safety First (Error Handling, Defer & Recovery)
     └── 10. Secret Envelopes (JSON Files & HTTP Networking)
                  │
                  ▼
  [ Part 2: Hands-On Cyber Mini-Projects ]
     ├── 🛡️  01. HashSentry (File Integrity & Checksum Auditor)
     ├── 🌐 02. HeaderAuditor (OWASP Security Headers & Grade Scoring)
     ├── 🚪 03. StealthPortScanner (Concurrent TCP Worker Pool Scanner)
     └── 🔎 04. PacketLens (Pure-Go Ethernet/IP/TCP Header Dissector)
                  │
                  ▼
  [ Part 3: The Capstone Flagship Project ]
     └── 🦫 GopherRecon (Passive OSINT & Attack Surface Discovery Engine)
            * Standalone Go module ready to be published as its own repository!
```

---

## 🧒 Part 1: Go for 5-Year-Olds (ELI5 Basics)

Every lesson lives in `01-basics-eli5/` and contains a beginner-friendly `README.md` and an interactive `main.go`.

| # | Lesson Folder | The 5-Year-Old Analogy | What You Learn in Go | Command to Run |
| :-: | :--- | :--- | :--- | :--- |
| **01** | `01-hello-gopher` | **The Recipe Card & Megaphone** | Packages, imports, `main()` entrypoint, `fmt.Println` | `go run 01-basics-eli5/01-hello-gopher/main.go` |
| **02** | `02-variables-and-boxes` | **Strict Toy Boxes** | `string`, `int`, `bool`, `float64`, Walrus `:=`, zero values, constants | `go run 01-basics-eli5/02-variables-and-boxes/main.go` |
| **03** | `03-traffic-lights-and-loops` | **Traffic Cop & Merry-Go-Round** | `if/else`, short if statements, `switch/case`, `for` loops | `go run 01-basics-eli5/03-traffic-lights-and-loops/main.go` |
| **04** | `04-magic-vending-machines-functions` | **Dual-Dispense Vending Machine** | Parameter passing, multiple return values `(val, err)`, named returns, variadic args | `go run 01-basics-eli5/04-magic-vending-machines-functions/main.go` |
| **05** | `05-treasure-maps-pointers` | **Treasure Maps vs. Gold Chests** | Memory addresses `&`, dereferencing `*`, pass-by-value vs pass-by-pointer | `go run 01-basics-eli5/05-treasure-maps-pointers/main.go` |
| **06** | `06-backpacks-and-blueprints-slices-structs` | **Stretchy Backpacks & Blueprints** | Slices, `append`, maps (key-value dictionaries), custom structs | `go run 01-basics-eli5/06-backpacks-and-blueprints-slices-structs/main.go` |
| **07** | `07-superhero-badges-interfaces` | **Universal USB-C Plugs** | Struct methods, interface contracts, polymorphism without inheritance | `go run 01-basics-eli5/07-superhero-badges-interfaces/main.go` |
| **08** | `08-army-of-ants-goroutines-channels` | **Army of Helper Ants & Walkie-Talkies** | `go` keyword, Goroutines, Channels, `sync.WaitGroup` | `go run 01-basics-eli5/08-army-of-ants-goroutines-channels/main.go` |
| **09** | `09-safety-first-errors-and-defer` | **Checking Shoes & Toy Cleanup** | Explicit error checks `if err != nil`, `defer` cleanup, panic & recover | `go run 01-basics-eli5/09-safety-first-errors-and-defer/main.go` |
| **10** | `10-secret-messages-json-and-web` | **Secret Notes & Messenger Pigeons** | JSON serialization (`Marshal`/`Unmarshal`), File I/O, `net/http` client & server | `go run 01-basics-eli5/10-secret-messages-json-and-web/main.go` |

---

## 🛡️ Part 2: Step-by-Step Cyber Mini-Projects

Located in `02-cyber-projects/`, each project is an educational, practical cyber tool featuring unit tests (`go test -v ./...`):

### 1. [HashSentry (`02-cyber-projects/01-hash-sentry`)](02-cyber-projects/01-hash-sentry/)
- **Concept:** Digital fingerprints and integrity monitoring.
- **Modules Used:** `crypto/sha256`, `crypto/md5`, `encoding/hex`, `io`, `os`.
- **Capabilities:** Computes SHA-256 and MD5 hashes with streaming memory-efficient buffers; verifies file integrity against expected baseline hashes to detect malware tampering or unauthorized modifications.

### 2. [HeaderAuditor (`02-cyber-projects/02-header-auditor`)](02-cyber-projects/02-header-auditor/)
- **Concept:** Web server armor inspection.
- **Modules Used:** `net/http`, `crypto/tls`, `strings`.
- **Capabilities:** Audits websites for critical OWASP defensive headers (HSTS, CSP, X-Frame-Options, X-Content-Type-Options, Referrer-Policy, Permissions-Policy). Calculates a security score (0-100), awards letter grades (**A+** through **F**), and flags sensitive server banner leaks.

### 3. [StealthPortScanner (`02-cyber-projects/03-stealth-port-scanner`)](02-cyber-projects/03-stealth-port-scanner/)
- **Concept:** High-speed network recon with concurrency.
- **Modules Used:** `net.DialTimeout`, `sync.WaitGroup`, buffered channels.
- **Capabilities:** Scans ranges of TCP ports concurrently using the Go **Worker Pool pattern**. Captures live service banners (SSH, HTTP) and handles offline ports with strict connection timeouts.

### 4. [PacketLens (`02-cyber-projects/04-packet-lens`)](02-cyber-projects/04-packet-lens/)
- **Concept:** Low-level network protocol dissection.
- **Modules Used:** `encoding/binary`, bitwise shift/mask operations, `net.HardwareAddr`, `net.IP`.
- **Capabilities:** A pure-Go packet parser that dissects raw Ethernet frames, IPv4 packets, and TCP segments down to individual bit flags (SYN, ACK, FIN, RST, PSH, URG), sequence numbers, and payloads without needing external C-bindings (no libpcap required).

---

## 🦫 Part 3: Capstone Final Project — `GopherRecon`

Inside `03-final-project-gophercon/` is **GopherRecon**, a full-featured, modular **Passive OSINT & Attack Surface Discovery Engine**.

It is completely isolated as its own independent Go module (`github.com/CYBERWOLF333/gopherrecon`), complete with:
- Dedicated `go.mod`
- Modular architecture (`pkg/banner`, `pkg/dnsrecon`, `pkg/tlsrecon`, `pkg/headerrecon`, `pkg/ctrecon`, `pkg/reporter`)
- ANSI colored terminal output & progress banners
- Concurrent scanning engine using Go Goroutines
- Full unit test coverage (`make test`)
- Multi-format exports: structured JSON and executive Markdown audit reports
- Standalone `README.md`, `LICENSE` (MIT), and `Makefile`

### How to Run GopherRecon
```bash
# Enter project directory
cd 03-final-project-gophercon

# Run the interactive built-in simulation demo
go run ./cmd/gopherrecon -demo

# Or build the binary
make build
./bin/gopherrecon -domain github.com -json scan.json -markdown report.md
```

### How to Publish GopherRecon as Its Own Repository
Because `03-final-project-gophercon/` is already a self-contained module, you can publish it to GitHub in 3 steps:

```bash
# 1. Create a new empty repository on your GitHub account (e.g., GopherRecon)
# 2. Push this folder to that new repository:
cd 03-final-project-gophercon
git init
git add .
git commit -m "feat: initial commit of GopherRecon OSINT engine"
git remote add origin https://github.com/CYBERWOLF333/GopherRecon.git
git branch -M main
git push -u origin main
```

---

## ⚡ Master Quickstart & Commands

You can run and test everything from the root of this repository using the provided `Makefile`:

```bash
# Run all unit tests across the entire course and all projects
make test

# Run through all 10 ELI5 Go fundamentals lessons
make run-basics

# Run all 4 cyber mini-projects
make run-cyber

# Build and run the flagship GopherRecon engine in demo mode
make run-final
```

---

## 🤝 Contributing & Next Steps
Feel free to add your own cyber modules, port scanners, fuzzers, or cryptanalysis scripts! Have fun exploring Go and stay ethical! 🚀
