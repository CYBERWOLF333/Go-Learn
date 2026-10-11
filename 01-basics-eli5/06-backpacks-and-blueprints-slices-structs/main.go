package main

import "fmt"

// Blueprint for a Target in a cyber security assessment
type TargetServer struct {
	HostName      string
	IPAddress     string
	OpenPorts     []int
	IsProtected   bool
	RiskScore     int
}

func main() {
	fmt.Println("=== 🎒 Lesson 06: Slices, Maps & Structs ===")

	// 1. Slices: The stretchy backpack
	var openPorts []int
	openPorts = append(openPorts, 22)
	openPorts = append(openPorts, 80, 443, 8080)

	fmt.Println("Ports in backpack:", openPorts)
	fmt.Println("Total ports found:", len(openPorts))
	fmt.Println("First two ports (slicing [0:2]):", openPorts[0:2])

	// Iterating over a slice with 'range'
	fmt.Println("\n--- 🔍 Iterating through Ports with range ---")
	for idx, port := range openPorts {
		fmt.Printf("  [%d] Scanned Port: %d\n", idx, port)
	}

	// 2. Maps: The fast lookup dictionary
	fmt.Println("\n--- 📖 Port Services Map (Lookup Dictionary) ---")
	portServices := map[int]string{
		21:   "FTP",
		22:   "SSH",
		80:   "HTTP",
		443:  "HTTPS",
		3306: "MySQL",
	}

	// Looking up a known port
	fmt.Println("Port 443 service is:", portServices[443])

	// The "Comma OK" idiom: checking if a key exists!
	unknownPort := 8888
	serviceName, exists := portServices[unknownPort]
	if exists {
		fmt.Printf("Port %d is %s\n", unknownPort, serviceName)
	} else {
		fmt.Printf("Port %d is NOT in our dictionary (serviceName=%q, exists=%t)\n",
			unknownPort, serviceName, exists)
	}

	// 3. Structs: The custom Lego blueprint
	fmt.Println("\n--- 🧱 Target Server Structs ---")
	prodServer := TargetServer{
		HostName:    "api.securecorp.internal",
		IPAddress:   "10.14.2.100",
		OpenPorts:   openPorts,
		IsProtected: true,
		RiskScore:   25,
	}

	fmt.Printf("Server Host: %s (%s)\n", prodServer.HostName, prodServer.IPAddress)
	fmt.Printf("Firewall Active: %t | Risk Score: %d/100\n", prodServer.IsProtected, prodServer.RiskScore)
	fmt.Printf("Open Ports Count: %d\n", len(prodServer.OpenPorts))
}
