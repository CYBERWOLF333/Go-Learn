package main

import (
	"os"
	"testing"
)

func TestHashString(t *testing.T) {
	input := "hello world"
	// Known standard SHA-256 for "hello world"
	expectedSHA256 := "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"
	gotSHA256, err := HashString(input, "sha256")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotSHA256 != expectedSHA256 {
		t.Errorf("expected %s, got %s", expectedSHA256, gotSHA256)
	}

	// Known standard MD5 for "hello world"
	expectedMD5 := "5eb63bbbe01eeed093cb22bb8f5acdc3"
	gotMD5, err := HashString(input, "md5")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMD5 != expectedMD5 {
		t.Errorf("expected %s, got %s", expectedMD5, gotMD5)
	}
}

func TestHashFile(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "hash-test-*.txt")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	content := "secure binary payload"
	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	expectedStrHash, _ := HashString(content, "sha256")
	fileHash, err := HashFile(tmpFile.Name(), "sha256")
	if err != nil {
		t.Fatalf("HashFile failed: %v", err)
	}

	if fileHash != expectedStrHash {
		t.Errorf("file hash %s != string hash %s", fileHash, expectedStrHash)
	}
}
