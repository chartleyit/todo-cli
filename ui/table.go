package ui

import (
	"bytes"
	"fmt"
	"os"
	"text/tabwriter"
)

type TabTable struct {
	writer *tabwriter.Writer
}

func New() *TabTable {
	return &TabTable{
		writer: tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0),
	}
}

// TODO handle header lines

func (t *TabTable) AddLine(args ...interface{}) {
	formatString := t.buildFormatString(args)
	fmt.Fprintf(t.writer, formatString, args...)
}

func (t *TabTable) Print() {
	t.writer.Flush()
}

func (t *TabTable) buildFormatString(args []interface{}) string {
	var b bytes.Buffer
	for idx := range args {
		b.WriteString("%v")
		if idx+1 != len(args) {
			b.WriteString("\t")
		}
	}
	b.WriteString("\n")
	return b.String()
}
