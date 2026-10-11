# 🛡️ Cyber Project 02: HeaderAuditor (HTTP Security Headers Inspector)

## 🧒 What are we building? (Explain Like I'm 5)
Imagine a brave medieval knight going out into battle 🛡️.
- Does he have a helmet? (That's **Content-Security-Policy**!)
- Does he have chainmail? (That's **Strict-Transport-Security**!)
- Does he have a shield? (That's **X-Frame-Options**!)

If he walks outside wearing only pajamas and sneakers, any attacker can shoot arrows at him! 🏹

Whenever a web server sends a website to your browser, it includes hidden metadata lines called **HTTP Headers**. Good web servers include defensive security headers that tell your browser to turn on shield defenses against XSS, Clickjacking, and packet sniffing.

`HeaderAuditor` is a Go tool that:
1. Connects to any web server over HTTP/HTTPS.
2. Checks for the top 6 essential defensive headers.
3. Warns if the server leaks sensitive software version banners (`Server`, `X-Powered-By`).
4. Calculates a Security Score (0 to 100) and awards a Report Card Grade: **A+**, **A**, **B**, **C**, **D**, or **F**!

---

## 🛠️ Go Concepts & Modules Used
- `net/http`: Custom HTTP clients with redirects handling, user agents, and timeouts.
- `crypto/tls`: TLS configuration and inspection.
- `strings`, `fmt`: Formatting colored terminal audit reports.
- `net/http/httptest`: Testing HTTP responses with mock servers.

---

## 🏃 Quickstart & Examples

### 1. Run in Demo Mode (Audits an internal secure vs insecure mock server)
```bash
go run main.go
```

### 2. Audit a live website
```bash
go run main.go -url https://github.com
```

### 3. Run automated tests
```bash
go test -v ./...
```
