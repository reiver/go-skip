# go-skip

Package **skip** provides tools for skipping leading input, for the Go programming language.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-skip

[![GoDoc](https://godoc.org/github.com/reiver/go-skip?status.svg)](https://godoc.org/github.com/reiver/go-skip)

## Example

Here is an example:

```golang
import "github.com/reiver/go-skip"

// ...

// In this example, we are skipping and leading space or tab characters.
//
// When you use it, you can choose other characters to skip.
err := skip.SkipRunes(runescanner, ' ', '\t')
```

## Import

To import package **skip** use `import` code like the follownig:
```
import "github.com/reiver/go-skip"
```

## Installation

To install package **skip** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-skip
```

## Author

Package **skip** was written by [Charles Iliya Krempeaux](http://reiver.link)
