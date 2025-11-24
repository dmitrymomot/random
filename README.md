# random

[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/dmitrymomot/random)](https://github.com/dmitrymomot/random/v2)
[![Tests](https://github.com/dmitrymomot/random/v2/actions/workflows/tests.yml/badge.svg)](https://github.com/dmitrymomot/random/v2/actions/workflows/tests.yml)
[![CodeQL Analysis](https://github.com/dmitrymomot/random/v2/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/dmitrymomot/random/v2/actions/workflows/codeql-analysis.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/dmitrymomot/random/v2)](https://goreportcard.com/report/github.com/dmitrymomot/random/v2)
[![Go Reference](https://pkg.go.dev/badge/github.com/dmitrymomot/random/v2.svg)](https://pkg.go.dev/github.com/dmitrymomot/random/v2)
[![License](https://img.shields.io/github/license/dmitrymomot/random)](https://github.com/dmitrymomot/random/v2/blob/main/LICENSE)

Utilities for generating random strings, cryptographically secure one-time passwords, and performing weighted random selections from slices, structs, and maps.

## Installation

```bash
go get -u github.com/dmitrymomot/random/v2
```

## Security Warning

This package provides two different random generation approaches. Choose the right one for your use case:

- **`String()`** - Uses `math/rand` for fast, deterministic random string generation. **NOT cryptographically secure**. Use only for non-security purposes like IDs, test data, and display strings.
- **`OTP()`** - Uses `crypto/rand` for cryptographically secure one-time password generation. Suitable for security-sensitive operations like authentication tokens, password reset links, and MFA codes.
- **Probability functions** - Use `math/rand` for performance. NOT suitable for security-sensitive applications. Use only for game mechanics, loot tables, simulations, and other non-security purposes.

## Quick Start

```go
package main

import (
	"fmt"
	"log"

	"github.com/dmitrymomot/random/v2"
)

func main() {
	// Generate a random alphanumeric string (default charset)
	randomID := random.String(16)
	fmt.Println("Random ID:", randomID)

	// Generate a cryptographically secure OTP
	otp, err := random.OTP()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("6-digit OTP:", otp)

	// Weighted random selection
	items := []any{"apple", "banana", "cherry"}
	probs := []float64{0.5, 0.3, 0.2}
	selected := random.GetRandomWithProbabilities(items, probs)
	fmt.Println("Selected:", selected)
}
```

## Random String Generation

The `String()` function generates random strings using `math/rand`. It is fast and suitable for non-security purposes.

### Basic Usage

```go
package main

import (
	"fmt"
	"github.com/dmitrymomot/random/v2"
)

func main() {
	// Default alphanumeric string of length 16
	str := random.String(16)
	fmt.Println(str) // Output: Cb0ajMig6N7l9Fzf
}
```

### Custom Character Sets

```go
// Uppercase only
str := random.String(10, random.Uppercase)
fmt.Println(str) // Output: ABCDEFGHIJ

// Lowercase only
str := random.String(10, random.Lowercase)
fmt.Println(str) // Output: abcdefghij

// Numeric only
str := random.String(8, random.Numeric)
fmt.Println(str) // Output: 12345678

// Hexadecimal
str := random.String(16, random.Hex)
fmt.Println(str) // Output: a1b2c3d4e5f6a7b8

// Custom symbols
str := random.String(12, random.Symbols)
fmt.Println(str) // Output: !@#$%^&*()-_

// Combining multiple charsets
str := random.String(20, random.Uppercase, random.Numeric)
fmt.Println(str) // Output: A1B2C3D4E5F6G7H8I9J0
```

## Cryptographically Secure OTP Generation

The `OTP()` function generates cryptographically secure one-time passwords using `crypto/rand`. It is suitable for security-sensitive operations.

### Default 6-Digit OTP

```go
package main

import (
	"fmt"
	"log"
	"github.com/dmitrymomot/random/v2"
)

func main() {
	otp, err := random.OTP()
	if err != nil {
		log.Fatal("OTP generation failed:", err)
	}
	fmt.Println("OTP:", otp) // Output: 123456
}
```

### Custom Length OTP

```go
// 8-digit OTP
otp, err := random.OTP(8)
if err != nil {
	log.Fatal(err)
}
fmt.Println("OTP:", otp) // Output: 12345678

// 10-digit OTP
otp, err := random.OTP(10)
if err != nil {
	log.Fatal(err)
}
fmt.Println("OTP:", otp) // Output: 1234567890
```

### Error Handling

```go
otp, err := random.OTP()
if err != nil {
	// Handle cryptographic entropy failure
	// This should rarely happen on properly configured systems
	log.Fatal("Failed to generate OTP:", err)
}
```

## Weighted Random Selection

### Slice with Custom Probabilities

Use `GetRandomWithProbabilities()` to select a random item from a slice with custom probability weights:

```go
package main

import (
	"fmt"
	"github.com/dmitrymomot/random/v2"
)

func main() {
	items := []any{"common", "uncommon", "rare", "epic"}
	probabilities := []float64{0.5, 0.3, 0.15, 0.05}

	// Simulate 10 drops
	for i := 0; i < 10; i++ {
		selected := random.GetRandomWithProbabilities(items, probabilities)
		fmt.Println("Drop:", selected)
	}
}
```

### Struct with Probability Interface

Use `GetRandomStructWithProbabilities()` when items implement `GetProbability() float64`:

```go
package main

import (
	"fmt"
	"github.com/dmitrymomot/random/v2"
)

type LootItem struct {
	Name string
	Rarity float64
}

func (l LootItem) GetProbability() float64 {
	return l.Rarity
}

func main() {
	loot := []interface{ GetProbability() float64 }{
		LootItem{"Iron Sword", 0.6},
		LootItem{"Silver Sword", 0.3},
		LootItem{"Gold Sword", 0.1},
	}

	selected := random.GetRandomStructWithProbabilities(loot)
	fmt.Println("Loot:", selected)
}
```

### Map with Probability Weights

Use `GetRandomMapItemWithProbabilities()` to select from a map where values are probability weights:

```go
package main

import (
	"fmt"
	"github.com/dmitrymomot/random/v2"
)

func main() {
	items := map[string]float64{
		"common":   0.7,
		"uncommon": 0.2,
		"rare":     0.1,
	}

	selected := random.GetRandomMapItemWithProbabilities(items)
	fmt.Println("Selected:", selected)
}
```

### Map with Percentage Drops

Use `GetRandomMapItemWithPercent()` when map values are percentages (0-100):

```go
package main

import (
	"fmt"
	"github.com/dmitrymomot/random/v2"
)

func main() {
	dropTable := map[string]float64{
		"common_drop":   50.0,
		"uncommon_drop": 30.0,
		"rare_drop":     15.0,
		"epic_drop":     5.0,
	}

	// Each key has a percentage chance to drop
	selected := random.GetRandomMapItemWithPercent(dropTable)
	fmt.Println("Drop:", selected)
}
```

## Available Charset Constants

The package provides predefined character set constants for common use cases:

```go
random.Uppercase    // "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
random.Lowercase    // "abcdefghijklmnopqrstuvwxyz"
random.Alphabetic   // Uppercase + Lowercase
random.Numeric      // "0123456789"
random.Alphanumeric // Alphabetic + Numeric (default)
random.Symbols      // "`~!@#$%^&*()-_+={}[]|\;:"<>,./?`"
random.Hex          // "0123456789abcdef"
```

## API Reference

### String(length uint8, charsets ...string) string

Generates a random string of the specified length using the provided character sets.

- `length`: Length of the generated string (0-255)
- `charsets`: Optional character sets to use (default: Alphanumeric)
- Returns: Random string

### OTP(length ...int) (string, error)

Generates a cryptographically secure one-time password.

- `length`: Optional OTP length (default: 6)
- Returns: OTP string and error if generation fails

### GetRandomWithProbabilities(items []any, probabilities []float64) any

Selects a random item from a slice with custom probability weights.

- `items`: Slice of items to select from
- `probabilities`: Corresponding probability weights
- Returns: Selected item or nil if invalid input

### GetRandomStructWithProbabilities(items []interface{ GetProbability() float64 }) any

Selects a random item from a slice of structs that implement `GetProbability()`.

- `items`: Slice of items implementing GetProbability()
- Returns: Selected item or nil if invalid input

### GetRandomMapItemWithProbabilities(items map[string]float64) string

Selects a random key from a map where values are probability weights.

- `items`: Map with string keys and float64 probability values
- Returns: Selected key or empty string if invalid input

### GetRandomMapItemWithPercent(items map[string]float64) string

Selects a random key from a map where values are percentages (0-100).

- `items`: Map with string keys and float64 percentage values
- Returns: Selected key or empty string if invalid input

## Breaking Changes (v2)

This is v2 with breaking changes from v1:

- `OTP()` now returns `(string, error)` instead of just `string`
- `Random` struct was removed - use package-level functions directly
- Function renamed: `GetRandomMapItemWithPrecent()` → `GetRandomMapItemWithPercent()`

## License

Licensed under [Apache License 2.0](https://github.com/dmitrymomot/random/v2/blob/main/LICENSE)
