# 📚 Go Arrays Cheat Sheet

Arrays in Go are **fixed-length** sequences of elements of the **same type**.

---

## 🔹 Declaring Arrays

```go
var arr [5]int          // Array of 5 integers, default initialized to 0
var names [3]string     // Array of 3 strings
```

---

## 🔹 Initializing Arrays

### Inline Initialization:

```go
numbers := [3]int{1, 2, 3}
```

### Let Go infer the length:

```go
colors := [...]string{"red", "green", "blue"}
```

---

## 🔹 Accessing and Modifying Elements

```go
arr[0] = 10             // Set value
fmt.Println(arr[0])     // Get value
```

---

## 🔹 Length of Array

```go
len(arr) // Returns the number of elements
```

---

## 🔹 Iterating Over Arrays

### Using `for` loop with index:

```go
for i := 0; i < len(arr); i++ {
    fmt.Println(arr[i])
}
```

### Using `range`:

```go
for i, val := range arr {
    fmt.Printf("Index %d: %d\n", i, val)
}
```

---

## 🔹 Multidimensional Arrays

```go
var matrix [2][3]int

matrix[0][0] = 1
matrix[1][2] = 5

fmt.Println(matrix)
```

---

## 🔹 Array Copying (by Value)

Arrays are **copied by value** (deep copy).

```go
a := [3]int{1, 2, 3}
b := a

b[0] = 100

fmt.Println(a) // [1 2 3]
fmt.Println(b) // [100 2 3]
```

---

## 🛠️ Tips

- Arrays have **fixed size**. If you need a **resizable array**, use **slices** instead.
- The type `[3]int` is **not** the same as `[4]int`. The length is part of the type.

---

## ✅ Summary

| Concept           | Syntax/Example                   |
| ----------------- | -------------------------------- |
| Declare           | `var a [5]int`                   |
| Initialize        | `b := [3]string{"x", "y", "z"}`  |
| Access            | `a[0] = 10`                      |
| Iterate (`for`)   | `for i := 0; i < len(a); i++ {}` |
| Iterate (`range`) | `for i, val := range a {}`       |
| Length            | `len(a)`                         |
| Copying           | Arrays are copied by value       |
| Multidimensional  | `var mtx [2][2]int`              |
| Fixed Size        | Length is part of the array type |

---

> 💡 Use arrays when you know the exact number of elements you need, otherwise go for slices!
