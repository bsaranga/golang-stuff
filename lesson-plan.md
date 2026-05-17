# Go: A to Z Lesson Plan

A staged curriculum based on `go-primer.md`. Each lesson is self-contained: concepts → hands-on exercise → checkpoint. Lessons live under `ch{N}/` directories with runnable `.go` files. You already have `ch1/hello.go` — that's lesson 0.

The arc: surface syntax → idiomatic Go → concurrency → runtime/perf → systems & arcane internals.

---

## Stage I — Fundamentals (syntax & types)

**L01. Hello, packages, and `go run`**
Packages, `main`, imports, `go run` vs `go build`, `go.mod` basics. Exercise: split hello into two packages.

**L02. Variables, constants, basic types**
`var`, `:=`, `const`, `iota`, numeric types, strings vs `[]byte`, runes. Exercise: temperature converter with typed constants.

**L03. Control flow**
`if` with init clause, `for` (the only loop), `switch` (incl. type switch teaser), labels. Exercise: FizzBuzz + prime sieve.

**L04. Functions & multiple return values**
Named returns, variadic, first-class functions, closures. Exercise: build a `map`/`filter`/`reduce` over `[]int`.

**L05. Pointers**
`&`, `*`, no pointer arithmetic, when to use value vs pointer. Exercise: in-place vs returning-new linked list ops.

**L06. Arrays, slices, the underlying array**
Length vs capacity, `append`, slice aliasing pitfalls, `copy`. Exercise: implement a ring buffer using a slice.

**L07. Maps**
Declaration, zero value, comma-ok idiom, deletion, iteration randomness. Exercise: word-frequency counter on a file.

**L08. Structs & methods**
Value vs pointer receivers, method sets, struct tags (intro). Exercise: `Vector2` with arithmetic methods.

**L09. Strings, runes, UTF-8**
`len` is bytes not runes, `range` over string, `strings`/`unicode` packages. Exercise: palindrome checker that handles multi-byte runes.

---

## Stage II — Idiomatic Go

**L10. Interfaces I — structural typing**
Defining, satisfying, `error`, `fmt.Stringer`. Exercise: shape area calculator with multiple types.

**L11. Interfaces II — small interfaces, consumer-side**
`io.Reader`/`io.Writer`, composition, accepting interfaces / returning structs. Exercise: implement a counting `io.Writer`.

**L12. Errors as values**
`errors.New`, `fmt.Errorf` with `%w`, `errors.Is`, `errors.As`, sentinels, typed errors. Exercise: parse-int with rich error wrapping.

**L13. `defer`, `panic`, `recover`**
Defer ordering, LIFO, panic semantics, when (rarely) to recover. Exercise: safe-divide using recover at an API boundary only.

**L14. Embedding & composition**
Field/method promotion, why it's NOT inheritance. Exercise: refactor an "inheritance" attempt into composition.

**L15. Generics (1.18+)**
Type parameters, constraints, `comparable`, when NOT to use. Exercise: generic `Set[T]` and `Map[K,V]`.

**L16. Project layout, `go.mod`, modules**
`go mod init/tidy`, semver, replace directives, internal packages. Exercise: split lessons into a multi-module workspace.

**L17. Testing fundamentals**
`testing.T`, table-driven tests, `t.Run`, `testing.TB`, fixtures, `go test -cover`. Exercise: TDD a tiny CSV parser.

**L18. Tooling tour**
`go fmt`, `go vet`, `golangci-lint`, `go doc`, `staticcheck`. Exercise: set up `.golangci.yml` and fix every flagged issue.

---

## Stage III — Concurrency

**L19. Goroutines — birth and death**
`go` keyword, lifecycle, "how does this goroutine exit?" Exercise: parallel sum with `sync.WaitGroup`.

**L20. Channels I — unbuffered, sends/receives, closing**
Synchronous handoff, close semantics, `v, ok := <-ch`. Exercise: producer/consumer with explicit close.

**L21. Channels II — buffered, `select`, `default`**
Backpressure, non-blocking ops, timeouts via `time.After` (and its alloc cost). Exercise: rate-limiter using a token-bucket channel.

**L22. `context.Context`**
Cancellation, deadlines, values (sparingly), propagation. Exercise: cancellable HTTP fetcher fan-out.

**L23. `sync` package**
`Mutex`, `RWMutex`, `WaitGroup`, `Once`, `Cond`. Exercise: thread-safe LRU cache.

**L24. `sync/atomic`**
Lock-free counters, `atomic.Value`, memory ordering basics. Exercise: high-throughput counter benchmarked against `Mutex`.

**L25. `errgroup` & structured concurrency**
`golang.org/x/sync/errgroup`, first-error-cancels-all. Exercise: parallel URL fetch with error propagation.

**L26. The race detector**
`go test -race`, identifying real data races, fixing them. Exercise: deliberately racy code → detect → fix.

**L27. Common concurrency patterns**
Fan-in/fan-out, pipelines, worker pools, semaphores via buffered channels, `sync.Pool`. Exercise: multi-stage pipeline with bounded parallelism.

**L28. Goroutine leak hunting**
`runtime.NumGoroutine()`, pprof goroutine profile, `go.uber.org/goleak`. Exercise: plant a leak, find it with pprof.

---

## Stage IV — Stdlib mastery

