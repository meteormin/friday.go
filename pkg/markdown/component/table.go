package component

import "strings"

type Table struct {
	header []string
	rows   [][]string
}

func (t *Table) Header(header []string) {
	t.header = header
}

func (t *Table) Append(row []string) {
	t.rows = append(t.rows, row)
}

func (t *Table) String() string {
	var sb strings.Builder
	var line strings.Builder

	if len(t.header) != 0 {
		sb.WriteString("|")
		line.WriteString("|")
		for _, h := range t.header {
			sb.WriteString(" " + h + " |")
			line.WriteString(" --- |")
		}

		sb.WriteString("\n")
		sb.WriteString(line.String() + "\n")
	}

	for _, row := range t.rows {
		sb.WriteString("|")
		for _, r := range row {
			sb.WriteString(" " + r + " |")
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func (t *Table) Clear() {
	t.header = []string{}
	t.rows = [][]string{}
}

func NewTable() *Table {
	return &Table{
		rows: make([][]string, 0),
	}
}
