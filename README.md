# spanneranalyzerwasm2go

**Generated code. Do not edit by hand.**

This repository holds the pure-Go form of the DDL-validation part of the
[Cloud Spanner Emulator](https://github.com/GoogleCloudPlatform/cloud-spanner-emulator).
The emulator's C++ is built for `wasm32-wasip1` with
[wasmify](https://github.com/goccy/wasmify), then the wasm binary is compiled
ahead-of-time into Go source (plus Go assembly) with
[wasm2go](https://github.com/goccy/wasm2go). No cgo, no C toolchain, and no
wasm runtime are needed at build or run time — it is a Go package like any
other. The API entry point lives in
[go-spanner-analyzer](https://github.com/tyzerrr/go-spanner-analyzer); this
repository only carries the generated body.

This follows the same pattern as
[goccy/googlesqlwasm2go](https://github.com/goccy/googlesqlwasm2go), which
compiles Google's GoogleSQL (formerly ZetaSQL) engine to wasm and then to
pure Go the same way. That project covers query analysis for BigQuery/
GoogleSQL; this one covers DDL parsing and semantic validation for Spanner,
built from the emulator instead.

## Provenance

| Source | Version |
|---|---|
| cloud-spanner-emulator | `fc811a1` (patched to drop the PostgreSQL dialect, gRPC, and google-cloud-cpp — see `patches/` in go-spanner-analyzer) |
| ICU | 76.1 (data trimmed to case folding, basic normalization, and the root collation only — see `tools/icu-keep.txt`) |
| go-spanner-analyzer (generation config) | v0.2.0 — adds `AnalyzeQuery`, so queries and DML are resolved and type-checked against the schema the DDL describes |
| wasm input | 17.4 MB → this Go output |

The generation pipeline (wasmify config, patches, ICU build, wasm2go
invocation) lives in go-spanner-analyzer's `Makefile` and `tools/`. This
repository is the pipeline's output, published on its own so consumers don't
need Bazel, wasmify, or wasm2go to `go get` the result.

## Layout

```
data.bin           wasm linear-memory initializer (embedded)
alias.go           cross-package function aliases (go:linkname)
wasm2go.go         module constructor, memory setup
base/              runtime support shared by every part (imports, memory helpers)
p0/ .. p11/        the converted code, split into 12 parts
```

`base`, `p0`–`p11`, and the repository root are each a **separate Go module**,
not just separate packages. A single Go module is capped at 500 MiB of
source (`module source tree too large`), and the converted output here is
about 749 MB. Splitting it was the only way to publish it as one `go get`-able
unit. Dependencies run one direction only — `p0`..`p11` depend on `base`,
and the root depends on `base` and `p11` — so there is no cycle to worry
about when one part imports a symbol defined in another (`alias.go` re-exports
those via `go:linkname`).

Each converted part carries three implementations of the same functions,
selected by Go's own build tags — nothing project-specific:

| File | Build constraint | Used when |
|---|---|---|
| `arm64.s` / `decls_arm64.go` | `arm64` | Apple Silicon, arm64 Linux |
| `amd64.s` / `decls_amd64.go` | `amd64 && amd64.v2` | amd64 with `GOAMD64=v2`+ |
| `*_pure.go` | everything else | default amd64 (`GOAMD64=v1`), other CPUs |

wasm2go emits hand-tuned Go assembly where it can, because it runs several
times faster than the equivalent portable Go. The `_pure.go` files are the
fallback for architectures without a hand-tuned path.

## An upstream wasm2go bug found while building this

Early arm64 builds produced corrupted assembly: stray text like
`github.com/tyzerrr/go-spanner-N(RSP)` appeared inside `p0/arm64.s`, tens of
thousands of lines of it, while `amd64.s` stayed clean. The first hypothesis —
that placing the generated code under a Go module's own path triggered it —
turned out to be wrong. Isolating the variable with a small throwaway project
(four combinations: import path inside vs. outside a module, with vs.
without a hyphen) showed the real trigger is a **hyphen in the wasm2go
output's import path**; nesting under a module is unrelated. That is why this
package's import path is `spanneranalyzerwasm2go`, not
`go-spanner-analyzer-wasm2go` or similar. We intend to report this to
[goccy/wasm2go](https://github.com/goccy/wasm2go) with the repro.

## Regenerating

This repository is not meant to be edited directly. To reproduce or update
it, run the pipeline in go-spanner-analyzer (`make classify build headers
bridge proto wasm go`), then split the output into nested modules under the
500 MiB limit and re-tag each one at a matching version. See
go-spanner-analyzer's `PROGRESS.md` for the exact sequence used to cut this
version.

## License

Apache-2.0. Upstream (Google) attribution is in `NOTICE` and
`THIRD_PARTY_NOTICES.txt`. "Spanner" is a trademark of Google LLC; this
project is independent and not affiliated with or endorsed by Google.
