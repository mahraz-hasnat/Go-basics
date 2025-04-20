# 📌 Go Pointers Cheat Sheet

Pointers in Go allow you to **reference memory addresses** directly — powerful for optimizing memory usage and modifying values inside functions.

---

## 🧠 What is a Pointer?

A pointer stores the **memory address** of a value.

```go
var a int = 42
var p *int = &a // p points to a
```

---

## 🔍 Declaring Pointers

```go
var x int = 10
var p *int = &x
```

- `&x` gives the **address of `x`**
- `*p` dereferences the pointer (gets the value at the address)

---

## 🧪 Example

```go
x := 5
ptr := &x         // pointer to x
fmt.Println(*ptr) // prints: 5
*ptr = 10         // modifies x
fmt.Println(x)    // prints: 10
```

---

## 🎯 Pointer vs Value Function Parameters

### Pass by value (copy):

```go
func modify(val int) {
    val = 100
}
```

### Pass by reference (pointer):

```go
func modify(val *int) {
    *val = 100
}
```

```go
x := 10
modify(&x)
fmt.Println(x) // 100
```

---

## 🛠️ New Keyword

Creates a pointer to a zero-value of the type.

```go
p := new(int) // *int, initialized to 0
*p = 25
fmt.Println(*p) // 25
```

---

## 💡 Struct with Pointers

```go
type Person struct {
    Name string
    Age  int
}

func (p *Person) Birthday() {
    p.Age++
}
```

> Use pointer receivers to avoid copying and to modify the struct.

---

## 🚫 Nil Pointers

Uninitialized pointers are `nil`.

```go
var p *int
if p == nil {
    fmt.Println("Pointer is nil")
}
```

---

## ✅ Summary

| Concept           | Syntax / Description              |
| ----------------- | --------------------------------- |
| Declare pointer   | `var p *int = &x`                 |
| Dereference       | `*p`                              |
| Address-of        | `&x`                              |
| `new` function    | `p := new(int)`                   |
| Pass by reference | `func update(x *int) { *x = 10 }` |
| Pointer to struct | `p := &Person{Name: "Mehraz"}`    |
| Nil check         | `if p == nil {}`                  |

---

> 📌 **Use pointers when you want to modify a value or avoid copying large data.**
