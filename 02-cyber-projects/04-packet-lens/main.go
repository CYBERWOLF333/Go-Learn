package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"strings"
)

// EthernetFrame represents Layer 2 Header (14 bytes)
type EthernetFrame struct {
	DestMAC   net.HardwareAddr
	SourceMAC net.HardwareAddr
	EtherType uint16
	Payload   []byte
}

// IPv4Header represents Layer 3 IPv4 Header (minimum 20 bytes)
type IPv4Header struct {
	Version        uint8
	HeaderLength   int // in bytes
	TypeOfService  uint8
	TotalLength    uint16
	Identification uint16
	Flags          uint8
	TTL            uint8
	Protocol       uint8
	Checksum       uint16
	SourceIP       net.IP
	DestIP         net.IP
	Payload        []byte
}

// TCPHeader represents Layer 4 TCP Header (minimum 20 bytes)
type TCPHeader struct {
	SourcePort uint16
	DestPort   uint16
	SeqNum     uint32
	AckNum     uint32
	DataOffset int // in bytes
	Flags      TCPFlags
	WindowSize uint16
	Checksum   uint16
	Payload    []byte
}

type TCPFlags struct {
	URG, ACK, PSH, RST, SYN, FIN bool
}

func (f TCPFlags) String() string {
	var active []string
	if f.SYN {
		active = append(active, "SYN")
	}
	if f.ACK {
		active = append(active, "ACK")
	}
	if f.FIN {
		active = append(active, "FIN")
	}
	if f.RST {
		active = append(active, "RST")
	}
	if f.PSH {
		active = append(active, "PSH")
	}
	if f.URG {
		active = append(active, "URG")
	}
	if len(active) == 0 {
		return "NONE"
	}
	return strings.Join(active, "|")
}

// ParseEthernet decodes the 14-byte Ethernet frame
func ParseEthernet(b []byte) (*EthernetFrame, error) {
	if len(b) < 14 {
		return nil, fmt.Errorf("packet too short for Ethernet frame: %d bytes (min 14)", len(b))
	}

	frame := &EthernetFrame{
		DestMAC:   net.HardwareAddr(b[0:6]),
		SourceMAC: net.HardwareAddr(b[6:12]),
		EtherType: binary.BigEndian.Uint16(b[12:14]),
		Payload:   b[14:],
	}
	return frame, nil
}

// ParseIPv4 decodes an IPv4 packet header
func ParseIPv4(b []byte) (*IPv4Header, error) {
	if len(b) < 20 {
		return nil, fmt.Errorf("packet too short for IPv4 header: %d bytes (min 20)", len(b))
	}

	versionAndIHL := b[0]
	version := versionAndIHL >> 4
	ihlWords := versionAndIHL & 0x0F
	headerLenBytes := int(ihlWords) * 4

	if headerLenBytes < 20 || headerLenBytes > len(b) {
		return nil, fmt.Errorf("invalid IPv4 header length: %d bytes", headerLenBytes)
	}

	header := &IPv4Header{
		Version:        version,
		HeaderLength:   headerLenBytes,
		TypeOfService:  b[1],
		TotalLength:    binary.BigEndian.Uint16(b[2:4]),
		Identification: binary.BigEndian.Uint16(b[4:6]),
		Flags:          b[6] >> 5,
		TTL:            b[8],
		Protocol:       b[9],
		Checksum:       binary.BigEndian.Uint16(b[10:12]),
		SourceIP:       net.IP(b[12:16]),
		DestIP:         net.IP(b[16:20]),
		Payload:        b[headerLenBytes:],
	}
	return header, nil
}

// ParseTCP decodes a TCP segment header
func ParseTCP(b []byte) (*TCPHeader, error) {
	if len(b) < 20 {
		return nil, fmt.Errorf("packet too short for TCP header: %d bytes (min 20)", len(b))
	}

	dataOffsetWords := b[12] >> 4
	dataOffsetBytes := int(dataOffsetWords) * 4
	if dataOffsetBytes < 20 || dataOffsetBytes > len(b) {
		return nil, fmt.Errorf("invalid TCP header offset: %d bytes", dataOffsetBytes)
	}

	flagsByte := b[13]
	flags := TCPFlags{
		URG: (flagsByte & 0x20) != 0,
		ACK: (flagsByte & 0x10) != 0,
		PSH: (flagsByte & 0x08) != 0,
		RST: (flagsByte & 0x04) != 0,
		SYN: (flagsByte & 0x02) != 0,
		FIN: (flagsByte & 0x01) != 0,
	}

	header := &TCPHeader{
		SourcePort: binary.BigEndian.Uint16(b[0:2]),
		DestPort:   binary.BigEndian.Uint16(b[2:4]),
		SeqNum:     binary.BigEndian.Uint32(b[4:8]),
		AckNum:     binary.BigEndian.Uint32(b[8:12]),
		DataOffset: dataOffsetBytes,
		Flags:      flags,
		WindowSize: binary.BigEndian.Uint16(b[14:16]),
		Checksum:   binary.BigEndian.Uint16(b[16:18]),
		Payload:    b[dataOffsetBytes:],
	}
	return header, nil
}

