.PHONY: all test run-basics run-cyber run-final clean

all: test run-final

test:
	@echo "🧪 Running unit tests across all lessons and projects..."
	go test -v ./02-cyber-projects/...
	@cd 03-final-project-gophercon && go test -v ./...

run-basics:
	@echo "--- 🐹 Running Lesson 01: Hello Gopher ---"
	@go run 01-basics-eli5/01-hello-gopher/main.go
	@echo "\n--- 📦 Running Lesson 02: Variables & Boxes ---"
	@go run 01-basics-eli5/02-variables-and-boxes/main.go
	@echo "\n--- 🚦 Running Lesson 03: Traffic Lights & Loops ---"
	@go run 01-basics-eli5/03-traffic-lights-and-loops/main.go
	@echo "\n--- 🥤 Running Lesson 04: Vending Machines (Functions) ---"
	@go run 01-basics-eli5/04-magic-vending-machines-functions/main.go
	@echo "\n--- 🗺️ Running Lesson 05: Treasure Maps (Pointers) ---"
	@go run 01-basics-eli5/05-treasure-maps-pointers/main.go
	@echo "\n--- 🎒 Running Lesson 06: Backpacks & Blueprints (Slices, Maps, Structs) ---"
	@go run 01-basics-eli5/06-backpacks-and-blueprints-slices-structs/main.go
	@echo "\n--- 🎖️ Running Lesson 07: Superhero Badges (Interfaces) ---"
	@go run 01-basics-eli5/07-superhero-badges-interfaces/main.go
	@echo "\n--- 🐜 Running Lesson 08: Army of Ants (Goroutines & Channels) ---"
	@go run 01-basics-eli5/08-army-of-ants-goroutines-channels/main.go
	@echo "\n--- 🛡️ Running Lesson 09: Safety First (Errors & Defer) ---"
	@go run 01-basics-eli5/09-safety-first-errors-and-defer/main.go
	@echo "\n--- ✉️ Running Lesson 10: Secret Envelopes (JSON & HTTP) ---"
	@go run 01-basics-eli5/10-secret-messages-json-and-web/main.go

run-cyber:
	@echo "\n=============================================="
	@echo "  CYBER PROJECT 1: HashSentry (File Integrity)"
	@echo "=============================================="
	@go run 02-cyber-projects/01-hash-sentry/main.go
	@echo "\n=============================================="
	@echo "  CYBER PROJECT 2: HeaderAuditor (HTTP Security)"
	@echo "=============================================="
	@go run 02-cyber-projects/02-header-auditor/main.go
	@echo "\n=============================================="
	@echo "  CYBER PROJECT 3: StealthPortScanner (TCP Recon)"
	@echo "=============================================="
	@go run 02-cyber-projects/03-stealth-port-scanner/main.go
	@echo "\n=============================================="
	@echo "  CYBER PROJECT 4: PacketLens (Packet Inspection)"
	@echo "=============================================="
	@go run 02-cyber-projects/04-packet-lens/main.go

run-final:
	@echo "\n=============================================="
	@echo "  FINAL PROJECT: GopherRecon (OSINT Engine)"
	@echo "=============================================="
	@cd 03-final-project-gophercon && go run ./cmd/gopherrecon -demo

clean:
	@rm -rf 03-final-project-gophercon/bin
	@rm -f 03-final-project-gophercon/demo_recon_report.*
