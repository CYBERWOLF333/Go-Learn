# 🥤 Lesson 04: Magic Vending Machines (Functions & Multiple Returns)

## 🧒 Explain Like I'm 5 (The Big Idea)

Imagine a magical vending machine:
- You drop in **flour, sugar, and chocolate** (these are **Parameters** / Inputs).
- The machine whirls and beeps...
- Out pops a **warm chocolate cupcake!** (this is the **Return Value** / Output).

### 🌟 Go's Superpower: Dual Dispenser! (Multiple Returns)
In most older programming languages, a vending machine can only give you ONE thing at a time.
In Go, functions can pop out **two or more things at once**!

For example, when Go tries to fetch a website, it gives you:
1. The website data 📄
2. An error note (if something went wrong, like the Wi-Fi being unplugged!) ❌

```go
func fetchWebsite(url string) (string, error) {
    // Returns BOTH the content AND any error!
}
```

This is the foundation of how Go avoids chaotic crashes: instead of throwing messy exceptions that blow up your app, Go calmly hands you `(result, error)` so you can inspect it!

---

## 🏃 How to Run It

```bash
go run main.go
```