// InspectPacket performs full-stack dissection
func InspectPacket(raw []byte, label string) {
	fmt.Println("==================================================================")
	fmt.Printf(" 🔎 PACKET LENS INSPECTOR: %s (%d bytes)\n", label, len(raw))
	fmt.Println("==================================================================")

	// Layer 2: Ethernet
	eth, err := ParseEthernet(raw)
	if err != nil {
		fmt.Println("[-] L2 Error:", err)
		return
	}
	fmt.Printf("[L2 - Ethernet] %s -> %s (Type: 0x%04x)\n", eth.SourceMAC, eth.DestMAC, eth.EtherType)

	// 0x0800 is IPv4
	if eth.EtherType != 0x0800 {
		fmt.Printf("   Non-IPv4 EtherType: 0x%04x\n", eth.EtherType)
		return
	}

	// Layer 3: IPv4
	ip, err := ParseIPv4(eth.Payload)
	if err != nil {
		fmt.Println("[-] L3 Error:", err)
		return
	}
	fmt.Printf("[L3 - IPv4]     %s -> %s | TTL: %d | Proto: %d | Len: %d bytes\n",
		ip.SourceIP, ip.DestIP, ip.TTL, ip.Protocol, ip.TotalLength)

	// Protocol 6 = TCP, 17 = UDP
	switch ip.Protocol {
	case 6: // TCP
		tcp, err := ParseTCP(ip.Payload)
		if err != nil {
			fmt.Println("[-] L4 Error:", err)
			return
		}
		fmt.Printf("[L4 - TCP]      Port %d -> %d | Flags: [%s] | Seq: %d | Ack: %d\n",
			tcp.SourcePort, tcp.DestPort, tcp.Flags.String(), tcp.SeqNum, tcp.AckNum)
		if len(tcp.Payload) > 0 {
			fmt.Printf("[L7 - Payload]  %d bytes: %q\n", len(tcp.Payload), string(tcp.Payload))
		} else {
			fmt.Println("[L7 - Payload]  Empty (Control/Handshake packet)")
		}
	case 17: // UDP
		if len(ip.Payload) >= 8 {
			srcPort := binary.BigEndian.Uint16(ip.Payload[0:2])
			dstPort := binary.BigEndian.Uint16(ip.Payload[2:4])
			udpLen := binary.BigEndian.Uint16(ip.Payload[4:6])
			fmt.Printf("[L4 - UDP]      Port %d -> %d | Length: %d bytes\n", srcPort, dstPort, udpLen)
			payload := ip.Payload[8:]
			if len(payload) > 0 {
				fmt.Printf("[L7 - Payload]  %d bytes: %q\n", len(payload), string(payload))
			}
		}
	default:
		fmt.Printf("[L4 - Unknown]  Protocol ID %d\n", ip.Protocol)
	}
	fmt.Println()
}

func main() {
	// Sample 1: A raw synthesized TCP SYN handshake packet
	// L2: Dest 00:11:22:33:44:55, Src 66:77:88:99:aa:bb, EtherType 0x0800
	// L3: IPv4, 192.168.1.100 -> 93.184.216.34 (example.com), TTL 64, Proto TCP(6)
	// L4: TCP Src 54321 -> Dst 443 (HTTPS), Seq 1000, SYN flag set
	tcpSynPacket := []byte{
		// Ethernet Header (14 bytes)
		0x00, 0x11, 0x22, 0x33, 0x44, 0x55, // Dest MAC
		0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, // Src MAC
		0x08, 0x00, // EtherType IPv4

		// IPv4 Header (20 bytes)
		0x45,       // Version 4, IHL 5 (20 bytes)
		0x00,       // DSCP/ECN
		0x00, 0x28, // Total Length: 40 bytes (20 IP + 20 TCP)
		0x1a, 0x2b, // Identification
		0x40, 0x00, // Flags (Don't Fragment)
		0x40,       // TTL: 64
		0x06,       // Protocol: 6 (TCP)
		0x00, 0x00, // Checksum
		192, 168, 1, 100, // Source IP
		93, 184, 216, 34, // Dest IP (example.com)

		// TCP Header (20 bytes)
		0xd4, 0x31, // Src Port: 54321
		0x01, 0xbb, // Dst Port: 443 (HTTPS)
		0x00, 0x00, 0x03, 0xe8, // Sequence Number: 1000
		0x00, 0x00, 0x00, 0x00, // Ack Number: 0
		0x50,       // Data Offset: 5 words (20 bytes)
		0x02,       // Flags: SYN (bit 1)
		0x72, 0x10, // Window Size: 29200
		0x00, 0x00, // Checksum
		0x00, 0x00, // Urgent Pointer
	}

	// Sample 2: An HTTP GET request payload packet
	httpGetPacket := []byte{
		// Ethernet Header (14 bytes)
		0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff,
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66,
		0x08, 0x00,

		// IPv4 Header (20 bytes)
		0x45, 0x00,
		0x00, 0x47, // Total Length: 71 bytes
		0x3c, 0x4d,
		0x40, 0x00,
		0x80, // TTL: 128
		0x06, // TCP
		0x00, 0x00,
		10, 0, 0, 15,
		172, 217, 16, 206,

		// TCP Header (20 bytes)
		0x1f, 0x90, // Src Port: 8080
		0x00, 0x50, // Dst Port: 80
		0x00, 0x01, 0x23, 0x45, // Seq
		0x00, 0x02, 0x34, 0x56, // Ack
		0x50,       // Offset 20 bytes
		0x18,       // Flags: PSH + ACK
		0x10, 0x00,
		0x00, 0x00,
		0x00, 0x00,

		// Payload: "GET /admin HTTP/1.1\r\n\r\n"
		'G', 'E', 'T', ' ', '/', 'a', 'd', 'm', 'i', 'n', ' ', 'H', 'T', 'T', 'P', '/', '1', '.', '1', '\r', '\n', '\r', '\n',
	}

	InspectPacket(tcpSynPacket, "TCP SYN Connection Probe")
	InspectPacket(httpGetPacket, "In-Flight HTTP GET Request")
}
