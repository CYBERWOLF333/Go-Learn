# 🛡️ Cyber Project 01: HashSentry (File Integrity & IOC Verifier)

## 🧒 What are we building? (Explain Like I'm 5)
Have you ever pressed your thumb into wet paint and stamped a white sheet of paper?
That fingerprint belongs to **you and only you**.

In the computer world, files have digital fingerprints called **Cryptographic Hashes** (like SHA-256 or MD5).
- If a file is 10 Gigabytes long, its SHA-256 fingerprint is still a tiny 64-character text string!
- If a hacker sneaks into a system and changes even **one single letter** inside a file, the whole fingerprint changes completely! (This is called the **Avalanche Effect** 🏔️).

`HashSentry` is a command-line tool that:
1. Calculates cryptographic checksums (SHA-256, MD5) for any file or raw text.
2. Compares files against a known **Integrity Baseline** to spot unauthorized tampering.
3. Checks hashes against a list of known malware hashes (Indicators of Compromise / IOCs).

---

## 🛠️ Go Concepts & Modules Used
- `crypto/sha256`, `crypto/md5`: High-performance Go crypto primitives.
- `encoding/hex`: Turning raw binary hash digests into readable hex strings.
- `io`, `os`: Efficient streaming file I/O (handling gigabyte files without crashing memory!).
- `flag`: Command line arguments parsing.

---

## 🏃 Quickstart & Examples

### 1. Hash a text message
```bash
go run main.go -text "SuperSecretPassword123"
```

### 2. Hash a file using SHA-256
```bash
go run main.go -file main.go
```

### 3. Verify a file against an expected hash
```bash
go run main.go -file main.go -check "YOUR_EXPECTED_HASH_HERE"
```

### 4. Run automated tests
```bash
go test -v ./...
```
