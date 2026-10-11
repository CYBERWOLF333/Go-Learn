package main

import "fmt"

// Package-level constants (fixed values that never change)
const (
	ToolName    = "CyberSentry"
	DefaultPort = 8080
	MaxRetries  = 3
)

func main() {
	fmt.Println("=== 📦 Lesson 02: Variables & Data Types ===")

	// 1. Long explicit declaration with 'var'
	var targetHost string = "scanme.nmap.org"
	var openPortsCount int = 4
	var packetLoss float64 = 0.05
	var firewallActive bool = true

	fmt.Printf("[Long Form] Target: %s, Ports: %d, Loss: %.2f%%, Firewall: %t\n",
		targetHost, openPortsCount, packetLoss*100, firewallActive)

	// 2. Short declaration with the Walrus `:=`
	agentAlias := "ZeroCool"
	discoveredVulnerabilities := 0
	signalStrength := 98.75
	isEncrypted := true

	fmt.Printf("[Walrus :=] Agent: %s | Vulns: %d | Signal: %.1f | Encrypted: %t\n",
		agentAlias, discoveredVulnerabilities, signalStrength, isEncrypted)

	// 3. Zero Values Demonstration (What happens when we don't assign anything?)
	var blankString string
	var blankInt int
	var blankBool bool
	var blankFloat float64

	fmt.Println("\n--- 👻 The Zero Values (Default Clean Slate) ---")
	fmt.Printf("Empty string: %q\n", blankString)
	fmt.Printf("Empty int:    %d\n", blankInt)
	fmt.Printf("Empty bool:   %t\n", blankBool)
	fmt.Printf("Empty float:  %.2f\n", blankFloat)

	// 4. Type conversions (Go never converts automatically for safety!)
	var packetsSent int = 250
	var packetsReceived int = 237
	// Must explicitly convert int to float64 to do float math!
	successRate := (float64(packetsReceived) / float64(packetsSent)) * 100.0

	fmt.Println("\n--- 🧮 Math with Strict Types ---")
	fmt.Printf("Packets: %d/%d (Success Rate: %.2f%%)\n",
		packetsReceived, packetsSent, successRate)

	fmt.Println("\n--- 🔒 Constants ---")
	fmt.Printf("Tool: %s (Default Port: %d, Max Retries: %d)\n",
		ToolName, DefaultPort, MaxRetries)
}
