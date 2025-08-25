# Concurrency and Parallelism in Go

### Concurrency vs Parallelism

- **Concurrency**: Allowing yourself to handle multiple tasks, but **not at the exact same time**.  
- **Parallelism**: Executing multiple tasks **at the same time**, not one after another.  

**Example (Instagram analogy):**  
Suppose you are eating rice and watching Instagram reels. Suddenly, you get a notification, and you also want to turn on the A.C.

- **Concurrency**:  
  You first check the notification, then turn on the A.C., and finally continue eating rice. Tasks are interleaved, but not simultaneous.  

- **Parallelism**:  
  You **simultaneously** eat rice, check the notification, and turn on the A.C. at the same time.  

<img src="./concurrencyvsparallel.png" alt="Concurrency vs Parallelism" width="400" />

---

### Goroutines

- **Goroutines**: Lightweight threads managed by the Go runtime.  
  - Flexible stack (starts at ~2KB).  
- **Threads**: Managed by the operating system.  
  - Fixed stack (usually ~1MB).  

✅ **In Go, goroutines are the way we achieve concurrency (and parallelism when multiple CPU cores are available).**  

The Go runtime can create many goroutines without asking the OS for new threads, giving Go more control and efficiency.

---

### Go’s Motto

> **"Do not communicate by sharing memory; instead, share memory by communicating."**

This philosophy is at the core of Go’s concurrency model.

---

### WaitGroups

When using goroutines, you often need to **wait for all of them to finish**. For this, Go provides **WaitGroups**.

A WaitGroup is like an advanced version of `time.Sleep()`, but more precise and safe.  

**Functions of WaitGroup:**
1. **Add(int)** → Increments the counter by the given number of goroutines.  
2. **Done()** → Decrements the counter by one when a goroutine finishes.  
3. **Wait()** → Blocks until the counter becomes zero (all goroutines finished).  

---

### Mutex (Mutual Exclusion)

- A **mutex** is a lock that ensures only **one goroutine** can access a shared resource at a time.  
- The zero value of a mutex is an unlocked mutex.  
- A mutex must not be copied after first use.  

**Lock/Unlock behavior:**  
- When one goroutine locks the mutex, no other goroutine can access that memory until it is unlocked.  

**Read-Write Mutex:**  
- Multiple goroutines can **read** a resource simultaneously.  
- But **only one goroutine can write**, and during writing, no one else can read or write.

---

### Race Conditions

A **race condition** happens when multiple goroutines access the same memory at the same time, and at least one of them writes to it.

