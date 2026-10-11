package main

import "fmt"

func main() {
	fmt.Println("=== 🚦 Lesson 03: Control Flow & Loops ===")

	// 1. Classic if / else if / else
	threatLevel := 75 // scale from 1 to 100

	if threatLevel > 80 {
		fmt.Println("🚨 CRITICAL: Intrusion detected! Lockdown initiated!")
	} else if threatLevel >= 50 {
		fmt.Println("⚠️  WARNING: Suspicious activity logged. Monitoring closely.")
	} else {
		fmt.Println("✅ SAFE: All network perimeters normal.")
	}

	// 2. Go Special Feature: 'if' with a short statement!
	// You can create a temporary variable right inside the if header!
	if pingMs := 42; pingMs < 50 {
		fmt.Printf("⚡ Superfast connection: %dms ping\n", pingMs)
	}

	// 3. Switch statement: The elevator panel
	protocol := "HTTPS"

	switch protocol {
	case "HTTP":
		fmt.Println("⚠️  Warning: HTTP is unencrypted cleartext! Packets can be sniffed.")
	case "HTTPS":
		fmt.Println("🔒 Secure: TLS encryption enabled (Port 443).")
	case "SSH":
		fmt.Println("🛡️ Secure Shell remote admin session (Port 22).")
	default:
		fmt.Println("❓ Unknown protocol:", protocol)
	}

	// 4. The Classic 'for' Loop (Counting)
	fmt.Println("\n--- 🔄 Port Scanning Simulation (Classic For Loop) ---")
	for port := 80; port <= 83; port++ {
		fmt.Printf("Testing TCP port %d... [Simulated Probe Sent]\n", port)
	}

	// 5. While-style Loop (using 'for' with only a condition)
	fmt.Println("\n--- 🔋 While-style Battery Depletion ---")
	droneBattery := 3
	for droneBattery > 0 {
		fmt.Printf("Recon drone in flight... Battery: %d\n", droneBattery)
		droneBattery--
	}
	fmt.Println("Drone landed safely at base.")

	// 6. Infinite Loop with 'break' and 'continue'
	fmt.Println("\n--- 🛑 Infinite Loop with Break & Continue ---")
	attempt := 0
	for {
		attempt++
		if attempt == 2 {
			fmt.Printf("Attempt %d had packet collision! Retrying (continue)...\n", attempt)
			continue
		}
		fmt.Printf("Attempt %d succeeded!\n", attempt)
		if attempt >= 4 {
			fmt.Println("Handshake complete. Breaking out of loop!")
			break
		}
	}
}
