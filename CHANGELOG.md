


# Changelog

## v0.5.0 (2025-10-02)

- Added `validations` submodule: robust validators (IsEmail, IsURL, IsUUID, IsAlpha, IsNumeric, IsAlnum, IsHex, IsIP, IsLower, IsUpper, IsASCII, IsPrintable, IsPhone, IsCreditCard) and type conversion helpers (ToString, ToInt, ToFloat64), with full documentation and tests.
- Added `concurrency` submodule: ParallelMap, ParallelForEach, Debounce, Throttle for easy parallelism and rate-limiting, with tests and documentation.
- Modularized project: each submodule has its own `go.mod` for clean imports and monorepo support.
- Added real-world usage examples in `examples/main.go` (validation, parallelism, conversion, analytics, throttling, etc.).
- Added `examples/bench_test.go` with benchmarks comparing fnkit utilities to stdlib approaches (map, filter, foreach, groupby, unique, validation).
- Improved documentation: cross-linked module READMEs, usage, and discoverability.

## v0.4.0 (2025-10-02)

- Added comprehensive string utility functions: Repeat, ReplaceAll, HasPrefix, HasSuffix, Contains, Count, ReverseString, IsAlpha, IsNumeric, Capitalize, StripLeft, StripRight, Partition, Rpartition, Words, CamelCase, SnakeCase, KebabCase, IsUpper, IsLower, SwapCase, Remove, Keep, PadCenter.
- All string utilities are Unicode-safe and fully tested.

## v0.3.0 (2025-10-02)

- Added Option[T] type for safe, idiomatic optional values (Rust-like Some/None).
- Added GroupBy, Chunk, Unique, and Flatten slice utilities.
- Added tests and documentation for these new utilities.



## v0.2.0 (2025-10-02)

- Added `Result[T]` type for Rust-like error handling in Go.
- Added robust tests for `Result[T]` including database and network call mocks.
- Improved documentation and usage examples for all utilities.



## v0.1.0

- Initial release: functional slice utilities (Map, Filter, Reduce, etc.) with generics.
