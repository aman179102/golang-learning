//Functions are values too. They can be passed around just like other values. 
package main

import "fmt"

// -------------------------------------------------------------
// 1. COUNTER CLOSURE
// -------------------------------------------------------------
// This closure creates a counter that remembers how many times
// it has been called. Each counter has its OWN independent memory.
func makeCounter() func() int {
    count := 0 // local variable remembered by the closure
    return func() int {
        count++
        return count
    }
}

// -------------------------------------------------------------
// 2. MULTIPLIER CLOSURE
// -------------------------------------------------------------
// This closure generates a NEW function which remembers the "factor"
// we gave it. Each function gets its own multiplier value.
func makeMultiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor // closure remembers 'factor'
    }
}

// -------------------------------------------------------------
// 3. BANK ACCOUNT CLOSURE
// -------------------------------------------------------------
// This is a powerful real-world example.
// A closure stores and updates the account balance privately.
func createBankAccount() func(int) int {
    balance := 0 // private variable, not accessible outside
    
    return func(amount int) int {
        balance += amount // deposit (+) or withdraw (-)
        return balance
    }
}

// -------------------------------------------------------------
// 4. CACHE / MEMOIZATION CLOSURE
// -------------------------------------------------------------
// This closure remembers previous calculations.
// Extremely powerful for performance.
func makeSquareCache() func(int) int {
    cache := make(map[int]int) // remembers results
    
    return func(n int) int {
        if val, ok := cache[n]; ok {
            fmt.Println("Returning cached value for", n)
            return val
        }
        result := n * n // expensive operation (pretend)
        cache[n] = result
        return result
    }
}

// -------------------------------------------------------------
// MAIN PROGRAM → TEST ALL CLOSURES
// -------------------------------------------------------------
func main() {

    // 1. COUNTERS
    counterA := makeCounter()
    counterB := makeCounter()

    fmt.Println("Counter A:", counterA()) // 1
    fmt.Println("Counter A:", counterA()) // 2
    fmt.Println("Counter B:", counterB()) // 1 (independent!)

    fmt.Println("------------------------------------")

    // 2. MULTIPLIERS
    double := makeMultiplier(2)
    triple := makeMultiplier(3)

    fmt.Println("Double 10 =", double(10)) // 20
    fmt.Println("Triple 10 =", triple(10)) // 30

    fmt.Println("------------------------------------")

    // 3. BANK ACCOUNT
    account := createBankAccount()

    fmt.Println("Balance after deposit 100 =", account(100))
    fmt.Println("Balance after deposit 50  =", account(50))
    fmt.Println("Balance after withdraw 30 =", account(-30))

    fmt.Println("------------------------------------")

    // 4. CACHED FUNCTION
    square := makeSquareCache()

    fmt.Println("Square(5) =", square(5))
    fmt.Println("Square(10) =", square(10))
    fmt.Println("Square(5) Again =", square(5)) // will use cache
}
