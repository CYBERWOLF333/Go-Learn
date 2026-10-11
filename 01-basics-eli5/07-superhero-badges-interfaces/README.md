# 🎖️ Lesson 07: Superhero Badges and Universal Plugs (Methods & Interfaces)

## 🧒 Explain Like I'm 5 (The Big Idea)

### 1. Methods: Giving Your Toys Special Action Buttons! 🎮
When you build a Lego robot, you don't just want it to sit there. You want to press a button on its chest and hear it speak!
In Go, a **Method** is a function that is permanently attached to a specific struct.

```go
func (r Robot) Speak() {
    fmt.Println("Beep boop!")
}
```

---

### 2. Interfaces: The Universal USB Plug! 🔌
Look at the side of your laptop. It has a USB-C port.
Does the laptop care if you plug in:
- A mechanical keyboard? ⌨️
- A gaming mouse? 🖱️
- A little desk fan? 🌀

**No!** As long as the device has the standard **USB-C plug**, the laptop happily talks to it!

In Go, an **Interface** is just a list of behaviors. 
If any struct can perform those behaviors, it automatically fits into the socket!

```go
type Scanner interface {
    RunScan(target string) string
}
```

If your `PortScanner` has a `RunScan(target string)` method, and your `DnsScanner` also has a `RunScan(target string)` method, they BOTH count as `Scanner`!
You don't even need to type `implements` like in Java! Go is smart enough to see they have the matching plug!

---

## 🏃 How to Run It

```bash
go run main.go
```
