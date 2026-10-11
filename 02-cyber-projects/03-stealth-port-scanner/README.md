# 🚪 Cyber Project 03: StealthPortScanner (Concurrent TCP Port & Service Scanner)

## 🧒 What are we building? (Explain Like I'm 5)
Imagine a massive apartment building with **65,535 doors** (ports!).
Behind some doors, someone is awake and serving tea (open port ☕). Behind other doors, the lights are off and the lock is turned (closed port 🔒).

If one detective walked up and knocked on all 65,535 doors one by one, it would take **days**!
Instead, what if you hired a team of **100 speedy detective gophers (Goroutines)** who each knock on different doors simultaneously?
The whole building can be inspected in just a few seconds!

`StealthPortScanner` is a Go tool that:
1. Implements the **Worker Pool Pattern**: a squad of goroutines pulling port numbers from a shared job channel.
2. Performs TCP handshakes with strict timeouts (so closed ports don't stall the scan).
3. If a door opens, it performs **Banner Grabbing** to see what service is running inside!

---

## 🛠️ Go Concepts & Modules Used
- `net.DialTimeout`: Safe non-blocking TCP network connection attempts.
- `sync.WaitGroup`: Synchronizing worker goroutines.
- `channels`: Dispatching work and collecting findings safely across threads.
- `sort`: Sorting open ports in numerical order for clean reporting.

---

## 🏃 Quickstart & Examples

### 1. Run in Demo Mode (Spins up mock local services and scans them!)
```bash
go run main.go
```

### 2. Scan a specific host (e.g., localhost or internal IP)
```bash
go run main.go -host 127.0.0.1 -start 20 -end 1024 -workers 50 -timeout 300ms
```

### 3. Run automated tests
```bash
go test -v ./...
```
