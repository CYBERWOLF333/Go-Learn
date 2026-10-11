package main

import (
	"fmt"
	"sync"
	"time"
)

// Simulated target probe
func probeTarget(workerID int, targetHost string, results chan<- string, wg *sync.WaitGroup) {
	// defer wg.Done() guarantees this runs when the function finishes!
	defer wg.Done()

	fmt.Printf("[Ant #%d] Starting probe on %s...\n", workerID, targetHost)
	// Simulate doing some work (like pinging a server)
	time.Sleep(100 * time.Millisecond)

	// Send message back through the walkie-talkie (channel)
	results <- fmt.Sprintf("Target %s is ALIVE (Checked by Ant #%d)", targetHost, workerID)
}

func main() {
	fmt.Println("=== 🐜 Lesson 08: Goroutines & Channels ===")

	targets := []string{
		"192.168.1.1",
		"192.168.1.25",
		"192.168.1.50",
		"192.168.1.100",
	}

	// 1. Create a Walkie-Talkie (Channel)
	// Buffered channel with capacity equal to targets count
	resultsChan := make(chan string, len(targets))

	// 2. Create the Chore Clipboard (WaitGroup)
	var wg sync.WaitGroup

	startTime := time.Now()

	// 3. Dispatch the ants!
	for i, target := range targets {
		wg.Add(1) // Tell clipboard: one more ant is working!
		go probeTarget(i+1, target, resultsChan, &wg)
	}

	// Wait for all ants to finish their job
	wg.Wait()
	close(resultsChan) // Close walkie-talkie when nobody has more messages to send

	// Read all messages sent back by the ants
	fmt.Println("\n--- 📻 Messages Received Over Walkie-Talkie ---")
	for report := range resultsChan {
		fmt.Println("📢", report)
	}

	elapsed := time.Since(startTime)
	fmt.Printf("\n⚡ All 4 probes completed concurrently in: %v!\n", elapsed)
	fmt.Println("(Notice: If run one by one, it would take ~400ms. With ants, ~100ms!)")
}
