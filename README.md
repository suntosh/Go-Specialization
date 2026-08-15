# Programming with Google Go — Specialization Coursework

Solutions to the programming assignments from the three-course **Programming with Google Go** specialization (University of California, Irvine, via Coursera).

The specialization introduces the Go programming language from Google and covers its distinguishing features — static typing with type inference, composition over inheritance, implicit interface satisfaction, and CSP-style concurrency built on goroutines and channels. Completing the sequence covers the ground needed to write concise, efficient, and clean Go applications.

![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go&logoColor=white)
![License](https://img.shields.io/badge/license-MIT-blue)

---

## Repository Layout

```
.
├── 01-getting-started/       # Course 1 — Getting Started with Go
├── 02-functions-methods/     # Course 2 — Functions, Methods, and Interfaces in Go
├── 03-concurrency/           # Course 3 — Concurrency in Go
├── go.mod
└── README.md
```

Each assignment lives in its own directory with a `main.go` and, where the assignment reads input, a sample data file.

---

## Course 1 — Getting Started with Go

Language fundamentals: the compilation and execution model, variables and type inference, the built-in type system, control flow, and Go's composite types.

**Topics covered**

- Compilation, the `go` toolchain, and workspace/module organization
- Variables, constants, type declarations, and type conversion
- Integers, floats, strings, runes, and the `strings` package
- Arrays, slices, maps, and their aliasing/growth semantics
- Structs and the `encoding/json` package
- Protocol-level basics: RFCs, Unicode, and text encoding

**Assignments**

| # | Assignment | Concepts |
|---|---|---|
| 1 | Prompt for a name and greet the user | I/O with `bufio`, string trimming |
| 2 | Read a name and address, emit a JSON object | Maps, `encoding/json`, marshalling |
| 3 | Read names from a file and sort by last then first name | File I/O, structs, slices, bubble sort |
| 4 | Animal query program (cow / bird / snake) | Maps of structs, command loops, string parsing |

---

## Course 2 — Functions, Methods, and Interfaces in Go

Moving from procedural code to Go's composition-based design: first-class functions, method sets, and interfaces satisfied implicitly rather than by declaration.

**Topics covered**

- Functions as values, closures, variadic parameters, and multiple returns
- Pointers, call-by-value semantics, and when a pointer receiver is required
- Methods, method sets, and the pointer/value receiver distinction
- Struct embedding and composition in place of inheritance
- Interfaces, implicit satisfaction, type assertions, and the empty interface
- Error handling as values

**Assignments**

| # | Assignment | Concepts |
|---|---|---|
| 1 | `GenDisplaceFn` — generate a displacement function from acceleration, initial velocity, and initial position | Closures, function values |
| 2 | In-place bubble sort on a slice of integers | Slices as reference types, call semantics |
| 3 | Read a file of people and print records via methods | Structs, methods, pointer receivers |
| 4 | `Animal` interface implemented by `Cow`, `Bird`, and `Snake` | Interfaces, implicit satisfaction, polymorphic dispatch |

---

## Course 3 — Concurrency in Go

Go's concurrency model and the synchronization primitives that make it safe: goroutines, channels, `select`, and the `sync` package.

**Topics covered**

- Processes, threads, and the goroutine scheduler
- Race conditions, mutual exclusion, and `sync.Mutex` / `sync.WaitGroup`
- Channels as typed conduits; blocking, buffering, and directionality
- `select`, timeouts, and default cases
- Classic synchronization problems and deadlock avoidance
- Sharing memory by communicating rather than communicating by sharing memory

**Assignments**

| # | Assignment | Concepts |
|---|---|---|
| 1 | Concurrent sort — partition a slice across four goroutines, sort each partition, then merge | Goroutines, `WaitGroup`, partitioning |
| 2 | Dining Philosophers with a host limiting concurrent diners | Mutexes, deadlock avoidance, resource ordering |

---

## Building and Running

Requires Go 1.22 or later.

```bash
# Run a single assignment
cd 03-concurrency/dining-philosophers
go run main.go

# Build a binary
go build -o bin/philosophers ./03-concurrency/dining-philosophers

# Vet and format everything
go vet ./...
gofmt -l .
```

For assignments that read input files, run from inside the assignment directory so relative paths resolve.

The concurrency assignments are worth running under the race detector, which is the fastest way to see why the synchronization is written the way it is:

```bash
go run -race main.go
```

---

## Notes

Solutions are written to the assignment specification as given, which occasionally means using a simpler algorithm than Go's standard library provides — the sorting assignments call for hand-written bubble and merge implementations rather than `sort.Slice`, because the point is the language mechanics, not the algorithm.

Where the grader accepted a narrower solution than the problem implied, the code here reflects the stated requirements rather than a generalized version.

---

## License

MIT. Course materials and assignment prompts remain the property of the University of California, Irvine and Coursera; only the solution code in this repository is covered by this license.

If you are currently taking this specialization, consider Coursera's Honor Code before reading solutions to assignments you have not yet submitted.
