package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"hash"
	"io"
	"os"
	"strings"
)

// HashString calculates the cryptographic digest of an input string
func HashString(input string, algorithm string) (string, error) {
	var h hash.Hash
	switch strings.ToLower(algorithm) {
	case "sha256", "sha-256":
		h = sha256.New()
	case "md5":
		h = md5.New()
	default:
		return "", fmt.Errorf("unsupported algorithm: %s (use 'sha256' or 'md5')", algorithm)
	}

	h.Write([]byte(input))
	return hex.EncodeToString(h.Sum(nil)), nil
}

// HashFile calculates the cryptographic digest of a file using a streaming buffer
func HashFile(filePath string, algorithm string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("cannot open file: %w", err)
	}
	defer file.Close()

	var h hash.Hash
	switch strings.ToLower(algorithm) {
	case "sha256", "sha-256":
		h = sha256.New()
	case "md5":
		h = md5.New()
	default:
		return "", fmt.Errorf("unsupported algorithm: %s", algorithm)
	}

	// io.Copy streams the file in small chunks directly to the hasher!
	// This uses almost ZERO memory even on massive multi-gigabyte files!
	if _, err := io.Copy(h, file); err != nil {
		return "", fmt.Errorf("failed to read file contents: %w", err)
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

func main() {
	filePath := flag.String("file", "", "Path to the file to inspect")
	textInput := flag.String("text", "", "Raw text string to hash")
	algorithm := flag.String("algo", "sha256", "Hashing algorithm: sha256 | md5")
	expectedHash := flag.String("check", "", "Expected hash digest to verify integrity against")
	flag.Parse()

	fmt.Println("==================================================")
	fmt.Println("  🛡️  HashSentry - File Integrity & Hash Auditor  ")
	fmt.Println("==================================================")

	if *filePath == "" && *textInput == "" {
		// Default demo run
		fmt.Println("ℹ️  No input specified. Running interactive demo mode...")
		demoText := "Attack at dawn via perimeter port 443"
		s256, _ := HashString(demoText, "sha256")
		m5, _ := HashString(demoText, "md5")

		fmt.Printf("Sample Payload: %q\n", demoText)
		fmt.Printf("  [SHA-256] -> %s\n", s256)
		fmt.Printf("  [MD-5]    -> %s\n\n", m5)
		fmt.Println("Tip: Pass '-file <path>' or '-text <str>' to audit custom files.")
		return
	}

	var computedHash string
	var err error

	if *filePath != "" {
		fmt.Printf("[+] Auditing File: %s (Algorithm: %s)\n", *filePath, *algorithm)
		computedHash, err = HashFile(*filePath, *algorithm)
		if err != nil {
			fmt.Printf("[-] Error: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Printf("[+] Auditing Raw Text (Algorithm: %s)\n", *algorithm)
		computedHash, err = HashString(*textInput, *algorithm)
		if err != nil {
			fmt.Printf("[-] Error: %v\n", err)
			os.Exit(1)
		}
	}

	fmt.Printf("[*] Computed Digest: %s\n", computedHash)

	if *expectedHash != "" {
		expectedClean := strings.ToLower(strings.TrimSpace(*expectedHash))
		computedClean := strings.ToLower(strings.TrimSpace(computedHash))

		fmt.Println("\n--- 🔍 INTEGRITY VERIFICATION ---")
		if expectedClean == computedClean {
			fmt.Println("✅ PASS: Hash matches expected digest! Integrity verified 100%.")
		} else {
			fmt.Println("🚨 ALERT: HASH MISMATCH! The file has been modified, corrupted, or tampered with!")
			fmt.Printf("   Expected: %s\n", expectedClean)
			fmt.Printf("   Actual:   %s\n", computedClean)
			os.Exit(2)
		}
	}
}
