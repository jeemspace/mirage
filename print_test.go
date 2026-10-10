package mirage

import (
	"os"
	"testing"
)

func Test(t *testing.T) {
	table := NewTable(4, 4)

	table.Separator = ""
	table.Padding = 1
	table.DoHorizontalBars = false
	table.TrimRows = true

	err := table.Set(0, 0, "Employee Name")
	err = table.Set(0, 1, "Department")
	err = table.Set(0, 2, "Salary ($/year)")
	err = table.Set(0, 3, "Years of service")

	err = table.Set(1, 0, "Alex Smith")
	err = table.Set(1, 1, "Engineering")
	err = table.Set(1, 2, "$75,000")
	err = table.Set(1, 3, "3")

	err = table.Set(2, 0, "Maria Garcia")
	err = table.Set(2, 1, "Marketing")
	err = table.Set(2, 2, "$68,000")
	err = table.Set(2, 3, "5")

	err = table.Set(3, 0, "Sam Chen")
	err = table.Set(3, 1, "Finance")
	err = table.Set(3, 2, "$82,000")
	err = table.Set(3, 3, "2")

	if err != nil {
		t.Error(err)
		os.Exit(-1)
	}

	table.Print()
}

func Benchmark(b *testing.B) {
	for b.Loop() {
		table := NewTable(4, 4)

		err := table.Set(0, 0, "Employee Name")
		err = table.Set(0, 1, "Department")
		err = table.Set(0, 2, "Salary ($/year)")
		err = table.Set(0, 3, "Years of service")

		err = table.Set(1, 0, "Alex Smith")
		err = table.Set(1, 1, "Engineering")
		err = table.Set(1, 2, "$75,000")
		err = table.Set(1, 3, "3")

		err = table.Set(2, 0, "Maria Garcia")
		err = table.Set(2, 1, "Marketing")
		err = table.Set(2, 2, "$68,000")
		err = table.Set(2, 3, "5")

		err = table.Set(3, 0, "Sam Chen")
		err = table.Set(3, 1, "Finance")
		err = table.Set(3, 2, "$82,000")
		err = table.Set(3, 3, "2")

		if err != nil {
			b.Error(err)
			os.Exit(-1)
		}

		table.Print()
	}
}
