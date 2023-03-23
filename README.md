# random

[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/dmitrymomot/random)](https://github.com/dmitrymomot/random)
[![Tests](https://github.com/dmitrymomot/random/actions/workflows/go.yml/badge.svg)](https://github.com/dmitrymomot/random/actions/workflows/go.yml)
[![CodeQL Analysis](https://github.com/dmitrymomot/random/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/dmitrymomot/random/actions/workflows/codeql-analysis.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/dmitrymomot/random)](https://goreportcard.com/report/github.com/dmitrymomot/random)
[![Go Reference](https://pkg.go.dev/badge/github.com/dmitrymomot/random.svg)](https://pkg.go.dev/github.com/dmitrymomot/random)
[![License](https://img.shields.io/github/license/dmitrymomot/random)](https://github.com/dmitrymomot/random/blob/main/LICENSE)

Small and simple helper to generate "random" string.

## Usage

Installation:
```
go get -u github.com/dmitrymomot/random
```
Use
```golang
import "github.com/dmitrymomot/random"

str := random.String(16)
log.Println(str)
```
Output:
```
Cb0ajMig6N7l9Fzf
```

---

Licensed under [Apache License 2.0](https://github.com/dmitrymomot/random/blob/master/LICENSE)
