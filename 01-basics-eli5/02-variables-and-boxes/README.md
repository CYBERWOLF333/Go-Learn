# 📦 Lesson 02: Variables and Toy Boxes

## 🧒 Explain Like I'm 5 (The Big Idea)

Imagine you have a bunch of clear plastic toy boxes.
On each box, there's a strict label:
- 🧸 **`string` box**: Can only hold words, letters, and sentences (always wrapped in `"quotes"`).
- 🔢 **`int` box**: Can only hold whole counting numbers (like `10`, `42`, or `-5`). No crumbs or decimals!
- 🥧 **`float64` box**: Holds numbers with decimal crumbs (like `3.14` or `99.9%`).
- 🔘 **`bool` box**: Can only hold one of two coins: `true` (YES) or `false` (NO).

If you try to jam a teddy bear (`"hacker"`) into the `int` number box, the Go compiler slaps your hand and says: **NOPE! That doesn't fit!** That is why Go is called a **statically typed** language: it keeps your toy room neat so things never break by surprise.

---

## 🎩 The Two Ways to Make a Box

### 1. The Polite Long Way: `var`
```go
var agentName string = "CyberWolf"
var securityLevel int = 99
```

### 2. The Speedy Shortcut: `:=` (The Walrus Operator!)
When you are inside a function, you can use `:=` (looks like a walrus with tusks `:=`). Go looks at what you put on the right side and guesses the box type automatically!
```go
agentName := "CyberWolf"  // Go knows this is a string!
securityLevel := 99       // Go knows this is an int!
isStealthMode := true     // Go knows this is a bool!
```

---

## 👻 The Magic "Zero Value"
In some messy languages (like C), if you make a box and put nothing in it, it's filled with random trash from the computer's memory.
In Go, every box gets scrubbed clean with a **Zero Value**:
- Numbers start at: `0`
- Strings start at: `""` (empty string)
- Booleans start at: `false`

---

## 🔒 Constants (`const`): Superglue on the Box!
If a box should NEVER be changed (like the speed of light or the default port for HTTPS), use `const`:
```go
const DefaultHTTPSPort = 443
```
If someone tries to change it later in the code, Go refuses to run!

---

## 🏃 How to Run It

```bash
go run main.go
```
