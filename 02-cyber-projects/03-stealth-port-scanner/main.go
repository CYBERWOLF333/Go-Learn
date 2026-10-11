package main

import (
	"flag"
	"fmt"
	"net"
	"sort"
	"strings"
	"sync"
	"time"
)

type PortResult struct {
	Port    int
	IsOpen  bool
	Banner  string
	Service string
}

// Well-known default ports mapping
var CommonServices = map[int]string{
	21:   "FTP",
	22:   "SSH",
	23:   "Telnet",
	25:   "SMTP",
	53:   "DNS",
	80:   "HTTP",
	110:  "POP3",
	143:  "IMAP",
	443:  "HTTPS",
	3306: "MySQL",
	5432: "PostgreSQL",
	6379: "Redis",
	8080: "HTTP-Proxy",
	8443: "HTTPS-Alt",
}

// GrabBanner sends an initial probe to see if the service greets us
func GrabBanner(conn net.Conn) string {
	conn.SetDeadline(time.Now().Add(500 * time.Millisecond))
	// Some protocols greet immediately (SSH, FTP, SMTP)
	// For web servers, send a minimal HEAD request
	fmt.Fprintf(conn, "HEAD / HTTP/1.0\r\n\r\n")

	buf := make([]byte, 256)
	n, err := conn.Read(buf)
	if err != nil || n == 0 {
		return ""
	}

	banner := string(buf[:n])
	banner = strings.ReplaceAll(banner, "\r", "")
	lines := strings.Split(banner, "\n")
	if len(lines) > 0 && strings.TrimSpace(lines[0]) != "" {
		return strings.TrimSpace(lines[0])
	}
	return ""
}

// Worker function that pulls ports from the jobs channel
func worker(host string, timeout time.Duration, grabBanner bool, jobs <-chan int, results chan<- PortResult, wg *sync.WaitGroup) {
	defer wg.Done()

	for port := range jobs {
		target := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.DialTimeout("tcp", target, timeout)
		if err != nil {
			// Port is closed or filtered
			continue
		}

		serviceName := CommonServices[port]
		if serviceName == "" {
			serviceName = "Unknown"
		}

		bannerText := ""
		if grabBanner {
			bannerText = GrabBanner(conn)
		}
		conn.Close()

		results <- PortResult{
			Port:    port,
			IsOpen:  true,
			Banner:  bannerText,
			Service: serviceName,
		}
	}
}

// ScanPorts coordinates the worker pool and returns open ports
func ScanPorts(host string, startPort, endPort int, workerCount int, timeout time.Duration, grabBanner bool) []PortResult {
	totalPorts := endPort - startPort + 1
	jobs := make(chan int, workerCount*2)
	results := make(chan PortResult, totalPorts)
	var wg sync.WaitGroup

	// 1. Launch worker pool
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(host, timeout, grabBanner, jobs, results, &wg)
	}

	// 2. Feed ports into the jobs channel in a background goroutine
	go func() {
		for port := startPort; port <= endPort; port++ {
			jobs <- port
		}
		close(jobs)
	}()

	// 3. Close results channel once all workers finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// 4. Collect results
	var openPorts []PortResult
	for res := range results {
		openPorts = append(openPorts, res)
	}

	// Sort ports ascending
	sort.Slice(openPorts, func(i, j int) bool {
		return openPorts[i].Port < openPorts[j].Port
	})

	return openPorts
}

func main() {
	hostFlag := flag.String("host", "", "Target host to scan (e.g. 127.0.0.1)")
	startPort := flag.Int("start", 1, "Starting port number")
	endPort := flag.Int("end", 1024, "Ending port number")
	workers := flag.Int("workers", 50, "Number of concurrent worker goroutines")
	timeoutStr := flag.String("timeout", "400ms", "Connection timeout duration (e.g. 300ms, 1s)")
	banner := flag.Bool("banner", true, "Attempt banner grabbing on open ports")
	flag.Parse()

	timeout, err := time.ParseDuration(*timeoutStr)
	if err != nil {
		fmt.Println("Invalid timeout:", err)
		return
	}

	if *hostFlag == "" {
		fmt.Println("ℹ️  No target host specified. Running Local Mock Demonstration...")

		// Spin up a couple of mock local listeners to show real scanning!
		l1, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			fmt.Println("Error creating test listener:", err)
			return
		}
		defer l1.Close()

		l2, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			fmt.Println("Error creating test listener 2:", err)
			return
		}
		defer l2.Close()

		p1 := l1.Addr().(*net.TCPAddr).Port
		p2 := l2.Addr().(*net.TCPAddr).Port

		// Responder for listener 1
		go func() {
			for {
				c, err := l1.Accept()
				if err != nil {
					return
				}
				c.Write([]byte("SSH-2.0-MockSecureShell_1.0\n"))
				c.Close()
			}
		}()

		fmt.Printf("[+] Spawned mock services on 127.0.0.1 ports: %d, %d\n", p1, p2)
		minPort := p1
		if p2 < minPort {
			minPort = p2
		}
		maxPort := p1
		if p2 > maxPort {
			maxPort = p2
		}

		fmt.Printf("[*] Scanning range %d to %d with 20 workers...\n", minPort-2, maxPort+2)
		startT := time.Now()
		found := ScanPorts("127.0.0.1", minPort-2, maxPort+2, 20, 200*time.Millisecond, true)
		elapsed := time.Since(startT)

		fmt.Printf("\n--- 🏁 SCAN RESULTS (%d open ports found in %v) ---\n", len(found), elapsed)
		for _, p := range found {
			bannerDisplay := p.Banner
			if bannerDisplay == "" {
				bannerDisplay = "(No banner response)"
			}
			fmt.Printf("  [OPEN] Port %-5d | Service: %-10s | Banner: %s\n", p.Port, p.Service, bannerDisplay)
		}

		fmt.Println("\nTip: Run with '-host <ip> -start <num> -end <num>' to scan custom targets!")
		return
	}

	fmt.Printf("[*] Starting concurrent scan on %s (Ports %d - %d, Workers: %d, Timeout: %v)...\n",
		*hostFlag, *startPort, *endPort, *workers, timeout)

	startT := time.Now()
	openPorts := ScanPorts(*hostFlag, *startPort, *endPort, *workers, timeout, *banner)
	elapsed := time.Since(startT)

	fmt.Printf("\n--- 🏁 SCAN RESULTS (%d open ports found in %v) ---\n", len(openPorts), elapsed)
	if len(openPorts) == 0 {
		fmt.Println("  No open ports discovered in this range.")
	} else {
		for _, p := range openPorts {
			bannerDisplay := p.Banner
			if bannerDisplay == "" {
				bannerDisplay = "(No banner response)"
			}
			fmt.Printf("  [OPEN] Port %-5d | Service: %-10s | Banner: %s\n", p.Port, p.Service, bannerDisplay)
		}
	}
}