**L29. `io` & composition**
`io.Reader`, `io.Writer`, `io.Copy`, `io.Pipe`, multi-readers/writers, `io.TeeReader`. Exercise: gzip+checksum streaming via composed writers.

**L30. `bufio`, `bytes`, `strings.Builder`**
Buffering, scanners, efficient string assembly. Exercise: line-oriented log parser.

**L31. `encoding/binary` & `encoding/json`**
Fixed-width binary encoding, byte order, JSON struct tags, custom marshalers. Exercise: encode/decode a binary record format with CRC.

**L32. `os`, `path/filepath`, file I/O**
`os.OpenFile` flags, `Sync`, atomic rename, temp files. Exercise: atomic file writer (write-temp-then-rename).

**L33. `net/http` — clients and servers**
Handlers, middleware, `http.ServeMux` 1.22+, contexts, timeouts. Exercise: tiny REST API with graceful shutdown.

**L34. `log/slog` (1.21+)**
Structured logging, handlers, attributes, log levels. Exercise: replace `fmt.Println` everywhere with `slog`.

**L35. `time`**
`time.Time`, monotonic clocks, `Ticker` vs `Timer`, why `time.After` allocates. Exercise: backoff/retry helper.

---

## Stage V — Runtime, performance, & the dark arts

**L36. Memory model & happens-before**
Read the official Go Memory Model spec. Why atomics + channels work, why naked reads/writes don't. Exercise: write 3 programs the model says are racy; race-detect them.

**L37. Escape analysis**
`go build -gcflags="-m"`, value vs pointer perf, when interfaces force escape. Exercise: refactor an allocation-heavy function based on escape output.

**L38. GC tuning**
`GOGC`, `GOMEMLIMIT`, allocation rate, `sync.Pool` patterns. Exercise: reduce allocations 10× in a JSON parser microbench.

**L39. Profiling — CPU, heap, mutex, block**
`runtime/pprof`, `net/http/pprof`, flamegraphs via `go tool pprof -http`. Exercise: optimize a deliberately-slow program; document each win.

**L40. Microbenchmarks**
`go test -bench`, `-benchmem`, `benchstat`, avoiding compiler elimination. Exercise: bench `[]byte` vs `string` conversions; explain results.

**L41. The scheduler**
M:N model, P/M/G, work stealing, preemption (1.14+), `GOMAXPROCS`, `runtime.Gosched`. Exercise: build a CPU-bound goroutine and observe scheduling with trace.

**L42. `runtime/trace`**
Execution tracer, viewing in browser, diagnosing scheduler latency. Exercise: trace a chatty channel workload, identify a stall.

**L43. `unsafe`, `reflect`, `cgo` (sparingly)**
What they cost, when (rarely) they're justified, layout/alignment, `unsafe.Pointer` rules. Exercise: implement a zero-copy `[]byte`↔`string` conversion and discuss the hazards.

**L44. Build, link, embed**
Build tags, `//go:embed`, cross-compilation, ldflags for version stamping, trimpath. Exercise: single-binary tool that embeds an HTML asset and stamps git SHA at build.

**L45. Advanced testing**
Fuzzing (`go test -fuzz`), `testing/quick`, `testing/synctest` (Go 1.24+), golden files, deterministic fakes. Exercise: fuzz the binary record decoder from L31 to a crash.

---

## Stage VI — Systems & capstone (Spitfire-aligned)

**L46. File durability & fsync**
`O_DIRECT`, `fsync` vs `fdatasync`, write barriers, torn writes. Read: "Files Are Hard". Exercise: append-only log with explicit fsync; measure throughput vs durability.

**L47. Write-Ahead Log v0**
Record framing, CRC32C (Castagnoli), length-prefixing, recovery on partial tail. Exercise: implement and crash-test a tiny WAL.

**L48. Networking deep dive: TCP → QUIC**
`net.Conn`, deadlines, half-close, then `quic-go`. Exercise: echo server in TCP, port to QUIC.

**L49. Protobuf & wire protocols**
`google.golang.org/protobuf` v2, schema evolution, oneof, unknown fields. Exercise: design a minimal job submit/ack proto.

**L50. Raft via `hashicorp/raft`**
`LogStore`/`StableStore`/`FSM`, snapshots, joint config (peer changes). Exercise: 3-node in-memory KV via raft.

**L51. Observability — metrics, traces, pprof in prod**
OpenTelemetry, `expvar`, pprof endpoints behind auth. Exercise: instrument the WAL with traces + counters.

**L52. Fault injection & chaos**
Simulated disk errors, partial writes, clock skew, network drops. Exercise: property-test crash recovery with injected faults.

**L53. Capstone**
Stitch L47 + L50 + L48 into a tiny replicated append-only log served over QUIC with a CLI client. This is the seed of Spitfire.

---

## How we'll work each lesson

1. I open a `chNN/` directory with a `README.md` (objective, key idioms, gotchas) and a starter `.go` file.
2. You read the README, then attempt the exercise. I leave `// TODO` markers.
3. You ask for hints or a review when ready. I review against the idioms in `go-primer.md` §6, §9, §12.
4. We commit each completed lesson before moving on, so progress is auditable in `git log`.

Pace target: ~2–3 lessons/week part-time → roughly 5–6 months end-to-end, which lines up with the Spitfire schedule in §11 of the primer.

Tell me when to scaffold L01.
