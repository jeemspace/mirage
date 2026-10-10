package mirage

import (
	"fmt"
	"strings"
)

type Table struct {
	rows, columns    int
	grid             [][]string
	Separator        string
	Padding          int
	DoHorizontalBars bool
	TrimRows         bool
}

// Creates a new table with the specified amount of empty rows and columns.
func NewTable(rows int, columns int) Table {
	// Create the table's grid.
	grid := make([][]string, rows)

	for i := range rows {
		grid[i] = make([]string, columns)
	}

	return Table{
		rows:    rows,
		columns: columns,
		grid:    grid,

		// The separator to use in between each cell in a column.
		Separator: "|",
		// The number of spaces to use between each cell's value and the separator
		Padding: 1,
		// Whether to add a horizontal bar on the top and bottom of the table.
		DoHorizontalBars: true,
		// Whether to trim the leading and trailing padding for each rows.
		// Best for when no separator is being used.
		TrimRows: false,
	}
}

// Sets the value of a specific cell in the table.
func (table *Table) Set(row int, column int, value string) error {
	if table.InBounds(row, column) {
		table.grid[row][column] = value
		return nil
	}

	return fmt.Errorf("Table position is out of bounds.")
}

// Returns true if the specified cell is within the bounds of the table, or false if it does not exist.
func (table *Table) InBounds(row int, column int) bool {
	// Make sure the row isn't out of bounds
	if 0 <= row && row < table.rows {
		// Make sure the column isn't out of bounds
		if 0 <= column && column < table.columns {
			return true
		}
	}

	return false
}

// Formats and prints the stylized table.
func (table Table) Print() {
	fmt.Print(table.Format())
}

// Formats and returns a stylized and printable table.
func (table Table) Format() string {
	// Find the width of the table
	width := table.getWidth()
	var builder strings.Builder

	if table.DoHorizontalBars {
		// Print the first horizontal bar
		for range width {
			builder.WriteByte('-')
		}
		builder.WriteByte('\n')
	}

	for row := range table.rows {

		for column := range table.columns {
			columnWidth, _ := table.getColumnWidth(column)
			str := table.grid[row][column]

			builder.WriteString(table.Separator)

			// Add the leading padding to the cell
			if !table.TrimRows || column > 0 {
				for range table.Padding {
					builder.WriteByte(' ')
				}
			}

			if str == "" {
				// The cell's value is empty, so fill it in with spaces.
				for range columnWidth {
					builder.WriteByte(' ')
				}
			} else {
				builder.WriteString(str)

				// If the cell's value is not as long as the longest string in the column, then fill in the difference
				// with spaces.
				if len(str) < columnWidth {
					fill := columnWidth - len(str)

					for range fill {
						builder.WriteByte(' ')
					}
				}
			}

			// Add the trailing padding to the cell
			if !(column == (table.columns-1) && table.TrimRows) {
				for range table.Padding {
					builder.WriteByte(' ')
				}
			}
		}

		builder.WriteString(table.Separator)
		builder.WriteByte('\n')
	}

	if table.DoHorizontalBars {
		// Print the second horizontal bar
		for range width {
			builder.WriteByte('-')
		}
		builder.WriteByte('\n')
	}

	return builder.String()
}

func (table Table) getWidth() int {
	separatorLength := len(table.Separator)

	// Immediately count the separator length into the width because there is always a separator on the left side
	// before anything is printed.
	width := separatorLength

	// Find the biggest length of string in each column to calculate the width of that column,
	// then add them all together to calculate the total width of the table.
	for column := range table.columns {
		columnWidth, _ := table.getColumnWidth(column)

		if table.TrimRows && (column == 0 || column == (table.columns-1)) {
			// If trimming rows is enabled, and if the first of last column is being counted, then only count
			// the padding once, because there will be no trailing padding.
			width += columnWidth + separatorLength + table.Padding
		} else {
			// Multiply padding by two, because there is padding on both sides of a cell.
			width += columnWidth + separatorLength + (table.Padding * 2)
		}
	}

	return width
}

func (table Table) getColumnWidth(column int) (int, error) {
	width := 0

	// TODO: Check if the column is within bounds

	biggestItem := 0

	for row := range table.rows {
		str := table.grid[row][column]
		length := len(str)

		if length > biggestItem {
			biggestItem = length
		}
	}

	width += biggestItem

	return width, nil
}
