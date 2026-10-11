# 🔍 Cyber Project 04: PacketLens (Pure-Go Packet Header Inspector)

## 🧒 What are we building? (Explain Like I'm 5)
Imagine you want to mail a secret birthday postcard to your friend 📬:
1. **The Post Office Van (Ethernet Layer)**: Moves the envelope between local post offices (uses MAC addresses).
2. **The Street Address (IP Layer)**: Has your friend's house address and your return address (IP addresses, TTL / Time-to-Live).
3. **The Bedroom Door Number (TCP / UDP Layer)**: Tells your friend which room to deliver the letter to (Port numbers) and includes a handshake ticket (SYN / ACK flags).
4. **The Letter Message (Payload)**: The actual text: `"Happy Birthday!"` or `"GET /index.html"`.

In Wireshark or tcpdump, programs slice up raw bytes into these logical layers.
`PacketLens` is a **Pure-Go Packet Parser** that decodes raw packet byte buffers from scratch without needing external C libraries (like libpcap).

---

## 🛠️ Go Concepts & Modules Used
- `encoding/binary`: Reading Big-Endian network byte order (`binary.BigEndian.Uint16`).
- Bitwise operators (`&`, `>>`): Extracting sub-byte header flags like IP Version, Header Length, and TCP Flags (SYN, ACK, FIN, RST, PSH, URG).
- `net.IP`: Formatting raw 4-byte and 16-byte arrays into clean dotted-decimal IP strings (`192.168.1.1`).
- Byte slicing & offset tracking.

---

## 🏃 Quickstart & Examples

### 1. Run the Built-in Inspection Demo (Inspects a simulated TCP SYN Handshake & UDP DNS packet)
```bash
go run main.go
```

### 2. Run automated tests
```bash
go test -v ./...
```
