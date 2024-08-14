# Dotenv

Dotenv is a Go library that can load environment variables from .env files. This will allow for easier configuration management across different environments.

## Installation
To install the package, use `go get`:

```bash
go get github.com/codescalersinternships/dotenv-amr
```

## Usage

### Import the dotenv package

```go
import dotenv "github.com/codescalersinternships/dotenv-amr/pkg"
```

### Loading Environment Variables
Use the Load function to load environment variables from one or more .env files into the current process. If no filenames are provided, it defaults to using .env at the root of your project.
```go
err := dotenv.Load() // load from .env at the root of your project

err := dotenv.load("<fileName1.env>", "<fileName2.env>", ...) // load environment variables from one or more .env files
```
### Parsing a `.env` File

To parse a `.env` file and retrieve its contents as a map of environment variables, use the `Parse` function.

```go
envVars, err := dotenv.Parse(".env")
```

## Testing
To run the tests for this package, use the following command:

```bash
make test
```
