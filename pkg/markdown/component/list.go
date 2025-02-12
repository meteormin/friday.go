package component

import "strings"

type List struct {
	items []Component
	level int
}

func (l *List) Clear() {
	l.items = []Component{}
}

func (l *List) Append(item Component) {
	l.items = append(l.items, item)
}

func (l *List) string(tab, bullet string) string {
	var sb strings.Builder
	for _, item := range l.items {
		if li, ok := item.(*List); ok {
			sb.WriteString(li.String())
		} else {
			for i := 0; i < l.level; i++ {
				sb.WriteString(tab)
			}
			sb.WriteString(bullet)
			sb.WriteString(" ")
			sb.WriteString(item.String())
		}
	}

	return sb.String()
}

func (l *List) String() string {
	return l.string("\t", "-")
}

func NewList(items ...Component) *List {
	return &List{
		items: items,
		level: 0,
	}
}

type TaskList struct {
	*List
}

func (tl *TaskList) String() string {
	return tl.string("\t", "- [ ]")
}

func NewTaskList(items ...Component) *TaskList {
	return &TaskList{
		List: &List{
			items: items,
			level: 0,
		},
	}
}
