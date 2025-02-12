package markdown

import (
	"github.com/meteormin/friday.go/pkg/markdown/component"
	"reflect"
	"testing"
)

func TestBuilder_Append(t *testing.T) {
	type args struct {
		c component.Component
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Alert: NOTE",
			args: args{
				c: component.NewAlert(component.NOTE, "This is a note"),
			},
		},
		{
			name: "Alert: TIP",
			args: args{
				c: component.NewAlert(component.TIP, "This is a tip"),
			},
		},
		{
			name: "Alert: WARNING",
			args: args{
				c: component.NewAlert(component.WARNING, "This is a warning"),
			},
		},
		{
			name: "Alert: IMPORTANT",
			args: args{
				c: component.NewAlert(component.IMPORTANT, "This is important"),
			},
		},
	}
	for _, tt := range tests {
		b := &Builder{
			components: []component.Component{},
		}

		t.Run(tt.name, func(t *testing.T) {
			b.Append(tt.args.c)
		})

		t.Log(b.String())
	}
}

func TestBuilder_Clear(t *testing.T) {
	type fields struct {
		components []component.Component
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
		components []component.Component
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "HorizontalRule",
			fields: fields{
				components: []component.Component{},
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
		components []component.Component
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "NewLine",
			fields: fields{
				components: []component.Component{},
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
		components []component.Component
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Empty",
			fields: fields{
				components: []component.Component{},
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
		components []component.Component
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
				components: []component.Component{},
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
				components: []component.Component{},
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
		components: []component.Component{},
	}

	table := component.NewTable()
	table.Header([]string{"ID", "Name", "Age", "Address"})
	table.Append([]string{"1", "Alice", "20", "New York"})
	table.Append([]string{"2", "Bob", "25", "Los Angeles"})
	table.Append([]string{"3", "Charlie", "30", "Chicago"})
	table.Append([]string{"4", "David", "35", "Houston"})
	table.Append([]string{"5", "Eve", "40", "Philadelphia"})

	b.Append(table)

	t.Log("\n" + b.String())
}