👉 Use:  
```bash
go run --race .
````

This checks for race condition errors in a Go program (exit status 66 if found).
Race conditions can be solved using **mutexes**.

---

### Channels

* **Channels** allow goroutines to communicate with each other by passing values.
* A send operation (`ch <- value`) will block until another goroutine is ready to receive from the channel.
* Similarly, a receive operation (`<-ch`) will block until there is data to receive.
* This ensures synchronization but can lead to **deadlocks** if nobody is listening.

⚠️ Always design your channel communication carefully, often together with WaitGroups.

---


# 📝 Golang Interview Prep Q\&A Sheet

---

## **1. Basics**

**Q:** Difference between `var`, `:=`, and `const` in Go?
**A:**

* `var` → declares variable, type optional, works globally & locally.
* `:=` → short-hand declaration, **only inside functions**, type inferred.
* `const` → immutable, compile-time constant.

---

**Q:** What are zero values in Go?
**A:**

* `int → 0`
* `string → ""`
* `bool → false`
* `pointer/map/slice/channel/function → nil`

---

**Q:** Difference between array and slice?
**A:**

* Array = fixed length, value type, `[3]int{1,2,3}`.
* Slice = dynamic, reference type, internally stores (pointer + length + capacity).

---

## **2. Memory & Pointers**

**Q:** What is the difference between passing by value and passing by pointer in Go?
**A:**

* Go **always passes by value**.
* But when you pass a pointer (address), the callee can mutate the underlying object.

---

**Q:** What’s the difference between `new` and `make`?
**A:**

* `new(T)` → allocates zeroed storage for type `T`, returns `*T`.
* `make(T, args)` → used only for slices, maps, and channels. Initializes internal data structures.

---

## **3. Concurrency**

**Q:** Buffered vs Unbuffered channel?
**A:**

* **Unbuffered** → send blocks until receiver ready.
* **Buffered** → send blocks only when buffer full, receive blocks only when empty.

---

**Q:** What happens if you send to a channel with no receiver?
**A:**

* On **unbuffered channel** → deadlock.
* On **buffered channel** → blocks only if buffer is full.

---

**Q:** How do you prevent goroutines from leaking?
**A:**

* Always provide a way to cancel goroutines (`context.Context`, closing channels).
* Example:

```go
ctx, cancel := context.WithCancel(context.Background())
go func() {
    select {
    case <-ctx.Done():
        return
    }
}()
cancel()
```

---

**Q:** What’s the Go memory model guarantee for goroutines?
**A:**

* Without synchronization, goroutines may see stale values.
* Use **channels** or **sync primitives** (`Mutex`, `WaitGroup`, `atomic`) to guarantee visibility of writes across goroutines.

---

**Q:** What’s wrong with this code?

```go
for i := 0; i < 5; i++ {
    go func() {
        fmt.Println(i)
    }()
}
```

**A:** Closure captures loop variable → prints `5` multiple times.
✅ Fix:

```go
for i := 0; i < 5; i++ {
    go func(n int) {
        fmt.Println(n)
    }(i)
}
```

---

## **4. Context & Cancellation**

**Q:** Why is `context.Context` important?
**A:**

* For **cancellation, timeouts, deadlines, request-scoped values** across goroutines.
* Common in DB queries, API calls, servers.

Example:

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()
```

---

**Q:** Why is it bad to pass `context.Background()` inside functions directly?
**A:**

* Because you lose cancellation signals from parent.
* Always pass `ctx` from the caller.

---

## **5. Interfaces & OOP**

**Q:** How does Go handle polymorphism?
**A:**

* Interfaces are **satisfied implicitly**.
* Any type that has the required methods implements the interface.

---

**Q:** Difference between `interface{}` and `any` in Go?
**A:**

* Both are the same (Go 1.18+ introduced `any` as alias for `interface{}`).

---

**Q:** What is the difference between value receiver vs pointer receiver in methods?
**A:**

* Value receiver → works on a copy, cannot mutate original.
* Pointer receiver → can modify original, avoids copying large structs.

---

## **6. Errors & Testing**

**Q:** How does Go handle errors?
**A:**

* Errors are values implementing `error` interface.
* Idiomatic Go uses:

```go
if err != nil {
    return err
}
```

* Prefer wrapping errors (`fmt.Errorf("msg: %w", err)`).

---

**Q:** How do you test in Go?
**A:**

* Write `_test.go` files.
* Use `testing` package.

```go
func TestAdd(t *testing.T) {
    got := Add(2, 3)
    want := 5
    if got != want {
        t.Errorf("got %d, want %d", got, want)
    }
}
```

---

## **7. Go Runtime & Advanced**

**Q:** How does Go manage memory (GC)?
**A:**

* Go uses a **concurrent, tri-color, mark-and-sweep garbage collector**.
* Low pause times, optimized for latency.

---

**Q:** How are goroutines scheduled?
**A:**

* Go uses an **M\:N scheduler** (Goroutines mapped to OS threads).
* Components:

  * **M** = OS threads
  * **P** = processors (logical CPU context)
  * **G** = goroutines

---

**Q:** How to limit concurrency?
**A:**

* Use buffered channels as semaphores.
* Example:

```go
sem := make(chan struct{}, 5) // allow 5 goroutines max
for _, job := range jobs {
    sem <- struct{}{}
    go func(job Job) {
        defer func() { <-sem }()
        process(job)
    }(job)
}
```
