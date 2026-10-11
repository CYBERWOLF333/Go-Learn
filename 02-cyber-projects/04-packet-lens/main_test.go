package main

import (
	"bytes"
	"net"
	"testing"
)

func TestParseEthernet(t *testing.T) {
	raw := []byte{
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, // Dest MAC
		0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, // Src MAC
		0x08, 0x00, // IPv4
		0xff, 0xfe, // payload
	}

	eth, err := ParseEthernet(raw)
	if err != nil {
		t.Fatalf("ParseEthernet failed: %v", err)
	}

	expectedDest := net.HardwareAddr{0x01, 0x02, 0x03, 0x04, 0x05, 0x06}
	if !bytes.Equal(eth.DestMAC, expectedDest) {
		t.Errorf("expected DestMAC %s, got %s", expectedDest, eth.DestMAC)
	}

	if eth.EtherType != 0x0800 {
		t.Errorf("expected EtherType 0x0800, got 0x%04x", eth.EtherType)
	}
}

func TestParseIPv4(t *testing.T) {
	rawIP := []byte{
		0x45, 0x00, 0x00, 0x14, 0x00, 0x01, 0x00, 0x00,
		0x40, 0x06, 0x00, 0x00,
		192, 168, 1, 5,
		10, 0, 0, 1,
	}

	ip, err := ParseIPv4(rawIP)
	if err != nil {
		t.Fatalf("ParseIPv4 failed: %v", err)
	}

	if ip.Version != 4 {
		t.Errorf("expected IPv4, got %d", ip.Version)
	}
	if ip.HeaderLength != 20 {
		t.Errorf("expected HeaderLength 20, got %d", ip.HeaderLength)
	}
	if ip.SourceIP.String() != "192.168.1.5" {
		t.Errorf("expected SourceIP 192.168.1.5, got %s", ip.SourceIP)
	}
	if ip.DestIP.String() != "10.0.0.1" {
		t.Errorf("expected DestIP 10.0.0.1, got %s", ip.DestIP)
	}
	if ip.Protocol != 6 {
		t.Errorf("expected Protocol 6 (TCP), got %d", ip.Protocol)
	}
}

func TestParseTCP(t *testing.T) {
	rawTCP := []byte{
		0x00, 0x50, // Port 80
		0x1f, 0x90, // Port 8080
		0x00, 0x00, 0x00, 0x01, // Seq
		0x00, 0x00, 0x00, 0x00, // Ack
		0x50, // 20 bytes offset
		0x02, // SYN
		0x20, 0x00,
		0x00, 0x00,
		0x00, 0x00,
	}

	tcp, err := ParseTCP(rawTCP)
	if err != nil {
		t.Fatalf("ParseTCP failed: %v", err)
	}

	if tcp.SourcePort != 80 || tcp.DestPort != 8080 {
		t.Errorf("unexpected ports: src=%d, dst=%d", tcp.SourcePort, tcp.DestPort)
	}
	if !tcp.Flags.SYN {
		t.Errorf("expected SYN flag to be true")
	}
	if tcp.Flags.ACK {
		t.Errorf("expected ACK flag to be false")
	}
}
