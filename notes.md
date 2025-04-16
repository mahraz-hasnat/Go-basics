# 🧠 Go Scope Cheat Sheet

## 🔹 1. Local Scope

- Declared **inside functions**, loops, or blocks (`{}`).
- **Accessible only** within that block.
- Commonly used for temporary or function-specific logic.

### Example:
```go
func greet() {
    message := "Hello, Go!" // local scope
    fmt.Println(message)
}
```

---

## 🔹 2. Global Scope (Package Level)

- Declared **outside of any function**, but **within the same package file**.
- Accessible to **all functions within that file/package**.
- Use sparingly to avoid unintended side-effects.

### Example:
```go
var appName = "FileCare" // global within the package

func printApp() {
    fmt.Println(appName)
}
```

---

## 🔹 3. Package Scope

- Any identifier (variable, constant, function, struct, etc.) declared **outside functions** is **visible within the entire package**.
- If the identifier starts with a **capital letter**, it is **exported** (visible from other packages).
- If it starts with a **lowercase letter**, it is **unexported** (visible only inside the same package).

### Example:
```go
// Inside filecare.go
package filecare

var internalVar = "Visible only in this package" // unexported
var PublicVar = "Accessible from other packages" // exported

func helperFunc() {}       // unexported
func ExportedFunc() {}     // exported
```

---

## ⚠️ Quick Tips

- Go does **not** use keywords like `private`, `public`, or `protected`.
- Instead, use **capitalization** to control visibility:
  - `myVar` → private/unexported
  - `MyVar` → public/exported
- Global variables should be avoided unless necessary. Prefer passing values through function parameters and return values.
