package mirage

import (
	"testing"
)

func Test(t *testing.T) {
	table := NewTable(3, 4)

	table.SetTitle(0, "Employee Name")
	table.SetTitle(1, "Department")
	table.SetTitle(2, "Salary ($/year)")
	table.SetTitle(3, "Years of service")

	table.Set(0, 0, "Alex Smith")
	table.Set(0, 1, "Engineering")
	table.Set(0, 2, "$75,000")
	table.Set(0, 3, "3")

	table.Set(1, 0, "Maria Garcia")
	table.Set(1, 1, "Marketing")
	table.Set(1, 2, "$68,000")
	table.Set(1, 3, "5")

	table.Set(2, 0, "Sam Chen")
	table.Set(2, 1, "Finance")
	table.Set(2, 2, "$82,000")
	table.Set(2, 3, "2")

	table.Print()
}

func Benchmark(b *testing.B) {
	for b.Loop() {
		table := NewTable(3, 4)

		table.SetTitle(0, "Employee Name")
		table.SetTitle(1, "Department")
		table.SetTitle(2, "Salary ($/year)")
		table.SetTitle(3, "Years of service")

		table.Set(0, 0, "Alex Smith")
		table.Set(0, 1, "Engineering")
		table.Set(0, 2, "$75,000")
		table.Set(0, 3, "3")

		table.Set(1, 0, "Maria Garcia")
		table.Set(1, 1, "Marketing")
		table.Set(1, 2, "$68,000")
		table.Set(1, 3, "5")

		table.Set(2, 0, "Sam Chen")
		table.Set(2, 1, "Finance")
		table.Set(2, 2, "$82,000")
		table.Set(2, 3, "2")

		table.Print()
	}
}
