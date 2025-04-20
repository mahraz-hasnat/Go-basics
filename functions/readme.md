# 🔧 Types of Functions in Go

Go provides several types of functions to suit different programming needs. Below are the most common function types with examples.

---

## 1️⃣ **Basic Function**

A function with fixed parameters and a single return value.

```go
func greet(name string) {
    fmt.Println("Hello", name)
}
```

---

## 2️⃣ **Function with Return Value**

Functions can return values of any type.

```go
func add(a int, b int) int {
    return a + b
}
```

---

## 3️⃣ **Function with Multiple Return Values**

Useful when you want to return multiple results from a function.

```go
func divide(a, b int) (int, int) {
    return a / b, a % b
}

q, r := divide(10, 3) // q = 3, r = 1
```

---

## 4️⃣ **Named Return Values**

You can name return variables for clarity and use a naked return.

```go
func rectInfo(length, width float64) (area, perimeter float64) {
    area = length * width
    perimeter = 2 * (length + width)
    return
}
```

---

## 5️⃣ **Variadic Function**

Accepts a variable number of arguments of the same type.

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

total := sum(1, 2, 3, 4) // total = 10
```

---

## 6️⃣ **Anonymous Function (Function Literal)**

Functions defined without a name, often assigned to variables or called inline.

```go
func() {
    fmt.Println("I am anonymous")
}()
```

```go
multiply := func(a, b int) int {
    return a * b
}
fmt.Println(multiply(2, 3)) // 6
```

---

## 7️⃣ **Closure**

A function that captures variables from its surrounding scope.

```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

next := counter()
fmt.Println(next()) // 1
fmt.Println(next()) // 2
```

---

## 8️⃣ **Recursive Function**

A function that calls itself, often used in algorithms.

```go
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}
```

---

## 9️⃣ **Function as a Parameter**

Functions can be passed as arguments to other functions.

```go
func operate(a, b int, fn func(int, int) int) int {
    return fn(a, b)
}

result := operate(5, 3, func(x, y int) int {
    return x - y
})
// result = 2
```

---

## 🔟 **Function as a Return Value**

Functions can return other functions.

```go
func makeMultiplier(factor int) func(int) int {
    return func(n int) int {
        return n * factor
    }
}

double := makeMultiplier(2)
fmt.Println(double(4)) // 8
```

---

## 🔁 **init() Function**

- `init()` is a special function in Go that runs **automatically before `main()`**.
- Every Go file can have **multiple `init()` functions**.
- Commonly used for **setup**, like initializing variables, checking envs, etc.
- It **cannot take arguments or return values**.

```go
package main

import "fmt"

func init() {
    fmt.Println("init() called before main()")
}

func main() {
    fmt.Println("main() function")
}
```

**Output:**

```
init() called before main()
main() function
```

---

## 🏁 Summary Table

| Type                  | Description                             |
| --------------------- | --------------------------------------- |
| Basic Function        | Simple function with fixed parameters   |
| Return Function       | Returns one value                       |
| Multiple Returns      | Returns multiple values                 |
| Named Returns         | Named return variables and naked return |
| Variadic Function     | Accepts variable number of arguments    |
| Anonymous Function    | Function without a name                 |
| Closure               | Captures variables from outer scope     |
| Recursive Function    | Calls itself                            |
| Function as Parameter | Pass function as argument               |
| Function as Return    | Return a function from another function |
| `init()` Function     | Runs before `main()`, used for setup    |

---

> 🧠 **Tip:** Treat `init()` like a "constructor" for your package — great for initializing state or validating config before your program starts.
