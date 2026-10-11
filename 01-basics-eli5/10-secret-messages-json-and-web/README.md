# ✉️ Lesson 10: Secret Envelopes and Web Talk (Files, JSON & HTTP)

## 🧒 Explain Like I'm 5 (The Big Idea)

In the real world of cybersecurity, computers don't just calculate numbers in a vacuum. They talk to each other across the globe!

To talk nicely, computers need:
1. **A Common Secret Language: `JSON`** 📜
   Imagine folded paper notes passed between friends in class. JSON is the universal format that every computer understands:
   `{"agent": "CyberWolf", "status": "stealth"}`.
   In Go, we turn our Go structs into JSON envelopes using `json.Marshal`, and unpack them back into structs using `json.Unmarshal`.

2. **A Messenger Pigeon: `net/http`** 🕊️
   How does your computer visit a website or fetch threat intelligence?
   The Go standard library includes an industrial-strength web engine right inside `net/http`!
   - You can fetch any web page with `http.Get("https://...")`.
   - You can even run your OWN mini web server in just 5 lines of code!

---

## 🏃 How to Run It

```bash
go run main.go
```
