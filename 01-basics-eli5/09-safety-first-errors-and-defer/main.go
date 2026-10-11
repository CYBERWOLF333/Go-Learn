package main

import (
	"errors"
	"fmt"
	"os"
)

// Custom function that returns a meaningful error
func simulateLogin(username string, password string) error {
	if username == "" {
		return errors.New("empty username not permitted")
	}
	if password != "GopherCyber2026!" {
		return fmt.Errorf("authentication failed: invalid credentials for user '%s'", username)
	}
	return nil // nil means NO ERROR! Everything succeeded!
}

// Function demonstrating 'defer' clean up
func cleanupPlayroom() {
	fmt.Println("\n--- 🧹 Testing defer: Cleaning Up The Playroom ---")
	defer fmt.Println("  [defer 1] Locking the front door (runs LAST!)")
	defer fmt.Println("  [defer 2] Turning off the lights")
	defer fmt.Println("  [defer 3] Putting the Lego bricks back in the bin (runs FIRST!)")

	fmt.Println("Playing with Lego robots inside the room...")
	fmt.Println("Playtime is over, exiting the function now...")
}

// Function demonstrating panic and recover (Safety Net)
func safeDivide(a, b int) (result int, err error) {
	// Setup safety trampoline (recover) in case something panics!
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered from catastrophic failure: %v", r)
		}
	}()

	if b == 0 {
		panic("division by zero attempted!") // This triggers the panic alarm!
	}
	return a / b, nil
}

func main() {
	fmt.Println("=== 🛡️ Lesson 09: Errors, Defer & Recovery ===")

	// 1. Handling errors cleanly
	usersToTest := []struct {
		user string
		pass string
	}{
		{"", "12345"},
		{"neo", "wrongpass"},
		{"cyberwolf", "GopherCyber2026!"},
	}

	for _, u := range usersToTest {
		err := simulateLogin(u.user, u.pass)
		if err != nil {
			fmt.Printf("❌ Login rejected for '%s': %v\n", u.user, err)
		} else {
			fmt.Printf("✅ Login APPROVED for '%s'! Access granted.\n", u.user)
		}
	}

	// 2. Demonstration of defer
	cleanupPlayroom()

	// 3. Panic & Recover demonstration
	fmt.Println("\n--- 🦺 Testing Panic & Recover (The Safety Net) ---")
	res, err := safeDivide(10, 2)
	fmt.Printf("10 / 2 = %d (error: %v)\n", res, err)

	resBad, errBad := safeDivide(10, 0)
	fmt.Printf("10 / 0 = %d (error: %v)\n", resBad, errBad)

	// Writing a temporary test file and closing with defer
	fmt.Println("\n--- 📝 Safe File Handling with defer ---")
	tmpFile, err := os.CreateTemp("", "gopher-audit-*.txt")
	if err != nil {
		fmt.Println("Failed to create temp file:", err)
		return
	}
	// Always defer closing and removing the temporary file!
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	tmpFile.WriteString("Audit Log: System secure at 100% integrity.\n")
	fmt.Printf("Wrote safe audit log to: %s (will auto-clean on exit!)\n", tmpFile.Name())
}
