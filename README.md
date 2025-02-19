# Go-Tutorial
[install go]{https://go.dev/doc/install}

1. Initialize a New Go Module
  go mod init <module-path>
  Example:
  "go mod init github.com/yourusername/yourproject"
  This command creates a go.mod file, initializing a new module with the specified path.

2. Build the Project
  "go build ./..."
  This command compiles the packages in your module without installing the binaries. The ./... pattern tells Go to include all packages in the module recursively.
3. Run the Project
  "go run ."
  This command compiles and runs the main package in the current directory. Ensure you have a main package with a main function.
4. Test the Project
  "go test ./..."
  This command runs all tests in your module recursively.
5. Benchmark the Project
  "go test -bench=. ./..."
  This command runs all benchmarks in your module. The -bench=. flag tells Go to run all benchmarks, and ./... includes all packages recursively.
  Note: Ensure your test files are named with the _test.go suffix and contain functions starting with Test for tests and Benchmark for benchmarks, following Go's testing conventions.
