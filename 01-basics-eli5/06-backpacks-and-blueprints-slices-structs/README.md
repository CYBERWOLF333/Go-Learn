# 🎒 Lesson 06: Backpacks, Decoders & Blueprints (Slices, Maps, Structs)

## 🧒 Explain Like I'm 5 (The Big Idea)

As your cybersecurity tools grow, single numbers and strings aren't enough. You need ways to organize piles of things!

---

### 1. The Stretchy Backpack: `Slices` 🎒
Imagine an egg carton with exactly 6 holes. That's an **Array**—stiff and unchangeable.
Now imagine a **Magic Backpack** that stretches every time you find new loot. That is a **Slice**!
- You can add things with `append(myBackpack, "new item")`.
- You can count how many things are inside with `len(myBackpack)`.
- You can peek at the first 3 items with `myBackpack[0:3]`.

---

### 2. The Secret Decoder Ring: `Maps` 📖
Imagine a dictionary or a spy phonebook:
You ask for a **Key** (like `"admin"`) and it immediately gives you the **Value** (like `"192.168.1.50"`).
No need to search line by line; maps give you instant answers!
```go
servers := map[string]string{
    "web": "10.0.0.1",
    "db":  "10.0.0.2",
}
```

---

### 3. The Lego Blueprint: `Structs` 🧱
What if you want to describe a whole Target Server? It has:
- A hostname (string)
- An IP address (string)
- A list of open ports (slice of ints)
- An active firewall (bool)

A **Struct** bundles all these different items together into one custom toy!

---

## 🏃 How to Run It

```bash
go run main.go
```
