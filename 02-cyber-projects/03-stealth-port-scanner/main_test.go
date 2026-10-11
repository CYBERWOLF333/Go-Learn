package main

import (
	"net"
	"testing"
	"time"
)

func TestScanPortsWithLocalListener(t *testing.T) {
	// Spin up an ephemeral listener
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}
	defer listener.Close()

	port := listener.Addr().(*net.TCPAddr).Port

	// Scan just this port
	results := ScanPorts("127.0.0.1", port, port, 1, 200*time.Millisecond, false)
	if len(results) != 1 {
		t.Fatalf("expected 1 open port, got %d", len(results))
	}

	if results[0].Port != port {
		t.Errorf("expected port %d, got %d", port, results[0].Port)
	}
	if !results[0].IsOpen {
		t.Errorf("expected port to be open")
	}
}
