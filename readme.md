# Mirage
Elegant command line interfaces for Go.

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
table := mirage.NewTable(4, 2)

table.Set(0, 0, "Year")
table.Set(0, 1, "Total revenue")

table.Set(1, 0, "2020")
table.Set(1, 1, "$68,400")

table.Set(2, 0, "2021")
table.Set(2, 1, "$99,200")

table.Set(3, 0, "2022")
table.Set(3, 1, "$55,600")

table.Print()
```

Result:
```
------------------------
| Year | Total revenue |
| 2020 | $68,400       |
| 2021 | $99,200       |
| 2022 | $55,600       |
------------------------
```