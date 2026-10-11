package main

import (
	"errors"
	"fmt"
	"strings"
)

// 1. Simple function taking inputs and returning a single value
func createCodename(prefix string, id int) string {
	return fmt.Sprintf("%s-%03d", prefix, id)
}

// 2. Go's Famous Feature: Multiple Return Values!
// Splits an address like "192.168.1.1:8080" into ("192.168.1.1", 8080, error)
func parseEndpoint(endpoint string) (string, int, error) {
	parts := strings.Split(endpoint, ":")
	if len(parts) != 2 {
		return "", 0, errors.New("invalid endpoint format, must be host:port")
	}

	host := parts[0]
	var port int
	_, err := fmt.Sscanf(parts[1], "%d", &port)
	if err != nil || port < 1 || port > 65535 {
		return "", 0, errors.New("port must be a valid number between 1 and 65535")
	}

	return host, port, nil
}

// 3. Named Return Values (Documented at the top of the function)
func calculateStats(scores []int) (min int, max int, avg float64) {
	if len(scores) == 0 {
		return 0, 0, 0.0
	}
	min, max = scores[0], scores[0]
	total := 0
	for _, score := range scores {
		if score < min {
			min = score
		}
		if score > max {
			max = score
		}
		total += score
	}
	avg = float64(total) / float64(len(scores))
	// A bare return returns whatever min, max, and avg currently hold!
	return
}

// 4. Variadic Function: Accepts ANY number of arguments (...)
func auditFirewallRules(firewallName string, ports ...int) {
	fmt.Printf("Auditing Firewall [%s] for %d rules...\n", firewallName, len(ports))
	for _, p := range ports {
		fmt.Printf("  -> Rule allowing Port %d\n", p)
	}
}

func main() {
	fmt.Println("=== 🥤 Lesson 04: Functions & Multiple Returns ===")

	// Calling single return function
	agent := createCodename("CYBER-GOPHER", 7)
	fmt.Println("Registered Agent:", agent)

	// Calling multiple return function
	fmt.Println("\n--- 🌐 Parsing Network Endpoints ---")
	validTarget := "10.0.0.1:443"
	host, port, err := parseEndpoint(validTarget)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Parsed Successfully! Host: %s, Port: %d\n", host, port)
	}

	// Calling with bad input to see error handling
	badTarget := "broken-endpoint-without-port"
	_, _, err = parseEndpoint(badTarget)
	if err != nil {
		fmt.Println("Handled expected error nicely:", err)
	}

	// Named returns demo
	fmt.Println("\n--- 📊 Statistics with Named Returns ---")
	latencyLog := []int{24, 18, 55, 32, 19, 88}
	lowest, highest, average := calculateStats(latencyLog)
	fmt.Printf("Latencies: Min=%dms, Max=%dms, Avg=%.1fms\n", lowest, highest, average)

	// Variadic function demo
	fmt.Println("\n--- 🛡️ Variadic Firewall Audit ---")
	auditFirewallRules("DMZ-Perimeter", 22, 80, 443, 8443)
}
