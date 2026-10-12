# Mirage
Elegant command line interfaces for Go.

[![Go Reference](https://pkg.go.dev/badge/github.com/jeemspace/mirage.svg)](https://pkg.go.dev/github.com/jeemspace/mirage)

## Getting Started

### Installation
Import the repository in Go:
```go
import "github.com/jeemspace/mirage"
```

Then, run `go mod tidy` or `go get github.com/jeemspace/mirage` in your command line.

## Usage
A simple table:
```go
// Creates a table with 3 rows and 2 columns.
table := NewTable(3, 2)

// Set the title for each column
table.SetTitle(0, "Year")
table.SetTitle(1, "Total revenue")

// Set the values for each cell of the table.
table.Set(0, 0, "2020")
table.Set(0, 1, "$68,400")

table.Set(1, 0, "2021")
table.Set(1, 1, "$99,200")

table.Set(2, 0, "2022")
table.Set(2, 1, "$55,600")

// Formats and prints the table.
table.Print()
```

Result:
```
  Year   Total revenue 
------------------------
| 2020 | $68,400       |
| 2021 | $99,200       |
| 2022 | $55,600       |
------------------------
```