# 🚦 Lesson 03: Traffic Lights and Merry-Go-Rounds (If/Else & For Loops)

## 🧒 Explain Like I'm 5 (The Big Idea)

Computers are super obedient robots, but they don't have brains unless you give them choices and repetitions:

### 1. The Traffic Cop (`if` and `else`)
- 🟢 **IF** the light is green -> Drive!
- 🟡 **ELSE IF** the light is yellow -> Slow down!
- 🔴 **ELSE** -> Stop!

In Go, you don't even need parentheses `()` around your condition! It's super clean and uncluttered.

### 2. The Multi-Door Elevator (`switch`)
If you have 10 different floors to visit, a long chain of `if/else` looks like a messy spiderweb. A `switch` statement is like an elevator panel: press floor 1, 2, or 3, and you go straight there! And unlike C or Java, you don't need messy `break` statements on every line—Go breaks automatically!

### 3. The One Loop to Rule Them All (`for`)
In Python, Javascript, and C++, you have `while`, `do-while`, `for`, `for-each`... it's exhausting!
**Go only has ONE keyword for loops: `for`!**
- Want a classic count-up loop? `for i := 0; i < 5; i++ { ... }`
- Want a `while` loop? `for hungry { eatCookie() }`
- Want an infinite loop? `for { ... }` (use `break` when you want to get off the merry-go-round!)

---

## 🏃 How to Run It

```bash
go run main.go
```
