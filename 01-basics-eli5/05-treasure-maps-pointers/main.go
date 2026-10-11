package main

import "fmt"

// Function that receives a COPY (Photocopy)
func tryToHealCopy(shieldHealth int) {
	shieldHealth = 100 // Only changes the local photocopy!
	fmt.Printf("Inside tryToHealCopy: shieldHealth is %d\n", shieldHealth)
}

// Function that receives a POINTER (Treasure Map)
func healWithPointer(shieldHealth *int) {
	// The * means: follow the map to the real address, and set its value!
	*shieldHealth = 100
	fmt.Println("Inside healWithPointer: followed map and set actual memory to 100!")
}

type CyberDrone struct {
	Model       string
	BatteryPct  int
	Coordinates string
}

// Struct pointer method
func rechargeDrone(drone *CyberDrone) {
	drone.BatteryPct = 100
	drone.Coordinates = "Home Base Dock"
}

func main() {
	fmt.Println("=== 🗺️ Lesson 05: Pointers & Memory Maps ===")

	// 1. Basic Variable & Address
	secretCode := 1337

	// The & symbol gets the memory address (the treasure map)
	codePointer := &secretCode

	fmt.Printf("The Value:           %d\n", secretCode)
	fmt.Printf("The Memory Address:  %p\n", codePointer)
	fmt.Printf("Dereferencing map:   %d\n", *codePointer)

	// 2. Modifying through the pointer
	*codePointer = 9999
	fmt.Printf("After updating via pointer, secretCode is now: %d\n", secretCode)

	// 3. Pass by Value vs Pass by Pointer
	fmt.Println("\n--- 🛡️ Shield Repair Test ---")
	shield := 25
	fmt.Println("Initial shield:", shield)

	// Try with copy
	tryToHealCopy(shield)
	fmt.Println("After tryToHealCopy, original shield is STILL:", shield) // Unchanged!

	// Now with pointer
	healWithPointer(&shield)
	fmt.Println("After healWithPointer, original shield is NOW:", shield) // Fixed!

	// 4. Working with Structs and Pointers
	fmt.Println("\n--- 🛸 Drone Mission ---")
	myDrone := CyberDrone{
		Model:       "GopherRecon-X1",
		BatteryPct:  12,
		Coordinates: "Sector 7G",
	}

	fmt.Printf("Before Recharge: %+v\n", myDrone)
	rechargeDrone(&myDrone)
	fmt.Printf("After Recharge:  %+v\n", myDrone)
}
