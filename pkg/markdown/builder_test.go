package markdown

import (
	"bytes"
	"github.com/meteormin/friday.go/pkg/markdown/component"
	"github.com/stretchr/testify/assert"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"reflect"
	"testing"
)

func TestBuilder_Append(t *testing.T) {
	type args struct {
		c components.Component
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Alert: NOTE",
			args: args{
				c: components.NewAlert(components.NOTE, "This is a note"),
			},
		},
		{
			name: "Alert: TIP",
			args: args{
				c: components.NewAlert(components.TIP, "This is a tip"),
			},
		},
		{
			name: "Alert: WARNING",
			args: args{
				c: components.NewAlert(components.WARNING, "This is a warning"),
			},
		},
		{
			name: "Alert: IMPORTANT",
			args: args{
				c: components.NewAlert(components.IMPORTANT, "This is important"),
			},
		},
	}
	for _, tt := range tests {
		b := &Builder{
			components: []components.Component{},
		}

		t.Run(tt.name, func(t *testing.T) {
			b.Append(tt.args.c)
		})

		t.Log(b.String())
	}
}

func TestBuilder_Clear(t *testing.T) {
	type fields struct {
		components []components.Component
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Builder{
				components: tt.fields.components,
			}
			b.Clear()
		})
	}
}

func TestBuilder_HorizontalRule(t *testing.T) {
	type fields struct {
		components []components.Component
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "HorizontalRule",
			fields: fields{
				components: []components.Component{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Builder{
				components: tt.fields.components,
			}
			b.HorizontalRule()
			t.Log(b.String())
		})
	}
}

func TestBuilder_NewLine(t *testing.T) {
	type fields struct {
		components []components.Component
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "NewLine",
			fields: fields{
				components: []components.Component{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Builder{
				components: tt.fields.components,
			}
			b.NewLine()
			t.Log(b.String())
		})
	}
}

func TestBuilder_String(t *testing.T) {
	type fields struct {
		components []components.Component
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Empty",
			fields: fields{
				components: []components.Component{},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Builder{
				components: tt.fields.components,
			}
			if got := b.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuilder_WriteString(t *testing.T) {
	type fields struct {
		components []components.Component
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "WriteString",
			fields: fields{
				components: []components.Component{},
			},
			args: args{
				s: "Hello, World!",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Builder{
				components: tt.fields.components,
			}
			b.WriteString(tt.args.s)
		})
	}
}

func TestNewBuilder(t *testing.T) {
	tests := []struct {
		name string
		want *Builder
	}{
		{
			name: "NewBuilder",
			want: &Builder{
				components: []components.Component{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBuilder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuilder_Table(t *testing.T) {
	b := &Builder{
		components: []components.Component{},
	}

	table := components.NewTable()
	table.Header([]string{"ID", "Name", "Age", "Address"})
	table.Append([]string{"1", "Alice", "20", "New York"})
	table.Append([]string{"2", "Bob", "25", "Los Angeles"})
	table.Append([]string{"3", "Charlie", "30", "Chicago"})
	table.Append([]string{"4", "David", "35", "Houston"})
	table.Append([]string{"5", "Eve", "40", "Philadelphia"})

	b.Append(table)

	t.Log("\n" + b.String())
}

func TestMdToHtml(t *testing.T) {
	b := &Builder{
		components: []components.Component{},
	}

	table := components.NewTable()
	table.Header([]string{"ID", "Name", "Age", "Address"})
	table.Append([]string{"1", "Alice", "20", "New York"})
	table.Append([]string{"2", "Bob", "25", "Los Angeles"})
	table.Append([]string{"3", "Charlie", "30", "Chicago"})
	table.Append([]string{"4", "David", "35", "Houston"})
	table.Append([]string{"5", "Eve", "40", "Philadelphia"})

	b.Append(components.NewHeader(0, "Table"))
	b.Append(table)

	var buf bytes.Buffer
	md := goldmark.New(
		goldmark.WithExtensions(extension.Table),
		goldmark.WithParserOptions(parser.WithAutoHeadingID()),
	)

	err := md.Convert([]byte(b.String()), &buf)
	assert.Nil(t, err)

	t.Log("\n" + buf.String())
}
