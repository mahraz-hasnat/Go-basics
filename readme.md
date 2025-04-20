# 🧱 Go `struct` and Receiver Function Cheat Sheet

---

## 📦 What is a `struct`?

A `struct` is a composite data type that groups together variables under a single name.

```go
type Person struct {
    Name string
    Age  int
}
```

---

## 🧪 Creating and Using a `struct`

```go
p := Person{Name: "Mehraz", Age: 25}
fmt.Println(p.Name) // "Mehraz"
```

### Using the `var` keyword:

```go
var p Person
p.Name = "John"
p.Age = 30
```

---

## 🧰 Struct with Pointer

```go
p := &Person{Name: "Bob", Age: 28}
fmt.Println(p.Name) // Go automatically dereferences
```

---

## 🔁 Anonymous Struct

Useful for quick, inline use.

```go
user := struct {
    ID   int
    Name string
}{ID: 1, Name: "Test"}
```

---

## 🧩 Nested Structs

```go
type Address struct {
    City  string
    State string
}

type Employee struct {
    Name    string
    Address Address
}
```

---

## 📍 Receiver Functions (Methods on Structs)

Go doesn't have classes, but you can define **methods** on structs using **receiver functions**.

### 🔹 Value Receiver

```go
func (p Person) Greet() {
    fmt.Println("Hello, my name is", p.Name)
}
```

> Use when the method does not need to modify the struct.

---

### 🔹 Pointer Receiver

```go
func (p *Person) HaveBirthday() {
    p.Age++
}
```

> Use when you want to **modify** the struct or avoid copying large structs.

---

## 🎯 Example: Struct + Method

```go
type Circle struct {
    Radius float64
}

// Value receiver (read-only)
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

// Pointer receiver (modifies struct)
func (c *Circle) Resize(factor float64) {
    c.Radius *= factor
}
```

---

## 🧠 When to Use Value vs Pointer Receiver

| Use Case                 | Value Receiver  | Pointer Receiver    |
| ------------------------ | --------------- | ------------------- |
| Reading only             | ✅              | ✅                  |
| Modifying the struct     | ❌              | ✅                  |
| Large struct             | ❌              | ✅ (avoids copying) |
| Interface implementation | ✅/❌ (depends) | ✅ (preferred)      |

---

## 🔄 Struct Comparison

Structs can be compared **only if all fields are comparable**:

```go
type Point struct {
    X, Y int
}

p1 := Point{1, 2}
p2 := Point{1, 2}

fmt.Println(p1 == p2) // true
```

---

## 📑 Summary

| Concept          | Syntax / Description                 |
| ---------------- | ------------------------------------ |
| Define Struct    | `type Person struct { Name string }` |
| Create Struct    | `p := Person{Name: "John"}`          |
| Anonymous Struct | `s := struct{...}{}`                 |
| Nested Struct    | One struct inside another            |
| Value Receiver   | `(s Struct) MethodName()`            |
| Pointer Receiver | `(s *Struct) MethodName()`           |
| Method on Struct | Works like OOP class methods         |

---

> 💡 Use `struct` + receiver functions in Go as an alternative to OOP-style class behavior.
