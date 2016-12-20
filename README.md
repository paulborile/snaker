# snaker

[![Build Status](https://travis-ci.org/serenize/snaker.svg?branch=master)](https://travis-ci.org/serenize/snaker)
[![GoDoc](https://godoc.org/github.com/serenize/snaker?status.svg)](https://godoc.org/github.com/serenize/snaker)

This is a small utility to convert camel cased strings to snake case and back, except some defined words.

```
package snaker // import "snaker"
```

Package snaker provides methods to convert CamelCase names to snake_case and
back. It considers the list of allowed initialsms used by
github.com/golang/lint/golint (e.g. ID or HTTP)

```
func CamelToSnake(s string) string
func SnakeToCamel(s string) string
```

