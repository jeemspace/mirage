package mirage

import (
	"fmt"
	"strings"
)

type Table struct {
	rows, columns int

	grid [][]string
}

func NewTable(rows int, columns int) Table {
	// Create the table's grid.
	grid := make([][]string, rows)

	for i := range rows {
		grid[i] = make([]string, columns)
	}

	return Table{
		rows,
		columns,
		grid,
	}
}

func (table *Table) Set(row int, column int, value string) error {
	if table.InBounds(row, column) {
		table.grid[row][column] = value
		return nil
	}

	return fmt.Errorf("Table position is out of bounds.")
}

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

func (table Table) Print() {
	fmt.Print(table.Format())
}

// Formats and returns a stylized and printable table.
func (table Table) Format() string {
	// Find the width of the table
	width := table.getWidth()
	var builder strings.Builder

	for range width {
		builder.WriteByte('-')
	}
	builder.WriteByte('\n')

	for row := range table.rows {

		for column := range table.columns {
			columnWidth, _ := table.getColumnWidth(column)
			str := table.grid[row][column]

			if str == "" {
				builder.WriteString("| ")
				for range columnWidth {
					builder.WriteString(" ")
				}
				builder.WriteString(" ")
			} else {
				builder.WriteString("| ")
				builder.WriteString(str)

				if len(str) < columnWidth {
					padding := columnWidth - len(str)

					for range padding {
						builder.WriteString(" ")
					}
				}

				builder.WriteString(" ")
			}
		}

		builder.Write([]byte{'|', '\n'})
	}

	for range width {
		builder.WriteByte('-')
	}
	builder.WriteByte('\n')

	return builder.String()
}

func (table Table) getWidth() int {
	width := 0

	// Find the biggest length of string in each column to calculate the width of that column,
	// then add them all together to calculate the total width of the table.
	for column := range table.columns {
		cWidth, _ := table.getColumnWidth(column)
		width += cWidth + 3
	}

	return width + 1 // Add one for the final bar on the side
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
